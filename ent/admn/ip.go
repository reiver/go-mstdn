package admn

import (
	"sourcecode.social/reiver/go-opt"
)

// IP represents a Mastodon API "Admin::Ip".
//
// See:
// https://docs.joinmastodon.org/entities/Admin_Ip/
type IP struct {
	IP     opt.Optional[string] `json:"ip"` // ip-address
	UsedAt opt.Optional[string] `json:"used_at"`
}
