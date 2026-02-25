package public

import (
	"fmt"
	"net/http"
	"net/url"

	"codeberg.org/reiver/go-httpsse"
)

func DialHostUsingBearerToken(host string, bearerToken string,) (Client, error) {
	var urloc = url.URL{
		Scheme:"https",
		Host:host,
		Path:Path,
	}

	httprequest, err := http.NewRequest("GET", urloc.String(), nil)
	if nil != err {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	httprequest.Header.Set("Authorization", "Bearer " + bearerToken)

	return Dial(httprequest)
}

func DialHost(host string) (Client, error) {
	var urloc = url.URL{
		Scheme:"https",
		Host:host,
		Path:Path,
	}

	httprequest, err := http.NewRequest("GET", urloc.String(), nil)
	if nil != err {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	return Dial(httprequest)
}

func Dial(req *http.Request) (Client, error) {
	if nil == req {
		return nil, errNilHTTPRequest
	}

	if nil == req.URL {
		req.URL = new(url.URL)
	}

	req.URL.Path = Path

	sseclient, err := httpsse.Dial(req)
	if nil != err {
		return nil, fmt.Errorf("failed to dial HTTP SSE connection: %w", err)
	}

	return &internalClient{
		sseclient:sseclient,
	}, nil
}
