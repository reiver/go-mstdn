package ent

import (
	"encoding/json"

	"github.com/reiver/go-erorr"
	"github.com/reiver/go-opt"
)

var _ json.Marshaler = CustomEmoji{}

// CustomEmoji represents a Mastodon API "CustomEmoji".
//
// See:
// https://docs.joinmastodon.org/entities/CustomEmoji/
type CustomEmoji struct {
	ShortCode       opt.Optional[string] `json:"shortcode"`
	URL             opt.Optional[string] `json:"url"`
	StaticURL       opt.Optional[string] `json:"static_url"`
	VisibleInPicker opt.Optional[bool]   `json:"visible_in_picker"`
	Category        opt.Optional[string] `json:"category"`
}

func (receiver CustomEmoji) MarshalJSON() ([]byte, error) {

	var buffer []byte

	buffer = append(buffer, "{"...)

	{
                buffer = append(buffer, `"shortcode":`...)

                marshaled, err := json.Marshal(receiver.ShortCode)
                if nil != err {
                        return nil, erorr.Errorf("mstdn/ent: could not marshal ent.CustomEmoji.ShortCode as JSON: %w", err)
                }

                buffer = append(buffer, marshaled...)
	}

	{
                buffer = append(buffer, `,"url":`...)

                marshaled, err := json.Marshal(receiver.URL)
                if nil != err {
                        return nil, erorr.Errorf("mstdn/ent: could not marshal ent.CustomEmoji.URL as JSON: %w", err)
                }

                buffer = append(buffer, marshaled...)
	}

	{
                buffer = append(buffer, `,"static_url":`...)

                marshaled, err := json.Marshal(receiver.StaticURL)
                if nil != err {
                        return nil, erorr.Errorf("mstdn/ent: could not marshal ent.CustomEmoji.StaticURL as JSON: %w", err)
                }

                buffer = append(buffer, marshaled...)
	}

	{
                buffer = append(buffer, `,"visible_in_picker":`...)

                marshaled, err := json.Marshal(receiver.VisibleInPicker)
                if nil != err {
                        return nil, erorr.Errorf("mstdn/ent: could not marshal ent.CustomEmoji.VisibleInPicker as JSON: %w", err)
                }

                buffer = append(buffer, marshaled...)
	}

	{
		if opt.Nothing[string]() != receiver.Category {

	                buffer = append(buffer, `,"category":`...)

        	        marshaled, err := json.Marshal(receiver.Category)
                	if nil != err {
                        	return nil, erorr.Errorf("mstdn/ent: could not marshal ent.CustomEmoji.Category as JSON: %w", err)
	                }

        	        buffer = append(buffer, marshaled...)
		}
	}

	buffer = append(buffer, "}"...)

	return buffer, nil
}
