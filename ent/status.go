package ent

import (
	gojson "encoding/json"
	"time"

	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-jsonint"
	"github.com/reiver/go-nul"
	"github.com/reiver/go-opt"
)

type Status struct {
	ID                 opt.Optional[string]          `json:"id,omitempty"`
	URI                nul.Nullable[string]          `json:"uri,omitempty"`
	URL                nul.Nullable[string]          `json:"url,omitempty"`
	CreatedAt          nul.Nullable[string]          `json:"created_at,omitempty"`
	Account            Account                       `json:"account,omitempty"`
	Content            opt.Optional[string]          `json:"content,omitempty"`
	Visibility         nul.Nullable[string]          `json:"visibility,omitempty"`
	Sensitive          nul.Nullable[bool]            `json:"sensitive,omitempty"`
	SpoilerText        nul.Nullable[string]          `json:"spoiler_text,omitempty"`
	MediaAttachments   []MediaAttachment             `json:"media_attachments,omitempty"`
	Application        nul.Nullable[Application]     `json:"application,omitempty"`
	Mentions           []Mention                     `json:"mentions,omitempty"`
	Tags               []Tag                         `json:"tags,omitempty"`
	Emojis             []CustomEmoji                 `json:"emojis,omitempty"`
	ReblogsCount       nul.Nullable[jsonint.Numeric] `json:"reblogs_count,omitempty"`
	FavouritesCount    nul.Nullable[jsonint.Numeric] `json:"favourites_count,omitempty"`
	RepliesCount       nul.Nullable[jsonint.Numeric] `json:"replies_count,omitempty"`
	InReplyToID        nul.Nullable[string]          `json:"in_reply_to_id,omitempty"`
	InReplyToAccountID nul.Nullable[string]          `json:"in_reply_to_account_id,omitempty"`
	Reblog             gojson.RawMessage             `json:"reblog,omitempty"`
	Poll               nul.Nullable[Poll]            `json:"poll,omitempty"`
	Card               nul.Nullable[PreviewCard]     `json:"card,omitempty"`
	Language           nul.Nullable[string]          `json:"language,omitempty"`
	Text               nul.Nullable[string]          `json:"text,omitempty"`
	EditedAt           nul.Nullable[string]          `json:"edited_at,omitempty"`
	Favourited         nul.Nullable[bool]            `json:"favourited,omitempty"`
	Reblogged          nul.Nullable[bool]            `json:"reblogged,omitempty"`
	Muted              nul.Nullable[bool]            `json:"muted,omitempty"`
	Bookmarked         nul.Nullable[bool]            `json:"bookmarked,omitempty"`
	Pinned             nul.Nullable[bool]            `json:"pinned,omitempty"`
	Filtered           gojson.RawMessage             `json:"filtered,omitempty"`
}

func (receiver *Status) ActivityNote(note *activitypub.Note) error {
	if nil == receiver {
		panic(ErrNilReceiver)
	}

	if nil == note {
		return ErrNilNote
	}

	if receiver.Account.ID.IsSomething() {
		attributedTo, found := receiver.Account.ID.Get()
		if found {
			note.AttributedTo = append(note.AttributedTo, activitypub.ObjectID(attributedTo))
		}
	}

	if receiver.Content.IsSomething() {
		note.Content = receiver.Content
	}

	if receiver.CreatedAt.IsSomething() {
		//@TODO: Do we need to change the time format for this conversion?
		note.Published = receiver.CreatedAt.Optional()
	}

	if receiver.URI.IsSomething() || receiver.URL.IsSomething() {
		ref := receiver.URI.Optional()
		if ref.IsNothing() {
			ref = receiver.URL.Optional()
		}

		note.ID = ref
		note.URL = ref
	}

	if 0 < len(receiver.MediaAttachments) {

		//@TODO

	}

	if 0 < len(receiver.Mentions) {

		//@TODO

	}

	if receiver.SpoilerText.IsSomething() {
		note.Summary = nul.Nothing[string]()

		value, found := receiver.SpoilerText.Get()
		if found {
			note.Summary = nul.Something(value)
		}
	}

	if 0 < len(receiver.Tags) {
		for _, tag := range receiver.Tags {
			var hashtag activitypub.HashTag
			hashtag.Name = tag.Name
			hashtag.HRef = tag.URL

			note.Tags = append(note.Tags, hashtag)
		}
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
