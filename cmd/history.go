package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/wataru0225/sreq/snippet"
)

var historyCmd = &cobra.Command{
	Use:     "history",
	Aliases: []string{"h"},
	Short:   "Search history (short-cut alias: \"h\")",
	Long:    "Search history (short-cut alias: \"h\")",
	Run: func(cmd *cobra.Command, args []string) {
		var snippets snippet.Snippets
		snippets.Load()
		for _, snip := range snippets.Snippets {
			fmt.Println(color.YellowString("url:     " + snip.URL))
			fmt.Println(color.GreenString("title: 	 " + snip.Title))
			fmt.Println(color.CyanString("keyword: " + snip.SearchKeyword))
			fmt.Println(color.WhiteString("-------------------------------"))
		}
	},
}

func init() {
	RootCmd.AddCommand(historyCmd)
}
