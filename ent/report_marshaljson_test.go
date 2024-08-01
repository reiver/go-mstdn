package ent_test

import (
	"testing"

	"encoding/json"

	"sourcecode.social/reiver/go-jsonint"
	"sourcecode.social/reiver/go-jsonpp"
	"sourcecode.social/reiver/go-jsonstr"
	"sourcecode.social/reiver/go-nul"
	"sourcecode.social/reiver/go-opt"

	"github.com/reiver/go-mstdn/ent"
)

func TestReport_MarshalJSON(t *testing.T) {

	tests := []struct{
		Report ent.Report
		Expected string
	}{
		{
			Report: ent.Report{
				ID:            opt.Something("0918273645"),
				ActionTaken:   opt.Something(false),
				ActionTakenAt: nul.Null[string](),
				Category:      opt.Something("spam"),
				Comment:       opt.Something("this f-er spamed me"),
				Forwarded:     opt.Something(false),
				CreatedAt:     opt.Something("2023-09-27T00:31:59Z"),
				StatusIDs:     nul.Null[jsonstr.Strings](),
				RuleIDs:       nul.Null[jsonstr.Strings](),
				TargetAccount: ent.Account{
					ID:             opt.Something("12345"),
					UserName:       opt.Something("joeblow"),
					Acct:           opt.Something("joeblow@example.com"),
					URL:            opt.Something("https://example.com/@joeblow"),
					URI:            opt.Something("https://example.com/users/joeblow"),
					DisplayName:    opt.Something("Joe Blow :-)"),
					Note:           opt.Something("<p>I came. I saw. I conquered!</p>"),
					Avatar:         opt.Something("https://files.example.com/avatar/joeblow.png"),
					AvatarStatic:   opt.Something("https://files.example.com/avatar-static/joeblow.png"),
					Header:         opt.Something("https://files.example.com/header/joeblow.png"),
					HeaderStatic:   opt.Something("https://files.example.com/header-static/joeblow/png"),
					//Fields       []Field                     `json:"fields"`
					//Emojis       []CustomEmoji               `json:"emojis"`
					Locked:         opt.Something(false),
					Bot:            opt.Something(false),
					Group:          opt.Something(false),
					Discoverable:   nul.Something(true),
					CreatedAt:      opt.Something("2017-02-08T02:00:53.274Z"),
					LastStatusAt:   nul.Something("2023-09-27T00:31:59Z"),
					StatusesCount:  opt.Something(jsonint.Int64(9876543210)),
					FollowersCount: opt.Something(jsonint.Int64(12345)),
					FollowingCount: opt.Something(jsonint.Int64(210)),
					NoIndex:        nul.Nothing[bool](),
					//Roles        []Role                      `json:"roles"`
				},
			},
			Expected :
			`{`+
				`"id":"0918273645"`+
				`,`+
				`"action_taken":false`+
				`,`+
				`"action_taken_at":null`+
				`,`+
				`"category":"spam"`+
				`,`+
				`"comment":"this f-er spamed me"`+
				`,`+
				`"forwarded":false`+
				`,`+
				`"created_at":"2023-09-27T00:31:59Z"`+
				`,`+
				`"status_ids":null`+
				`,`+
				`"rule_ids":null`+
				`,`+
				`"target_account":{`+
					`"id":"12345"`+
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
					`"note":"\u003cp\u003eI came. I saw. I conquered!\u003c/p\u003e"`+
					`,`+
					`"avatar":"https://files.example.com/avatar/joeblow.png"`+
					`,`+
					`"avatar_static":"https://files.example.com/avatar-static/joeblow.png"`+
					`,`+
					`"header":"https://files.example.com/header/joeblow.png"`+
					`,`+
					`"header_static":"https://files.example.com/header-static/joeblow/png"`+
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
					`"created_at":"2017-02-08T02:00:53.274Z"`+
					`,`+
					`"last_status_at":"2023-09-27T00:31:59Z"`+
					`,`+
					`"statuses_count":9876543210`+
					`,`+
					`"followers_count":12345`+
					`,`+
					`"following_count":210`+
				`}`+
			`}`,
		},









		{
			Report: ent.Report{
				ID:            opt.Something("0918273645"),
				ActionTaken:   opt.Something(true),
				ActionTakenAt: nul.Null[string](),
				Category:      opt.Something("spam"),
				Comment:       opt.Something("this f-er spamed me"),
				Forwarded:     opt.Something(false),
				CreatedAt:     opt.Something("2023-09-27T00:31:59Z"),
				StatusIDs:     nul.Null[jsonstr.Strings](),
				RuleIDs:       nul.Null[jsonstr.Strings](),
				TargetAccount: ent.Account{
					ID:             opt.Something("12345"),
					UserName:       opt.Something("joeblow"),
					Acct:           opt.Something("joeblow@example.com"),
					URL:            opt.Something("https://example.com/@joeblow"),
					URI:            opt.Something("https://example.com/users/joeblow"),
					DisplayName:    opt.Something("Joe Blow :-)"),
					Note:           opt.Something("<p>I came. I saw. I conquered!</p>"),
					Avatar:         opt.Something("https://files.example.com/avatar/joeblow.png"),
					AvatarStatic:   opt.Something("https://files.example.com/avatar-static/joeblow.png"),
					Header:         opt.Something("https://files.example.com/header/joeblow.png"),
					HeaderStatic:   opt.Something("https://files.example.com/header-static/joeblow/png"),
					//Fields       []Field                     `json:"fields"`
					//Emojis       []CustomEmoji               `json:"emojis"`
					Locked:         opt.Something(false),
					Bot:            opt.Something(false),
					Group:          opt.Something(false),
					Discoverable:   nul.Something(true),
					CreatedAt:      opt.Something("2017-02-08T02:00:53.274Z"),
					LastStatusAt:   nul.Something("2023-09-27T00:31:59Z"),
					StatusesCount:  opt.Something(jsonint.Int64(9876543210)),
					FollowersCount: opt.Something(jsonint.Int64(12345)),
					FollowingCount: opt.Something(jsonint.Int64(210)),
					NoIndex:        nul.Nothing[bool](),
					//Roles        []Role                      `json:"roles"`
				},
			},
			Expected :
			`{`+
				`"id":"0918273645"`+
				`,`+
				`"action_taken":true`+
				`,`+
				`"action_taken_at":null`+
				`,`+
				`"category":"spam"`+
				`,`+
				`"comment":"this f-er spamed me"`+
				`,`+
				`"forwarded":false`+
				`,`+
				`"created_at":"2023-09-27T00:31:59Z"`+
				`,`+
				`"status_ids":null`+
				`,`+
				`"rule_ids":null`+
				`,`+
				`"target_account":{`+
					`"id":"12345"`+
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
					`"note":"\u003cp\u003eI came. I saw. I conquered!\u003c/p\u003e"`+
					`,`+
					`"avatar":"https://files.example.com/avatar/joeblow.png"`+
					`,`+
					`"avatar_static":"https://files.example.com/avatar-static/joeblow.png"`+
					`,`+
					`"header":"https://files.example.com/header/joeblow.png"`+
					`,`+
					`"header_static":"https://files.example.com/header-static/joeblow/png"`+
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
					`"created_at":"2017-02-08T02:00:53.274Z"`+
					`,`+
					`"last_status_at":"2023-09-27T00:31:59Z"`+
					`,`+
					`"statuses_count":9876543210`+
					`,`+
					`"followers_count":12345`+
					`,`+
					`"following_count":210`+
				`}`+
			`}`,
		},









		{
			Report: ent.Report{
				ID:            opt.Something("0918273645"),
				ActionTaken:   opt.Something(false),
				ActionTakenAt: nul.Something("2023-09-28T07:10:19Z"),
				Category:      opt.Something("spam"),
				Comment:       opt.Something("this f-er spamed me"),
				Forwarded:     opt.Something(false),
				CreatedAt:     opt.Something("2023-09-27T00:31:59Z"),
				StatusIDs:     nul.Null[jsonstr.Strings](),
				RuleIDs:       nul.Null[jsonstr.Strings](),
				TargetAccount: ent.Account{
					ID:             opt.Something("12345"),
					UserName:       opt.Something("joeblow"),
					Acct:           opt.Something("joeblow@example.com"),
					URL:            opt.Something("https://example.com/@joeblow"),
					URI:            opt.Something("https://example.com/users/joeblow"),
					DisplayName:    opt.Something("Joe Blow :-)"),
					Note:           opt.Something("<p>I came. I saw. I conquered!</p>"),
					Avatar:         opt.Something("https://files.example.com/avatar/joeblow.png"),
					AvatarStatic:   opt.Something("https://files.example.com/avatar-static/joeblow.png"),
					Header:         opt.Something("https://files.example.com/header/joeblow.png"),
					HeaderStatic:   opt.Something("https://files.example.com/header-static/joeblow/png"),
					//Fields       []Field                     `json:"fields"`
					//Emojis       []CustomEmoji               `json:"emojis"`
					Locked:         opt.Something(false),
					Bot:            opt.Something(false),
					Group:          opt.Something(false),
					Discoverable:   nul.Something(true),
					CreatedAt:      opt.Something("2017-02-08T02:00:53.274Z"),
					LastStatusAt:   nul.Something("2023-09-27T00:31:59Z"),
					StatusesCount:  opt.Something(jsonint.Int64(9876543210)),
					FollowersCount: opt.Something(jsonint.Int64(12345)),
					FollowingCount: opt.Something(jsonint.Int64(210)),
					NoIndex:        nul.Nothing[bool](),
					//Roles        []Role                      `json:"roles"`
				},
			},
			Expected :
			`{`+
				`"id":"0918273645"`+
				`,`+
				`"action_taken":false`+
				`,`+
				`"action_taken_at":"2023-09-28T07:10:19Z"`+
				`,`+
				`"category":"spam"`+
				`,`+
				`"comment":"this f-er spamed me"`+
				`,`+
				`"forwarded":false`+
				`,`+
				`"created_at":"2023-09-27T00:31:59Z"`+
				`,`+
				`"status_ids":null`+
				`,`+
				`"rule_ids":null`+
				`,`+
				`"target_account":{`+
					`"id":"12345"`+
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
					`"note":"\u003cp\u003eI came. I saw. I conquered!\u003c/p\u003e"`+
					`,`+
					`"avatar":"https://files.example.com/avatar/joeblow.png"`+
					`,`+
					`"avatar_static":"https://files.example.com/avatar-static/joeblow.png"`+
					`,`+
					`"header":"https://files.example.com/header/joeblow.png"`+
					`,`+
					`"header_static":"https://files.example.com/header-static/joeblow/png"`+
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
					`"created_at":"2017-02-08T02:00:53.274Z"`+
					`,`+
					`"last_status_at":"2023-09-27T00:31:59Z"`+
					`,`+
					`"statuses_count":9876543210`+
					`,`+
					`"followers_count":12345`+
					`,`+
					`"following_count":210`+
				`}`+
			`}`,
		},
	}

	for testNumber, test := range tests {

		actualBytes, err := json.Marshal(test.Report)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("REPORT: %#v", test.Report)
			t.Logf("EXPECTED:\n%s", test.Expected)
			t.Logf("EXPECTED:\n%s", jsonpp.SPrettyPrint([]byte(test.Expected)))
			continue
		}

		{
			actual := string(actualBytes)
			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
				t.Logf("EXPECTED:\n%s", expected)
				t.Logf("ACTUAL:\n%s", actual)
				t.Logf("EXPECTED:\n%s", jsonpp.SPrettyPrint([]byte(expected)))
				t.Logf("ACTUAL:\n%s", jsonpp.SPrettyPrint([]byte(actual)))
				t.Logf("REPORT: %#v", test.Report)
				continue
			}
		}
	}
}
