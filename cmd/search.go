package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/wasanx25/sreq/manager"

	"github.com/spf13/cobra"
)

var sort string

var searchCmd = &cobra.Command{
	Use:     "search",
	Aliases: []string{"s"},
	Short:   "Search on Qiita (short-cut alias: \"s\")",
	Run: func(cmd *cobra.Command, args []string) {
		m, err := manager.New(strings.Join(args, ","), sort)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = m.Run()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringVar(&sort, "sort", "rel", "Select rel or created or stock for sort")
}
