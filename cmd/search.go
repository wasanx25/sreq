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
)

type Qiita struct {
	Title string `json: "title"`
	Url   string `json: "url"`
}

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Call Api of Qiita",
	Long:  "Call Api of Qiita",
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get("http://qiita.com/api/v2/items?page=1&per_page=10&query=" + strings.Join(args, ","))
		if err == nil {
			puts(resp)
		} else {
			fmt.Println(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(searchCmd)
}

func puts(resp *http.Response) {
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		var content interface{}
		json.Unmarshal(b, &content)
		for i := 0; i < 10; i++ {
			fmt.Print(color.YellowString(strconv.Itoa(i) + " -> "))
			fmt.Println(content.([]interface{})[i].(map[string]interface{})["title"].(string))
		}
		var num int
		fmt.Print("SELECT > ")
		_, errr := fmt.Scanf("%d", &num)
		if errr == nil {
			errrr := exec.Command("open", content.([]interface{})[num].(map[string]interface{})["url"].(string)).Run()
			if errrr != nil {
				fmt.Println(errrr)
			}
		}
	} else {
		fmt.Println(err)
	}
}
