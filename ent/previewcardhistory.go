package ent

import (
	"sourcecode.social/reiver/go-opt"
)

type PreviewCardHistory struct {
	Day      opt.Optional[string] `json:"day"`
	Accounts opt.Optional[string] `json:"accounts"`
	Uses     opt.Optional[string] `json:"uses"`
}
