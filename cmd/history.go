package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/wataru0225/sreq/snippet"
)

var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Search keyword history",
	Long:  "Search keyword history",
	Run: func(cmd *cobra.Command, args []string) {
		var snippets snippet.Snippets
		snippets.Load()
		for _, snip := range snippets.Snippets {
			fmt.Println(color.YellowString(snip.Url))
			fmt.Println(color.GreenString(snip.Title))
			fmt.Println(color.GreenString(snip.SearchKeyword))
			fmt.Println(color.WhiteString("-------------------------------"))
		}
	},
}

func init() {
	RootCmd.AddCommand(historyCmd)
}