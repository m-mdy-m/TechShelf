package command

import (
	"encoding/json"
	"os"

	"github.com/m-mdy-m/TechShelf/internal/common"
	"github.com/spf13/cobra"
)

var (
	Version   = "dev"
	GitCommit = "unknown"
	BuildDate = "unknown"
)

type globalFlags struct {
	catalogPath string
	shelvesDir  string
}

func Execute() error {
	gf := &globalFlags{}

	root := &cobra.Command{
		Use:   "shelf",
		Short: "Manage your personal book catalog",
		Long:  `shelf — a personal, opinionated bookshelf organized like a library.`,
	}

	root.PersistentFlags().StringVar(&gf.catalogPath, "catalog", "", "path to catalog.json (default: ./catalog.json)")
	root.PersistentFlags().StringVar(&gf.shelvesDir, "shelves", ShelvesDir, "path to shelves directory")

	root.AddCommand(
		AddCmd(gf),
		FindCmd(gf),
		RemoveCmd(gf),
		StatsCmd(gf),
		SyncCmd(gf),
		CatalogCmd(gf),
		VersionCmd(),
	)

	return root.Execute()
}

func VersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print build version",
		Run: func(cmd *cobra.Command, args []string) {
			os.Stdout.WriteString("version=" + Version + " commit=" + GitCommit + " build_date=" + BuildDate + "\n")
		},
	}
}

func CatalogCmd(gf *globalFlags) *cobra.Command {
	return &cobra.Command{
		Use:   "catalog",
		Short: "Print catalog as JSON",
		RunE: func(cmd *cobra.Command, args []string) error {
			path, err := common.CatalogPath(gf.catalogPath)
			if err != nil {
				return err
			}
			c, err := Load(path)
			if err != nil {
				return err
			}
			b, err := json.MarshalIndent(c, "", "  ")
			if err != nil {
				return err
			}
			cmd.Println(string(b))
			return nil
		},
	}
}

func SyncCmd(gf *globalFlags) *cobra.Command {
	return &cobra.Command{
		Use:   "sync",
		Short: "Rebuild all shelf markdown files from catalog",
		Long: `sync rebuilds every shelves/<category>/README.md from the current catalog.

Run this after manual edits to catalog.json, or to reset drifted shelf files.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			path, err := common.CatalogPath(gf.catalogPath)
			if err != nil {
				return err
			}
			c, err := Load(path)
			if err != nil {
				return err
			}
			if err := SyncAllShelves(gf.shelvesDir, c.Books); err != nil {
				return err
			}
			cmd.Printf("synced %d shelves in %s\n", len(uniqueCategories(c.Books)), gf.shelvesDir)
			return nil
		},
	}
}

func uniqueCategories(books []Book) []string {
	seen := map[string]bool{}
	var out []string
	for _, b := range books {
		if !seen[b.Category] {
			seen[b.Category] = true
			out = append(out, b.Category)
		}
	}
	return out
}
