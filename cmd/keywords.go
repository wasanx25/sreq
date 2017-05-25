package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/wataru0225/sreq/config"
	"github.com/wataru0225/sreq/snippet"
)

var keywordsCmd = &cobra.Command{
	Use:     "keywords",
	Aliases: []string{"k"},
	Short:   "Show Keywords History (short-cut alias: \"k\")",
	Long:    "Show Keywords History (short-cut alias: \"k\")",
	Run: func(cmd *cobra.Command, args []string) {
		var snippets snippet.Snippets
		file := config.HistoryFile()
		snippets.Load(file)
		for _, snip := range snippets.Snippets {
			fmt.Println(color.CyanString("keyword: " + snip.SearchKeyword))
		}
	},
}

func init() {
	RootCmd.AddCommand(keywordsCmd)
}
