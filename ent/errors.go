package ent

import (
	"github.com/reiver/go-erorr"
)

const (
	errNilReceiver = erorr.Error("mstdn/ent: nil receiver")
	errNothingID   = erorr.Error("mstdn/ent: nothing id")
)
