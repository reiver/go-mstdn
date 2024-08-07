package ent

import (
	"github.com/reiver/go-opt"
)

// Tag represents a Mastodon API "Tag".
//
// See:
// https://docs.joinmastodon.org/entities/Status/#Tag
type Tag struct {
	Name opt.Optional[string] `json:"name"`
	URL  opt.Optional[string] `json:"url"`
}
