package tagname

import (
	"testing"
)

func TestRSSURL(t *testing.T) {
	tests := []struct {
		Host     string
		Tag      string
		Expected string
	}{
		{},



		{
			Host:             "example.com",
			Tag:                               "fedidev",
			Expected: "https://example.com/tags/fedidev",
		},
		{
			Host:             "Example.COM",
			Tag:                               "FediDev",
			Expected: "https://Example.COM/tags/FediDev",
		},



		{
			Host:             "mastodon.social",
			Tag:                                   "fediverse",
			Expected: "https://mastodon.social/tags/fediverse",
		},
	}

	for testNumber, test := range tests {

		actual := rssURL(test.Host, test.Tag)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual RSS URL is not what was expected", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("HOST:            %q", test.Host)
			t.Logf("TAG:             %q", test.Tag)
			continue
		}
	}
}
