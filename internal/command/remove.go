package command

import (
	"errors"

	"github.com/m-mdy-m/TechShelf/internal/common"
	"github.com/m-mdy-m/TechShelf/internal/logger"
	"github.com/spf13/cobra"
)

func RemoveCmd(gf *globalFlags) *cobra.Command {
	var assumeYes bool
	cmd := &cobra.Command{Use: "remove [book-id]", Args: cobra.MaximumNArgs(1), Short: "Remove a book", RunE: func(cmd *cobra.Command, args []string) error {
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
			id, err = p.AskRequired("Book ID")
			if err != nil {
				return err
			}
		}

		if !assumeYes {
			p := common.NewPrompter()
			ok, askErr := p.AskConfirm("Remove book ID: "+id, false)
			if askErr != nil {
				return askErr
			}
			if !ok {
				logger.Warnf("remove", "remove operation cancelled for %s", id)
				cmd.Println("cancelled")
				return nil
			}
		}

		if !c.RemoveBook(id) {
			return errors.New("book id not found")
		}
		if err := c.Save(path); err != nil {
			return err
		}
		logger.Infof("remove", "book %s removed", id)
		cmd.Println("removed:", id)
		return nil
	}}
	cmd.Flags().BoolVarP(&assumeYes, "yes", "y", false, "skip confirmation")
	return cmd
}
