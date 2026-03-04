package ent

import (
	"github.com/reiver/go-opt"
	"github.com/reiver/go-nul"
)

// Appication represents a Mastodon API "Appication".
//
// See:
// https://docs.joinmastodon.org/entities/Application/
type Application struct {
	Name         opt.Optional[string] `json:"name,omitempty"`
	WebSite      nul.Nullable[string] `json:"website,omitempty"`       // optional — field has JSON null value in JSON if not set
	VapidKey     opt.Optional[string] `json:"vapid_key,omitempty"`     // optional — field not included in JSON if not set
	ClientID     opt.Optional[string] `json:"client_id,omitempty"`     // optional — field not included in JSON if not set
	ClientSecret opt.Optional[string] `json:"client_secret,omitempty"` // optional — field not included in JSON if not set
}
