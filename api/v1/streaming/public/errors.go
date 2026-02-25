package public

import (
	"codeberg.org/reiver/go-erorr"
)

const (
	errEmptyData        = erorr.Error("mstdn: empty data")
	errNilDestination   = erorr.Error("mstdn: nil destination")
	errNilHTTPRequest   = erorr.Error("mstdn: nil http-request")
	errNilHTTPSSEClient = erorr.Error("mstdn: nil http-sse-client")
	errNilReceiver      = erorr.Error("mstdn: nil receiver")
)
