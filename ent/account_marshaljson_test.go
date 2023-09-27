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
