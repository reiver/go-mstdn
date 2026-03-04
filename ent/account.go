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
	ID              opt.Optional[string]        `json:"id"`
	UserName        opt.Optional[string]        `json:"username,omitempty"`
	Acct            opt.Optional[string]        `json:"acct,omitempty"`
	URL             opt.Optional[string]        `json:"url,omitempty"`
	URI             opt.Optional[string]        `json:"uri,omitempty"` // currently undocumented, but is included in in account JSON returned from the Mastodon API on mastodon.social.
	DisplayName     opt.Optional[string]        `json:"display_name,omitempty"`
	Note            opt.Optional[string]        `json:"note,omitempty"`
	Avatar          opt.Optional[string]        `json:"avatar,omitempty"`
	AvatarStatic    opt.Optional[string]        `json:"avatar_static,omitempty"`
	Header          opt.Optional[string]        `json:"header,omitempty"`
	HeaderStatic    opt.Optional[string]        `json:"header_static,omitempty"`
	Locked          opt.Optional[bool]          `json:"locked,nullempty"`
	Fields        []Field                       `json:"fields,omitempty"`
	Emojis        []CustomEmoji                 `json:"emojis,omitempty"`
	Bot             opt.Optional[bool]          `json:"bot,nullempty"`
	Group           opt.Optional[bool]          `json:"group,nullempty"`
	Discoverable    nul.Nullable[bool]          `json:"discoverable,nullempty"`
	Indexable       nul.Nullable[bool]          `json:"indexable,nullempty"`
	NoIndex         nul.Nullable[bool]          `json:"noindex,nullempty"`
	Moved           nul.Nullable[AccountHolder] `json:"moved"`
	Memorial        nul.Nullable[bool]          `json:"memorial,nullempty"`
	Suspended       opt.Optional[bool]          `json:"suspended,nullempty"`
	Limited         opt.Optional[bool]          `json:"limited,nullempty"`
	CreatedAt       opt.Optional[string]        `json:"created_at,omitempty"`
	LastStatusAt    nul.Nullable[string]        `json:"last_status_at"`
	StatusesCount   opt.Optional[jsonint.Numeric]   `json:"statuses_count,omitempty"`
	FollowersCount  opt.Optional[jsonint.Numeric]   `json:"followers_count,omitempty"`
	FollowingCount  opt.Optional[jsonint.Numeric]   `json:"following_count,omitempty"`
	HideCollections opt.Optional[bool]          `json:"hide_collections,nullempty"`
	Roles         []Role                        `json:"roles,omitempty"`
	MuteExpiresAt   nul.Nullable[string]        `json:"mute_expires_at"`
}
