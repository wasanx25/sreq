package search

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
	"github.com/wasanx25/sreq/config"
	"github.com/wasanx25/sreq/src"
)

// Content is structure that scraping content from Qiita
type Content struct {
	ID    string
	Title string
	Desc  string
}

func search(argument string, pagenation int, sort string) ([]*Content, error) {
	doc, err := goquery.NewDocument(config.GetPageURL(argument, sort, strconv.Itoa(pagenation)))
	if err != nil {
		return nil, err
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

	return contents, nil
}

func viewList(contents []*Content) {
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

func scan(contents []*Content, argument string, lynx bool) bool {
	var num string
	if _, err := fmt.Scanf("%s", &num); err != nil {
		fmt.Println(err)
	}

	if num == "n" {
		return false
	}

	index, _ := strconv.Atoi(num)
	target := contents[index]

	resp, err := http.Get(config.GetAPIURL(target.ID))
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

	writeHistory(qiita, argument)

	if lynx {
		openFile(qiita.HTML, "/tmp/sreq.html", "lynx", "-display_charset=utf-8", "-assume_charset=utf-8")
		return true
	}

	openFile(qiita.Markdown, "/tmp/sreq.txt", "less")
	return true
}

func openFile(body string, file string, cmdName ...string) {
	text := []byte(body)
	ioutil.WriteFile(file, text, os.ModePerm)
	cmdName = append(cmdName, file)
	cmd := exec.Command(cmdName[0], cmdName[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func writeHistory(content *config.Qiita, argument string) {
	var snippets src.Snippets
	snippets.Load()
	url := content.URL
	newSnippet := src.Snippet{
		SearchKeyword: argument,
		URL:           url,
		Title:         content.Title,
	}
	snippets.Snippets = append(snippets.Snippets, newSnippet)
	if err := snippets.Save(); err != nil {
		fmt.Printf("Failed. %v", err)
		os.Exit(2)
	}
}
