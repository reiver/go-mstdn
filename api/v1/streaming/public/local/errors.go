package local

import (
	"github.com/reiver/go-erorr"
)

const (
	errNilDestination   = erorr.Error("mstdn: nil destination")
	errNilHTTPRequest   = erorr.Error("mstdn: nil http-request")
	errNilHTTPSSEClient = erorr.Error("mstdn: nil http-sse-client")
	errNilReceiver      = erorr.Error("mstdn: nil receiver")
)
