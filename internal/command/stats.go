package command

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/fatih/color"
	"github.com/m-mdy-m/TechShelf/internal/common"
	"github.com/m-mdy-m/TechShelf/internal/logger"
	"github.com/spf13/cobra"
)

func StatsCmd(gf *globalFlags) *cobra.Command {
	var jsonOut bool
	cmd := &cobra.Command{
		Use:   "stats",
		Short: "Show catalog statistics",
		RunE: func(cmd *cobra.Command, args []string) error {
			path, err := common.CatalogPath(gf.catalogPath)
			if err != nil {
				return err
			}
			c, err := Load(path)
			if err != nil {
				return err
			}
			logger.Infof("stats", "computing stats for %d books", len(c.Books))

			if jsonOut {
				return printStatsJSON(cmd, c)
			}
			printStatsFormatted(c)
			return nil
		},
	}
	cmd.Flags().BoolVar(&jsonOut, "json", false, "output as JSON")
	return cmd
}

func printStatsFormatted(c *Catalog) {
	books := c.Books
	header := color.New(color.FgCyan, color.Bold)
	dim := color.New(color.FgHiBlack)
	bold := color.New(color.Bold)

	fmt.Println()
	header.Printf("  📚 %s\n", c.Meta.Name)
	dim.Printf("  %s\n", c.Meta.Description)
	dim.Printf("  last updated: %s\n", c.Meta.LastUpdated)
	fmt.Println()

	if len(books) == 0 {
		color.New(color.FgYellow).Println("  No books yet. Run `shelf add` to get started.\n")
		return
	}

	bold.Printf("  Total: %d books\n\n", len(books))

	byCat := countBy(books, func(b Book) string { return b.Category })
	bold.Println("  By category:")
	printCountTable(byCat, color.New(color.FgMagenta), dim)
	fmt.Println()

	byStatus := countBy(books, func(b Book) string { return b.Status })
	bold.Println("  By status:")
	for _, s := range Statuses {
		n := byStatus[s]
		if n == 0 {
			continue
		}
		statusDisplayColor(s).Printf("    %-12s", s)
		dim.Printf(" %d\n", n)
	}
	fmt.Println()

	byLevel := countBy(books, func(b Book) string { return b.Level })
	bold.Println("  By level:")
	for _, l := range Levels {
		n := byLevel[l]
		if n == 0 {
			continue
		}
		dim.Printf("    %-14s %d\n", l, n)
	}
	fmt.Println()

	completed := byStatus["completed"]
	reading := byStatus["reading"]
	if len(books) > 0 {
		pct := float64(completed) / float64(len(books)) * 100
		bold.Println("  Progress:")
		dim.Printf("    completed  %d / %d  (%.0f%%)\n", completed, len(books), pct)
		dim.Printf("    reading    %d\n", reading)
		fmt.Println()
	}
}

func statusDisplayColor(s string) *color.Color {
	switch s {
	case "completed":
		return color.New(color.FgGreen)
	case "reading":
		return color.New(color.FgCyan)
	case "unread":
		return color.New(color.FgWhite)
	case "paused":
		return color.New(color.FgYellow)
	default:
		return color.New(color.FgHiBlack)
	}
}

func countBy(books []Book, key func(Book) string) map[string]int {
	m := map[string]int{}
	for _, b := range books {
		m[key(b)]++
	}
	return m
}

func printCountTable(m map[string]int, labelClr, numClr *color.Color) {
	keys := make([]string, 0, len(m))
	maxLen := 0
	for k := range m {
		keys = append(keys, k)
		if len(k) > maxLen {
			maxLen = len(k)
		}
	}
	sort.Slice(keys, func(i, j int) bool { return m[keys[i]] > m[keys[j]] })
	for _, k := range keys {
		pad := strings.Repeat(" ", maxLen-len(k))
		labelClr.Printf("    %s%s", k, pad)
		numClr.Printf("  %d\n", m[k])
	}
}

func printStatsJSON(cmd *cobra.Command, c *Catalog) error {
	type statsOut struct {
		Total      int            `json:"total"`
		ByCategory map[string]int `json:"by_category"`
		ByStatus   map[string]int `json:"by_status"`
		ByLevel    map[string]int `json:"by_level"`
	}
	out := statsOut{
		Total:      len(c.Books),
		ByCategory: countBy(c.Books, func(b Book) string { return b.Category }),
		ByStatus:   countBy(c.Books, func(b Book) string { return b.Status }),
		ByLevel:    countBy(c.Books, func(b Book) string { return b.Level }),
	}
	b, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		return err
	}
	cmd.Println(string(b))
	return nil
}
