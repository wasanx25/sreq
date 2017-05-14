package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/wataru0225/sreq/config"
)

const (
	version = "0.1.0"
)

var RootCmd = &cobra.Command{
	Use:   "sreq",
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
	cobra.OnInitialize(initConfig)
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

func initConfig() {
	cfgFile := config.ConfigFile()
	if _, err := os.Stat(cfgFile); err != nil {
		dir := filepath.Join(os.Getenv("HOME"), ".config", "sreq")
		if err := os.MkdirAll(dir, 0700); err != nil {
			fmt.Errorf("cannot create directory: %v", err)
		}

		var cfg config.Config
		cfg.Load()
		newCfg := config.GeneralConfig{
			OutputType: "editor",
			Editor:     "vim",
		}

		cfg.General = newCfg
		if err := cfg.Save(); err != nil {
			fmt.Errorf("Failed. %v", err)
		}
	}
}
