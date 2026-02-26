package ent

import (
	gojson "encoding/json"
	"time"

	"codeberg.org/reiver/go-asns"
	"github.com/reiver/go-jsonint"
	"github.com/reiver/go-opt"
	"github.com/reiver/go-nul"
)

type Status struct {
	ID                 opt.Optional[string]      `json:"id,omitempty"`
	URI                nul.Nullable[string]      `json:"uri,omitempty"`
	URL                nul.Nullable[string]      `json:"url,omitempty"`
	CreatedAt          opt.Optional[string]      `json:"created_at,omitempty"`
	Account            Account                   `json:"account,omitempty"`
	Content            opt.Optional[string]      `json:"content,omitempty"`
	Visibility         opt.Optional[string]      `json:"visibility,omitempty"`
	Sensitive          opt.Optional[bool]        `json:"sensitive,omitempty"`
	SpoilerText        opt.Optional[string]      `json:"spoiler_text,omitempty"`
	MediaAttachments []MediaAttachment           `json:"media_attachments,omitempty"`
	Application        nul.Nullable[Application] `json:"application,omitempty"`
	Mentions         []Mention                   `json:"mentions,omitempty"`
	Tags             []Tag                       `json:"tags,omitempty"`
	Emojis           []CustomEmoji               `json:"emojis,omitempty"`
	ReblogsCount       opt.Optional[jsonint.Int] `json:"reblogs_count,omitempty"`
	FavouritesCount    opt.Optional[jsonint.Int] `json:"favourites_count,omitempty"`
	RepliesCount       opt.Optional[jsonint.Int] `json:"replies_count,omitempty"`
	InReplyToID        nul.Nullable[string]      `json:"in_reply_to_id,omitempty"`
	InReplyToAccountID nul.Nullable[string]      `json:"in_reply_to_account_id,omitempty"`
	Reblog             gojson.RawMessage         `json:"reblog,omitempty"`
	Poll               nul.Nullable[Poll]        `json:"poll,omitempty"`
	Card               nul.Nullable[PreviewCard] `json:"card,omitempty"`
	Language           nul.Nullable[string]      `json:"language,omitempty"`
	Text               nul.Nullable[string]      `json:"text,omitempty"`
	EditedAt           nul.Nullable[string]      `json:"edited_at,omitempty"`
	Favourited         opt.Optional[bool]        `json:"favourited,omitempty"`
	Reblogged          opt.Optional[bool]        `json:"reblogged,omitempty"`
	Muted              opt.Optional[bool]        `json:"muted,omitempty"`
	Bookmarked         opt.Optional[bool]        `json:"bookmarked,omitempty"`
	Pinned             opt.Optional[bool]        `json:"pinned,omitempty"`
	Filtered           gojson.RawMessage         `json:"filtered,omitempty"`
}

func (receiver *Status) ActivityNote(note *asns.Note) error {
	if nil == receiver {
		panic(ErrNilReceiver)
	}

	if nil == note {
		return ErrNilNote
	}

	if receiver.Account.ID.IsSomething() {
		attributedTo, found := receiver.Account.ID.Get()
		if found {
			note.AttributedTo = asns.SomeString(attributedTo)
		}
	}

	if receiver.Content.IsSomething() {
		note.Content = receiver.Content
	}

	if receiver.CreatedAt.IsSomething() {
		//@TODO: Do we need to change the time format for this conversion?
		note.Published = receiver.CreatedAt
	}

	if receiver.URI.IsSomething() {
		note.ID  = receiver.URI.Optional()
		note.URL = receiver.URI.Optional()
	}

	if 0 < len(receiver.MediaAttachments) {
		
//@TODO
		
	}

	if 0 < len(receiver.Mentions) {
		
//@TODO
		
	}

	if receiver.SpoilerText.IsSomething() {
		note.Summary = receiver.SpoilerText
	}

	if 0 < len(receiver.Tags) {
		
//@TODO
		
	}

	return nil
}

func (receiver *Status) ParseCreatedAt() (time.Time, error) {
	if nil == receiver {
		panic(ErrNilReceiver)
	}

	createdAt, found := receiver.CreatedAt.Get()
	if !found {
		var nada time.Time
		return nada, errNothingCreatedAt
	}

	return time.Parse(time.RFC3339Nano, createdAt)
}
