package local

import (
	"encoding/json"
	"fmt"

	"github.com/reiver/go-httpsse"
	"github.com/reiver/go-opt"
)

// Client represenrs a client for interacting with the "/api/v1/streaming/public/local" API end-point.
//
// That API end-point will return a series of events.
//
// Client will let you iterate through those events.
type Client = httpsse.Client

type internalClient struct {
	sseclient httpsse.Client
	data []string
	err error
}

var _ Client = &internalClient{}

func (receiver *internalClient) Close() error {
	if nil == receiver {
		return errNilReceiver
	}

	var sseclient httpsse.Client = receiver.sseclient

	if nil == sseclient {
		return errNilHTTPSSEClient
	}

	return sseclient.Close()
}

func (receiver *internalClient) Decode(dst interface{}) error {

	if nil == dst {
		return errNilDestination
	}

	var event *Event
	{
		var casted bool
		event, casted = dst.(*Event)
		if !casted {
			
		}
	}

	if len(receiver.data) <= 0 {
		
	}

	var datum string
	{
		datum, receiver.data = receiver.data[0], receiver.data[1:]
	}
fmt.Printf("DATUM:\n%s \n", datum)

	if len(datum) <= 0 {

	}

	switch datum[0] {
	case '0','1','2','3','4','5','6','7','8','9':
		event.Name="delete"
		event.Status.ID = opt.Something(datum)
	default:
		event.Name="update"

		var bytes []byte = []byte(datum)

		{
			err := json.Unmarshal(bytes, &(event.Status))
			if nil != err {
				
				return err
			}
		}
	}


	return nil
}

func (receiver *internalClient) Err() error {
	if nil == receiver {
		return errNilReceiver
	}

	if nil != receiver.err {
		return receiver.err
	}

	return receiver.sseclient.Err()
}

func (receiver *internalClient) Next() bool {
	if nil == receiver {
		return false
	}

	if 0 < len(receiver.data) {
		return true
	}

	more := receiver.sseclient.Next()
	if !more {
		return false
	}

	var event httpsse.Event
	{
		err := receiver.sseclient.Decode(&event)
		if nil != err {
			receiver.err = err
			return false
		}
	}

	receiver.data = append(receiver.data, event.EventData()...)

//@TODO: what if event.EventData() returns an empty array?
//       which could happen.

	return true
}
