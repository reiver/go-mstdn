package ent

import (
	"encoding/json"

	"sourcecode.social/reiver/go-erorr"
	"sourcecode.social/reiver/go-nul"
	"sourcecode.social/reiver/go-opt"
	"sourcecode.social/reiver/go-jsonint"
)

var _ json.Marshaler = Account{}

// Account represents a Mastodon API "Account".
//
// See:
// https://docs.joinmastodon.org/entities/Account/
type Account struct {
	ID             opt.Optional[string]        `json:"id"`
	UserName       opt.Optional[string]        `json:"username"`
	Acct           opt.Optional[string]        `json:"acct"`
	URL            opt.Optional[string]        `json:"url"`
	URI            opt.Optional[string]        `json:"uri"` // currently undocumented, but is included in in account JSON returned from the Mastodon API on mastodon.social.
	DisplayName    opt.Optional[string]        `json:"display_name"`
	Note           opt.Optional[string]        `json:"note"`
	Avatar         opt.Optional[string]        `json:"avatar"`
	AvatarStatic   opt.Optional[string]        `json:"avatar_static"`
	Header         opt.Optional[string]        `json:"header"`
	HeaderStatic   opt.Optional[string]        `json:"header_static"`
	Fields       []Field                       `json:"fields"`
	Emojis       []CustomEmoji                 `json:"emojis"`
	Locked         opt.Optional[bool]          `json:"locked"`
	Bot            opt.Optional[bool]          `json:"bot"`
	Group          opt.Optional[bool]          `json:"group"`
	Discoverable   nul.Nullable[bool]          `json:"discoverable"`
	NoIndex        nul.Nullable[bool]          `json:"noindex"`
	Moved          nul.Nullable[AccountHolder] `json:"moved"`
	Suspended      opt.Optional[bool]          `json:"suspended"`
	Limited        opt.Optional[bool]          `json:"limited"`
	CreatedAt      opt.Optional[string]        `json:"created_at"`
	LastStatusAt   nul.Nullable[string]        `json:"last_status_at"`
	StatusesCount  opt.Optional[jsonint.Int]   `json:"statuses_count"`
	FollowersCount opt.Optional[jsonint.Int]   `json:"followers_count"`
	FollowingCount opt.Optional[jsonint.Int]   `json:"following_count"`
	Roles        []Role                        `json:"roles"`
	MuteExpiresAt  nul.Nullable[string]        `json:"mute_expires_at"`
}

func (receiver Account) MarshalJSON() ([]byte, error) {

	var buffer []byte

	buffer = append(buffer, "{"...)

	{
		buffer = append(buffer, `"id":`...)

		marshaled, err := json.Marshal(receiver.ID)
		if nil != err {
			return nil, erorr.Errorf("mstdn/ent: could not marshal ent.Account.ID as JSON: %w", err)
		}

		buffer = append(buffer, marshaled...)
	}

	{
		buffer = append(buffer, `,"username":`...)

		marshaled, err := json.Marshal(receiver.UserName)
		if nil != err {
			return nil, erorr.Errorf("mstdn/ent: could not marshal ent.Account.UserName as JSON: %w", err)
		}

		buffer = append(buffer, marshaled...)
	}

	{
		buffer = append(buffer, `,"acct":`...)

		marshaled, err := json.Marshal(receiver.Acct)
		if nil != err {
			return nil, erorr.Errorf("mstdn/ent: could not marshal ent.Account.Acct as JSON: %w", err)
		}

		buffer = append(buffer, marshaled...)
	}

	{
		buffer = append(buffer, `,"url":`...)

		marshaled, err := json.Marshal(receiver.URL)
		if nil != err {
			return nil, erorr.Errorf("mstdn/ent: could not marshal ent.Account.URL as JSON: %w", err)
		}

		buffer = append(buffer, marshaled...)
	}

	{
		buffer = append(buffer, `,"uri":`...)

		marshaled, err := json.Marshal(receiver.URI)
		if nil != err {
			return nil, erorr.Errorf("mstdn/ent: could not marshal ent.Account.URI as JSON: %w", err)
		}

		buffer = append(buffer, marshaled...)
	}

	{
		buffer = append(buffer, `,"display_name":`...)

		marshaled, err := json.Marshal(receiver.DisplayName)
		if nil != err {
			return nil, erorr.Errorf("mstdn/ent: could not marshal ent.Account.DisplayName as JSON: %w", err)
		}

		buffer = append(buffer, marshaled...)
	}

	{
		buffer = append(buffer, `,"note":`...)

		marshaled, err := json.Marshal(receiver.Note)
		if nil != err {
			return nil, erorr.Errorf("mstdn/ent: could not marshal ent.Account.Note as JSON: %w", err)
		}

		buffer = append(buffer, marshaled...)
	}

	{
		buffer = append(buffer, `,"avatar":`...)

		marshaled, err := json.Marshal(receiver.Avatar)
		if nil != err {
			return nil, erorr.Errorf("mstdn/ent: could not marshal ent.Account.Avatar as JSON: %w", err)
		}

		buffer = append(buffer, marshaled...)
	}

	{
		buffer = append(buffer, `,"avatar_static":`...)

		marshaled, err := json.Marshal(receiver.AvatarStatic)
		if nil != err {
			return nil, erorr.Errorf("mstdn/ent: could not marshal ent.Account.AvatarStatic as JSON: %w", err)
		}

		buffer = append(buffer, marshaled...)
	}

	{
		buffer = append(buffer, `,"header":`...)

		marshaled, err := json.Marshal(receiver.Header)
		if nil != err {
			return nil, erorr.Errorf("mstdn/ent: could not marshal ent.Account.Header as JSON: %w", err)
		}

		buffer = append(buffer, marshaled...)
	}

	{
		buffer = append(buffer, `,"header_static":`...)

		marshaled, err := json.Marshal(receiver.HeaderStatic)
		if nil != err {
			return nil, erorr.Errorf("mstdn/ent: could not marshal ent.Account.HeaderStatic as JSON: %w", err)
		}

		buffer = append(buffer, marshaled...)
	}

	{
		buffer = append(buffer, `,"fields":`...)

		var src interface{} = receiver.Fields
		if nil == receiver.Fields {
			src = []Field{}
		}

		marshaled, err := json.Marshal(src)
		if nil != err {
			return nil, erorr.Errorf("mstdn/ent: could not marshal ent.Account.Fields as JSON: %w", err)
		}

		buffer = append(buffer, marshaled...)
	}

	{
		buffer = append(buffer, `,"emojis":`...)

		var src interface{} = receiver.Emojis
		if nil == receiver.Emojis {
			src = []CustomEmoji{}
		}

		marshaled, err := json.Marshal(src)
		if nil != err {
			return nil, erorr.Errorf("mstdn/ent: could not marshal ent.Account.Emojis as JSON: %w", err)
		}

		buffer = append(buffer, marshaled...)
	}

	{
		buffer = append(buffer, `,"locked":`...)

		marshaled, err := json.Marshal(receiver.Locked)
		if nil != err {
			return nil, erorr.Errorf("mstdn/ent: could not marshal ent.Account.Locked as JSON: %w", err)
		}

		buffer = append(buffer, marshaled...)
	}

	{
		buffer = append(buffer, `,"bot":`...)

		marshaled, err := json.Marshal(receiver.Bot)
		if nil != err {
			return nil, erorr.Errorf("mstdn/ent: could not marshal ent.Account.Bot as JSON: %w", err)
		}

		buffer = append(buffer, marshaled...)
	}

	{
		buffer = append(buffer, `,"group":`...)

		marshaled, err := json.Marshal(receiver.Group)
		if nil != err {
			return nil, erorr.Errorf("mstdn/ent: could not marshal ent.Account.Group as JSON: %w", err)
		}

		buffer = append(buffer, marshaled...)
	}

	{
		buffer = append(buffer, `,"discoverable":`...)

		marshaled, err := json.Marshal(receiver.Discoverable)
		if nil != err {
			return nil, erorr.Errorf("mstdn/ent: could not marshal ent.Account.Discoverable as JSON: %w", err)
		}

		buffer = append(buffer, marshaled...)
	}

	{
		if nul.Nothing[bool]() != receiver.NoIndex {

			buffer = append(buffer, `,"noindex":`...)

			marshaled, err := json.Marshal(receiver.NoIndex)
			if nil != err {
				return nil, erorr.Errorf("mstdn/ent: could not marshal ent.Account.NoIndex as JSON: %w", err)
			}

			buffer = append(buffer, marshaled...)
		}
	}

	{
		if nul.Nothing[AccountHolder]() != receiver.Moved {

			buffer = append(buffer, `,"moved":`...)

			marshaled, err := json.Marshal(receiver.Moved)
			if nil != err {
				return nil, erorr.Errorf("mstdn/ent: could not marshal ent.Account.Moved as JSON: %w", err)
			}

			buffer = append(buffer, marshaled...)
		}
	}

	{
		if opt.Nothing[bool]() != receiver.Suspended {

			buffer = append(buffer, `,"suspended":`...)

			marshaled, err := json.Marshal(receiver.Suspended)
			if nil != err {
				return nil, erorr.Errorf("mstdn/ent: could not marshal ent.Account.Suspended as JSON: %w", err)
			}

			buffer = append(buffer, marshaled...)
		}
	}

	{
		if opt.Nothing[bool]() != receiver.Limited {

			buffer = append(buffer, `,"limited":`...)

			marshaled, err := json.Marshal(receiver.Limited)
			if nil != err {
				return nil, erorr.Errorf("mstdn/ent: could not marshal ent.Account.Limited as JSON: %w", err)
			}

			buffer = append(buffer, marshaled...)
		}
	}

	{
		buffer = append(buffer, `,"created_at":`...)

		marshaled, err := json.Marshal(receiver.CreatedAt)
		if nil != err {
			return nil, erorr.Errorf("mstdn/ent: could not marshal ent.Account.CreatedAt as JSON: %w", err)
		}

		buffer = append(buffer, marshaled...)
	}

	{
		buffer = append(buffer, `,"last_status_at":`...)

		marshaled, err := json.Marshal(receiver.LastStatusAt)
		if nil != err {
			return nil, erorr.Errorf("mstdn/ent: could not marshal ent.Account.LastStatusAt as JSON: %w", err)
		}

		buffer = append(buffer, marshaled...)
	}

	{
		buffer = append(buffer, `,"statuses_count":`...)

		marshaled, err := json.Marshal(receiver.StatusesCount)
		if nil != err {
			return nil, erorr.Errorf("mstdn/ent: could not marshal ent.Account.StatusesCount as JSON: %w", err)
		}

		buffer = append(buffer, marshaled...)
	}

	{
		buffer = append(buffer, `,"followers_count":`...)

		marshaled, err := json.Marshal(receiver.FollowersCount)
		if nil != err {
			return nil, erorr.Errorf("mstdn/ent: could not marshal ent.Account.FollowersCount as JSON: %w", err)
		}

		buffer = append(buffer, marshaled...)
	}

	{
		buffer = append(buffer, `,"following_count":`...)

		marshaled, err := json.Marshal(receiver.FollowingCount)
		if nil != err {
			return nil, erorr.Errorf("mstdn/ent: could not marshal ent.Account.FollowingCount as JSON: %w", err)
		}

		buffer = append(buffer, marshaled...)
	}

	{
		if nil != receiver.Roles {

			buffer = append(buffer, `,"roles":`...)

			marshaled, err := json.Marshal(receiver.Roles)
			if nil != err {
				return nil, erorr.Errorf("mstdn/ent: could not marshal ent.Account.Roles as JSON: %w", err)
			}

			buffer = append(buffer, marshaled...)
		}
	}

	{
		if nul.Nothing[string]() != receiver.MuteExpiresAt {

			buffer = append(buffer, `,"mute_expires_at":`...)

			marshaled, err := json.Marshal(receiver.MuteExpiresAt)
			if nil != err {
				return nil, erorr.Errorf("mstdn/ent: could not marshal ent.Account.MuteExpiresAt as JSON: %w", err)
			}

			buffer = append(buffer, marshaled...)
		}
	}

	buffer = append(buffer, "}"...)

	return buffer, nil
}
