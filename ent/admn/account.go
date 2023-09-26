package admn

import (
	"sourcecode.social/reiver/go-opt"
	"sourcecode.social/reiver/go-nul"

	"sourcecode.social/reiver/go-mstdn/ent"
)

// Account represents a Mastodon API "Admin::Account".
//
// Note that this is NOT an admin-account.
// But instead this is an administrated-account.
// I.e., the administrator's view of an account.
//
// See:
// https://docs.joinmastodon.org/entities/Admin_Account/
type Account struct {
	ID                     opt.Optional[string] `json:"id"`
	UserName               opt.Optional[string] `json:"username"`
	Domain                 nul.Nullable[string] `json:"domain"`
	CreatedAt              opt.Optional[string] `json:"created_at"`
	EMail                  opt.Optional[string] `json:"email"`
	IP                     nul.Nullable[string] `json:"ip"`
	IPs                  []IP                   `json:"ips"`
	Locale                 nul.Nullable[string] `json:"locale"`
	InviteRequest          nul.Nullable[string] `json:"invite_request"`
	Role                   ent.Role             `json:"role"`
	Confirmed              opt.Optional[bool]   `json:"confirmed"`
	Approved               opt.Optional[bool]   `json:"approved"`
	Disabled               opt.Optional[bool]   `json:"disabled"`
	Silenced               opt.Optional[bool]   `json:"silenced"`
	Suspended              opt.Optional[bool]   `json:"suspended"`
	Account                ent.Account          `json:"account"`
	CreatedByApplicationID opt.Optional[string] `json:"created_by_application_id"`
	InvitedByAccountID     opt.Optional[string] `json:"invited_by_account_id"`
}

