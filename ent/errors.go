package ent

import (
	"github.com/reiver/go-erorr"
)

const (
	ErrNilReceiver = erorr.Error("mstdn/ent: nil receiver")
)

const (
	errNothingID        = erorr.Error("mstdn/ent: nothing id")
	errNothingCreatedAt = erorr.Error("mstdn/ent: nothing created_at")
)
