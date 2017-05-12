package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/wataru0225/sreq/config"
	"github.com/wataru0225/sreq/snippet"
)

var keywordsCmd = &cobra.Command{
	Use:   "keywords",
	Short: "Show Keywords History",
	Long:  "Show Keywords History",
	Run: func(cmd *cobra.Command, args []string) {
		var snippets snippet.Snippets
		file := config.KeywordFile()
		snippets.Load(file)
		for _, snip := range snippets.Snippets {
			fmt.Println(color.CyanString("keyword: " + snip.SearchKeyword))
		}
	},
}

func init() {
	RootCmd.AddCommand(keywordsCmd)
}
