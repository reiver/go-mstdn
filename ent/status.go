package ent

import (
	"encoding/json"

	"github.com/reiver/go-erorr"
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

func (receiver *Status) MarshalJSON() ([]byte, error) {
	if nil == receiver {
		return nil, errNilReceiver
	}

	var buffer [1024]byte
	var p []byte = buffer[0:0]

	p = append(p, '{')

	{
		val, found := receiver.ID.Get()
		if !found {
			return nil, errNothingID
		}

		const name =    "id"
		p  = append(p, `"id":`...)

		bytes, err := json.Marshal(val)
		if nil != err {
			return nil, erorr.Errorf("problem marshaling %q: %w", name, err)
		}

		p = append(p, bytes...)
	}

	{
		val, found :=    receiver.URI.Get()
		if found {
			const name =     "uri"
			p  = append(p, `,"uri":`...)

			bytes, err := json.Marshal(val)
			if nil != err {
				return nil, erorr.Errorf("problem marshaling %q: %w", name, err)
			}

			p = append(p, bytes...)
		}
	}

	{
		val, found :=    receiver.URL.Get()
		if found {
			const name =     "url"
			p  = append(p, `,"url":`...)

			bytes, err := json.Marshal(val)
			if nil != err {
				return nil, erorr.Errorf("problem marshaling %q: %w", name, err)
			}

			p = append(p, bytes...)
		}
	}

	{
		val, found :=    receiver.CreatedAt.Get()
		if found {
			const name =     "created_at"
			p  = append(p, `,"created_at":`...)

			bytes, err := json.Marshal(val)
			if nil != err {
				return nil, erorr.Errorf("problem marshaling %q: %w", name, err)
			}

			p = append(p, bytes...)
		}
	}

//@TODO: Account

	{
		val, found :=    receiver.Content.Get()
		if found {
			const name =     "content"
			p  = append(p, `,"content":`...)

			bytes, err := json.Marshal(val)
			if nil != err {
				return nil, erorr.Errorf("problem marshaling %q: %w", name, err)
			}

			p = append(p, bytes...)
		}
	}

	{
		val, found :=    receiver.Visibility.Get()
		if found {
			const name =     "visibility"
			p  = append(p, `,"visibility":`...)

			bytes, err := json.Marshal(val)
			if nil != err {
				return nil, erorr.Errorf("problem marshaling %q: %w", name, err)
			}

			p = append(p, bytes...)
		}
	}

	{
		val, found :=    receiver.Sensitive.Get()
		if found {
			const name =     "sensitive"
			p  = append(p, `,"sensitive":`...)

			bytes, err := json.Marshal(val)
			if nil != err {
				return nil, erorr.Errorf("problem marshaling %q: %w", name, err)
			}

			p = append(p, bytes...)
		}
	}

	{
		val, found :=    receiver.SpoilerText.Get()
		if found {
			const name =     "spoiler_text"
			p  = append(p, `,"spoiler_text":`...)

			bytes, err := json.Marshal(val)
			if nil != err {
				return nil, erorr.Errorf("problem marshaling %q: %w", name, err)
			}

			p = append(p, bytes...)
		}
	}

//@TODO: MediaAttachments

//@TODO: Application

//@TODO: Mentions

	{
		vals := receiver.Tags
		if 0 < len(vals) {
			const name =     "tags"
			p  = append(p, `,"tags":`...)

			bytes, err := json.Marshal(vals)
			if nil != err {
				return nil, erorr.Errorf("problem marshaling %q: %w", name, err)
			}

			p = append(p, bytes...)
		}
	}

//@TODO: Emojis

//@TODO: ReblogsCount

//@TODO: FavouritesCount

//@TODO: RepliesCount

//@TODO: InReplyToID

//@TODO: InReplyToAccountID

//@TODO: Reblog

//@TODO: Poll

//@TODO: Card

	{
		val, found :=    receiver.Language.Get()
		if found {
			const name =     "language"
			p  = append(p, `,"language":`...)

			bytes, err := json.Marshal(val)
			if nil != err {
				return nil, erorr.Errorf("problem marshaling %q: %w", name, err)
			}

			p = append(p, bytes...)
		}
	}

	{
		val, found :=    receiver.Text.Get()
		if found {
			const name =     "text"
			p  = append(p, `,"text":`...)

			bytes, err := json.Marshal(val)
			if nil != err {
				return nil, erorr.Errorf("problem marshaling %q: %w", name, err)
			}

			p = append(p, bytes...)
		}
	}

	{
		val, found :=    receiver.EditedAt.Get()
		if found {
			const name =     "edited_at"
			p  = append(p, `,"edited_at":`...)

			bytes, err := json.Marshal(val)
			if nil != err {
				return nil, erorr.Errorf("problem marshaling %q: %w", name, err)
			}

			p = append(p, bytes...)
		}
	}

	{
		val, found :=    receiver.Favourited.Get()
		if found {
			const name =     "favourited"
			p  = append(p, `,"favourited":`...)

			bytes, err := json.Marshal(val)
			if nil != err {
				return nil, erorr.Errorf("problem marshaling %q: %w", name, err)
			}

			p = append(p, bytes...)
		}
	}

	{
		val, found :=    receiver.Reblogged.Get()
		if found {
			const name =     "reblogged"
			p  = append(p, `,"reblogged":`...)

			bytes, err := json.Marshal(val)
			if nil != err {
				return nil, erorr.Errorf("problem marshaling %q: %w", name, err)
			}

			p = append(p, bytes...)
		}
	}

	{
		val, found :=    receiver.Muted.Get()
		if found {
			const name =     "muted"
			p  = append(p, `,"muted":`...)

			bytes, err := json.Marshal(val)
			if nil != err {
				return nil, erorr.Errorf("problem marshaling %q: %w", name, err)
			}

			p = append(p, bytes...)
		}
	}

	{
		val, found :=    receiver.Bookmarked.Get()
		if found {
			const name =     "bookmarked"
			p  = append(p, `,"bookmarked":`...)

			bytes, err := json.Marshal(val)
			if nil != err {
				return nil, erorr.Errorf("problem marshaling %q: %w", name, err)
			}

			p = append(p, bytes...)
		}
	}

	{
		val, found :=    receiver.Pinned.Get()
		if found {
			const name =     "pinned"
			p  = append(p, `,"pinned":`...)

			bytes, err := json.Marshal(val)
			if nil != err {
				return nil, erorr.Errorf("problem marshaling %q: %w", name, err)
			}

			p = append(p, bytes...)
		}
	}

	{
		bytes :=  []byte(receiver.Filtered)
		if 0 < len(bytes) {
			const name =     "filtered"
			p  = append(p, `,"filtered":`...)

			p = append(p, bytes...)
		}
	}

	p = append(p, '}')

	return p, nil
}
