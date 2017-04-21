package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Call Api of Qiita",
	Long:  "Call Api of Qiita",
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get("http://qiita.com/api/v2/items?page=1&per_page=2&query=ruby")
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
		fmt.Println(string(b))
	} else {
		fmt.Println(err)
	}
}
