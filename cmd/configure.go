package cmd

import (
	"fmt"
	"os"

	"github.com/Songmu/prompter"
	"github.com/spf13/cobra"
	"github.com/wataru0225/sreq/config"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:     "configure",
	Aliases: []string{"c"},
	Short:   "Setting configuration (short-cut alias: \"c\")",
	Long:    "Setting configuration (short-cut alias: \"c\")",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("configure called")
		output := (&prompter.Prompter{
			Choices:    []string{"editor", "browse"},
			Default:    "editor",
			Message:    "please select OutputType (default: editor)",
			IgnoreCase: true,
		}).Prompt()

		var editor string
		if output == "editor" {
			editor = (&prompter.Prompter{
				Choices:    []string{"vim", "emacs", "nano", "less"},
				Default:    "vim",
				Message:    "please select Editor (default: vim)",
				IgnoreCase: true,
			}).Prompt()
		} else {
			editor = "vim"
		}

		var cfg config.Config
		newCfg := config.GeneralConfig{
			OutputType: output,
			Editor:     editor,
		}

		cfg.General = newCfg
		if err := cfg.Save(); err != nil {
			fmt.Printf("Failed. %v", err)
			os.Exit(2)
		} else {
			fmt.Println("your setting OutputType is " + output)
			fmt.Println("your setting Editor is " + editor)
		}
	},
}

func init() {
	RootCmd.AddCommand(configureCmd)
}
