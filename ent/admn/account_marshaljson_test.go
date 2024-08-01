package admn_test

import (
	"testing"

	"encoding/json"

	"sourcecode.social/reiver/go-jsonint"
	"sourcecode.social/reiver/go-nul"
	"github.com/reiver/go-opt"

	"github.com/reiver/go-mstdn/ent"
	"github.com/reiver/go-mstdn/ent/admn"
)

func TestAdmin_MarshalJSON(t *testing.T) {

	tests := []struct{
		Account admn.Account
		Expected string
	}{
		{
			Account: admn.Account{
				ID:            opt.Something("973249706"),
				UserName:      opt.Something("joeblow"),
				Domain:        nul.Null[string](),
				CreatedAt:     opt.Something("2022-09-08T23:03:26.762Z"),
				EMail:         opt.Something("joeblow1924@canmail.tld"),
				IP:            nul.Null[string](),
				//IPs
				Locale:        opt.Something("fa"),
				InviteRequest: nul.Null[string](),
				Role:          ent.Role{
					ID:          opt.Something(jsonint.Int64(3)),
					Name:        opt.Something("admin"),
					Color:       opt.Something("#ff2400"),
					Permissions: opt.Something(jsonint.Int64(4294967295)),
					Highlighted: opt.Something(true),
				},
				Confirmed:     opt.Something(false),
				Approved:      opt.Something(false),
				Disabled:      opt.Something(false),
				Silenced:      opt.Something(false),
				Suspended:     opt.Something(false),
				Account:       ent.Account{
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
			},
			Expected: `{`+
				`"id":"973249706"`+
				`,`+
				`"username":"joeblow"`+
				`,`+
				`"domain":null`+
				`,`+
				`"created_at":"2022-09-08T23:03:26.762Z"`+
				`,`+
				`"email":"joeblow1924@canmail.tld"`+
				`,`+
				`"ip":null`+
				`,`+
				`"ips":[]`+
				`,`+
				`"locale":"fa"`+
				`,`+
				`"invite_request":null`+
				`,`+
				`"role":{`+
					`"id":3`+
					`,`+
					`"name":"admin"`+
					`,`+
					`"color":"#ff2400"`+
					`,`+
					`"permissions":4294967295`+
					`,`+
					`"highlighted":true`+
				`}`+
				`,`+
				`"confirmed":false`+
				`,`+
				`"approved":false`+
				`,`+
				`"disabled":false`+
				`,`+
				`"silenced":false`+
				`,`+
				`"suspended":false`+
				`,`+
				`"account":{`+
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
				`}`+
			`}`,
		},









		{
			Account: admn.Account{
				ID:            opt.Something("973249706"),
				UserName:      opt.Something("joeblow"),
				Domain:        nul.Null[string](),
				CreatedAt:     opt.Something("2022-09-08T23:03:26.762Z"),
				EMail:         opt.Something("joeblow1924@canmail.tld"),
				IP:            nul.Null[string](),
				IPs:           nil,
				Locale:        opt.Something("fa"),
				InviteRequest: nul.Null[string](),
				Role:          ent.Role{
					ID:          opt.Something(jsonint.Int64(3)),
					Name:        opt.Something("admin"),
					Color:       opt.Something("#ff2400"),
					Permissions: opt.Something(jsonint.Int64(4294967295)),
					Highlighted: opt.Something(true),
				},
				Confirmed:     opt.Something(false),
				Approved:      opt.Something(false),
				Disabled:      opt.Something(false),
				Silenced:      opt.Something(false),
				Suspended:     opt.Something(false),
				Account:       ent.Account{
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
			},
			Expected: `{`+
				`"id":"973249706"`+
				`,`+
				`"username":"joeblow"`+
				`,`+
				`"domain":null`+
				`,`+
				`"created_at":"2022-09-08T23:03:26.762Z"`+
				`,`+
				`"email":"joeblow1924@canmail.tld"`+
				`,`+
				`"ip":null`+
				`,`+
				`"ips":[]`+
				`,`+
				`"locale":"fa"`+
				`,`+
				`"invite_request":null`+
				`,`+
				`"role":{`+
					`"id":3`+
					`,`+
					`"name":"admin"`+
					`,`+
					`"color":"#ff2400"`+
					`,`+
					`"permissions":4294967295`+
					`,`+
					`"highlighted":true`+
				`}`+
				`,`+
				`"confirmed":false`+
				`,`+
				`"approved":false`+
				`,`+
				`"disabled":false`+
				`,`+
				`"silenced":false`+
				`,`+
				`"suspended":false`+
				`,`+
				`"account":{`+
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
				`}`+
			`}`,
		},



		{
			Account: admn.Account{
				ID:            opt.Something("973249706"),
				UserName:      opt.Something("joeblow"),
				Domain:        nul.Null[string](),
				CreatedAt:     opt.Something("2022-09-08T23:03:26.762Z"),
				EMail:         opt.Something("joeblow1924@canmail.tld"),
				IP:            nul.Null[string](),
				IPs:         []admn.IP{},
				Locale:        opt.Something("fa"),
				InviteRequest: nul.Null[string](),
				Role:          ent.Role{
					ID:          opt.Something(jsonint.Int64(3)),
					Name:        opt.Something("admin"),
					Color:       opt.Something("#ff2400"),
					Permissions: opt.Something(jsonint.Int64(4294967295)),
					Highlighted: opt.Something(true),
				},
				Confirmed:     opt.Something(false),
				Approved:      opt.Something(false),
				Disabled:      opt.Something(false),
				Silenced:      opt.Something(false),
				Suspended:     opt.Something(false),
				Account:       ent.Account{
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
			},
			Expected: `{`+
				`"id":"973249706"`+
				`,`+
				`"username":"joeblow"`+
				`,`+
				`"domain":null`+
				`,`+
				`"created_at":"2022-09-08T23:03:26.762Z"`+
				`,`+
				`"email":"joeblow1924@canmail.tld"`+
				`,`+
				`"ip":null`+
				`,`+
				`"ips":[]`+
				`,`+
				`"locale":"fa"`+
				`,`+
				`"invite_request":null`+
				`,`+
				`"role":{`+
					`"id":3`+
					`,`+
					`"name":"admin"`+
					`,`+
					`"color":"#ff2400"`+
					`,`+
					`"permissions":4294967295`+
					`,`+
					`"highlighted":true`+
				`}`+
				`,`+
				`"confirmed":false`+
				`,`+
				`"approved":false`+
				`,`+
				`"disabled":false`+
				`,`+
				`"silenced":false`+
				`,`+
				`"suspended":false`+
				`,`+
				`"account":{`+
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
				`}`+
			`}`,
		},



		{
			Account: admn.Account{
				ID:            opt.Something("973249706"),
				UserName:      opt.Something("joeblow"),
				Domain:        nul.Null[string](),
				CreatedAt:     opt.Something("2022-09-08T23:03:26.762Z"),
				EMail:         opt.Something("joeblow1924@canmail.tld"),
				IP:            nul.Null[string](),
				IPs:         []admn.IP{
					admn.IP{
						IP:     opt.Something("142.58.103.17"),
						UsedAt: opt.Something("2023-09-28T07:10:19Z"),
					},
				},
				Locale:        opt.Something("fa"),
				InviteRequest: nul.Null[string](),
				Role:          ent.Role{
					ID:          opt.Something(jsonint.Int64(3)),
					Name:        opt.Something("admin"),
					Color:       opt.Something("#ff2400"),
					Permissions: opt.Something(jsonint.Int64(4294967295)),
					Highlighted: opt.Something(true),
				},
				Confirmed:     opt.Something(false),
				Approved:      opt.Something(false),
				Disabled:      opt.Something(false),
				Silenced:      opt.Something(false),
				Suspended:     opt.Something(false),
				Account:       ent.Account{
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
			},
			Expected: `{`+
				`"id":"973249706"`+
				`,`+
				`"username":"joeblow"`+
				`,`+
				`"domain":null`+
				`,`+
				`"created_at":"2022-09-08T23:03:26.762Z"`+
				`,`+
				`"email":"joeblow1924@canmail.tld"`+
				`,`+
				`"ip":null`+
				`,`+
				`"ips":[`+
					`{`+
						`"ip":"142.58.103.17"`+
						`,`+
						`"used_at":"2023-09-28T07:10:19Z"`+
					`}`+
				`]`+
				`,`+
				`"locale":"fa"`+
				`,`+
				`"invite_request":null`+
				`,`+
				`"role":{`+
					`"id":3`+
					`,`+
					`"name":"admin"`+
					`,`+
					`"color":"#ff2400"`+
					`,`+
					`"permissions":4294967295`+
					`,`+
					`"highlighted":true`+
				`}`+
				`,`+
				`"confirmed":false`+
				`,`+
				`"approved":false`+
				`,`+
				`"disabled":false`+
				`,`+
				`"silenced":false`+
				`,`+
				`"suspended":false`+
				`,`+
				`"account":{`+
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
				`}`+
			`}`,
		},
	}

	for testNumber, test := range tests {

		actualBytes, err := json.Marshal(test.Account)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one." , testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("ACCOUNT:\n%#v", test.Account)
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
				t.Logf("ACCOUNT:\n%#v", test.Account)
				continue
			}
		}
	}
}
