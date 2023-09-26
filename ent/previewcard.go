package ent

import (
	"sourcecode.social/reiver/go-opt"
	"sourcecode.social/reiver/go-nul"
)

// See:
// https://docs.joinmastodon.org/entities/PreviewCard/
type PreviewCard struct {
	URL          opt.Optional[string] `json:"url"`
	Title        opt.Optional[string] `json:"title"`
	Description  opt.Optional[string] `json:"description"`
	Type         opt.Optional[string] `json:"type"`
	AuthorName   opt.Optional[string] `json:"author_name"`
	AuthorURL    opt.Optional[string] `json:"author_url"`
	ProviderName opt.Optional[string] `json:"provider_name"`
	ProviderURL  opt.Optional[string] `json:"provider_url"`
	HTML         opt.Optional[string] `json:"html"`
	Width        opt.Optional[string] `json:"width"`
	Height       opt.Optional[string] `json:"height"`
	Image        nul.Nullable[string] `json:"image"`
	EmbedURL     opt.Optional[string] `json:"embed_url"`
	BlurHash     nul.Nullable[string] `json:"blurhash"`
	History    []History              `json:"history"`
}
