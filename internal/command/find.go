package command

import (
	"encoding/json"
	"strings"

	"github.com/m-mdy-m/TechShelf/internal/common"
	"github.com/m-mdy-m/TechShelf/internal/logger"
	"github.com/spf13/cobra"
)

func FindCmd(gf *globalFlags) *cobra.Command {
	var author, category, tag, q string
	cmd := &cobra.Command{Use: "find", Short: "Search books", RunE: func(cmd *cobra.Command, args []string) error {
		if author == "" && category == "" && tag == "" && q == "" {
			p := common.NewPrompter()
			var err error
			q, err = p.Ask("Search text (title/description, optional)", "")
			if err != nil {
				return err
			}
			author, err = p.Ask("Author filter (optional)", "")
			if err != nil {
				return err
			}
			category, err = p.Ask("Category filter (optional)", "")
			if err != nil {
				return err
			}
			tag, err = p.Ask("Tag filter (optional)", "")
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
		out := make([]Book, 0)
		for _, b := range c.Books {
			if author != "" && !containsFold(b.Author, author) {
				continue
			}
			if category != "" && !strings.EqualFold(b.Category, category) {
				continue
			}
			if tag != "" && !containsFold(b.Tags, tag) {
				continue
			}
			if q != "" && !strings.Contains(strings.ToLower(b.Title+" "+b.Description), strings.ToLower(q)) {
				continue
			}
			out = append(out, b)
		}
		buf, err := json.MarshalIndent(out, "", "  ")
		if err != nil {
			return err
		}
		logger.Infof("find", "found %d books", len(out))
		cmd.Println(string(buf))
		return nil
	}}
	cmd.Flags().StringVar(&author, "author", "", "filter by author")
	cmd.Flags().StringVar(&category, "category", "", "filter by category")
	cmd.Flags().StringVar(&tag, "tag", "", "filter by tag")
	cmd.Flags().StringVar(&q, "q", "", "free-text title/description query")
	return cmd
}

func containsFold(items []string, val string) bool {
	for _, item := range items {
		if strings.EqualFold(item, val) {
			return true
		}
	}
	return false
}
