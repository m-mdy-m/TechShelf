package command

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/m-mdy-m/TechShelf/internal/common"
	"github.com/m-mdy-m/TechShelf/internal/logger"
	"github.com/spf13/cobra"
)

func FindCmd(gf *globalFlags) *cobra.Command {
	var author, category, tag, q, status string
	var jsonOut bool

	cmd := &cobra.Command{
		Use:   "find",
		Short: "Search books in the catalog",
		Long: `Search books by title, author, category, tag, or status.

Examples:
  shelf find --q "clean code"
  shelf find --category computer-science --status unread
  shelf find --author "Knuth" --json`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if author == "" && category == "" && tag == "" && q == "" && status == "" {
				p := common.NewPrompter()
				var err error
				q, err = p.Ask("Search text (title/description)", "")
				if err != nil {
					return err
				}
				author, err = p.Ask("Author filter", "")
				if err != nil {
					return err
				}

				existingCats, _ := common.ListCategories(gf.shelvesDir)
				if len(existingCats) > 0 {
					val, err := p.SelectOrNew("Category filter (leave blank to skip)", append([]string{""}, existingCats...))
					if err != nil {
						return err
					}
					category = strings.TrimSpace(val)
				} else {
					category, err = p.Ask("Category filter", "")
					if err != nil {
						return err
					}
				}

				tag, err = p.Ask("Tag filter", "")
				if err != nil {
					return err
				}
				status, err = p.Ask("Status filter (unread/reading/completed/paused)", "")
				if err != nil {
					return err
				}
			}

			path, err := common.CatalogPath(gf.catalogPath)
			if err != nil {
				return err
			}
			c, err := Load(path)
			if err != nil {
				return err
			}

			results := filterBooks(c.Books, q, author, category, tag, status)
			logger.Infof("find", "found %d books", len(results))

			if jsonOut {
				printBooksJSON(cmd, results)
				return nil
			}
			printBooksFormatted(results)
			return nil
		},
	}

	cmd.Flags().StringVar(&author, "author", "", "filter by author name")
	cmd.Flags().StringVar(&category, "category", "", "filter by category")
	cmd.Flags().StringVar(&tag, "tag", "", "filter by tag")
	cmd.Flags().StringVar(&q, "q", "", "free-text search in title and description")
	cmd.Flags().StringVar(&status, "status", "", "filter by status (unread/reading/completed/paused)")
	cmd.Flags().BoolVar(&jsonOut, "json", false, "output results as JSON")
	return cmd
}

func filterBooks(books []Book, q, author, category, tag, status string) []Book {
	out := make([]Book, 0)
	for _, b := range books {
		if author != "" && !containsFold(b.Author, author) {
			continue
		}
		if category != "" && !strings.EqualFold(b.Category, category) {
			continue
		}
		if tag != "" && !containsFold(b.Tags, tag) {
			continue
		}
		if status != "" && !strings.EqualFold(b.Status, status) {
			continue
		}
		if q != "" {
			haystack := strings.ToLower(b.Title + " " + b.Description)
			if !strings.Contains(haystack, strings.ToLower(q)) {
				continue
			}
		}
		out = append(out, b)
	}
	return out
}

func printBooksFormatted(books []Book) {
	if len(books) == 0 {
		color.New(color.FgYellow).Println("\n  No books found.\n")
		return
	}

	titleColor := color.New(color.FgCyan, color.Bold)
	metaColor := color.New(color.FgHiBlack)
	labelColor := color.New(color.FgWhite)
	tagColor := color.New(color.FgMagenta)

	fmt.Printf("\n")
	color.New(color.FgGreen).Printf("  %d book(s) found\n\n", len(books))

	for i, b := range books {
		titleColor.Printf("  %d. %s\n", i+1, b.Title)
		metaColor.Printf("     [%s]", b.Category)
		if b.Subcategory != "" {
			metaColor.Printf("/%s", b.Subcategory)
		}
		metaColor.Printf("  level: %s  status: %s", b.Level, b.Status)
		if b.YearPublished != nil {
			metaColor.Printf("  year: %d", *b.YearPublished)
		}
		fmt.Println()

		labelColor.Printf("     authors: %s\n", strings.Join(b.Author, ", "))

		if len(b.Tags) > 0 {
			tagColor.Printf("     tags: %s\n", strings.Join(b.Tags, ", "))
		}
		if b.Description != "" {
			desc := b.Description
			if len(desc) > 100 {
				desc = desc[:97] + "..."
			}
			metaColor.Printf("     %s\n", desc)
		}
		if b.Source.OnlineURL != nil {
			metaColor.Printf("     🔗 %s\n", *b.Source.OnlineURL)
		}
		fmt.Println()
	}
}

func containsFold(items []string, val string) bool {
	for _, item := range items {
		if strings.EqualFold(item, val) {
			return true
		}
	}
	return false
}

func printBooksJSON(cmd *cobra.Command, books []Book) {
	b, err := json.MarshalIndent(books, "", "  ")
	if err != nil {
		cmd.PrintErr(err)
		return
	}
	cmd.Println(string(b))
}
