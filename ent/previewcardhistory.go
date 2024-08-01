package ent

import (
	"github.com/reiver/go-opt"
)

// See:
// https://docs.joinmastodon.org/entities/PreviewCard/#history
type PreviewCardHistory struct {
	Day      opt.Optional[string] `json:"day"`
	Accounts opt.Optional[string] `json:"accounts"`
	Uses     opt.Optional[string] `json:"uses"`
}
