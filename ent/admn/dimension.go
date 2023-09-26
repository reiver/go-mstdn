package admn

import (
	"sourcecode.social/reiver/go-opt"
)

// Dimension represents a Mastodon API "Admin::Dimension".
//
// See:
// https://docs.joinmastodon.org/entities/Admin_Dimension/
type Dimension struct {
	Key    opt.Optional[string] `json:"key"`
	Data []DimensionData        `json:"data"`
}

type DimensionData struct {
	Key        opt.Optional[string] `json:"key"`
	HumanKey   opt.Optional[string] `json:"human_key"`
	Value      opt.Optional[string] `json:"value"`
	Unit       opt.Optional[string] `json:"unit"`
	HumanValue opt.Optional[string] `json:"human_value"`
}
