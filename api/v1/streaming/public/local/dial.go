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

	sseclient, err := httpsse.DialURL(urloc.String())
	if nil != err {
		return nil, err
	}

	return &internalClient{
		sseclient:sseclient,
	}, nil
}

func Dial(req *http.Request) (Client, error) {
	if nil == req {
		return nil, errNilHTTPRequest
	}

	if nil == req.URL {
		req.URL = new(url.URL)
	}

	req.URL.Path = Path

	sseclient, err :=httpsse.Dial(req)
	if nil != err {
		return nil, err
	}

	return &internalClient{
		sseclient:sseclient,
	}, nil
}
