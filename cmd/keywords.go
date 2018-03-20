package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/wataru0225/sreq/src"
)

var keywordsCmd = &cobra.Command{
	Use:     "keywords",
	Aliases: []string{"k"},
	Short:   "Show Keywords History (short-cut alias: \"k\")",
	Run: func(cmd *cobra.Command, args []string) {
		var snippets src.Snippets
		snippets.Load()
		for _, snip := range snippets.Snippets {
			fmt.Println(color.CyanString("keyword: " + snip.SearchKeyword))
		}
	},
}

func init() {
	RootCmd.AddCommand(keywordsCmd)
}
