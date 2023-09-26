package ent

import (
	"sourcecode.social/reiver/go-opt"
)

// See:
// https://docs.joinmastodon.org/entities/Status/#Tag
type Tag struct {
	Name opt.Optional[string] `json:"name"`
	URL  opt.Optional[string] `json:"url"`
}
