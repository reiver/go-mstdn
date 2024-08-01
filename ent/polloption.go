package ent

import (
	"github.com/reiver/go-jsonint"
	"github.com/reiver/go-opt"
	"sourcecode.social/reiver/go-nul"
)

type PollOption struct {
	Title      opt.Optional[string]      `json:"title"`
	VotesCount nul.Nullable[jsonint.Int] `json:"votes_count"`
}
