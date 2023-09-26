package mstdn

import (
	"encoding/json"
	"net/http"
)

func OK(resp http.ResponseWriter, value interface{}) {
	if nil == resp {
		return
	}

	resp.WriteHeader(http.StatusOK)
	resp.Header().Add("Content-Type", "application/json; charset=utf-8")

	err := json.NewEncoder(resp).Encode(value)
	if nil != err {
		// Nothing here.
	}
}
