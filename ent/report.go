package ent

import (
	"sourcecode.social/reiver/go-jsonstr"
	"sourcecode.social/reiver/go-nul"
	"github.com/reiver/go-opt"
)

// Report represents a Mastodon API "Report".
//
// See:
// https://docs.joinmastodon.org/entities/Report/
type Report struct {
	ID            opt.Optional[string]          `json:"id"`
	ActionTaken   opt.Optional[bool]            `json:"action_taken"`
	ActionTakenAt nul.Nullable[string]          `json:"action_taken_at"`
	Category      opt.Optional[string]          `json:"category"`
	Comment       opt.Optional[string]          `json:"comment"`
	Forwarded     opt.Optional[bool]            `json:"forwarded"`
	CreatedAt     opt.Optional[string]          `json:"created_at"`
	StatusIDs     nul.Nullable[jsonstr.Strings] `json:"status_ids"`
	RuleIDs       nul.Nullable[jsonstr.Strings] `json:"rule_ids"`
	TargetAccount Account                       `json:"target_account"`
}
