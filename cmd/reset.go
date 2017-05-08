package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset history",
	Long:  "Reset history",
	Run: func(cmd *cobra.Command, args []string) {
		var file = filepath.Join(os.Getenv("HOME"), ".config", "sreq", "sreq-history.toml")
		if err := os.Remove(file); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Reset!!")
		}
	},
}

func init() {
	RootCmd.AddCommand(resetCmd)
}
