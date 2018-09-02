package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/wasanx25/sreq/history"
)

var historyCmd = &cobra.Command{
	Use:     "history",
	Aliases: []string{"h"},
	Short:   "Search history (short-cut alias: \"h\")",
	Run: func(cmd *cobra.Command, args []string) {
		file := filepath.Join(os.Getenv("HOME"), ".config", "sreq", "sreq-history.toml")
		h := history.New(file)
		for _, snippet := range h.Snippets.Snippets {
			fmt.Println(color.YellowString("url:     " + snippet.URL))
			fmt.Println(color.GreenString("title: 	 " + snippet.Title))
			fmt.Println(color.CyanString("keyword: " + snippet.SearchKeyword))
			fmt.Println(color.WhiteString("-------------------------------"))
		}
	},
}

func init() {
	RootCmd.AddCommand(historyCmd)
}
