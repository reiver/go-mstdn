package ent

import (
	"github.com/reiver/go-jsonstr"
	"github.com/reiver/go-nul"
	"github.com/reiver/go-opt"
)

// Report represents a Mastodon API "Report".
//
// See:
// https://docs.joinmastodon.org/entities/Report/
type Report struct {
	ID            opt.Optional[string]          `json:"id,omitempty"`
	ActionTaken   opt.Optional[bool]            `json:"action_taken,omitempty"`
	ActionTakenAt nul.Nullable[string]          `json:"action_taken_at,omitempty"`
	Category      opt.Optional[string]          `json:"category,omitempty"`
	Comment       opt.Optional[string]          `json:"comment,omitempty"`
	Forwarded     opt.Optional[bool]            `json:"forwarded,omitempty"`
	CreatedAt     opt.Optional[string]          `json:"created_at,omitempty"`
	StatusIDs     nul.Nullable[jsonstr.Strings] `json:"status_ids,omitempty"`
	RuleIDs       nul.Nullable[jsonstr.Strings] `json:"rule_ids,omitempty"`
	TargetAccount opt.Optional[Account]         `json:"target_account,omitempty"`
}
