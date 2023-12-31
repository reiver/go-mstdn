package ent

import (
	"sourcecode.social/reiver/go-jsonint"
	"sourcecode.social/reiver/go-opt"
	"sourcecode.social/reiver/go-nul"
)

type PollOption struct {
	Title      opt.Optional[string]      `json:"title"`
	VotesCount nul.Nullable[jsonint.Int] `json:"votes_count"`
}
