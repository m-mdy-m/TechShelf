package command

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/m-mdy-m/TechShelf/internal/common"
	"github.com/m-mdy-m/TechShelf/internal/logger"
	"github.com/spf13/cobra"
)

func AddCmd(gf *globalFlags) *cobra.Command {
	var inlineJSON string
	var fileJSON string
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a book to the catalog and shelf",
		Long: `Add a book from inline JSON, a file, or interactive prompt.

Examples:
  shelf add
  shelf add --json '{"title":"SICP","author":["Abelson"],"category":"computer-science"}'
  shelf add --file ./book.json`,
		RunE: func(cmd *cobra.Command, args []string) error {
			path, err := common.CatalogPath(gf.catalogPath)
			if err != nil {
				return err
			}
			c, err := Load(path)
			if err != nil {
				return err
			}

			book, err := readBookInput(c, gf.shelvesDir, inlineJSON, fileJSON)
			if err != nil {
				if err.Error() == "cancelled" {
					logger.Warnf("add", "add cancelled by user")
					color.New(color.FgYellow).Println("  cancelled")
					return nil
				}
				return err
			}

			if err := c.AddBook(book); err != nil {
				return err
			}
			if err := c.Save(path); err != nil {
				return err
			}

			if syncErr := SyncShelfCategory(gf.shelvesDir, book.Category, c.BooksInCategory(book.Category)); syncErr != nil {
				logger.Warnf("add", "shelf sync failed: %v", syncErr)
			}

			logger.Infof("add", "book added: %s", book.ID)
			color.New(color.FgGreen, color.Bold).Printf("\n  ✔ Added: %s\n", book.Title)
			color.New(color.FgHiBlack).Printf("    id: %s  |  category: %s\n\n", book.ID, book.Category)
			return nil
		},
	}
	cmd.Flags().StringVar(&inlineJSON, "json", "", "book as inline JSON object")
	cmd.Flags().StringVar(&fileJSON, "file", "", "path to a JSON file containing one book object")
	return cmd
}

func readBookInput(c *Catalog, shelvesDir, inlineJSON, fileJSON string) (Book, error) {
	if inlineJSON != "" || fileJSON != "" {
		return bookFromJSON(inlineJSON, fileJSON)
	}
	return promptBookInput(c, shelvesDir)
}

func bookFromJSON(inline, file string) (Book, error) {
	raw := inline
	if file != "" {
		b, err := os.ReadFile(file)
		if err != nil {
			return Book{}, err
		}
		raw = string(b)
	}
	var book Book
	if err := json.Unmarshal([]byte(raw), &book); err != nil {
		return Book{}, err
	}
	normalizeBookFields(&book)
	return book, nil
}

func promptBookInput(c *Catalog, shelvesDir string) (Book, error) {
	p := common.NewPrompter()

	common.Section("Core info")

	title, err := p.AskRequired("Title")
	if err != nil {
		return Book{}, err
	}
	rawID, err := p.Ask("ID (slug)", Slugify(title))
	if err != nil {
		return Book{}, err
	}
	id := c.UniqueID(Slugify(rawID))

	authors, err := p.AskList("Authors")
	if err != nil {
		return Book{}, err
	}
	if len(authors) == 0 {
		a, e := p.AskRequired("Author")
		if e != nil {
			return Book{}, e
		}
		authors = []string{a}
	}

	existingCats, err := common.ListCategories(shelvesDir)
	if err != nil {
		return Book{}, err
	}
	category, err := p.SelectOrNew("Category", existingCats)
	if err != nil {
		return Book{}, err
	}
	category = strings.TrimSpace(category)

	subcategory, err := p.Ask("Subcategory", "")
	if err != nil {
		return Book{}, err
	}

	common.Section("Details")

	yearPublished, err := p.AskIntOptional("Year published")
	if err != nil {
		return Book{}, err
	}
	language, err := p.Ask("Language", DefaultLanguage)
	if err != nil {
		return Book{}, err
	}
	tags, err := p.AskTags("Tags")
	if err != nil {
		return Book{}, err
	}
	level, err := p.AskChoice("Level", Levels, DefaultLevel)
	if err != nil {
		return Book{}, err
	}
	description, err := p.Ask("Description", "")
	if err != nil {
		return Book{}, err
	}
	whyRead, err := p.Ask("Why read", "")
	if err != nil {
		return Book{}, err
	}
	whenToRead, err := p.Ask("When to read", "")
	if err != nil {
		return Book{}, err
	}
	importance, err := p.Ask("Importance", "")
	if err != nil {
		return Book{}, err
	}

	common.Section("Relations (optional)")

	prerequisites, err := p.AskList("Prerequisites")
	if err != nil {
		return Book{}, err
	}
	pairsWellWith, err := p.AskList("Pairs well with (book IDs)")
	if err != nil {
		return Book{}, err
	}

	common.Section("Source (at least one recommended)")

	onlineURL, err := p.Ask("Online URL", "")
	if err != nil {
		return Book{}, err
	}
	filePath, err := p.Ask("Local file path", "")
	if err != nil {
		return Book{}, err
	}

	common.Section("Reading status")

	status, err := p.AskChoice("Status", Statuses, DefaultStatus)
	if err != nil {
		return Book{}, err
	}
	notes, err := p.Ask("Notes", "")
	if err != nil {
		return Book{}, err
	}

	printBookSummary(title, id, authors, category, level, status)
	confirm, err := p.AskConfirm("Add this book", true)
	if err != nil {
		return Book{}, err
	}
	if !confirm {
		return Book{}, errors.New("cancelled")
	}

	book := Book{
		ID:            id,
		Title:         title,
		Author:        authors,
		YearPublished: yearPublished,
		Language:      language,
		Category:      category,
		Subcategory:   subcategory,
		Tags:          tags,
		Level:         level,
		Description:   description,
		WhyRead:       whyRead,
		WhenToRead:    whenToRead,
		Importance:    importance,
		Prerequisites: prerequisites,
		PairsWellWith: pairsWellWith,
		Status:        status,
		Notes:         notes,
	}
	if onlineURL != "" {
		book.Source.OnlineURL = strPtr(onlineURL)
	}
	if filePath != "" {
		book.Source.FilePath = strPtr(filePath)
	}
	return book, nil
}

func printBookSummary(title, id string, authors []string, category, level, status string) {
	fmt.Println()
	color.New(color.FgMagenta, color.Bold).Println("  ── Summary ──")
	fmt.Println()
	color.New(color.Bold).Printf("  %s\n", title)
	color.New(color.FgHiBlack).Printf("  id: %s | authors: %s\n", id, strings.Join(authors, ", "))
	color.New(color.FgHiBlack).Printf("  category: %s | level: %s | status: %s\n", category, level, status)
	fmt.Println()
}

func normalizeBookFields(b *Book) {
	b.ID = Slugify(strings.TrimSpace(b.ID))
	b.Title = strings.TrimSpace(b.Title)
	b.Category = strings.TrimSpace(b.Category)
	b.Subcategory = strings.TrimSpace(b.Subcategory)
	b.Language = strings.TrimSpace(b.Language)
	b.Level = NormalizeLevel(b.Level)
	b.Status = NormalizeStatus(b.Status)
	for i := range b.Author {
		b.Author[i] = strings.TrimSpace(b.Author[i])
	}
	for i := range b.Tags {
		b.Tags[i] = strings.TrimSpace(b.Tags[i])
	}
	if b.Source.FilePath != nil {
		v := strings.TrimSpace(*b.Source.FilePath)
		b.Source.FilePath = &v
	}
	if b.Source.OnlineURL != nil {
		v := strings.TrimSpace(*b.Source.OnlineURL)
		b.Source.OnlineURL = &v
	}
}
