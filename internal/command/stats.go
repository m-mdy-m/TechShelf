package command

import (
	"encoding/json"

	"github.com/m-mdy-m/TechShelf/internal/common"
	"github.com/m-mdy-m/TechShelf/internal/logger"
	"github.com/spf13/cobra"
)

type statsView struct {
	Total      int            `json:"total"`
	ByCategory map[string]int `json:"by_category"`
	ByStatus   map[string]int `json:"by_status"`
}

func StatsCmd(gf *globalFlags) *cobra.Command {
	return &cobra.Command{Use: "stats", Short: "Catalog summary", RunE: func(cmd *cobra.Command, args []string) error {
		path, err := common.CatalogPath(gf.catalogPath)
		if err != nil {
			return err
		}
		c, err := Load(path)
		if err != nil {
			return err
		}
		v := statsView{Total: len(c.Books), ByCategory: map[string]int{}, ByStatus: map[string]int{}}
		for _, b := range c.Books {
			v.ByCategory[b.Category]++
			v.ByStatus[b.Status]++
		}
		b, err := json.MarshalIndent(v, "", "  ")
		if err != nil {
			return err
		}

		logger.Infof("stats", "computed stats for %d books", len(c.Books))
		cmd.Println(string(b))
		return nil
	}}
}
