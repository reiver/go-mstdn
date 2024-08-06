package ent

import (
	"encoding/json"

	"github.com/reiver/go-jsonint"
	"github.com/reiver/go-opt"
	"github.com/reiver/go-nul"
)

type Status struct {
	ID                 opt.Optional[string]      `json:"id"`
	URI                nul.Nullable[string]      `json:"uri"`
	URL                nul.Nullable[string]      `json:"url"`
	CreatedAt          opt.Optional[string]      `json:"created_at"`
	Account            Account                   `json:"account"`
	Content            opt.Optional[string]      `json:"content"`
	Visibility         opt.Optional[string]      `json:"visibility"`
	Sensitive          opt.Optional[bool]        `json:"sensitive"`
	SpoilerText        opt.Optional[string]      `json:"spoiler_text"`
	MediaAttachments []MediaAttachment           `json:"media_attachments"`
	Application        nul.Nullable[Application] `json:"application"`
	Mentions         []Mention                   `json:"mentions"`
	Tags             []Tag                       `json:"tags"`
	Emojis           []CustomEmoji               `json:"emojis"`
	ReblogsCount       opt.Optional[jsonint.Int] `json:"reblogs_count"`
	FavouritesCount    opt.Optional[jsonint.Int] `json:"favourites_count"`
	RepliesCount       opt.Optional[jsonint.Int] `json:"replies_count"`
	InReplyToID        nul.Nullable[string]      `json:"in_reply_to_id"`
	InReplyToAccountID nul.Nullable[string]      `json:"in_reply_to_account_id"`
	Reblog             json.RawMessage           `json:"reblog"`
	Poll               nul.Nullable[Poll]        `json:"poll"`
	Card               nul.Nullable[PreviewCard] `json:"card"`
	Language           nul.Nullable[string]      `json:"language"`
	Text               nul.Nullable[string]      `json:"text"`
	EditedAt           nul.Nullable[string]      `json:"edited_at"`
	Favourited         opt.Optional[bool]        `json:"favourited"`
	Reblogged          opt.Optional[bool]        `json:"reblogged"`
	Muted              opt.Optional[bool]        `json:"muted"`
	Bookmarked         opt.Optional[bool]        `json:"bookmarked"`
	Pinned             opt.Optional[bool]        `json:"pinned"`
	Filtered           json.RawMessage           `json:"filtered"`
}
