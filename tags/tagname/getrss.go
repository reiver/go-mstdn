package tagname

import (
	"fmt"

	"codeberg.org/reiver/go-erorr"
	"github.com/reiver/go-rss"
)

const (
	ErrNilURLConstructor = erorr.Error("nil URL constructor")
)

// RSS2 represents and RSS 2.0 document.
type RSS2 = rss.RSS2

// GetRSS loads the RSS feed for a tag at a Mastodon host.
//
// Example usage:
//
//	host := "mastodon.social"
//	tag := "fedidev"
//	
//	var rss2 tags.RSS2
//	
//	err := tags.RSS(&rss2, host, tag)
//
// The previous example would make a call to:
//
//	https://mastodon.social/tags/fedidev.rss
//
// With the value of Accept HTTP request header being:
//
//	Accept: application/rss+xml
//
// GetRSS takes can of these details for you.
func GetRSS(dst *RSS2, host string, tag string) error {
	if nil == DefaultURLConstructor {
		return ErrNilURLConstructor
	}

	url := DefaultURLConstructor.ConstructURL(host, tag)
	if "" == url {
		return fmt.Errorf("failed to construct URL to Mastodon tags RSS web-feed for host %q and tag %q", host, tag)
	}

	err := rss.HTTPGetAndUnmarshal(url, dst)
	if nil != err {
		return err
	}

	return nil
}
