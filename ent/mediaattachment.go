package ent

import (
	"encoding/json"

	"github.com/reiver/go-opt"
	"sourcecode.social/reiver/go-nul"
)

// See:
// https://docs.joinmastodon.org/entities/MediaAttachment/
type MediaAttachment struct {
	ID          opt.Optional[string] `json:"id"`
	Type        opt.Optional[string] `json:"type"`
	URL         opt.Optional[string] `json:"url"`
	PreviewURL  opt.Optional[string] `json:"preview_url"`
	RemoteURL   nul.Nullable[string] `json:"remote_url"`
	Meta        json.RawMessage      `json:"meta"`
	Description opt.Optional[string] `json:"description"`
	BlurHash    opt.Optional[string] `json:"blurhash"`
}
