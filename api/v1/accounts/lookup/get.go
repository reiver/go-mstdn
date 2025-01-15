package lookup

import (
	"encoding/json"
	"io"
	"net/http"
	gourl "net/url"

	"github.com/reiver/go-erorr"

	"github.com/reiver/go-mstdn/ent"
)

const (
	errNilAccount = erorr.Error("mstdn: nil account")
)

// https://docs.joinmastodon.org/methods/accounts/#lookup
func Get(account *ent.Account, host string, acct string) error {
	if nil == account {
		return errNilAccount
	}

	var url string
	{
		var urloc = gourl.URL{
			Scheme: "https",
			Host:   host,
			Path:   Path,
			RawQuery: "acct=" + gourl.QueryEscape(acct),
		}

		url = urloc.String()
	}

	var httpResponse *http.Response
	{
		var err error
		httpResponse, err = http.Get(url)
		if nil != err {
			return erorr.Errorf("mstdn: problem making HTTP GET request to %q: %w", url, err)
		}
		defer httpResponse.Body.Close()
	}

	var bytes []byte
	{
		var err error
		bytes, err = io.ReadAll(httpResponse.Body)
		if nil != err {
			return erorr.Errorf("mstdn: problem reading (all) the body of the HTTP response for %q: %w", url, err)
		}
	}

	{
		err := json.Unmarshal(bytes, account)
		if nil != err {
			return erorr.Errorf("mstdn: problem json-unmarshaling content from %q: %w", url, err)
		}
	}

	return nil
}
