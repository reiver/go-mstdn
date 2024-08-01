package ent

import (
	"github.com/reiver/go-opt"
)

// Notification represents a Mastodon API "Notification".
//
// See:
// https://docs.joinmastodon.org/entities/Notification/
type Notification struct {
	ID        opt.Optional[string]  `json:"id"`
	Type      opt.Optional[string]  `json:"type"`
	CreatedAt opt.Optional[string]  `json:"created_at"`
	Account   Account               `json:"account"`
	Status    opt.Optional[Status]  `json:"account"`
	Report    opt.Optional[Report]  `json:"report"`
}
