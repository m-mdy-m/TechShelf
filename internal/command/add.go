package command

import (
	"encoding/json"
	"errors"
	"os"
	"strings"

	"github.com/m-mdy-m/TechShelf/internal/common"
	"github.com/m-mdy-m/TechShelf/internal/logger"
	"github.com/spf13/cobra"
)

func AddCmd(gf *globalFlags) *cobra.Command {
	var inlineJSON string
	var fileJSON string
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a book from JSON (inline/file) or interactive prompt",
		RunE: func(cmd *cobra.Command, args []string) error {
			path, err := common.CatalogPath(gf.catalogPath)
			if err != nil {
				return err
			}
			c, err := Load(path)
			if err != nil {
				return err
			}

			book, err := readBookInput(c, inlineJSON, fileJSON)
			if err != nil {
				if err.Error() == "cancelled" {
					logger.Warnf("add", "add operation cancelled")
					cmd.Println("cancelled")
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
			logger.Infof("add", "book %s added", book.ID)
			logger.Infof("add", "catalog updated: %s", path)
			cmd.Println("added:", book.ID)
			return nil
		},
	}
	cmd.Flags().StringVar(&inlineJSON, "json", "", "book JSON object")
	cmd.Flags().StringVar(&fileJSON, "file", "", "path to JSON file containing one book object")
	return cmd
}

func readBookInput(c *Catalog, inlineJSON, fileJSON string) (Book, error) {
	if inlineJSON != "" || fileJSON != "" {
		raw := inlineJSON
		if fileJSON != "" {
			b, err := os.ReadFile(fileJSON)
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
	return promptBookInput(c)
}

func promptBookInput(c *Catalog) (Book, error) {
	p := common.NewPrompter()
	logger.Infof("add", "starting interactive mode")

	title, err := p.AskRequired("Title")
	if err != nil {
		return Book{}, err
	}
	defaultID := Slugify(title)
	id, err := p.Ask("ID (slug)", defaultID)
	if err != nil {
		return Book{}, err
	}
	id = c.UniqueID(Slugify(id))

	authors, err := p.AskList("Authors")
	if err != nil {
		return Book{}, err
	}
	if len(authors) == 0 {
		author, askErr := p.AskRequired("Author")
		if askErr != nil {
			return Book{}, askErr
		}
		authors = []string{author}
	}

	yearPublished, err := p.AskIntOptional("Year published (optional)")
	if err != nil {
		return Book{}, err
	}
	language, err := p.Ask("Language", "English")
	if err != nil {
		return Book{}, err
	}
	category, err := p.AskRequired("Category")
	if err != nil {
		return Book{}, err
	}
	subcategory, err := p.Ask("Subcategory", "")
	if err != nil {
		return Book{}, err
	}
	tags, err := p.AskTags("Tags")
	if err != nil {
		return Book{}, err
	}
	level, err := p.AskChoice("Level", ValidLevels, "general")
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
	prerequisites, err := p.AskList("Prerequisites")
	if err != nil {
		return Book{}, err
	}
	pairsWellWith, err := p.AskList("Pairs well with (book IDs)")
	if err != nil {
		return Book{}, err
	}
	filePath, err := p.Ask("Local file path (optional)", "")
	if err != nil {
		return Book{}, err
	}
	onlineURL, err := p.Ask("Online URL (optional)", "")
	if err != nil {
		return Book{}, err
	}
	status, err := p.AskChoice("Status", ValidStatuses, "unread")
	if err != nil {
		return Book{}, err
	}
	notes, err := p.Ask("Notes", "")
	if err != nil {
		return Book{}, err
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
		Source:        Source{},
		Status:        status,
		Notes:         notes,
	}
	if filePath != "" {
		book.Source.FilePath = strPtr(filePath)
	}
	if onlineURL != "" {
		book.Source.OnlineURL = strPtr(onlineURL)
	}

	confirm, err := p.AskConfirm("Add this book", true)
	if err != nil {
		return Book{}, err
	}
	if !confirm {
		logger.Warnf("add", "interactive add canceled by user")
		return Book{}, errors.New("cancelled")
	}
	return book, nil
}

func normalizeBookFields(b *Book) {
	b.ID = Slugify(strings.TrimSpace(b.ID))
	b.Title = strings.TrimSpace(b.Title)
	for i := range b.Author {
		b.Author[i] = strings.TrimSpace(b.Author[i])
	}
	for i := range b.Tags {
		b.Tags[i] = strings.TrimSpace(b.Tags[i])
	}
	b.Category = strings.TrimSpace(b.Category)
	b.Subcategory = strings.TrimSpace(b.Subcategory)
	b.Language = strings.TrimSpace(b.Language)
	b.Level = normalizeLevel(b.Level)
	b.Status = normalizeStatus(b.Status)
	if b.Source.FilePath != nil {
		v := strings.TrimSpace(*b.Source.FilePath)
		b.Source.FilePath = &v
	}
	if b.Source.OnlineURL != nil {
		v := strings.TrimSpace(*b.Source.OnlineURL)
		b.Source.OnlineURL = &v
	}
}

func strPtr(v string) *string {
	return &v
}
