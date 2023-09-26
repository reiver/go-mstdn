package mstdn

import (
	"encoding/json"

	"sourcecode.social/reiver/go-opt"
	"sourcecode.social/reiver/go-nul"
)

var _ json.Marshaler = Field{}
var _ json.Unmarshaler = &Field{}

// Field represents a Mastodon API "Field".
//
// See:
// https://docs.joinmastodon.org/entities/Field/
type Field struct {
	Name         opt.Optional[string] `json:"name"`
	Value        opt.Optional[string] `json:"value"`
	VerifiedAt   nul.Nullable[string] `json:"verified_at"`
}

func FieldNameValue(name string, value string) Field {
	return Field{
		Name:  opt.Something(name),
		Value: opt.Something(value),
	}
}

func FieldVerifiedNameValue(when string, name string, value string) Field {
	return Field{
		Name:       opt.Something(name),
		Value:      opt.Something(value),
		VerifiedAt: nul.Something(when),
	}
}

func (receiver Field) MarshalJSON() ([]byte, error) {

	duplicate := receiver

	duplicate.VerifiedAt.WhenNothing(func(){
		duplicate.VerifiedAt = nul.Null[string]()
	})

	return json.Marshal(duplicate)
}

func (receiver *Field) UnmarshalJSON(data []byte) error {
	if nil == receiver {
		return errNilReceiver
	}

	err := json.Unmarshal(data, receiver)
	if nil != err {
		return err
	}

	if nul.Null[string]() == receiver.VerifiedAt {
		receiver.VerifiedAt = nul.Nothing[string]()
	}

	return nil
}
