package src

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
)

// Content is structure that scraping content from Qiita
type Content struct {
	ID    string
	Title string
	Desc  string
}

// ExecSearch is scraping and viewing contents and selecting contents
func ExecSearch(argument string, pagenation int, sort string, lynx bool) {
	for {
		contents, err := search(argument, pagenation, sort)
		if err != nil {
			fmt.Println(err)
			break
		}
		viewList(contents)
		endPhase := scan(contents, argument, lynx)
		if endPhase {
			break
		}
		pagenation++
	}
}

func search(argument string, pagenation int, sort string) ([]*Content, error) {
	doc, err := goquery.NewDocument(PageURL(argument, sort, strconv.Itoa(pagenation)))
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

	resp, err := http.Get(APIURL(target.ID))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var qiita *Qiita
	json.Unmarshal(b, &qiita)

	writeHistory(qiita, argument)

	if lynx {
		openFile(qiita.HTML, "lynx", "/tmp/sreq.html")
		return true
	}

	openFile(qiita.Markdown, "less", "/tmp/sreq.txt")
	return true
}

func openFile(body string, cmdName string, file string) {
	text := []byte(body)
	ioutil.WriteFile(file, text, os.ModePerm)
	cmd := exec.Command(cmdName, file)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func writeHistory(content *Qiita, argument string) {
	var snippets Snippets
	snippets.Load()
	url := content.URL
	newSnippet := Snippet{
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
