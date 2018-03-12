package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
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

		for {
			end := execute(argument, pagenation)
			if end {
				break
			}
			pagenation++
		}
	},
}

// Content is structure that scraping content from Qiita
type Content struct {
	ID    string
	Title string
	Desc  string
}

func init() {
	RootCmd.AddCommand(searchCmd)
}

func execute(argument string, pagenation int) bool {
	doc, err := goquery.NewDocument(config.PageURL(argument, "rel", strconv.Itoa(pagenation)))
	if err != nil {
		fmt.Printf("Scraping failed -> err: %v", err)
		return true
	}

	var contents []*Content

	doc.Find(".searchResult").Each(func(_ int, s *goquery.Selection) {
		itemID, _ := s.Attr("data-uuid")
		title := s.Find(".searchResult_itemTitle a").Text()
		desc := s.Find(".searchResult_snippet").Text()

		content := &Content{
			ID:    itemID,
			Title: title,
			Desc:  desc,
		}

		contents = append(contents, content)
	})

	rendering(contents)
	end := scan(contents, argument)

	return end
}

func rendering(contents []*Content) {
	for num, content := range contents {
		fmt.Print(color.YellowString(strconv.Itoa(num) + " -> "))
		fmt.Println(content.Title)
		fmt.Println(color.GreenString(content.Desc))
		fmt.Print("\n")
	}
	if len(contents) == 10 {
		fmt.Println(color.YellowString("n -> ") + "next page")
	}
	fmt.Print("SELECT > ")
}

func scan(contents []*Content, argument string) bool {
	var num string
	if _, err := fmt.Scanf("%s", &num); err != nil {
		fmt.Println(err)
	}

	if num == "n" {
		return false
	}

	index, _ := strconv.Atoi(num)
	target := contents[index]

	resp, err := http.Get(config.APIURL(target.ID))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var qiita *config.Qiita
	json.Unmarshal(b, &qiita)

	OpenEditor(qiita.Markdown, "less")

	writeHistory(qiita, argument)

	return true
}

func writeHistory(content *config.Qiita, argument string) {
	var snippets snippet.Snippets
	file := config.HistoryFile()
	snippets.Load(file)
	url := content.URL
	newSnippet := snippet.SnippetInfo{
		SearchKeyword: argument,
		Url:           url,
		Title:         content.Title,
	}
	snippets.Snippets = append(snippets.Snippets, newSnippet)
	if err := snippets.Save(file); err != nil {
		fmt.Printf("Failed. %v", err)
		os.Exit(2)
	}
}
