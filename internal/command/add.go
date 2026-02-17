package command

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/m-mdy-m/TechShelf/internal/common"
	"github.com/m-mdy-m/TechShelf/internal/logger"
	"github.com/spf13/cobra"
)

// ─── Wizard ──────────────────────────────────────────────────────────────────

// wizard wraps a bufio.Reader and provides structured prompt helpers.
// One wizard instance is created per interactive session so the reader
// buffer is never split across calls.
type wizard struct {
	r *bufio.Reader
}

func newWizard() *wizard {
	return &wizard{r: bufio.NewReader(os.Stdin)}
}

func (w *wizard) readLine() (string, error) {
	line, err := w.r.ReadString('\n')
	return strings.TrimSpace(line), err
}

// ask prompts the user and returns their input.
// If input is empty and defaultVal is set, defaultVal is returned.
// If required is true and no input (and no default) is given, it re-prompts.
func (w *wizard) ask(label, defaultVal string, required bool) string {
	label = color.New(color.FgCyan).Sprint(label)
	for {
		if defaultVal != "" {
			fmt.Printf("  %s [%s]: ", label, color.New(color.FgWhite, color.Faint).Sprint(defaultVal))
		} else {
			fmt.Printf("  %s: ", label)
		}
		input, err := w.readLine()
		if err != nil {
			return defaultVal
		}
		if input == "" {
			if defaultVal != "" {
				return defaultVal
			}
			if required {
				color.New(color.FgYellow).Println("    ⚠  This field is required.")
				continue
			}
			return ""
		}
		return input
	}
}

// askList prompts for a comma-separated list, returning a []string.
func (w *wizard) askList(label string) []string {
	raw := w.ask(label+" (comma-separated, or blank to skip)", "", false)
	if raw == "" {
		return nil
	}
	parts := strings.Split(raw, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		if p = strings.TrimSpace(p); p != "" {
			out = append(out, p)
		}
	}
	return out
}

// askChoice shows a numbered menu and returns the chosen value.
func (w *wizard) askChoice(label string, choices []string, defaultVal string) string {
	fmt.Printf("\n  %s\n", color.New(color.FgCyan).Sprint(label))
	for i, c := range choices {
		marker := ""
		if c == defaultVal {
			marker = color.New(color.FgGreen).Sprint(" ← default")
		}
		fmt.Printf("    %s. %s%s\n",
			color.New(color.FgWhite, color.Faint).Sprintf("%d", i+1),
			c,
			marker,
		)
	}
	for {
		fmt.Print("  Choice (number or name): ")
		input, err := w.readLine()
		if err != nil {
			return defaultVal
		}
		if input == "" && defaultVal != "" {
			return defaultVal
		}
		// Numeric?
		if n, err := strconv.Atoi(input); err == nil && n >= 1 && n <= len(choices) {
			return choices[n-1]
		}
		// Exact match?
		for _, c := range choices {
			if strings.EqualFold(input, c) {
				return c
			}
		}
		color.New(color.FgYellow).Println("    ⚠  Invalid choice. Try again.")
	}
}

// askYear prompts for an integer year (negative = BCE). Returns nil on blank.
func (w *wizard) askYear() *int {
	for {
		raw := w.ask("Year published (e.g. 1984, -300 for 300 BCE, blank to skip)", "", false)
		if raw == "" {
			return nil
		}
		n, err := strconv.Atoi(raw)
		if err != nil {
			color.New(color.FgYellow).Println("    ⚠  Enter a number.")
			continue
		}
		return intPtr(n)
	}
}

// ─── Interactive wizard ──────────────────────────────────────────────────────

func runWizard(c *Catalog) (Book, error) {
	bold := color.New(color.FgCyan, color.Bold)
	bold.Println("\n" + strings.Repeat("═", 58))
	bold.Println("  📚  techShelf — Add a book")
	bold.Println(strings.Repeat("═", 58))
	color.New(color.FgWhite, color.Faint).Println("  Press Enter to skip optional fields.\n")

	wiz := newWizard()

	// ── Required ─────────────────────────────────────────────────────────────
	title := wiz.ask("Title", "", true)
	baseID := Slugify(title)
	id := wiz.ask("ID (slug)", baseID, true)
	id = c.UniqueID(id)

	authorsRaw := wiz.askList("Author(s)")
	if len(authorsRaw) == 0 {
		a := wiz.ask("Author (at least one)", "", true)
		authorsRaw = []string{a}
	}

	// ── Optional base fields ──────────────────────────────────────────────────
	year := wiz.askYear()
	language := wiz.ask("Language", "English", false)

	// ── Category / subcategory ───────────────────────────────────────────────
	categoryChoices := make([]string, 0, len(c.Taxonomy))
	for cat := range c.Taxonomy {
		categoryChoices = append(categoryChoices, cat)
	}
	if len(categoryChoices) == 0 {
		categoryChoices = []string{"Computer Science", "Mathematics", "Philosophy", "Other"}
	}
	category := wiz.askChoice("Category", categoryChoices, "")

	var subcategory string
	if subs, ok := c.Taxonomy[category]; ok && len(subs) > 0 {
		subcategory = wiz.askChoice("Subcategory", append(subs, "(none)"), "(none)")
		if subcategory == "(none)" {
			subcategory = ""
		}
	} else {
		subcategory = wiz.ask("Subcategory", "", false)
	}

	tags := wiz.askList("Tags (lowercase, hyphenated)")
	level := wiz.askChoice("Level", ValidLevels, "general")

	// ── Descriptions ─────────────────────────────────────────────────────────
	color.New(color.FgCyan).Println("\n  — Descriptions —")
	description := wiz.ask("What is this book?", "", false)
	whyRead := wiz.ask("Why read it?", "", false)
	whenToRead := wiz.ask("When to read it?", "", false)
	importance := wiz.ask("Why is it notable?", "", false)

	prereqs := wiz.askList("Prerequisites (book IDs or concepts)")
	pairs := wiz.askList("Pairs well with (book IDs)")

	// ── Source ───────────────────────────────────────────────────────────────
	color.New(color.FgCyan).Println("\n  — Source —")
	filePathRaw := wiz.ask("Local file path (relative to TechShelf root)", "", false)
	onlineURLRaw := wiz.ask("Online URL (PDF link, project site)", "", false)

	status := wiz.askChoice("Reading status", ValidStatuses, "unread")
	notes := wiz.ask("Personal notes", "", false)

	src := Source{}
	if filePathRaw != "" {
		src.FilePath = strPtr(filePathRaw)
	}
	if onlineURLRaw != "" {
		src.OnlineURL = strPtr(onlineURLRaw)
	}

	book := Book{
		ID:            id,
		Title:         title,
		Author:        authorsRaw,
		YearPublished: year,
		Language:      language,
		Category:      category,
		Subcategory:   subcategory,
		Tags:          tags,
		Level:         level,
		Description:   description,
		WhyRead:       whyRead,
		WhenToRead:    whenToRead,
		Importance:    importance,
		Prerequisites: prereqs,
		PairsWellWith: pairs,
		Source:        src,
		Status:        status,
		Notes:         notes,
		AddedDate:     time.Now().Format("2006-01-02"),
	}
	if book.Prerequisites == nil {
		book.Prerequisites = []string{}
	}
	if book.PairsWellWith == nil {
		book.PairsWellWith = []string{}
	}
	if book.Tags == nil {
		book.Tags = []string{}
	}

	// ── Preview ───────────────────────────────────────────────────────────────
	color.New(color.FgWhite, color.Faint).Println("\n" + strings.Repeat("─", 58))
	color.New(color.FgCyan).Println("  Preview:")
	color.New(color.FgWhite, color.Faint).Println(strings.Repeat("─", 58))
	preview, _ := json.MarshalIndent(book, "  ", "  ")
	fmt.Println("  " + string(preview))
	color.New(color.FgWhite, color.Faint).Println(strings.Repeat("─", 58))

	confirm := newWizard().ask("\nAdd this book? [Y/n]", "Y", false)
	if strings.ToUpper(confirm) != "Y" {
		return Book{}, fmt.Errorf("cancelled")
	}
	return book, nil
}

// ─── fromJSON mode ────────────────────────────────────────────────────────────

// bookFromJSON parses a partial JSON object into a Book, applying defaults.
func bookFromJSON(c *Catalog, raw string) (Book, error) {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte(raw), &fields); err != nil {
		return Book{}, fmt.Errorf("invalid JSON: %w", err)
	}

	required := []string{"title", "author", "category"}
	for _, f := range required {
		if _, ok := fields[f]; !ok {
			return Book{}, fmt.Errorf("missing required field: %q", f)
		}
	}

	// Unmarshal into a full Book to get all fields.
	var book Book
	if err := json.Unmarshal([]byte(raw), &book); err != nil {
		return Book{}, fmt.Errorf("parse book: %w", err)
	}

	// Normalise author: accept both string and []string.
	if _, ok := fields["author"]; ok {
		var single string
		if err := json.Unmarshal(fields["author"], &single); err == nil {
			book.Author = []string{single}
		}
	}

	// Auto-assign id, defaults.
	if book.ID == "" {
		book.ID = c.UniqueID(Slugify(book.Title))
	} else if c.IDExists(book.ID) {
		return Book{}, fmt.Errorf("id %q already exists in catalog", book.ID)
	}
	if book.Language == "" {
		book.Language = "English"
	}
	if book.Level == "" {
		book.Level = "general"
	}
	if book.Status == "" {
		book.Status = "unread"
	}
	if book.Prerequisites == nil {
		book.Prerequisites = []string{}
	}
	if book.PairsWellWith == nil {
		book.PairsWellWith = []string{}
	}
	if book.Tags == nil {
		book.Tags = []string{}
	}
	book.AddedDate = time.Now().Format("2006-01-02")

	return book, nil
}

// ─── Command ─────────────────────────────────────────────────────────────────

func newAddCmd(gf *globalFlags) *cobra.Command {
	var jsonInput string

	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a book to the catalog",
		Long: `Add a book interactively (wizard) or inline with --json.

Interactive:
  techshelf add

Inline JSON (minimum required fields):
  techshelf add --json '{"title":"...","author":["..."],"category":"..."}'

All fields are described in SCHEMA.md.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			catalogPath, err := common.CatalogPath(gf.catalogPath)
			if err != nil {
				return err
			}

			c, err := Load(catalogPath)
			if err != nil {
				return err
			}

			var book Book
			if jsonInput != "" {
				book, err = bookFromJSON(c, jsonInput)
				if err != nil {
					return err
				}
			} else {
				book, err = runWizard(c)
				if err != nil {
					// "cancelled" is not an error to the user
					if err.Error() == "cancelled" {
						color.New(color.FgYellow).Println("  Cancelled.")
						return nil
					}
					return err
				}
			}

			if err := c.AddBook(book); err != nil {
				return err
			}
			if err := c.Save(catalogPath); err != nil {
				return err
			}

			logger.Infof("add", "book %q added to catalog", book.ID)
			color.New(color.FgGreen, color.Bold).Printf("\n  ✅  \"%s\" added.\n", book.Title)
			color.New(color.FgWhite, color.Faint).Println("  Run  techshelf generate  to rebuild shelves.\n")
			return nil
		},
	}

	cmd.Flags().StringVar(&jsonInput, "json", "", "add book from inline JSON string (non-interactive)")
	return cmd
}
