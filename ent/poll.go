package ent

import (
	"github.com/reiver/go-jsonint"
	"github.com/reiver/go-opt"
	"github.com/reiver/go-nul"
)

type Poll struct {
	ID          opt.Optional[string]      `json:"id,omitempty"`
	ExpiresAt   nul.Nullable[string]      `json:"expires_at,omitempty"`
	Expired     opt.Optional[bool]        `json:"expired,omitempty"`
	Multiple    opt.Optional[bool]        `json:"multiple,omitempty"`
	VotesCount  opt.Optional[jsonint.Int] `json:"votes_count,omitempty"`
	VotersCount opt.Optional[jsonint.Int] `json:"voters_count,omitempty"`
	Options   []PollOption                `json:"options,omitempty"`
	Emojis    []CustomEmoji               `json:"emojis,omitempty"`
	Voted       opt.Optional[bool]        `json:"voted,omitempty"`
	OwnVotes  []jsonint.Int               `json:"own_votes,omitempty"`
}
