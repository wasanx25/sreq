package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

var RootCmd = &cobra.Command{
	Use:   "srebq",
	Short: "Search reference on Qiita",
	Long:  "If you do not know or want to research, Search on Qiita.",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize()
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show Srebq version",
	Long:  "Show Srebq version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("srebq v1.0")
	},
}
