package lookup

import (
	"net/http"

	"sourcecode.social/reiver/go-mstdn"
	"sourcecode.social/reiver/go-mstdn/ent"
)

var _ http.Handler = Handler{}

const Path string = "/api/v1/accounts/lookup"

type Handler struct {
	LoaderFunc LoaderFunc
}

func (Handler) Path() string {
	return Path
}

func (receiver Handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	if nil == resp {
		return
	}

	if nil == req {
		mstdn.InternalServerError(resp)
		return
	}

	if http.MethodGet != req.Method {
		mstdn.MethodNotAllowed(resp)
		return
	}

	fn := receiver.LoaderFunc
	if nil == fn {
		mstdn.InternalServerError(resp)
		return
	}

	var acct string
	{
		const key string = "acct"

		url := req.URL
		if nil == url {
			mstdn.InternalServerError(resp)
			return
		}

		query := url.Query()
		if nil == query {
			mstdn.InternalServerError(resp)
			return
		}

		if !query.Has(key) {
			mstdn.BadRequest(resp)
			return
		}

		acct = query.Get("acct")
	}

	var account ent.Account
	{

		err := fn(&account, acct)
		if nil != err {
			mstdn.Error(resp, err.ErrHTTP())
			return
		}
	}

	mstdn.OK(resp, account)
}
