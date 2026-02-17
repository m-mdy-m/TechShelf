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
}

func Execute() error {
	gf := &globalFlags{}
	root := &cobra.Command{
		Use:   "techshelf",
		Short: "Manage your personal library catalog",
	}
	root.PersistentFlags().StringVar(&gf.catalogPath, "catalog", "", "path to catalog.json")
	root.AddCommand(AddCmd(gf), FindCmd(gf), RemoveCmd(gf), StatsCmd(gf), GenerateCmd(gf), CatalogCmd(gf), VersionCmd())
	return root.Execute()
}

func VersionCmd() *cobra.Command {
	return &cobra.Command{Use: "version", Run: func(cmd *cobra.Command, args []string) {
		_ = cmd
		os.Stdout.WriteString("version=" + Version + " commit=" + GitCommit + " build_date=" + BuildDate + "\n")
	}}
}

func CatalogCmd(gf *globalFlags) *cobra.Command {
	return &cobra.Command{Use: "catalog", Short: "Print catalog as JSON", RunE: func(cmd *cobra.Command, args []string) error {
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
	}}
}
