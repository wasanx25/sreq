package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/wataru0225/sreq/config"
	"github.com/wataru0225/sreq/snippet"
)

var pagenation int
var argument string
var editor string

type Qiita struct {
	Title string `json: "title"`
	Url   string `json: "url"`
	Body  string `json: "body"`
}

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search on Qiita",
	Long:  "Search on Qiita",
	Run: func(cmd *cobra.Command, args []string) {
		pagenation = 1
		argument = strings.Join(args, ",")
		execute()
	},
}

func init() {
	RootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringVar(&editor, "editor", "vim", "You choise editor")
	searchCmd.Flags().Bool("browse", false, "Use open browse")
}

func execute() {
	resp, err := http.Get("http://qiita.com/api/v2/items?page=" + strconv.Itoa(pagenation) + "&per_page=10&query=" + argument)
	if err == nil {
		defer resp.Body.Close()
		if b, err := ioutil.ReadAll(resp.Body); err == nil {
			rendering(b)
		}
	}
}

func rendering(b []byte) {
	var contents []Qiita
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
	scan(contents)
}

func scan(content []Qiita) {
	var num string
	_, err := fmt.Scanf("%s", &num)
	if err == nil {
		if num == "n" {
			pagenation++
			execute()
		} else {
			numb, _ := strconv.Atoi(num)
			url, body := writeHistory(content[numb])

			var cfg config.Config
			cfg.Load()
			if cfg.General.OutputType == "editor" {
				text := []byte(body)
				ioutil.WriteFile("/tmp/sreq.txt", text, os.ModePerm)
				cmd := exec.Command(editor, "/tmp/sreq.txt")
				cmd.Stdin = os.Stdin
				cmd.Stdout = os.Stdout
				cmd.Run()
			} else if cfg.General.OutputType == "browse" {
				exec.Command("open", url).Run()
			}
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}

func writeHistory(content Qiita) (string, string) {
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
