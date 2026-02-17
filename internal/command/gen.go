package command

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/m-mdy-m/TechShelf/internal/common"
	"github.com/m-mdy-m/TechShelf/internal/logger"
	"github.com/spf13/cobra"
)

func GenerateCmd(gf *globalFlags) *cobra.Command {
	var outDir string
	cmd := &cobra.Command{Use: "generate", Short: "Generate shelf markdown files from catalog", RunE: func(cmd *cobra.Command, args []string) error {
		path, err := common.CatalogPath(gf.catalogPath)
		if err != nil {
			return err
		}
		c, err := Load(path)
		if err != nil {
			return err
		}
		if outDir == "" {
			outDir = "shelves"
		}
		if err := common.EnsureDir(outDir); err != nil {
			return err
		}
		group := map[string][]Book{}
		for _, b := range c.Books {
			group[b.Category] = append(group[b.Category], b)
		}
		for category, books := range group {
			catDir := filepath.Join(outDir, Slugify(category))
			if err := common.EnsureDir(catDir); err != nil {
				return err
			}
			var sb strings.Builder
			sb.WriteString("# " + category + "\n\n")
			for _, b := range books {
				sb.WriteString("## " + b.Title + "\n")
				sb.WriteString("- **Authors**: " + strings.Join(b.Author, ", ") + "\n")
				sb.WriteString("- **Why read**: " + b.WhyRead + "\n")
				sb.WriteString("- **Description**: " + b.Description + "\n\n")
			}
			if err := os.WriteFile(filepath.Join(catDir, "README.md"), []byte(sb.String()), 0o644); err != nil {
				return err
			}
		}

		logger.Infof("generate", "generated %d category shelves in %s", len(group), outDir)
		cmd.Println("generated shelves in", outDir)
		return nil
	}}
	cmd.Flags().StringVar(&outDir, "out", "shelves", "output shelves directory")
	return cmd
}
