package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/wataru0225/sreq/config"
	"github.com/wataru0225/sreq/snippet"
)

var editor string
var browse bool

var searchCmd = &cobra.Command{
	Use:     "search",
	Aliases: []string{"s"},
	Short:   "Search on Qiita (short-cut alias: \"s\")",
	Long:    "Search on Qiita (short-cut alias: \"s\")",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Failed to not argument of search keyword.")
			os.Exit(2)
		}

		argument := strings.Join(args, ",")
		pagenation := 1
		end := false

		for {
			end = execute(argument, pagenation)
			if end {
				break
			}
			pagenation++
		}
	},
}

func init() {
	RootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringVar(&editor, "editor", "vim", "Open editor")
	searchCmd.Flags().Bool("browse", false, "Open browse")
}

func execute(argument string, pagenation int) bool {
	resp, err := http.Get(config.BaseURL(strconv.Itoa(pagenation), argument))
	end := true
	if err == nil {
		defer resp.Body.Close()
		if b, err := ioutil.ReadAll(resp.Body); err == nil {
			contents := rendering(b)
			end = scan(contents, argument)
		}
	}

	return end
}

func rendering(b []byte) []*config.Qiita {
	var contents []*config.Qiita
	json.Unmarshal(b, &contents)
	for i, c := range contents {
		fmt.Print(color.YellowString(strconv.Itoa(i) + " -> "))
		fmt.Println(c.Title)
		if count := len(c.Body); count > 256 {
			fmt.Println(color.GreenString(strings.Replace(c.Body, "\n", "", -1)[0:256]))
		} else {
			fmt.Println(color.GreenString(strings.Replace(c.Body, "\n", "", -1)))
		}
		fmt.Print("\n")
	}
	if len(contents) == 10 {
		fmt.Println(color.YellowString("n -> ") + "next page")
	}
	fmt.Print("SELECT > ")
	return contents
}

func scan(content []*config.Qiita, argument string) bool {
	var num string
	if _, err := fmt.Scanf("%s", &num); err == nil {
		if num == "n" {
			return false
		} else {
			numb, _ := strconv.Atoi(num)
			url, body := writeHistory(content[numb], argument)

			var cfg config.Config
			cfg.Load()

			if cfg.General.OutputType == "browse" || browse == true {
				OpenBrowse(url)
				return true
			}

			if editor == "" {
				editor = cfg.General.Editor
			}
			OpenEditor(body, editor)
		}
	} else {
		fmt.Println(err)
	}
	return true
}

func writeHistory(content *config.Qiita, argument string) (string, string) {
	var snippets snippet.Snippets
	file := config.HistoryFile()
	snippets.Load(file)
	url := content.Url
	newSnippet := snippet.SnippetInfo{
		SearchKeyword: argument,
		Url:           url,
		Title:         content.Title,
	}
	snippets.Snippets = append(snippets.Snippets, newSnippet)
	if err := snippets.Save(file); err != nil {
		fmt.Errorf("Failed. %v", err)
	}
	return url, content.Body
}
