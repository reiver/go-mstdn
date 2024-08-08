package ent

import (

	"github.com/reiver/go-opt"
	"github.com/reiver/go-nul"
)

// Field represents a Mastodon API "Field".
//
// See:
// https://docs.joinmastodon.org/entities/Field/
type Field struct {
	Name         opt.Optional[string] `json:"name,omitempty"`
	Value        opt.Optional[string] `json:"value,omitempty"`
	VerifiedAt   nul.Nullable[string] `json:"verified_at"`
}

func FieldNameValue(name string, value string) Field {
	return Field{
		Name:  opt.Something(name),
		Value: opt.Something(value),
		VerifiedAt: nul.Null[string](),
	}
}

func FieldVerifiedNameValue(when string, name string, value string) Field {
	return Field{
		Name:       opt.Something(name),
		Value:      opt.Something(value),
		VerifiedAt: nul.Something(when),
	}
}
