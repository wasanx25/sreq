package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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
		resp, err := http.Get("http://qiita.com/api/v2/items?page=1&per_page=10&query=ruby")
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
			fmt.Println(i)
			fmt.Println(content.([]interface{})[i].(map[string]interface{})["url"].(string))
		}
	} else {
		fmt.Println(err)
	}
}
