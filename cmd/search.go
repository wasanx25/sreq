package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/wataru0225/sreq/snippet"
)

var pagenation int
var argument string

type Qiita struct {
	Title string `json: "title"`
	Url   string `json: "url"`
}

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Call Api of Qiita",
	Long:  "Call Api of Qiita",
	Run: func(cmd *cobra.Command, args []string) {
		pagenation = 1
		argument = strings.Join(args, ",")
		execute()
	},
}

func init() {
	RootCmd.AddCommand(searchCmd)
}

func execute() {
	resp, err := http.Get("http://qiita.com/api/v2/items?page=" + strconv.Itoa(pagenation) + "&per_page=10&query=" + argument)
	if err == nil {
		defer resp.Body.Close()
		b, errr := ioutil.ReadAll(resp.Body)
		if errr == nil {
			rendering(b)
		}
	}
}

func rendering(b []byte) {
	var content interface{}
	json.Unmarshal(b, &content)
	for i := 0; i < 10; i++ {
		fmt.Print(color.YellowString(strconv.Itoa(i) + " -> "))
		fmt.Println(content.([]interface{})[i].(map[string]interface{})["title"].(string))
	}
	fmt.Println(color.YellowString("n -> ") + "next page")
	fmt.Print("SELECT > ")
	scan(content)
}

func scan(content interface{}) {
	var num string
	_, err := fmt.Scanf("%s", &num)
	if err == nil {
		if num == "n" {
			pagenation++
			execute()
		} else {
			var snippets snippet.Snippets
			snippets.Load()
			numb, _ := strconv.Atoi(num)
			url := content.([]interface{})[numb].(map[string]interface{})["url"].(string)
			title := content.([]interface{})[numb].(map[string]interface{})["title"].(string)
			newSnippet := snippet.SnippetInfo{
				SearchKeyword: argument,
				Url:           url,
				Title:         title,
			}
			snippets.Snippets = append(snippets.Snippets, newSnippet)
			errr := snippets.Save()
			if errr != nil {
				fmt.Errorf("Failed. %v", errr)
			}
			exec.Command("open", url).Run()
		}
	} else {
		fmt.Println(err)
	}
}
