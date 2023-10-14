package url

import (
	"fmt"
	"net/url"
)

var print = fmt.Println

func FormatUrl() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	print(u.Scheme)
	print(u.Host)
	print(u.Path)
	print(u.Fragment)
	print(u.User)
	queryParams, _ := url.ParseQuery(u.RawQuery)
	print(queryParams["k"][0])
}
