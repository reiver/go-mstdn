package ent

import (
	"encoding/json"
)

var _ json.Marshaler = AccountHolder{}
var _ json.Unmarshaler = new(AccountHolder)

// AccountHolder is a holder for a Mastodon API "Account".
//
// The way that Account and AccountHolder is different is that —
// AccountHolder just stores the JSON it receives from a UnmarshalJSON.
// It does not try to parse it or validate it.
//
// AccountHolder is useful for situations where you have a recursive type.
// Such as with Account.Moved.
//
// See:
// https://docs.joinmastodon.org/entities/Account/
//
// And in particular see:
// https://docs.joinmastodon.org/entities/Account/#moved
type AccountHolder struct {
	json string
}

//func (receiver AccountHolder) Account() (Account, error) {
//
//}

func (receiver AccountHolder) MarshalJSON() ([]byte, error) {
	if "" == receiver.json {
		return []byte("{}"), nil
	}

	return []byte(receiver.json), nil
}

func (receiver *AccountHolder) UnmarshalJSON(data []byte) error {
	if nil == receiver {
		return errNilReceiver
	}

	if 2 == len(data) && '{' == data[0] && '}' == data[1] {
		receiver.json = ""
		return nil
	}

	receiver.json = string(data)
	return nil
}
