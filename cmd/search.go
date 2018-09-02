package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/wasanx25/sreq/control"
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
		c := control.New(strings.Join(args, ","), sort, lynx)
		err := c.Exec()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringVar(&sort, "sort", "rel", "Select rel or created or stock for sort")
	searchCmd.Flags().BoolVar(&lynx, "lynx", false, "Use lynx CUI browse")
}
