package ent

import (
	"github.com/reiver/go-jsonint"
	"github.com/reiver/go-opt"
	"sourcecode.social/reiver/go-nul"
)

type Poll struct {
	ID          opt.Optional[string]      `json:"id"`
	ExpiresAt   nul.Nullable[string]      `json:"expires_at"`
	Expired     opt.Optional[bool]        `json:"expired"`
	Multiple    opt.Optional[bool]        `json:"multiple"`
	VotesCount  opt.Optional[jsonint.Int] `json:"votes_count"`
	VotersCount opt.Optional[jsonint.Int] `json:"voters_count"`
	Options   []PollOption                `json:"options"`
	Emojis    []CustomEmoji               `json:"emojis"`
	Voted       opt.Optional[bool]        `json:"voted"`
	OwnVotes  []jsonint.Int              `json:"own_votes"`
}
