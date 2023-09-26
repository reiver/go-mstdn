package ent

import (
	"encoding/json"

	"sourcecode.social/reiver/go-erorr"
	"sourcecode.social/reiver/go-opt"
	"sourcecode.social/reiver/go-nul"
)

var _ json.Marshaler = Application{}

const (
	errCannotMashalApplicationAsJSONNoName     = erorr.Error("cannot marshal ent.Application to JSON — no ‘name’ set")
	errCannotMashalApplicationAsJSONNoVapidKey = erorr.Error("cannot marshal ent.Application to JSON — no ‘vapid_key’ set")
)

// Appication represents a Mastodon API "Appication".
//
// See:
// https://docs.joinmastodon.org/entities/Application/
type Application struct {
	Name         opt.Optional[string] `json:"name"`
	WebSite      nul.Nullable[string] `json:"website"`
	VapidKey     opt.Optional[string] `json:"vapid_key"`
	ClientID     opt.Optional[string] `json:"client_id"`
	ClientSecret opt.Optional[string] `json:"client_secret"`
}

func (receiver Application) MarshalJSON() ([]byte, error) {
	var data map[string]interface{}

	{
		val, found := receiver.Name.Get()
		if !found {
			return nil, errCannotMashalApplicationAsJSONNoName
		}

		data["name"] = val
	}

	receiver.WebSite.WhenSomething(func(value string){
		data["website"] = value
	})
	receiver.WebSite.WhenNull(func(){
		data["website"] = nil
	})

	{
		val, found := receiver.VapidKey.Get()
		if !found {
			return nil, errCannotMashalApplicationAsJSONNoVapidKey
		}

		data["vapid_key"] = val
	}

	receiver.ClientID.WhenSomething(func(value string){
		data["client_id"] = value
	})

	receiver.ClientSecret.WhenSomething(func(value string){
		data["client_secret"] = value
	})

	return json.Marshal(data)
}
