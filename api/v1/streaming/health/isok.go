package health

import (
	"net/http"
	"net/url"
)

// IsOK return true if the host is "alive".
//
// The way to checks if the host is "alive" is by making
// an HTTP request to "/api/v1/streaming/health".
// If the HTTP-status-code of the HTTP-response is "200",
// it returns true.
//
// This is part of the Mastodon client-server API.
//
// See:
// https://docs.joinmastodon.org/methods/streaming/#health
func IsOK(host string) (bool, error) {
	var urloc = url.URL{
		Scheme:"https",
		Host:host,
		Path:Path,
	}

	resp, err := http.Get(urloc.String())
	if nil != err {
		return false, err
	}

	if nil == resp {
		return false, errNilHTTPResponse
	}

	return http.StatusOK == resp.StatusCode, nil
}
