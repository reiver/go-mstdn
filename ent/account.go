package ent

import (
	"github.com/reiver/go-nul"
	"github.com/reiver/go-jsonint"
	"github.com/reiver/go-opt"
)

// Account represents a Mastodon API "Account".
//
// See:
// https://docs.joinmastodon.org/entities/Account/
type Account struct {
	ID             opt.Optional[string]        `json:"id"`
	UserName       opt.Optional[string]        `json:"username,omitempty"`
	Acct           opt.Optional[string]        `json:"acct,omitempty"`
	URL            opt.Optional[string]        `json:"url,omitempty"`
	URI            opt.Optional[string]        `json:"uri,omitempty"` // currently undocumented, but is included in in account JSON returned from the Mastodon API on mastodon.social.
	DisplayName    opt.Optional[string]        `json:"display_name,omitempty"`
	Note           opt.Optional[string]        `json:"note,omitempty"`
	Avatar         opt.Optional[string]        `json:"avatar,omitempty"`
	AvatarStatic   opt.Optional[string]        `json:"avatar_static,omitempty"`
	Header         opt.Optional[string]        `json:"header,omitempty"`
	HeaderStatic   opt.Optional[string]        `json:"header_static,omitempty"`
	Fields       []Field                       `json:"fields,omitempty"`
	Emojis       []CustomEmoji                 `json:"emojis,omitempty"`
	Locked         opt.Optional[bool]          `json:"locked,omitempty"`
	Bot            opt.Optional[bool]          `json:"bot,omitempty"`
	Group          opt.Optional[bool]          `json:"group,omitempty"`
	Discoverable   nul.Nullable[bool]          `json:"discoverable,omitempty"`
	NoIndex        nul.Nullable[bool]          `json:"noindex,omitempty"`
	Moved          nul.Nullable[AccountHolder] `json:"moved,omitempty"`
	Suspended      opt.Optional[bool]          `json:"suspended,omitempty"`
	Limited        opt.Optional[bool]          `json:"limited,omitempty"`
	CreatedAt      opt.Optional[string]        `json:"created_at,omitempty"`
	LastStatusAt   nul.Nullable[string]        `json:"last_status_at,omitempty"`
	StatusesCount  opt.Optional[jsonint.Int]   `json:"statuses_count,omitempty"`
	FollowersCount opt.Optional[jsonint.Int]   `json:"followers_count,omitempty"`
	FollowingCount opt.Optional[jsonint.Int]   `json:"following_count,omitempty"`
	Roles        []Role                        `json:"roles,omitempty"`
	MuteExpiresAt  nul.Nullable[string]        `json:"mute_expires_at,omitempty"`
}
