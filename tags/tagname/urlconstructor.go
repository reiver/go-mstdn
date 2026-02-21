package tagname

import (
	"fmt"
	gourl "net/url"
)

type URLConstructor interface {
	ConstructURL(host string, tag string) string
}

type URLConstructorFunc func(host string, tag string) string

func (fn URLConstructorFunc) ConstructURL(host string, tag string) string {
	return fn(host, tag)
}

var DefaultURLConstructor URLConstructor = URLConstructorFunc(DefaultURLConstructorFunc)

// DefaultURLConstructor returns the URL to the RSS version of the Mastodon tag page.
func DefaultURLConstructorFunc(host string, tag string) string {
	if "" == tag {
		return ""
	}

	path := fmt.Sprintf("/tags/%s", tag)

	var url = gourl.URL{
		Scheme: "https",
		Host:   host,
		Path:   path,
	}

	return url.String()
}
