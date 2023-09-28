package admn

import (
	"sourcecode.social/reiver/go-opt"
)


// DimensionData represents a Mastodon API "Admin::Dimension Data".
//
// See:
// https://web.archive.org/web/20230716101711/https://docs.joinmastodon.org/entities/Admin_Dimension/#data
type DimensionData struct {
	Key        opt.Optional[string] `json:"key"`
	HumanKey   opt.Optional[string] `json:"human_key"`
	Value      opt.Optional[string] `json:"value"`
	Unit       opt.Optional[string] `json:"unit"`
	HumanValue opt.Optional[string] `json:"human_value"`
}
