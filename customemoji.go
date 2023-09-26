package mstdn

import (
	"encoding/json"

	"sourcecode.social/reiver/go-erorr"
	"sourcecode.social/reiver/go-opt"
)

var _ json.Marshaler = CustomEmoji{}

const (
	errCannotMashalCustomEmojiAsJSONNoShortCode       = erorr.Error("cannot marshal mstdn.CustomEmoji to JSON — no ‘shortcode’ set")
	errCannotMashalCustomEmojiAsJSONNoURL             = erorr.Error("cannot marshal mstdn.CustomEmoji to JSON — no ‘url’ set")
	errCannotMashalCustomEmojiAsJSONNoStaticURL       = erorr.Error("cannot marshal mstdn.CustomEmoji to JSON — no ‘static_url’ set")
	errCannotMashalCustomEmojiAsJSONNoVisibleInPicker = erorr.Error("cannot marshal mstdn.CustomEmoji to JSON — no ‘visible_in_picker’ set")
)

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

	data := map[string]interface{}{}

	{
		value, found := receiver.ShortCode.Get()
		if !found {
			return nil, errCannotMashalCustomEmojiAsJSONNoShortCode
		}

		data["shortcode"] = value
	}
	{
		value, found := receiver.URL.Get()
		if !found {
			return nil, errCannotMashalCustomEmojiAsJSONNoURL
		}

		data["url"] = value
	}
	{
		value, found := receiver.StaticURL.Get()
		if !found {
			return nil, errCannotMashalCustomEmojiAsJSONNoStaticURL
		}

		data["static_url"] = value
	}
	{
		value, found := receiver.VisibleInPicker.Get()
		if !found {
			return nil, errCannotMashalCustomEmojiAsJSONNoVisibleInPicker
		}

		data["visible_in_picker"] = value
	}
	{
		receiver.Category.WhenSomething(func(value string){
			data["category"] = value
		})
	}

	return json.Marshal(data)
}
