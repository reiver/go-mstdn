package account_id

import (
	"net/http"

	"sourcecode.social/reiver/go-mstdn"
	"sourcecode.social/reiver/go-mstdn/ent"
)

var _ http.Handler = Handler{}

const Path string = "/api/v1/accounts/{account_id}"

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

	var accountid string
	{
		const key string = "accountid"

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

		accountid = query.Get("accountid")
	}

	var account ent.Account
	{

		err := fn(&account, accountid)
		if nil != err {
			mstdn.Error(resp, err.ErrHTTP())
			return
		}
	}

	mstdn.OK(resp, account)
}
