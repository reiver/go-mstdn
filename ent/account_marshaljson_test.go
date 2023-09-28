package ent_test

import (
	"testing"

	"encoding/json"

	"sourcecode.social/reiver/go-jsonint"
	"sourcecode.social/reiver/go-nul"
	"sourcecode.social/reiver/go-opt"

	"sourcecode.social/reiver/go-mstdn/ent"
)

func TestAccount_MarshalJSON(t *testing.T) {

	tests := []struct{
		Account ent.Account
		Expected string
	}{
		{
			Account: ent.Account{
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
				CreatedAt:      opt.Something("2023-09-27T22:06:19Z"),
				LastStatusAt:   nul.Null[string](),
				StatusesCount:  opt.Something(jsonint.Int64(123)),
				FollowersCount: opt.Something(jsonint.Int64(24789)),
				FollowingCount: opt.Something(jsonint.Int64(355)),
				//Roles
				//MuteExpiresAt
			},

			Expected: `{`+
				`"id":"47"`+
				`,`+
				`"username":"joeblow"`+
				`,`+
				`"acct":"joeblow@example.com"`+
				`,`+
				`"url":"https://example.com/@joeblow"`+
				`,`+
				`"uri":"https://example.com/users/joeblow"`+
				`,`+
				`"display_name":"Joe Blow :-)"`+
				`,`+
				`"note":"\u003cp\u003eHello world!\u003c/p\u003e"`+
				`,`+
				`"avatar":"https://files.example.com/avatar/joeblow.png"`+
				`,`+
				`"avatar_static":"https://files.example.com/avatar-static/joeblow.png"`+
				`,`+
				`"header":"https://files.example.com/header/joeblow.png"`+
				`,`+
				`"header_static":"https://files.example.com/header-static/joeblow.png"`+
				`,`+
				`"fields":[]`+
				`,`+
				`"emojis":[]`+
				`,`+
				`"locked":false`+
				`,`+
				`"bot":false`+
				`,`+
				`"group":false`+
				`,`+
				`"discoverable":null`+
				`,`+
				`"created_at":"2023-09-27T22:06:19Z"`+
				`,`+
				`"last_status_at":null`+
				`,`+
				`"statuses_count":123`+
				`,`+
				`"followers_count":24789`+
				`,`+
				`"following_count":355`+
			`}`,
		},









		{
			Account: ent.Account{
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
				Locked:         opt.Something(true),
				Bot:            opt.Something(false),
				Group:          opt.Something(false),
				Discoverable:   nul.Null[bool](),
				CreatedAt:      opt.Something("2023-09-27T22:06:19Z"),
				LastStatusAt:   nul.Null[string](),
				StatusesCount:  opt.Something(jsonint.Int64(123)),
				FollowersCount: opt.Something(jsonint.Int64(24789)),
				FollowingCount: opt.Something(jsonint.Int64(355)),
				//Roles
				//MuteExpiresAt
			},

			Expected: `{`+
				`"id":"47"`+
				`,`+
				`"username":"joeblow"`+
				`,`+
				`"acct":"joeblow@example.com"`+
				`,`+
				`"url":"https://example.com/@joeblow"`+
				`,`+
				`"uri":"https://example.com/users/joeblow"`+
				`,`+
				`"display_name":"Joe Blow :-)"`+
				`,`+
				`"note":"\u003cp\u003eHello world!\u003c/p\u003e"`+
				`,`+
				`"avatar":"https://files.example.com/avatar/joeblow.png"`+
				`,`+
				`"avatar_static":"https://files.example.com/avatar-static/joeblow.png"`+
				`,`+
				`"header":"https://files.example.com/header/joeblow.png"`+
				`,`+
				`"header_static":"https://files.example.com/header-static/joeblow.png"`+
				`,`+
				`"fields":[]`+
				`,`+
				`"emojis":[]`+
				`,`+
				`"locked":true`+
				`,`+
				`"bot":false`+
				`,`+
				`"group":false`+
				`,`+
				`"discoverable":null`+
				`,`+
				`"created_at":"2023-09-27T22:06:19Z"`+
				`,`+
				`"last_status_at":null`+
				`,`+
				`"statuses_count":123`+
				`,`+
				`"followers_count":24789`+
				`,`+
				`"following_count":355`+
			`}`,
		},



		{
			Account: ent.Account{
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
				Bot:            opt.Something(true),
				Group:          opt.Something(false),
				Discoverable:   nul.Null[bool](),
				CreatedAt:      opt.Something("2023-09-27T22:06:19Z"),
				LastStatusAt:   nul.Null[string](),
				StatusesCount:  opt.Something(jsonint.Int64(123)),
				FollowersCount: opt.Something(jsonint.Int64(24789)),
				FollowingCount: opt.Something(jsonint.Int64(355)),
				//Roles
				//MuteExpiresAt
			},

			Expected: `{`+
				`"id":"47"`+
				`,`+
				`"username":"joeblow"`+
				`,`+
				`"acct":"joeblow@example.com"`+
				`,`+
				`"url":"https://example.com/@joeblow"`+
				`,`+
				`"uri":"https://example.com/users/joeblow"`+
				`,`+
				`"display_name":"Joe Blow :-)"`+
				`,`+
				`"note":"\u003cp\u003eHello world!\u003c/p\u003e"`+
				`,`+
				`"avatar":"https://files.example.com/avatar/joeblow.png"`+
				`,`+
				`"avatar_static":"https://files.example.com/avatar-static/joeblow.png"`+
				`,`+
				`"header":"https://files.example.com/header/joeblow.png"`+
				`,`+
				`"header_static":"https://files.example.com/header-static/joeblow.png"`+
				`,`+
				`"fields":[]`+
				`,`+
				`"emojis":[]`+
				`,`+
				`"locked":false`+
				`,`+
				`"bot":true`+
				`,`+
				`"group":false`+
				`,`+
				`"discoverable":null`+
				`,`+
				`"created_at":"2023-09-27T22:06:19Z"`+
				`,`+
				`"last_status_at":null`+
				`,`+
				`"statuses_count":123`+
				`,`+
				`"followers_count":24789`+
				`,`+
				`"following_count":355`+
			`}`,
		},



		{
			Account: ent.Account{
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
				Group:          opt.Something(true),
				Discoverable:   nul.Null[bool](),
				CreatedAt:      opt.Something("2023-09-27T22:06:19Z"),
				LastStatusAt:   nul.Null[string](),
				StatusesCount:  opt.Something(jsonint.Int64(123)),
				FollowersCount: opt.Something(jsonint.Int64(24789)),
				FollowingCount: opt.Something(jsonint.Int64(355)),
				//Roles
				//MuteExpiresAt
			},

			Expected: `{`+
				`"id":"47"`+
				`,`+
				`"username":"joeblow"`+
				`,`+
				`"acct":"joeblow@example.com"`+
				`,`+
				`"url":"https://example.com/@joeblow"`+
				`,`+
				`"uri":"https://example.com/users/joeblow"`+
				`,`+
				`"display_name":"Joe Blow :-)"`+
				`,`+
				`"note":"\u003cp\u003eHello world!\u003c/p\u003e"`+
				`,`+
				`"avatar":"https://files.example.com/avatar/joeblow.png"`+
				`,`+
				`"avatar_static":"https://files.example.com/avatar-static/joeblow.png"`+
				`,`+
				`"header":"https://files.example.com/header/joeblow.png"`+
				`,`+
				`"header_static":"https://files.example.com/header-static/joeblow.png"`+
				`,`+
				`"fields":[]`+
				`,`+
				`"emojis":[]`+
				`,`+
				`"locked":false`+
				`,`+
				`"bot":false`+
				`,`+
				`"group":true`+
				`,`+
				`"discoverable":null`+
				`,`+
				`"created_at":"2023-09-27T22:06:19Z"`+
				`,`+
				`"last_status_at":null`+
				`,`+
				`"statuses_count":123`+
				`,`+
				`"followers_count":24789`+
				`,`+
				`"following_count":355`+
			`}`,
		},








		{
			Account: ent.Account{
				ID:            opt.Something("2737523"),
				UserName:      opt.Something("janedoe"),
				Acct:          opt.Something("janedoe@something.example"),
				URL:           opt.Something("https://something.example/@janedoe"),
				URI:           opt.Something("https://something.example/~janedoe"),
				DisplayName:   opt.Something("Jane Doe &lt;3"),
				Note:          opt.Something("<p>Hello world!</p>"),
				Avatar:        opt.Something("https://files.example.com/avatar/janedoe.png"),
				AvatarStatic:  opt.Something("https://files.example.com/avatar-static/janedoe.png"),
				Header:        opt.Something("https://files.example.com/header/janedoe.png"),
				HeaderStatic:  opt.Something("https://files.example.com/header-static/janedoe.png"),
				//Fields
				//Emojis
				Locked:         opt.Something(false),
				Bot:            opt.Something(false),
				Group:          opt.Something(false),
				Discoverable:   nul.Something(false),
				CreatedAt:      opt.Something("2023-09-27T22:06:19Z"),
				LastStatusAt:   nul.Null[string](),
				StatusesCount:  opt.Something(jsonint.Int64(123)),
				FollowersCount: opt.Something(jsonint.Int64(24789)),
				FollowingCount: opt.Something(jsonint.Int64(355)),
				//Roles
				//MuteExpiresAt
			},

			Expected: `{`+
				`"id":"2737523"`+
				`,`+
				`"username":"janedoe"`+
				`,`+
				`"acct":"janedoe@something.example"`+
				`,`+
				`"url":"https://something.example/@janedoe"`+
				`,`+
				`"uri":"https://something.example/~janedoe"`+
				`,`+
				`"display_name":"Jane Doe \u0026lt;3"`+
				`,`+
				`"note":"\u003cp\u003eHello world!\u003c/p\u003e"`+
				`,`+
				`"avatar":"https://files.example.com/avatar/janedoe.png"`+
				`,`+
				`"avatar_static":"https://files.example.com/avatar-static/janedoe.png"`+
				`,`+
				`"header":"https://files.example.com/header/janedoe.png"`+
				`,`+
				`"header_static":"https://files.example.com/header-static/janedoe.png"`+
				`,`+
				`"fields":[]`+
				`,`+
				`"emojis":[]`+
				`,`+
				`"locked":false`+
				`,`+
				`"bot":false`+
				`,`+
				`"group":false`+
				`,`+
				`"discoverable":false`+
				`,`+
				`"created_at":"2023-09-27T22:06:19Z"`+
				`,`+
				`"last_status_at":null`+
				`,`+
				`"statuses_count":123`+
				`,`+
				`"followers_count":24789`+
				`,`+
				`"following_count":355`+
			`}`,
		},



		{
			Account: ent.Account{
				ID:            opt.Something("2737523"),
				UserName:      opt.Something("janedoe"),
				Acct:          opt.Something("janedoe@something.example"),
				URL:           opt.Something("https://something.example/@janedoe"),
				URI:           opt.Something("https://something.example/~janedoe"),
				DisplayName:   opt.Something("Jane Doe &lt;3"),
				Note:          opt.Something("<p>Hello world!</p>"),
				Avatar:        opt.Something("https://files.example.com/avatar/janedoe.png"),
				AvatarStatic:  opt.Something("https://files.example.com/avatar-static/janedoe.png"),
				Header:        opt.Something("https://files.example.com/header/janedoe.png"),
				HeaderStatic:  opt.Something("https://files.example.com/header-static/janedoe.png"),
				//Fields
				//Emojis
				Locked:         opt.Something(false),
				Bot:            opt.Something(false),
				Group:          opt.Something(false),
				Discoverable:   nul.Something(true),
				CreatedAt:      opt.Something("2023-09-27T22:06:19Z"),
				LastStatusAt:   nul.Null[string](),
				StatusesCount:  opt.Something(jsonint.Int64(123)),
				FollowersCount: opt.Something(jsonint.Int64(24789)),
				FollowingCount: opt.Something(jsonint.Int64(355)),
				//Roles
				//MuteExpiresAt
			},

			Expected: `{`+
				`"id":"2737523"`+
				`,`+
				`"username":"janedoe"`+
				`,`+
				`"acct":"janedoe@something.example"`+
				`,`+
				`"url":"https://something.example/@janedoe"`+
				`,`+
				`"uri":"https://something.example/~janedoe"`+
				`,`+
				`"display_name":"Jane Doe \u0026lt;3"`+
				`,`+
				`"note":"\u003cp\u003eHello world!\u003c/p\u003e"`+
				`,`+
				`"avatar":"https://files.example.com/avatar/janedoe.png"`+
				`,`+
				`"avatar_static":"https://files.example.com/avatar-static/janedoe.png"`+
				`,`+
				`"header":"https://files.example.com/header/janedoe.png"`+
				`,`+
				`"header_static":"https://files.example.com/header-static/janedoe.png"`+
				`,`+
				`"fields":[]`+
				`,`+
				`"emojis":[]`+
				`,`+
				`"locked":false`+
				`,`+
				`"bot":false`+
				`,`+
				`"group":false`+
				`,`+
				`"discoverable":true`+
				`,`+
				`"created_at":"2023-09-27T22:06:19Z"`+
				`,`+
				`"last_status_at":null`+
				`,`+
				`"statuses_count":123`+
				`,`+
				`"followers_count":24789`+
				`,`+
				`"following_count":355`+
			`}`,
		},









		{
			Account: ent.Account{
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
				CreatedAt:      opt.Something("2022-08-16T11:05:08Z"),
				LastStatusAt:   nul.Something("2023-09-27T22:06:19Z"),
				StatusesCount:  opt.Something(jsonint.Int64(123)),
				FollowersCount: opt.Something(jsonint.Int64(24789)),
				FollowingCount: opt.Something(jsonint.Int64(355)),
				//Roles
				//MuteExpiresAt
			},

			Expected: `{`+
				`"id":"47"`+
				`,`+
				`"username":"joeblow"`+
				`,`+
				`"acct":"joeblow@example.com"`+
				`,`+
				`"url":"https://example.com/@joeblow"`+
				`,`+
				`"uri":"https://example.com/users/joeblow"`+
				`,`+
				`"display_name":"Joe Blow :-)"`+
				`,`+
				`"note":"\u003cp\u003eHello world!\u003c/p\u003e"`+
				`,`+
				`"avatar":"https://files.example.com/avatar/joeblow.png"`+
				`,`+
				`"avatar_static":"https://files.example.com/avatar-static/joeblow.png"`+
				`,`+
				`"header":"https://files.example.com/header/joeblow.png"`+
				`,`+
				`"header_static":"https://files.example.com/header-static/joeblow.png"`+
				`,`+
				`"fields":[]`+
				`,`+
				`"emojis":[]`+
				`,`+
				`"locked":false`+
				`,`+
				`"bot":false`+
				`,`+
				`"group":false`+
				`,`+
				`"discoverable":null`+
				`,`+
				`"created_at":"2022-08-16T11:05:08Z"`+
				`,`+
				`"last_status_at":"2023-09-27T22:06:19Z"`+
				`,`+
				`"statuses_count":123`+
				`,`+
				`"followers_count":24789`+
				`,`+
				`"following_count":355`+
			`}`,
		},









		{
			Account: ent.Account{
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
				Fields:        nil,
				//Emojis
				Locked:         opt.Something(false),
				Bot:            opt.Something(false),
				Group:          opt.Something(false),
				Discoverable:   nul.Null[bool](),
				CreatedAt:      opt.Something("2023-09-27T22:06:19Z"),
				LastStatusAt:   nul.Null[string](),
				StatusesCount:  opt.Something(jsonint.Int64(123)),
				FollowersCount: opt.Something(jsonint.Int64(24789)),
				FollowingCount: opt.Something(jsonint.Int64(355)),
				//Roles
				//MuteExpiresAt
			},

			Expected: `{`+
				`"id":"47"`+
				`,`+
				`"username":"joeblow"`+
				`,`+
				`"acct":"joeblow@example.com"`+
				`,`+
				`"url":"https://example.com/@joeblow"`+
				`,`+
				`"uri":"https://example.com/users/joeblow"`+
				`,`+
				`"display_name":"Joe Blow :-)"`+
				`,`+
				`"note":"\u003cp\u003eHello world!\u003c/p\u003e"`+
				`,`+
				`"avatar":"https://files.example.com/avatar/joeblow.png"`+
				`,`+
				`"avatar_static":"https://files.example.com/avatar-static/joeblow.png"`+
				`,`+
				`"header":"https://files.example.com/header/joeblow.png"`+
				`,`+
				`"header_static":"https://files.example.com/header-static/joeblow.png"`+
				`,`+
				`"fields":[]`+
				`,`+
				`"emojis":[]`+
				`,`+
				`"locked":false`+
				`,`+
				`"bot":false`+
				`,`+
				`"group":false`+
				`,`+
				`"discoverable":null`+
				`,`+
				`"created_at":"2023-09-27T22:06:19Z"`+
				`,`+
				`"last_status_at":null`+
				`,`+
				`"statuses_count":123`+
				`,`+
				`"followers_count":24789`+
				`,`+
				`"following_count":355`+
			`}`,
		},



		{
			Account: ent.Account{
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
				Fields:        []ent.Field{},
				//Emojis
				Locked:         opt.Something(false),
				Bot:            opt.Something(false),
				Group:          opt.Something(false),
				Discoverable:   nul.Null[bool](),
				CreatedAt:      opt.Something("2023-09-27T22:06:19Z"),
				LastStatusAt:   nul.Null[string](),
				StatusesCount:  opt.Something(jsonint.Int64(123)),
				FollowersCount: opt.Something(jsonint.Int64(24789)),
				FollowingCount: opt.Something(jsonint.Int64(355)),
				//Roles
				//MuteExpiresAt
			},

			Expected: `{`+
				`"id":"47"`+
				`,`+
				`"username":"joeblow"`+
				`,`+
				`"acct":"joeblow@example.com"`+
				`,`+
				`"url":"https://example.com/@joeblow"`+
				`,`+
				`"uri":"https://example.com/users/joeblow"`+
				`,`+
				`"display_name":"Joe Blow :-)"`+
				`,`+
				`"note":"\u003cp\u003eHello world!\u003c/p\u003e"`+
				`,`+
				`"avatar":"https://files.example.com/avatar/joeblow.png"`+
				`,`+
				`"avatar_static":"https://files.example.com/avatar-static/joeblow.png"`+
				`,`+
				`"header":"https://files.example.com/header/joeblow.png"`+
				`,`+
				`"header_static":"https://files.example.com/header-static/joeblow.png"`+
				`,`+
				`"fields":[]`+
				`,`+
				`"emojis":[]`+
				`,`+
				`"locked":false`+
				`,`+
				`"bot":false`+
				`,`+
				`"group":false`+
				`,`+
				`"discoverable":null`+
				`,`+
				`"created_at":"2023-09-27T22:06:19Z"`+
				`,`+
				`"last_status_at":null`+
				`,`+
				`"statuses_count":123`+
				`,`+
				`"followers_count":24789`+
				`,`+
				`"following_count":355`+
			`}`,
		},



		{
			Account: ent.Account{
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
				Fields:        []ent.Field{
					ent.Field{
						Name: opt.Something("alt"),
						Value: opt.Something("@joeblow@greatape.social"),
					},
				},
				//Emojis
				Locked:         opt.Something(false),
				Bot:            opt.Something(false),
				Group:          opt.Something(false),
				Discoverable:   nul.Null[bool](),
				CreatedAt:      opt.Something("2023-09-27T22:06:19Z"),
				LastStatusAt:   nul.Null[string](),
				StatusesCount:  opt.Something(jsonint.Int64(123)),
				FollowersCount: opt.Something(jsonint.Int64(24789)),
				FollowingCount: opt.Something(jsonint.Int64(355)),
				//Roles
				//MuteExpiresAt
			},

			Expected: `{`+
				`"id":"47"`+
				`,`+
				`"username":"joeblow"`+
				`,`+
				`"acct":"joeblow@example.com"`+
				`,`+
				`"url":"https://example.com/@joeblow"`+
				`,`+
				`"uri":"https://example.com/users/joeblow"`+
				`,`+
				`"display_name":"Joe Blow :-)"`+
				`,`+
				`"note":"\u003cp\u003eHello world!\u003c/p\u003e"`+
				`,`+
				`"avatar":"https://files.example.com/avatar/joeblow.png"`+
				`,`+
				`"avatar_static":"https://files.example.com/avatar-static/joeblow.png"`+
				`,`+
				`"header":"https://files.example.com/header/joeblow.png"`+
				`,`+
				`"header_static":"https://files.example.com/header-static/joeblow.png"`+
				`,`+
				`"fields":[`+
					`{`+
						`"name":"alt"`+
						`,`+
						`"value":"@joeblow@greatape.social"`+
						`,`+
						`"verified_at":null`+
					`}`+
				`]`+
				`,`+
				`"emojis":[]`+
				`,`+
				`"locked":false`+
				`,`+
				`"bot":false`+
				`,`+
				`"group":false`+
				`,`+
				`"discoverable":null`+
				`,`+
				`"created_at":"2023-09-27T22:06:19Z"`+
				`,`+
				`"last_status_at":null`+
				`,`+
				`"statuses_count":123`+
				`,`+
				`"followers_count":24789`+
				`,`+
				`"following_count":355`+
			`}`,
		},



		{
			Account: ent.Account{
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
				Fields:        []ent.Field{
					ent.Field{
						Name: opt.Something("alt"),
						Value: opt.Something("@joeblow@greatape.social"),
					},
					ent.Field{
						Name: opt.Something("alt (long-form)"),
						Value: opt.Something("@joeblow@postfreely.social"),
						VerifiedAt: nul.Something("2022-09-29T08:41:46Z"),
					},
				},
				//Emojis
				Locked:         opt.Something(false),
				Bot:            opt.Something(false),
				Group:          opt.Something(false),
				Discoverable:   nul.Null[bool](),
				CreatedAt:      opt.Something("2023-09-27T22:06:19Z"),
				LastStatusAt:   nul.Null[string](),
				StatusesCount:  opt.Something(jsonint.Int64(123)),
				FollowersCount: opt.Something(jsonint.Int64(24789)),
				FollowingCount: opt.Something(jsonint.Int64(355)),
				//Roles
				//MuteExpiresAt
			},

			Expected: `{`+
				`"id":"47"`+
				`,`+
				`"username":"joeblow"`+
				`,`+
				`"acct":"joeblow@example.com"`+
				`,`+
				`"url":"https://example.com/@joeblow"`+
				`,`+
				`"uri":"https://example.com/users/joeblow"`+
				`,`+
				`"display_name":"Joe Blow :-)"`+
				`,`+
				`"note":"\u003cp\u003eHello world!\u003c/p\u003e"`+
				`,`+
				`"avatar":"https://files.example.com/avatar/joeblow.png"`+
				`,`+
				`"avatar_static":"https://files.example.com/avatar-static/joeblow.png"`+
				`,`+
				`"header":"https://files.example.com/header/joeblow.png"`+
				`,`+
				`"header_static":"https://files.example.com/header-static/joeblow.png"`+
				`,`+
				`"fields":[`+
					`{`+
						`"name":"alt"`+
						`,`+
						`"value":"@joeblow@greatape.social"`+
						`,`+
						`"verified_at":null`+
					`}`+
					`,`+
					`{`+
						`"name":"alt (long-form)"`+
						`,`+
						`"value":"@joeblow@postfreely.social"`+
						`,`+
						`"verified_at":"2022-09-29T08:41:46Z"`+
					`}`+
				`]`+
				`,`+
				`"emojis":[]`+
				`,`+
				`"locked":false`+
				`,`+
				`"bot":false`+
				`,`+
				`"group":false`+
				`,`+
				`"discoverable":null`+
				`,`+
				`"created_at":"2023-09-27T22:06:19Z"`+
				`,`+
				`"last_status_at":null`+
				`,`+
				`"statuses_count":123`+
				`,`+
				`"followers_count":24789`+
				`,`+
				`"following_count":355`+
			`}`,
		},









		{
			Account: ent.Account{
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
				Emojis:         nil,
				Locked:         opt.Something(false),
				Bot:            opt.Something(false),
				Group:          opt.Something(false),
				Discoverable:   nul.Null[bool](),
				CreatedAt:      opt.Something("2023-09-27T22:06:19Z"),
				LastStatusAt:   nul.Null[string](),
				StatusesCount:  opt.Something(jsonint.Int64(123)),
				FollowersCount: opt.Something(jsonint.Int64(24789)),
				FollowingCount: opt.Something(jsonint.Int64(355)),
				//Roles
				//MuteExpiresAt
			},

			Expected: `{`+
				`"id":"47"`+
				`,`+
				`"username":"joeblow"`+
				`,`+
				`"acct":"joeblow@example.com"`+
				`,`+
				`"url":"https://example.com/@joeblow"`+
				`,`+
				`"uri":"https://example.com/users/joeblow"`+
				`,`+
				`"display_name":"Joe Blow :-)"`+
				`,`+
				`"note":"\u003cp\u003eHello world!\u003c/p\u003e"`+
				`,`+
				`"avatar":"https://files.example.com/avatar/joeblow.png"`+
				`,`+
				`"avatar_static":"https://files.example.com/avatar-static/joeblow.png"`+
				`,`+
				`"header":"https://files.example.com/header/joeblow.png"`+
				`,`+
				`"header_static":"https://files.example.com/header-static/joeblow.png"`+
				`,`+
				`"fields":[]`+
				`,`+
				`"emojis":[]`+
				`,`+
				`"locked":false`+
				`,`+
				`"bot":false`+
				`,`+
				`"group":false`+
				`,`+
				`"discoverable":null`+
				`,`+
				`"created_at":"2023-09-27T22:06:19Z"`+
				`,`+
				`"last_status_at":null`+
				`,`+
				`"statuses_count":123`+
				`,`+
				`"followers_count":24789`+
				`,`+
				`"following_count":355`+
			`}`,
		},



		{
			Account: ent.Account{
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
				Emojis:       []ent.CustomEmoji{},
				Locked:         opt.Something(false),
				Bot:            opt.Something(false),
				Group:          opt.Something(false),
				Discoverable:   nul.Null[bool](),
				CreatedAt:      opt.Something("2023-09-27T22:06:19Z"),
				LastStatusAt:   nul.Null[string](),
				StatusesCount:  opt.Something(jsonint.Int64(123)),
				FollowersCount: opt.Something(jsonint.Int64(24789)),
				FollowingCount: opt.Something(jsonint.Int64(355)),
				//Roles
				//MuteExpiresAt
			},

			Expected: `{`+
				`"id":"47"`+
				`,`+
				`"username":"joeblow"`+
				`,`+
				`"acct":"joeblow@example.com"`+
				`,`+
				`"url":"https://example.com/@joeblow"`+
				`,`+
				`"uri":"https://example.com/users/joeblow"`+
				`,`+
				`"display_name":"Joe Blow :-)"`+
				`,`+
				`"note":"\u003cp\u003eHello world!\u003c/p\u003e"`+
				`,`+
				`"avatar":"https://files.example.com/avatar/joeblow.png"`+
				`,`+
				`"avatar_static":"https://files.example.com/avatar-static/joeblow.png"`+
				`,`+
				`"header":"https://files.example.com/header/joeblow.png"`+
				`,`+
				`"header_static":"https://files.example.com/header-static/joeblow.png"`+
				`,`+
				`"fields":[]`+
				`,`+
				`"emojis":[]`+
				`,`+
				`"locked":false`+
				`,`+
				`"bot":false`+
				`,`+
				`"group":false`+
				`,`+
				`"discoverable":null`+
				`,`+
				`"created_at":"2023-09-27T22:06:19Z"`+
				`,`+
				`"last_status_at":null`+
				`,`+
				`"statuses_count":123`+
				`,`+
				`"followers_count":24789`+
				`,`+
				`"following_count":355`+
			`}`,
		},



		{
			Account: ent.Account{
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
				Emojis:       []ent.CustomEmoji{
					ent.CustomEmoji{
						ShortCode: opt.Something("batman"),
						URL: opt.Something("https://files.example.com/emoji/batman.png"),
						StaticURL: opt.Something("https://files.example.com/emoji-static/batman.png"),
						VisibleInPicker: opt.Something(true),
					},
				},
				Locked:         opt.Something(false),
				Bot:            opt.Something(false),
				Group:          opt.Something(false),
				Discoverable:   nul.Null[bool](),
				CreatedAt:      opt.Something("2023-09-27T22:06:19Z"),
				LastStatusAt:   nul.Null[string](),
				StatusesCount:  opt.Something(jsonint.Int64(123)),
				FollowersCount: opt.Something(jsonint.Int64(24789)),
				FollowingCount: opt.Something(jsonint.Int64(355)),
				//Roles
				//MuteExpiresAt
			},

			Expected: `{`+
				`"id":"47"`+
				`,`+
				`"username":"joeblow"`+
				`,`+
				`"acct":"joeblow@example.com"`+
				`,`+
				`"url":"https://example.com/@joeblow"`+
				`,`+
				`"uri":"https://example.com/users/joeblow"`+
				`,`+
				`"display_name":"Joe Blow :-)"`+
				`,`+
				`"note":"\u003cp\u003eHello world!\u003c/p\u003e"`+
				`,`+
				`"avatar":"https://files.example.com/avatar/joeblow.png"`+
				`,`+
				`"avatar_static":"https://files.example.com/avatar-static/joeblow.png"`+
				`,`+
				`"header":"https://files.example.com/header/joeblow.png"`+
				`,`+
				`"header_static":"https://files.example.com/header-static/joeblow.png"`+
				`,`+
				`"fields":[]`+
				`,`+
				`"emojis":[`+
					`{`+
						`"shortcode":"batman"`+
						`,`+
						`"url":"https://files.example.com/emoji/batman.png"`+
						`,`+
						`"static_url":"https://files.example.com/emoji-static/batman.png"`+
						`,`+
						`"visible_in_picker":true`+
					`}`+
				`]`+
				`,`+
				`"locked":false`+
				`,`+
				`"bot":false`+
				`,`+
				`"group":false`+
				`,`+
				`"discoverable":null`+
				`,`+
				`"created_at":"2023-09-27T22:06:19Z"`+
				`,`+
				`"last_status_at":null`+
				`,`+
				`"statuses_count":123`+
				`,`+
				`"followers_count":24789`+
				`,`+
				`"following_count":355`+
			`}`,
		},



		{
			Account: ent.Account{
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
				Emojis:       []ent.CustomEmoji{
					ent.CustomEmoji{
						ShortCode: opt.Something("batman"),
						URL: opt.Something("https://files.example.com/emoji/batman.png"),
						StaticURL: opt.Something("https://files.example.com/emoji-static/batman.png"),
						VisibleInPicker: opt.Something(true),
					},
					ent.CustomEmoji{
						ShortCode: opt.Something("joker"),
						URL: opt.Something("https://files.example.com/emoji/joker.png"),
						StaticURL: opt.Something("https://files.example.com/emoji-static/joker.png"),
						VisibleInPicker: opt.Something(false),
					},
				},
				Locked:         opt.Something(false),
				Bot:            opt.Something(false),
				Group:          opt.Something(false),
				Discoverable:   nul.Null[bool](),
				CreatedAt:      opt.Something("2023-09-27T22:06:19Z"),
				LastStatusAt:   nul.Null[string](),
				StatusesCount:  opt.Something(jsonint.Int64(123)),
				FollowersCount: opt.Something(jsonint.Int64(24789)),
				FollowingCount: opt.Something(jsonint.Int64(355)),
				//Roles
				//MuteExpiresAt
			},

			Expected: `{`+
				`"id":"47"`+
				`,`+
				`"username":"joeblow"`+
				`,`+
				`"acct":"joeblow@example.com"`+
				`,`+
				`"url":"https://example.com/@joeblow"`+
				`,`+
				`"uri":"https://example.com/users/joeblow"`+
				`,`+
				`"display_name":"Joe Blow :-)"`+
				`,`+
				`"note":"\u003cp\u003eHello world!\u003c/p\u003e"`+
				`,`+
				`"avatar":"https://files.example.com/avatar/joeblow.png"`+
				`,`+
				`"avatar_static":"https://files.example.com/avatar-static/joeblow.png"`+
				`,`+
				`"header":"https://files.example.com/header/joeblow.png"`+
				`,`+
				`"header_static":"https://files.example.com/header-static/joeblow.png"`+
				`,`+
				`"fields":[]`+
				`,`+
				`"emojis":[`+
					`{`+
						`"shortcode":"batman"`+
						`,`+
						`"url":"https://files.example.com/emoji/batman.png"`+
						`,`+
						`"static_url":"https://files.example.com/emoji-static/batman.png"`+
						`,`+
						`"visible_in_picker":true`+
					`}`+
					`,`+
					`{`+
						`"shortcode":"joker"`+
						`,`+
						`"url":"https://files.example.com/emoji/joker.png"`+
						`,`+
						`"static_url":"https://files.example.com/emoji-static/joker.png"`+
						`,`+
						`"visible_in_picker":false`+
					`}`+
				`]`+
				`,`+
				`"locked":false`+
				`,`+
				`"bot":false`+
				`,`+
				`"group":false`+
				`,`+
				`"discoverable":null`+
				`,`+
				`"created_at":"2023-09-27T22:06:19Z"`+
				`,`+
				`"last_status_at":null`+
				`,`+
				`"statuses_count":123`+
				`,`+
				`"followers_count":24789`+
				`,`+
				`"following_count":355`+
			`}`,
		},



		{
			Account: ent.Account{
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
				Emojis:       []ent.CustomEmoji{
					ent.CustomEmoji{
						ShortCode: opt.Something("batman"),
						URL: opt.Something("https://files.example.com/emoji/batman.png"),
						StaticURL: opt.Something("https://files.example.com/emoji-static/batman.png"),
						VisibleInPicker: opt.Something(true),
					},
					ent.CustomEmoji{
						ShortCode: opt.Something("joker"),
						URL: opt.Something("https://files.example.com/emoji/joker.png"),
						StaticURL: opt.Something("https://files.example.com/emoji-static/joker.png"),
						VisibleInPicker: opt.Something(false),
					},
					ent.CustomEmoji{
						ShortCode: opt.Something("spider-man"),
						URL: opt.Something("https://files.example.com/emoji/spider-man.png"),
						StaticURL: opt.Something("https://files.example.com/emoji-static/spider-man.png"),
						VisibleInPicker: opt.Something(true),
						Category: opt.Something("spiders"),
					},
				},
				Locked:         opt.Something(false),
				Bot:            opt.Something(false),
				Group:          opt.Something(false),
				Discoverable:   nul.Null[bool](),
				CreatedAt:      opt.Something("2023-09-27T22:06:19Z"),
				LastStatusAt:   nul.Null[string](),
				StatusesCount:  opt.Something(jsonint.Int64(123)),
				FollowersCount: opt.Something(jsonint.Int64(24789)),
				FollowingCount: opt.Something(jsonint.Int64(355)),
				//Roles
				//MuteExpiresAt
			},

			Expected: `{`+
				`"id":"47"`+
				`,`+
				`"username":"joeblow"`+
				`,`+
				`"acct":"joeblow@example.com"`+
				`,`+
				`"url":"https://example.com/@joeblow"`+
				`,`+
				`"uri":"https://example.com/users/joeblow"`+
				`,`+
				`"display_name":"Joe Blow :-)"`+
				`,`+
				`"note":"\u003cp\u003eHello world!\u003c/p\u003e"`+
				`,`+
				`"avatar":"https://files.example.com/avatar/joeblow.png"`+
				`,`+
				`"avatar_static":"https://files.example.com/avatar-static/joeblow.png"`+
				`,`+
				`"header":"https://files.example.com/header/joeblow.png"`+
				`,`+
				`"header_static":"https://files.example.com/header-static/joeblow.png"`+
				`,`+
				`"fields":[]`+
				`,`+
				`"emojis":[`+
					`{`+
						`"shortcode":"batman"`+
						`,`+
						`"url":"https://files.example.com/emoji/batman.png"`+
						`,`+
						`"static_url":"https://files.example.com/emoji-static/batman.png"`+
						`,`+
						`"visible_in_picker":true`+
					`}`+
					`,`+
					`{`+
						`"shortcode":"joker"`+
						`,`+
						`"url":"https://files.example.com/emoji/joker.png"`+
						`,`+
						`"static_url":"https://files.example.com/emoji-static/joker.png"`+
						`,`+
						`"visible_in_picker":false`+
					`}`+
					`,`+
					`{`+
						`"shortcode":"spider-man"`+
						`,`+
						`"url":"https://files.example.com/emoji/spider-man.png"`+
						`,`+
						`"static_url":"https://files.example.com/emoji-static/spider-man.png"`+
						`,`+
						`"visible_in_picker":true`+
						`,`+
						`"category":"spiders"`+
					`}`+
				`]`+
				`,`+
				`"locked":false`+
				`,`+
				`"bot":false`+
				`,`+
				`"group":false`+
				`,`+
				`"discoverable":null`+
				`,`+
				`"created_at":"2023-09-27T22:06:19Z"`+
				`,`+
				`"last_status_at":null`+
				`,`+
				`"statuses_count":123`+
				`,`+
				`"followers_count":24789`+
				`,`+
				`"following_count":355`+
			`}`,
		},









		{
			Account: ent.Account{
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
				CreatedAt:      opt.Something("2023-09-27T22:06:19Z"),
				LastStatusAt:   nul.Null[string](),
				StatusesCount:  opt.Something(jsonint.Int64(123)),
				FollowersCount: opt.Something(jsonint.Int64(24789)),
				FollowingCount: opt.Something(jsonint.Int64(355)),
				Roles: nil,
				//MuteExpiresAt
			},

			Expected: `{`+
				`"id":"47"`+
				`,`+
				`"username":"joeblow"`+
				`,`+
				`"acct":"joeblow@example.com"`+
				`,`+
				`"url":"https://example.com/@joeblow"`+
				`,`+
				`"uri":"https://example.com/users/joeblow"`+
				`,`+
				`"display_name":"Joe Blow :-)"`+
				`,`+
				`"note":"\u003cp\u003eHello world!\u003c/p\u003e"`+
				`,`+
				`"avatar":"https://files.example.com/avatar/joeblow.png"`+
				`,`+
				`"avatar_static":"https://files.example.com/avatar-static/joeblow.png"`+
				`,`+
				`"header":"https://files.example.com/header/joeblow.png"`+
				`,`+
				`"header_static":"https://files.example.com/header-static/joeblow.png"`+
				`,`+
				`"fields":[]`+
				`,`+
				`"emojis":[]`+
				`,`+
				`"locked":false`+
				`,`+
				`"bot":false`+
				`,`+
				`"group":false`+
				`,`+
				`"discoverable":null`+
				`,`+
				`"created_at":"2023-09-27T22:06:19Z"`+
				`,`+
				`"last_status_at":null`+
				`,`+
				`"statuses_count":123`+
				`,`+
				`"followers_count":24789`+
				`,`+
				`"following_count":355`+
			`}`,
		},



		{
			Account: ent.Account{
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
				CreatedAt:      opt.Something("2023-09-27T22:06:19Z"),
				LastStatusAt:   nul.Null[string](),
				StatusesCount:  opt.Something(jsonint.Int64(123)),
				FollowersCount: opt.Something(jsonint.Int64(24789)),
				FollowingCount: opt.Something(jsonint.Int64(355)),
				Roles: []ent.Role{},
				//MuteExpiresAt
			},

			Expected: `{`+
				`"id":"47"`+
				`,`+
				`"username":"joeblow"`+
				`,`+
				`"acct":"joeblow@example.com"`+
				`,`+
				`"url":"https://example.com/@joeblow"`+
				`,`+
				`"uri":"https://example.com/users/joeblow"`+
				`,`+
				`"display_name":"Joe Blow :-)"`+
				`,`+
				`"note":"\u003cp\u003eHello world!\u003c/p\u003e"`+
				`,`+
				`"avatar":"https://files.example.com/avatar/joeblow.png"`+
				`,`+
				`"avatar_static":"https://files.example.com/avatar-static/joeblow.png"`+
				`,`+
				`"header":"https://files.example.com/header/joeblow.png"`+
				`,`+
				`"header_static":"https://files.example.com/header-static/joeblow.png"`+
				`,`+
				`"fields":[]`+
				`,`+
				`"emojis":[]`+
				`,`+
				`"locked":false`+
				`,`+
				`"bot":false`+
				`,`+
				`"group":false`+
				`,`+
				`"discoverable":null`+
				`,`+
				`"created_at":"2023-09-27T22:06:19Z"`+
				`,`+
				`"last_status_at":null`+
				`,`+
				`"statuses_count":123`+
				`,`+
				`"followers_count":24789`+
				`,`+
				`"following_count":355`+
				`,`+
				`"roles":[]`+
			`}`,
		},



		{
			Account: ent.Account{
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
				CreatedAt:      opt.Something("2023-09-27T22:06:19Z"),
				LastStatusAt:   nul.Null[string](),
				StatusesCount:  opt.Something(jsonint.Int64(123)),
				FollowersCount: opt.Something(jsonint.Int64(24789)),
				FollowingCount: opt.Something(jsonint.Int64(355)),
				Roles: []ent.Role{
					ent.Role{
						ID: opt.Something(jsonint.Int64(15)),
						Name: opt.Something("sponsor (PostFreely)"),
						Color: opt.Something("#ff2400"),
						Permissions:  opt.Something(jsonint.Int64(0)),
						Highlighted:  opt.Something(true),
					},
				},
				//MuteExpiresAt
			},

			Expected: `{`+
				`"id":"47"`+
				`,`+
				`"username":"joeblow"`+
				`,`+
				`"acct":"joeblow@example.com"`+
				`,`+
				`"url":"https://example.com/@joeblow"`+
				`,`+
				`"uri":"https://example.com/users/joeblow"`+
				`,`+
				`"display_name":"Joe Blow :-)"`+
				`,`+
				`"note":"\u003cp\u003eHello world!\u003c/p\u003e"`+
				`,`+
				`"avatar":"https://files.example.com/avatar/joeblow.png"`+
				`,`+
				`"avatar_static":"https://files.example.com/avatar-static/joeblow.png"`+
				`,`+
				`"header":"https://files.example.com/header/joeblow.png"`+
				`,`+
				`"header_static":"https://files.example.com/header-static/joeblow.png"`+
				`,`+
				`"fields":[]`+
				`,`+
				`"emojis":[]`+
				`,`+
				`"locked":false`+
				`,`+
				`"bot":false`+
				`,`+
				`"group":false`+
				`,`+
				`"discoverable":null`+
				`,`+
				`"created_at":"2023-09-27T22:06:19Z"`+
				`,`+
				`"last_status_at":null`+
				`,`+
				`"statuses_count":123`+
				`,`+
				`"followers_count":24789`+
				`,`+
				`"following_count":355`+
				`,`+
				`"roles":[`+
					`{`+
						`"id":15`+
						`,`+
						`"name":"sponsor (PostFreely)"`+
						`,`+
						`"color":"#ff2400"`+
						`,`+
						`"permissions":0`+
						`,`+
						`"highlighted":true`+
					`}`+
				`]`+
			`}`,
		},



		{
			Account: ent.Account{
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
				CreatedAt:      opt.Something("2023-09-27T22:06:19Z"),
				LastStatusAt:   nul.Null[string](),
				StatusesCount:  opt.Something(jsonint.Int64(123)),
				FollowersCount: opt.Something(jsonint.Int64(24789)),
				FollowingCount: opt.Something(jsonint.Int64(355)),
				Roles: []ent.Role{
					ent.Role{
						ID: opt.Something(jsonint.Int64(15)),
						Name: opt.Something("sponsor (PostFreely)"),
						Color: opt.Something("#ff2400"),
						Permissions:  opt.Something(jsonint.Int64(0)),
						Highlighted:  opt.Something(true),
					},
					ent.Role{
						ID: opt.Something(jsonint.Int64(28)),
						Name: opt.Something("sponsor (Firefish)"),
						Color: opt.Something("#30ff11"),
						Permissions:  opt.Something(jsonint.Int64(255)),
						Highlighted:  opt.Something(false),
					},
				},
				//MuteExpiresAt
			},

			Expected: `{`+
				`"id":"47"`+
				`,`+
				`"username":"joeblow"`+
				`,`+
				`"acct":"joeblow@example.com"`+
				`,`+
				`"url":"https://example.com/@joeblow"`+
				`,`+
				`"uri":"https://example.com/users/joeblow"`+
				`,`+
				`"display_name":"Joe Blow :-)"`+
				`,`+
				`"note":"\u003cp\u003eHello world!\u003c/p\u003e"`+
				`,`+
				`"avatar":"https://files.example.com/avatar/joeblow.png"`+
				`,`+
				`"avatar_static":"https://files.example.com/avatar-static/joeblow.png"`+
				`,`+
				`"header":"https://files.example.com/header/joeblow.png"`+
				`,`+
				`"header_static":"https://files.example.com/header-static/joeblow.png"`+
				`,`+
				`"fields":[]`+
				`,`+
				`"emojis":[]`+
				`,`+
				`"locked":false`+
				`,`+
				`"bot":false`+
				`,`+
				`"group":false`+
				`,`+
				`"discoverable":null`+
				`,`+
				`"created_at":"2023-09-27T22:06:19Z"`+
				`,`+
				`"last_status_at":null`+
				`,`+
				`"statuses_count":123`+
				`,`+
				`"followers_count":24789`+
				`,`+
				`"following_count":355`+
				`,`+
				`"roles":[`+
					`{`+
						`"id":15`+
						`,`+
						`"name":"sponsor (PostFreely)"`+
						`,`+
						`"color":"#ff2400"`+
						`,`+
						`"permissions":0`+
						`,`+
						`"highlighted":true`+
					`}`+
					`,`+
					`{`+
						`"id":28`+
						`,`+
						`"name":"sponsor (Firefish)"`+
						`,`+
						`"color":"#30ff11"`+
						`,`+
						`"permissions":255`+
						`,`+
						`"highlighted":false`+
					`}`+
				`]`+
			`}`,
		},



		{
			Account: ent.Account{
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
				CreatedAt:      opt.Something("2023-09-27T22:06:19Z"),
				LastStatusAt:   nul.Null[string](),
				StatusesCount:  opt.Something(jsonint.Int64(123)),
				FollowersCount: opt.Something(jsonint.Int64(24789)),
				FollowingCount: opt.Something(jsonint.Int64(355)),
				Roles: []ent.Role{
					ent.Role{
						ID: opt.Something(jsonint.Int64(15)),
						Name: opt.Something("sponsor (PostFreely)"),
						Color: opt.Something("#ff2400"),
						Permissions:  opt.Something(jsonint.Int64(0)),
						Highlighted:  opt.Something(true),
					},
					ent.Role{
						ID: opt.Something(jsonint.Int64(28)),
						Name: opt.Something("sponsor (Firefish)"),
						Color: opt.Something("#30ff11"),
						Permissions:  opt.Something(jsonint.Int64(255)),
						Highlighted:  opt.Something(false),
					},
					ent.Role{
						ID: opt.Something(jsonint.Int64(1)),
						Name: opt.Something("admin"),
						Permissions:  opt.Something(jsonint.Int64(4294967295)),
						Highlighted:  opt.Something(false),
					},
				},
				//MuteExpiresAt
			},

			Expected: `{`+
				`"id":"47"`+
				`,`+
				`"username":"joeblow"`+
				`,`+
				`"acct":"joeblow@example.com"`+
				`,`+
				`"url":"https://example.com/@joeblow"`+
				`,`+
				`"uri":"https://example.com/users/joeblow"`+
				`,`+
				`"display_name":"Joe Blow :-)"`+
				`,`+
				`"note":"\u003cp\u003eHello world!\u003c/p\u003e"`+
				`,`+
				`"avatar":"https://files.example.com/avatar/joeblow.png"`+
				`,`+
				`"avatar_static":"https://files.example.com/avatar-static/joeblow.png"`+
				`,`+
				`"header":"https://files.example.com/header/joeblow.png"`+
				`,`+
				`"header_static":"https://files.example.com/header-static/joeblow.png"`+
				`,`+
				`"fields":[]`+
				`,`+
				`"emojis":[]`+
				`,`+
				`"locked":false`+
				`,`+
				`"bot":false`+
				`,`+
				`"group":false`+
				`,`+
				`"discoverable":null`+
				`,`+
				`"created_at":"2023-09-27T22:06:19Z"`+
				`,`+
				`"last_status_at":null`+
				`,`+
				`"statuses_count":123`+
				`,`+
				`"followers_count":24789`+
				`,`+
				`"following_count":355`+
				`,`+
				`"roles":[`+
					`{`+
						`"id":15`+
						`,`+
						`"name":"sponsor (PostFreely)"`+
						`,`+
						`"color":"#ff2400"`+
						`,`+
						`"permissions":0`+
						`,`+
						`"highlighted":true`+
					`}`+
					`,`+
					`{`+
						`"id":28`+
						`,`+
						`"name":"sponsor (Firefish)"`+
						`,`+
						`"color":"#30ff11"`+
						`,`+
						`"permissions":255`+
						`,`+
						`"highlighted":false`+
					`}`+
					`,`+
					`{`+
						`"id":1`+
						`,`+
						`"name":"admin"`+
						`,`+
						`"color":""`+
						`,`+
						`"permissions":4294967295`+
						`,`+
						`"highlighted":false`+
					`}`+
				`]`+
			`}`,
		},









		{
			Account: ent.Account{
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
				CreatedAt:      opt.Something("2023-09-27T22:06:19Z"),
				LastStatusAt:   nul.Null[string](),
				StatusesCount:  opt.Something(jsonint.Int64(123)),
				FollowersCount: opt.Something(jsonint.Int64(24789)),
				FollowingCount: opt.Something(jsonint.Int64(355)),
				//Roles
				MuteExpiresAt: nul.Null[string](),
			},

			Expected: `{`+
				`"id":"47"`+
				`,`+
				`"username":"joeblow"`+
				`,`+
				`"acct":"joeblow@example.com"`+
				`,`+
				`"url":"https://example.com/@joeblow"`+
				`,`+
				`"uri":"https://example.com/users/joeblow"`+
				`,`+
				`"display_name":"Joe Blow :-)"`+
				`,`+
				`"note":"\u003cp\u003eHello world!\u003c/p\u003e"`+
				`,`+
				`"avatar":"https://files.example.com/avatar/joeblow.png"`+
				`,`+
				`"avatar_static":"https://files.example.com/avatar-static/joeblow.png"`+
				`,`+
				`"header":"https://files.example.com/header/joeblow.png"`+
				`,`+
				`"header_static":"https://files.example.com/header-static/joeblow.png"`+
				`,`+
				`"fields":[]`+
				`,`+
				`"emojis":[]`+
				`,`+
				`"locked":false`+
				`,`+
				`"bot":false`+
				`,`+
				`"group":false`+
				`,`+
				`"discoverable":null`+
				`,`+
				`"created_at":"2023-09-27T22:06:19Z"`+
				`,`+
				`"last_status_at":null`+
				`,`+
				`"statuses_count":123`+
				`,`+
				`"followers_count":24789`+
				`,`+
				`"following_count":355`+
				`,`+
				`"mute_expires_at":null`+
			`}`,
		},



		{
			Account: ent.Account{
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
				CreatedAt:      opt.Something("2023-09-27T22:06:19Z"),
				LastStatusAt:   nul.Null[string](),
				StatusesCount:  opt.Something(jsonint.Int64(123)),
				FollowersCount: opt.Something(jsonint.Int64(24789)),
				FollowingCount: opt.Something(jsonint.Int64(355)),
				//Roles
				MuteExpiresAt: nul.Something("2025-11-12T09:10:11Z"),
			},

			Expected: `{`+
				`"id":"47"`+
				`,`+
				`"username":"joeblow"`+
				`,`+
				`"acct":"joeblow@example.com"`+
				`,`+
				`"url":"https://example.com/@joeblow"`+
				`,`+
				`"uri":"https://example.com/users/joeblow"`+
				`,`+
				`"display_name":"Joe Blow :-)"`+
				`,`+
				`"note":"\u003cp\u003eHello world!\u003c/p\u003e"`+
				`,`+
				`"avatar":"https://files.example.com/avatar/joeblow.png"`+
				`,`+
				`"avatar_static":"https://files.example.com/avatar-static/joeblow.png"`+
				`,`+
				`"header":"https://files.example.com/header/joeblow.png"`+
				`,`+
				`"header_static":"https://files.example.com/header-static/joeblow.png"`+
				`,`+
				`"fields":[]`+
				`,`+
				`"emojis":[]`+
				`,`+
				`"locked":false`+
				`,`+
				`"bot":false`+
				`,`+
				`"group":false`+
				`,`+
				`"discoverable":null`+
				`,`+
				`"created_at":"2023-09-27T22:06:19Z"`+
				`,`+
				`"last_status_at":null`+
				`,`+
				`"statuses_count":123`+
				`,`+
				`"followers_count":24789`+
				`,`+
				`"following_count":355`+
				`,`+
				`"mute_expires_at":"2025-11-12T09:10:11Z"`+
			`}`,
		},
	}

	for testNumber, test := range tests {

		actualBytes, err := json.Marshal(test.Account)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("ACCOUNT: %#v", test.Account)
			t.Logf("EXPECTED:\n%s", test.Expected)
			continue
		}

		{
			actual := string(actualBytes)
			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
				t.Logf("EXPECTED:\n%s", expected)
				t.Logf("ACTUAL:\n%s", actual)
				t.Logf("ACCOUNT: %#v", test.Account)
				continue
			}
		}
	}
}
