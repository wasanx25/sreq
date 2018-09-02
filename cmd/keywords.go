package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/wasanx25/sreq/history"
)

var keywordsCmd = &cobra.Command{
	Use:     "keywords",
	Aliases: []string{"k"},
	Short:   "Show Keywords History (short-cut alias: \"k\")",
	Run: func(cmd *cobra.Command, args []string) {
		file := filepath.Join(os.Getenv("HOME"), ".config", "sreq", "sreq-history.toml")
		h := history.New(file)
		for _, snippet := range h.Snippets.Snippets {
			fmt.Println(color.CyanString("keyword: " + snippet.SearchKeyword))
		}
	},
}

func init() {
	RootCmd.AddCommand(keywordsCmd)
}
