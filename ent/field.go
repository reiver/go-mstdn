package ent

import (
	"encoding/json"

	"github.com/reiver/go-erorr"
	"github.com/reiver/go-opt"
	"github.com/reiver/go-nul"
)

var _ json.Marshaler = Field{}
var _ json.Unmarshaler = &Field{}

const (
	errCannotMashalFieldAsJSONNoName  = erorr.Error("cannot marshal ent.Field to JSON — no ‘name’ set")
	errCannotMashalFieldAsJSONNoValue = erorr.Error("cannot marshal ent.Field to JSON — no ‘value’ set")
)

// Field represents a Mastodon API "Field".
//
// See:
// https://docs.joinmastodon.org/entities/Field/
type Field struct {
	Name         opt.Optional[string]
	Value        opt.Optional[string]
	VerifiedAt   nul.Nullable[string]
}

type field struct {
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

	var data = map[string]interface{}{}

	{
		val, found := receiver.Name.Get()
		if !found {
				return nil, errCannotMashalFieldAsJSONNoName
		}

		data["name"] = val
	}
	{
		val, found := receiver.Value.Get()
		if !found {
			return nil, errCannotMashalFieldAsJSONNoValue
		}

		data["value"] = val
	}
	{
		val, found := receiver.VerifiedAt.Get()
		if !found {
			data["verified_at"] = nil
		} else {
			data["verified_at"] = val
		}
	}

	return json.Marshal(data)
}

func (receiver *Field) UnmarshalJSON(data []byte) error {
	if nil == receiver {
		return errNilReceiver
	}

	var f field

	err := json.Unmarshal(data, &f)
	if nil != err {
		return err
	}

	if nul.Null[string]() == f.VerifiedAt {
		f.VerifiedAt = nul.Nothing[string]()
	}

	*receiver = Field{
		Name: f.Name,
		Value: f.Value,
		VerifiedAt: f.VerifiedAt,
	}

	return nil
}
