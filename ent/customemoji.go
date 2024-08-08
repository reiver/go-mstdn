package ent

import (
	"github.com/reiver/go-opt"
)

// CustomEmoji represents a Mastodon API "CustomEmoji".
//
// See:
// https://docs.joinmastodon.org/entities/CustomEmoji/
type CustomEmoji struct {
	ShortCode       opt.Optional[string] `json:"shortcode,omitempty"`
	URL             opt.Optional[string] `json:"url,omitempty"`
	StaticURL       opt.Optional[string] `json:"static_url,omitempty"`
	VisibleInPicker opt.Optional[bool]   `json:"visible_in_picker,omitempty"`
	Category        opt.Optional[string] `json:"category,omitempty"`
}
