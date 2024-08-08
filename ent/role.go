package ent

import (
	"github.com/reiver/go-opt"
	"github.com/reiver/go-jsonint"
)

// Role represents a Mastodon API "Role".
//
// See:
// https://docs.joinmastodon.org/entities/Role/
type Role struct {
	ID          opt.Optional[jsonint.Int] `json:"id,omitempty"`
	Name        opt.Optional[string]      `json:"name,omitempty"`
	Color       opt.Optional[string]      `json:"color,omitempty"`
	Permissions opt.Optional[jsonint.Int] `json:"permissions,omitempty"`
	Highlighted opt.Optional[bool]        `json:"highlighted,omitempty"`
}
