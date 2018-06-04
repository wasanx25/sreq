package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/wasanx25/sreq/search"
)

var (
	editor string
	lynx   bool
	sort   string
)

var searchCmd = &cobra.Command{
	Use:     "search",
	Aliases: []string{"s"},
	Short:   "Search on Qiita (short-cut alias: \"s\")",
	Run: func(cmd *cobra.Command, args []string) {
		searcher := &search.Searcher{
			Keywords:   strings.Join(args, ","),
			Pagination: 1,
			Sort:       sort,
			Lynx:       lynx,
		}

		if err := searcher.Validate(); err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		searcher.Exec()
	},
}

func init() {
	RootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringVar(&sort, "sort", "rel", "Select rel or created or stock for sort")
	searchCmd.Flags().BoolVar(&lynx, "lynx", false, "Use lynx CUI browse")
}
