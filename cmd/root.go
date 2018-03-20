package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	version = "0.1.0"
)

// RootCmd initialize cmd base
var RootCmd = &cobra.Command{
	Use:   "sreq",
	Short: "Search reference on Qiita",
	Long:  "If you do not know or want to research, search on Qiita.",
}

// Execute RootCmd
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show Sreq version",
	Long:  "Show Sreq version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("sreq version %s\n", version)
	},
}
