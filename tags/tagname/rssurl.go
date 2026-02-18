package tagname

import (
	"fmt"
	gourl "net/url"
)

// rssURL returns the URL to the RSS version of the Mastodon page.
func rssURL(host string, tag string) string {
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
