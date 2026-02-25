package ent

import (
	"codeberg.org/reiver/go-erorr"
)

const (
	ErrNilNote     = erorr.Error("mstdn/ent: nil note")
	ErrNilReceiver = erorr.Error("mstdn/ent: nil receiver")
)

const (
	errNothingID        = erorr.Error("mstdn/ent: nothing id")
	errNothingCreatedAt = erorr.Error("mstdn/ent: nothing created_at")
)
