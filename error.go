package mstdn

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/reiver/go-opt"
)

func errorJSON(writer io.Writer, statusCode int, msg ...interface{}) {
	if nil == writer {
		return
	}

	// short error message.
	var short string
	{
		var shortErrorMessage opt.Optional[string]

		shortErrorMessage.WhenNothing(func(){
			if 1 <= len(msg) {
				shortErrorMessage = opt.Something(fmt.Sprint(msg[0]))
			}
		})
		shortErrorMessage.WhenNothing(func(){
			shortErrorMessage = opt.Something(http.StatusText(statusCode))
		})

		shortErrorMessage.WhenSomething(func(value string){
			short = value
		})
	}


	var long opt.Optional[string]
	{
		if 2 <= len(msg) {
			long = opt.Something(fmt.Sprint(msg[1:]...))
		}
	}

	var array [512]byte
	var buffer []byte = array[:0]
	{

		buffer = append(buffer, "{"...)
			buffer = append(buffer, `"error":`...)
			{
				// This should never return an error since we are passing a string.
				marshaled, _ := json.Marshal(short)

				buffer = append(buffer, marshaled...)
			}

			long.WhenSomething(func(value string){
				buffer = append(buffer, `,"error_description":`...)

				// This should never return an error since we are passing a string.
				marshaled, _ := json.Marshal(value)

				buffer = append(buffer, marshaled...)
			})
		buffer = append(buffer, "}"...)
	}

	writer.Write(buffer)
}

// Error responds to the request with the specified error message and HTTP code.
// The response will be in JSON (application/json; charset=utf-8).
//
// For example:
//
//	var statusCode int = http.StatusUnauthorized
//	var shortErrorMessage = http.StatusText(http.StatusUnauthorized)
//	var longErrorMessage = "Access Denied!!! — you can keep on knocking but you can't come in"
//	
//	mstdn.Error(resp, statusCode, shortErrorMessage, longErrorMessage)
//
// Responds with:
//
//	{
//		"error":"Unauthorized",
//		"error_description":"Access Defined!!!! — you can keep on knocking but you can't come in"
//	}
func Error(resp http.ResponseWriter, statusCode int, msg ...interface{}) {
	if nil == resp {
		return
	}

	resp.WriteHeader(statusCode)
	resp.Header().Add("Content-Type", "application/json; charset=utf-8")

	errorJSON(resp, statusCode, msg...)
}

// BadRequest replies to the request with an HTTP 400 Bad Request.
// The response will be in JSON (application/json; charset=utf-8).
func BadRequest(resp http.ResponseWriter, msg ...interface{}) {
	const statusCode int = http.StatusBadRequest

	Error(resp, statusCode, msg...)
}

// Unauthorized replies to the request with an HTTP 401 Unauthorized.
// The response will be in JSON (application/json; charset=utf-8).
func Unauthorized(resp http.ResponseWriter, msg ...interface{}) {
	const statusCode int = http.StatusUnauthorized

	Error(resp, statusCode, msg...)
}

// PaymentRequired replies to the request with an HTTP 402 Payment Required
// The response will be in JSON (application/json; charset=utf-8).
func PaymentRequired(resp http.ResponseWriter, msg ...interface{}) {
	const statusCode int = http.StatusPaymentRequired

	Error(resp, statusCode, msg...)
}

// Forbidden replies to the request with an HTTP 403 Forbidden.
// The response will be in JSON (application/json; charset=utf-8).
func Forbidden(resp http.ResponseWriter, msg ...interface{}) {
	const statusCode int = http.StatusForbidden

	Error(resp, statusCode, msg...)
}

// NotFound replies to the request with an HTTP 404 Not Found.
// The response will be in JSON (application/json; charset=utf-8).
func NotFound(resp http.ResponseWriter, msg ...interface{}) {
	const statusCode int = http.StatusNotFound

	Error(resp, statusCode, msg...)
}

// MethodNotAllowed replies to the request with an HTTP 405 Method Not Allowed.
// The response will be in JSON (application/json; charset=utf-8).
func MethodNotAllowed(resp http.ResponseWriter, msg ...interface{}) {
	const statusCode int = http.StatusMethodNotAllowed

	Error(resp, statusCode, msg...)
}

// RequestTimeout replies to the request with an HTTP 408 Request Timeout.
// The response will be in JSON (application/json; charset=utf-8).
func RequestTimeout(resp http.ResponseWriter, msg ...interface{}) {
	const statusCode int = http.StatusRequestTimeout

	Error(resp, statusCode, msg...)
}

// UnprocessableEntity replies to the request with an HTTP 422 Unprocessable Entity.
// The response will be in JSON (application/json; charset=utf-8).
func UnprocessableEntity(resp http.ResponseWriter, msg ...interface{}) {
	const statusCode int = http.StatusUnprocessableEntity

	Error(resp, statusCode, msg...)
}

// TooManyRequests replies to the request with an HTTP 429 Too Many Requests.
// The response will be in JSON (application/json; charset=utf-8).
func TooManyRequests (resp http.ResponseWriter, msg ...interface{}) {
	const statusCode int = http.StatusTooManyRequests 

	Error(resp, statusCode, msg...)
}

// InternalServerError replies to the request with an HTTP 500 Internal Server Error.
// The response will be in JSON (application/json; charset=utf-8).
func InternalServerError(resp http.ResponseWriter, msg ...interface{}) {
	const statusCode int = http.StatusInternalServerError

	Error(resp, statusCode, msg...)
}

// BadGateway replies to the request with an HTTP 502 Bad Gateway.
// The response will be in JSON (application/json; charset=utf-8).
func BadGateway(resp http.ResponseWriter, msg ...interface{}) {
	const statusCode int = http.StatusBadGateway

	Error(resp, statusCode, msg...)
}

// ServiceUnavailable replies to the request with an HTTP 503 Service Unavailable.
// The response will be in JSON (application/json; charset=utf-8).
func ServiceUnavailable(resp http.ResponseWriter, msg ...interface{}) {
	const statusCode int = http.StatusServiceUnavailable

	Error(resp, statusCode, msg...)
}

// Gateway Timeout replies to the request with an HTTP 504 Gateway Timeout.
// The response will be in JSON (application/json; charset=utf-8).
func GatewayTimeout(resp http.ResponseWriter, msg ...interface{}) {
	const statusCode int = http.StatusGatewayTimeout

	Error(resp, statusCode, msg...)
}
