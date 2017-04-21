package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Call Api of Qiita",
	Long:  "Call Api of Qiita",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("search called")
	},
}

func init() {
	RootCmd.AddCommand(searchCmd)
}
