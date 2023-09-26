package ent

import (
	"sourcecode.social/reiver/go-opt"
	"sourcecode.social/reiver/go-jsonint"
)

// Role represents a Mastodon API "Role".
//
// See:
// https://docs.joinmastodon.org/entities/Role/
type Role struct {
	ID          opt.Optional[jsonint.Int] `json:"id"`
	Name        opt.Optional[string]      `json:"name"`
	Color       opt.Optional[string]      `json:"color"`
	Permissions opt.Optional[jsonint.Int] `json:"permissions"`
	Highlighted opt.Optional[bool]        `json:"color"`
}
