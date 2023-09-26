package ent

import (
	"encoding/json"

	"sourcecode.social/reiver/go-opt"
	"sourcecode.social/reiver/go-jsonint"
)

var _ json.Marshaler = Role{}
var _ json.Unmarshaler = new(Role)

// Role represents a Mastodon API "Role".
//
// See:
// https://docs.joinmastodon.org/entities/Role/
type Role struct {
	ID          opt.Optional[jsonint.Int] `json:"id"`
	Name        opt.Optional[string]      `json:"name"`
	Color       opt.Optional[string]      `json:"color"`
	Permissions opt.Optional[jsonint.Int] `json:"permissions"`
	Highlighted opt.Optional[bool]        `json:"highlighted"`
}

type role struct {
	ID          opt.Optional[jsonint.Int] `json:"id"`
	Name        opt.Optional[string]      `json:"name"`
	Color       opt.Optional[string]      `json:"color"`
	Permissions opt.Optional[jsonint.Int] `json:"permissions"`
	Highlighted opt.Optional[bool]        `json:"Highlighted"`
}

func (receiver Role) MarshalJSON() ([]byte, error) {

	var src role

	src.ID          = receiver.ID
	src.Name        = receiver.Name
	src.Permissions = receiver.Permissions
	src.Highlighted = src.Highlighted

	src.Color = receiver.Color
	receiver.Color.WhenNothing(func(){
		src.Color = opt.Something("")
	})

	return json.Marshal(src)
}

func (receiver *Role) UnmarshalJSON(data []byte) error {
	if nil == receiver {
		return errNilReceiver
	}

	var dst role

	if err := json.Unmarshal(data, &dst); nil != err {
		return err
	}

	receiver.ID          = dst.ID
	receiver.Name        = dst.Name
	receiver.Permissions = dst.Permissions
	receiver.Highlighted = dst.Highlighted

	dst.Color.WhenSomething(func(value string){
		if "" == value {
			return
		}

		receiver.Color = opt.Something(value)
	})

	return nil
}
