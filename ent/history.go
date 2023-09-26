package ent

import (
	"sourcecode.social/reiver/go-opt"
)

type History struct {
	Day      opt.Optional[string] `json:"day"`
	Accounts opt.Optional[string] `json:"accounts"`
	Uses     opt.Optional[string] `json:"uses"`
}
