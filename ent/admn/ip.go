package admn

import (
	"github.com/reiver/go-opt"
)

// IP represents a Mastodon API "Admin::Ip".
//
// Note that "ip" means "ip-address" in this context.
//
// See:
// https://docs.joinmastodon.org/entities/Admin_Ip/
type IP struct {
	IP     opt.Optional[string] `json:"ip"` // ip-address
	UsedAt opt.Optional[string] `json:"used_at"`
}
