package account_id

import (
	"net/http"

	"sourcecode.social/reiver/go-pathmatch"

	"sourcecode.social/reiver/go-mstdn"
	"sourcecode.social/reiver/go-mstdn/ent"
)

var _ http.Handler = internalHandler{}

const Path string = "/api/v1/accounts/{account_id}"

var pattern *pathmatch.Pattern = pathmatch.MustCompile(Path)

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

	var accountid string
	{
		url := req.URL
		if nil == url {
			mstdn.InternalServerError(resp)
			return
		}

		path := url.Path

		success, err := pattern.Find(path, &accountid)
		if nil != err {
			mstdn.BadRequest(resp)
			return
		}
		if !success {
			mstdn.BadRequest(resp)
			return
		}
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
