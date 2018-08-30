package control

import (
	"fmt"
	"net/url"
	"path/filepath"
	"strconv"
)

type Control struct {
}

func (c *Control) GetURL(id string) string {
	u := url.URL{
		Scheme: "https",
		Host:   "qiita.com",
		Path:   filepath.Join("api", "v2", "items", id),
	}
	return u.String()
}
