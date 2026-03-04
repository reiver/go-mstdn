package ent_test

import (

	"github.com/reiver/go-json"
	"github.com/reiver/go-jsonint"
	"github.com/reiver/go-nul"
	"github.com/reiver/go-opt"

	"github.com/reiver/go-mstdn/ent"
)

func toJSON(value any) string {
	bytes, err := json.Marshal(value)
	if nil != err {
		panic(err)
	}

	return string(bytes)
}

func demoAccount1() ent.Account {
	return ent.Account{
		ID:            opt.Something("47"),
		UserName:      opt.Something("joeblow"),
		Acct:          opt.Something("joeblow@example.com"),
		URL:           opt.Something("https://example.com/@joeblow"),
		URI:           opt.Something("https://example.com/users/joeblow"),
		DisplayName:   opt.Something("Joe Blow :-)"),
		Note:          opt.Something("<p>Hello world!</p>"),
		Avatar:        opt.Something("https://files.example.com/avatar/joeblow.png"),
		AvatarStatic:  opt.Something("https://files.example.com/avatar-static/joeblow.png"),
		Header:        opt.Something("https://files.example.com/header/joeblow.png"),
		HeaderStatic:  opt.Something("https://files.example.com/header-static/joeblow.png"),
		//Fields
		//Emojis
		Locked:         opt.Something(false),
		Bot:            opt.Something(false),
		Group:          opt.Something(false),
		Discoverable:   nul.Null[bool](),
		NoIndex:        nul.Null[bool](),
		Moved:          nul.Null[ent.AccountHolder](),
		CreatedAt:      opt.Something("2023-09-27T22:06:19Z"),
		LastStatusAt:   nul.Null[string](),
		StatusesCount:  opt.Something(jsonint.NumericFromInt64(123)),
		FollowersCount: opt.Something(jsonint.NumericFromInt64(24789)),
		FollowingCount: opt.Something(jsonint.NumericFromInt64(355)),
		//Roles
		MuteExpiresAt:  nul.Null[string](),
	}
}
