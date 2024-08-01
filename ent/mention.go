package ent

import (
	"github.com/reiver/go-opt"
)

// See:
// https://docs.joinmastodon.org/entities/Status/#Mention
type Mention struct {
	ID       opt.Optional[string] `json:"id"`
	UserName opt.Optional[string] `json:"username"`
	URL      opt.Optional[string] `json:"url"`
	Acct     opt.Optional[string] `json:"acct"`
}
