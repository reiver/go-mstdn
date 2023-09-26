package ent

import (
	"sourcecode.social/reiver/go-opt"
	"sourcecode.social/reiver/go-jsonint"
)

// Account represents a Mastodon API "Account".
//
// See:
// https://docs.joinmastodon.org/entities/Account/
type Account struct {
	ID             opt.Optional[string]      `json:"id"`
	UserName       opt.Optional[string]      `json:"username"`
	Acct           opt.Optional[string]      `json:"acct"`
	DisplayName    opt.Optional[string]      `json:"display_name"`
	Locked         opt.Optional[bool]        `json:"locked"`
	Bot            opt.Optional[bool]        `json:"bot"`
	Discoverable   opt.Optional[bool]        `json:"discoverable"`
	CreatedAt      opt.Optional[string]      `json:"created_at"`
	Note           opt.Optional[string]      `json:"note"`
	URL            opt.Optional[string]      `json:"url"`
	URI            opt.Optional[string]      `json:"uri"`
	Avatar         opt.Optional[string]      `json:"avatar"`
	AvatarStatic   opt.Optional[string]      `json:"avatar_static"`
	Header         opt.Optional[string]      `json:"header"`
	HeaderStatic   opt.Optional[string]      `json:"header_static"`
	FollowersCount opt.Optional[jsonint.Int] `json:"followers_count"`
	FollowingCount opt.Optional[jsonint.Int] `json:"following_count"`
	StatusesCount  opt.Optional[jsonint.Int] `json:"statuses_count"`
	LastStatusAt   opt.Optional[string]      `json:"last_status_at"`
	NoIndex        opt.Optional[bool]        `json:"noindex"`
	Emojis       []CustomEmoji               `json:"emojis"`
	Roles        []Role                      `json:"roles"`
	Fields       []Field                     `json:"fields"`
}
