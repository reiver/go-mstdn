package ent

import (
	"encoding/json"

	"github.com/reiver/go-erorr"
	"github.com/reiver/go-opt"
	"github.com/reiver/go-nul"
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
	WebSite      nul.Nullable[string] `json:"website"`       // optional — field has JSON null value in JSON if not set
	VapidKey     opt.Optional[string] `json:"vapid_key"`
	ClientID     opt.Optional[string] `json:"client_id"`     // optional — field not included in JSON if not set
	ClientSecret opt.Optional[string] `json:"client_secret"` // optional — field not included in JSON if not set
}

func (receiver Application) MarshalJSON() ([]byte, error) {

	var array [512]byte
	var buffer []byte = array[0:0]

	buffer = append(buffer, "{"...)

	{
		val, found := receiver.Name.Get()
		if !found {
			return nil, errCannotMashalApplicationAsJSONNoName
		}

		buffer = append(buffer, `"name":`...)

		{
			marshaled, err := json.Marshal(val)
			if nil != err {
				return nil, erorr.Errorf("ent: could not marshal ent.Application.Name as JSON: %w", err)
			}

			buffer = append(buffer, marshaled...)
		}
	}

	{
		buffer = append(buffer, `,"website":`...)

		switch receiver.WebSite {
		case nul.Nothing[string](), nul.Null[string]():
			buffer = append(buffer, `null`...)
		default:
			website, found := receiver.WebSite.Get()
			if !found {
				return nil, erorr.Error("ent: could not marshal ent.Application.WebSite as JSON: internal error")
			}

			marshaled, err := json.Marshal(website)
			if nil != err {
				return nil, erorr.Errorf("ent: could not marshal ent.Application.WebSite as JSON: %w", err)
			}

			buffer = append(buffer, marshaled...)
		}
	}

	{
		val, found := receiver.VapidKey.Get()
		if !found {
			return nil, errCannotMashalApplicationAsJSONNoVapidKey
		}

		buffer = append(buffer, `,"vapid_key":`...)

		{
			marshaled, err := json.Marshal(val)
			if nil != err {
				return nil, erorr.Errorf("ent: could not marshal ent.Application.VapidKey as JSON: %w", err)
			}

			buffer = append(buffer, marshaled...)
		}
	}

	{
		switch receiver.ClientID {
		case opt.Nothing[string]():
			// Nothing here.
		default:
			clientID, found := receiver.ClientID.Get()
			if !found {
				return nil, erorr.Error("ent: could not marshal ent.Application.ClientID as JSON: internal error")
			}

			buffer = append(buffer, `,"client_id":`...)

			marshaled, err := json.Marshal(clientID)
			if nil != err {
				return nil, erorr.Errorf("ent: could not marshal ent.Application.ClientID as JSON: %w", err)
			}

			buffer = append(buffer, marshaled...)
		}
	}

	{
		switch receiver.ClientSecret {
		case opt.Nothing[string]():
			// Nothing here.
		default:
			clientSecret, found := receiver.ClientSecret.Get()
			if !found {
				return nil, erorr.Error("ent: could not marshal ent.Application.ClientSecret as JSON: internal error")
			}

			buffer = append(buffer, `,"client_secret":`...)

			marshaled, err := json.Marshal(clientSecret)
			if nil != err {
				return nil, erorr.Errorf("ent: could not marshal ent.Application.ClientSecret as JSON: %w", err)
			}

			buffer = append(buffer, marshaled...)
		}
	}

	buffer = append(buffer, "}"...)


	return buffer, nil
}
