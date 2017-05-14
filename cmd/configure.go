package cmd

import (
	"fmt"

	"github.com/Songmu/prompter"
	"github.com/spf13/cobra"
	"github.com/wataru0225/sreq/config"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Setting configuration",
	Long:  "Setting configuration",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("configure called")
		input := (&prompter.Prompter{
			Choices:    []string{"vim", "emacs", "nano", "less"},
			Default:    "vim",
			Message:    "please select Editor (default: vim)",
			IgnoreCase: true,
		}).Prompt()

		var cfg config.Config
		newCfg := config.GeneralConfig{
			OutputType: "editor",
			Editor:     input,
		}

		cfg.General = newCfg
		if err := cfg.Save(); err != nil {
			fmt.Errorf("Failed. %v", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(configureCmd)
}
