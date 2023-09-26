package mstdn

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"sourcecode.social/reiver/go-opt"
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
