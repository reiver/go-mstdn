package local

import (
	"net/http"
	"net/url"

	"github.com/reiver/go-httpsse"
)

func DialHost(host string) (Client, error) {
	var urloc = url.URL{
		Scheme:"https",
		Host:host,
		Path:Path,
	}

	return httpsse.DialURL(urloc.String())
}

func Dial(req *http.Request) (Client, error) {
	return httpsse.Dial(req)
}
