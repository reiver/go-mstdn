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
