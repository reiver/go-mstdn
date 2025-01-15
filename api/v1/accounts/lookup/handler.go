package lookup

import (
	"net/http"

	"github.com/reiver/go-mstdn"
	"github.com/reiver/go-mstdn/ent"
)

var _ http.Handler = internalHandler{}

func Handler(fn LoaderFunc) http.Handler {
	return internalHandler{
		loaderFunc:fn,
	}
}

type internalHandler struct {
	loaderFunc LoaderFunc
}

func (internalHandler) Path() string {
	return Path
}

func (receiver internalHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
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

	fn := receiver.loaderFunc
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
