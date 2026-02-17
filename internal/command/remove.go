package command

import (
	"errors"

	"github.com/fatih/color"
	"github.com/m-mdy-m/TechShelf/internal/common"
	"github.com/m-mdy-m/TechShelf/internal/logger"
	"github.com/spf13/cobra"
)

func RemoveCmd(gf *globalFlags) *cobra.Command {
	var assumeYes bool
	cmd := &cobra.Command{
		Use:   "remove [book-id]",
		Args:  cobra.MaximumNArgs(1),
		Short: "Remove a book from the catalog and shelf",
		RunE: func(cmd *cobra.Command, args []string) error {
			path, err := common.CatalogPath(gf.catalogPath)
			if err != nil {
				return err
			}
			c, err := Load(path)
			if err != nil {
				return err
			}

			id := ""
			if len(args) > 0 {
				id = args[0]
			} else {
				p := common.NewPrompter()
				id, err = p.AskRequired("Book ID to remove")
				if err != nil {
					return err
				}
			}

			// Show what will be removed
			var found *Book
			for i := range c.Books {
				if c.Books[i].ID == id {
					found = &c.Books[i]
					break
				}
			}
			if found == nil {
				return errors.New("book not found: " + id)
			}
			color.New(color.FgHiBlack).Printf("\n  Will remove: %s (%s)\n\n", found.Title, found.Category)

			if !assumeYes {
				p := common.NewPrompter()
				ok, askErr := p.AskConfirm("Remove this book", false)
				if askErr != nil {
					return askErr
				}
				if !ok {
					logger.Warnf("remove", "remove cancelled for %s", id)
					color.New(color.FgYellow).Println("  cancelled")
					return nil
				}
			}

			removed, ok := c.RemoveBook(id)
			if !ok {
				return errors.New("book not found: " + id)
			}
			if err := c.Save(path); err != nil {
				return err
			}

			// Rebuild shelf for the affected category
			if syncErr := SyncShelfCategory(gf.shelvesDir, removed.Category, c.BooksInCategory(removed.Category)); syncErr != nil {
				logger.Warnf("remove", "shelf sync failed: %v", syncErr)
			}

			logger.Infof("remove", "book removed: %s", id)
			color.New(color.FgGreen, color.Bold).Printf("\n  ✔ Removed: %s\n\n", removed.Title)
			return nil
		},
	}
	cmd.Flags().BoolVarP(&assumeYes, "yes", "y", false, "skip confirmation prompt")
	return cmd
}
