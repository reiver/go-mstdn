package ent_test

import (
	"testing"

	"encoding/json"

	"sourcecode.social/reiver/go-jsonint"
	"sourcecode.social/reiver/go-jsonpp"
	"sourcecode.social/reiver/go-jsonstr"
	"sourcecode.social/reiver/go-nul"
	"sourcecode.social/reiver/go-opt"

	"sourcecode.social/reiver/go-mstdn/ent"
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
					DisplayName:    opt.Something("Joe Blow :-)"),
					Locked:         opt.Something(false),
					Bot:            opt.Something(false),
					Discoverable:   opt.Something(true),
					CreatedAt:      opt.Something("2017-02-08T02:00:53.274Z"),
					Note:           opt.Something("<p>I came. I saw. I conquered!</p>"),
					URL:            opt.Something("https://example.com/@joeblow"),
					URI:            opt.Something("https://example.com/users/joeblow"),
					Avatar:         opt.Something("https://files.example.com/avatar/joeblow.png"),
					AvatarStatic:   opt.Something("https://files.example.com/avatar-static/joeblow.png"),
					Header:         opt.Something("https://files.example.com/header/joeblow.png"),
					HeaderStatic:   opt.Something("https://files.example.com/header-static/joeblow/png"),
					FollowersCount: opt.Something(jsonint.Int64(12345)),
					FollowingCount: opt.Something(jsonint.Int64(210)),
					StatusesCount:  opt.Something(jsonint.Int64(9876543210)),
					LastStatusAt:   opt.Something("2023-09-27T00:31:59Z"),
					NoIndex:        opt.Something(false),
					//Emojis       []CustomEmoji               `json:"emojis"`
					//Roles        []Role                      `json:"roles"`
					//Fields       []Field                     `json:"fields"`
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
					`"display_name":"Joe Blow :-)"`+
					`,`+
					`"locked":false`+
					`,`+
					`"bot":false`+
					`,`+
					`"discoverable":true`+
					`,`+
					`"created_at":"2017-02-08T02:00:53.274Z"`+
					`,`+
					`"note":"\u003cp\u003eI came. I saw. I conquered!\u003c/p\u003e"`+
					`,`+
					`"url":"https://example.com/@joeblow"`+
					`,`+
					`"uri":"https://example.com/users/joeblow"`+
					`,`+
					`"avatar":"https://files.example.com/avatar/joeblow.png"`+
					`,`+
					`"avatar_static":"https://files.example.com/avatar-static/joeblow.png"`+
					`,`+
					`"header":"https://files.example.com/header/joeblow.png"`+
					`,`+
					`"header_static":"https://files.example.com/header-static/joeblow/png"`+
					`,`+
					`"followers_count":12345`+
					`,`+
					`"following_count":210`+
					`,`+
					`"statuses_count":9876543210`+
					`,`+
					`"last_status_at":"2023-09-27T00:31:59Z"`+
					`,`+
					`"noindex":false`+
					`,`+
					`"emojis":null`+
					`,`+
					`"roles":null`+
					`,`+
					`"fields":null`+
				`}`+
			`}`,
		},



		{
			Report: ent.Report{
				ID:            opt.Something("0918273645"),
				ActionTaken:   opt.Something(true),
				ActionTakenAt: nul.Something("2024-02-27T00:31:59Z"),
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
					DisplayName:    opt.Something("Joe Blow :-)"),
					Locked:         opt.Something(false),
					Bot:            opt.Something(false),
					Discoverable:   opt.Something(true),
					CreatedAt:      opt.Something("2017-02-08T02:00:53.274Z"),
					Note:           opt.Something("<p>I came. I saw. I conquered!</p>"),
					URL:            opt.Something("https://example.com/@joeblow"),
					URI:            opt.Something("https://example.com/users/joeblow"),
					Avatar:         opt.Something("https://files.example.com/avatar/joeblow.png"),
					AvatarStatic:   opt.Something("https://files.example.com/avatar-static/joeblow.png"),
					Header:         opt.Something("https://files.example.com/header/joeblow.png"),
					HeaderStatic:   opt.Something("https://files.example.com/header-static/joeblow/png"),
					FollowersCount: opt.Something(jsonint.Int64(12345)),
					FollowingCount: opt.Something(jsonint.Int64(210)),
					StatusesCount:  opt.Something(jsonint.Int64(9876543210)),
					LastStatusAt:   opt.Something("2023-09-27T00:31:59Z"),
					NoIndex:        opt.Something(false),
					//Emojis       []CustomEmoji               `json:"emojis"`
					//Roles        []Role                      `json:"roles"`
					//Fields       []Field                     `json:"fields"`
				},
			},
			Expected :
			`{`+
				`"id":"0918273645"`+
				`,`+
				`"action_taken":true`+
				`,`+
				`"action_taken_at":"2024-02-27T00:31:59Z"`+
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
					`"display_name":"Joe Blow :-)"`+
					`,`+
					`"locked":false`+
					`,`+
					`"bot":false`+
					`,`+
					`"discoverable":true`+
					`,`+
					`"created_at":"2017-02-08T02:00:53.274Z"`+
					`,`+
					`"note":"\u003cp\u003eI came. I saw. I conquered!\u003c/p\u003e"`+
					`,`+
					`"url":"https://example.com/@joeblow"`+
					`,`+
					`"uri":"https://example.com/users/joeblow"`+
					`,`+
					`"avatar":"https://files.example.com/avatar/joeblow.png"`+
					`,`+
					`"avatar_static":"https://files.example.com/avatar-static/joeblow.png"`+
					`,`+
					`"header":"https://files.example.com/header/joeblow.png"`+
					`,`+
					`"header_static":"https://files.example.com/header-static/joeblow/png"`+
					`,`+
					`"followers_count":12345`+
					`,`+
					`"following_count":210`+
					`,`+
					`"statuses_count":9876543210`+
					`,`+
					`"last_status_at":"2023-09-27T00:31:59Z"`+
					`,`+
					`"noindex":false`+
					`,`+
					`"emojis":null`+
					`,`+
					`"roles":null`+
					`,`+
					`"fields":null`+
				`}`+
			`}`,
		},



		{
			Report: ent.Report{
				ID:            opt.Something("0918273645"),
				ActionTaken:   opt.Something(true),
				ActionTakenAt: nul.Something("2024-02-27T00:31:59Z"),
				Category:      opt.Something("spam"),
				Comment:       opt.Something("this f-er spamed me"),
				Forwarded:     opt.Something(false),
				CreatedAt:     opt.Something("2023-09-27T00:31:59Z"),
				StatusIDs:     nul.Something(jsonstr.CompileStrings("7823598731", "7916234158", "5218091242")),
				RuleIDs:       nul.Null[jsonstr.Strings](),
				TargetAccount: ent.Account{
					ID:             opt.Something("12345"),
					UserName:       opt.Something("joeblow"),
					Acct:           opt.Something("joeblow@example.com"),
					DisplayName:    opt.Something("Joe Blow :-)"),
					Locked:         opt.Something(false),
					Bot:            opt.Something(false),
					Discoverable:   opt.Something(true),
					CreatedAt:      opt.Something("2017-02-08T02:00:53.274Z"),
					Note:           opt.Something("<p>I came. I saw. I conquered!</p>"),
					URL:            opt.Something("https://example.com/@joeblow"),
					URI:            opt.Something("https://example.com/users/joeblow"),
					Avatar:         opt.Something("https://files.example.com/avatar/joeblow.png"),
					AvatarStatic:   opt.Something("https://files.example.com/avatar-static/joeblow.png"),
					Header:         opt.Something("https://files.example.com/header/joeblow.png"),
					HeaderStatic:   opt.Something("https://files.example.com/header-static/joeblow/png"),
					FollowersCount: opt.Something(jsonint.Int64(12345)),
					FollowingCount: opt.Something(jsonint.Int64(210)),
					StatusesCount:  opt.Something(jsonint.Int64(9876543210)),
					LastStatusAt:   opt.Something("2023-09-27T00:31:59Z"),
					NoIndex:        opt.Something(false),
					//Emojis       []CustomEmoji               `json:"emojis"`
					//Roles        []Role                      `json:"roles"`
					//Fields       []Field                     `json:"fields"`
				},
			},
			Expected :
			`{`+
				`"id":"0918273645"`+
				`,`+
				`"action_taken":true`+
				`,`+
				`"action_taken_at":"2024-02-27T00:31:59Z"`+
				`,`+
				`"category":"spam"`+
				`,`+
				`"comment":"this f-er spamed me"`+
				`,`+
				`"forwarded":false`+
				`,`+
				`"created_at":"2023-09-27T00:31:59Z"`+
				`,`+
				`"status_ids":["7823598731","7916234158","5218091242"]`+
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
					`"display_name":"Joe Blow :-)"`+
					`,`+
					`"locked":false`+
					`,`+
					`"bot":false`+
					`,`+
					`"discoverable":true`+
					`,`+
					`"created_at":"2017-02-08T02:00:53.274Z"`+
					`,`+
					`"note":"\u003cp\u003eI came. I saw. I conquered!\u003c/p\u003e"`+
					`,`+
					`"url":"https://example.com/@joeblow"`+
					`,`+
					`"uri":"https://example.com/users/joeblow"`+
					`,`+
					`"avatar":"https://files.example.com/avatar/joeblow.png"`+
					`,`+
					`"avatar_static":"https://files.example.com/avatar-static/joeblow.png"`+
					`,`+
					`"header":"https://files.example.com/header/joeblow.png"`+
					`,`+
					`"header_static":"https://files.example.com/header-static/joeblow/png"`+
					`,`+
					`"followers_count":12345`+
					`,`+
					`"following_count":210`+
					`,`+
					`"statuses_count":9876543210`+
					`,`+
					`"last_status_at":"2023-09-27T00:31:59Z"`+
					`,`+
					`"noindex":false`+
					`,`+
					`"emojis":null`+
					`,`+
					`"roles":null`+
					`,`+
					`"fields":null`+
				`}`+
			`}`,
		},



		{
			Report: ent.Report{
				ID:            opt.Something("0918273645"),
				ActionTaken:   opt.Something(true),
				ActionTakenAt: nul.Something("2024-02-27T00:31:59Z"),
				Category:      opt.Something("spam"),
				Comment:       opt.Something("this f-er spamed me"),
				Forwarded:     opt.Something(false),
				CreatedAt:     opt.Something("2023-09-27T00:31:59Z"),
				StatusIDs:     nul.Something(jsonstr.CompileStrings("7823598731", "7916234158", "5218091242")),
				RuleIDs:       nul.Something(jsonstr.CompileStrings("23","1","13","42")),
				TargetAccount: ent.Account{
					ID:             opt.Something("12345"),
					UserName:       opt.Something("joeblow"),
					Acct:           opt.Something("joeblow@example.com"),
					DisplayName:    opt.Something("Joe Blow :-)"),
					Locked:         opt.Something(false),
					Bot:            opt.Something(false),
					Discoverable:   opt.Something(true),
					CreatedAt:      opt.Something("2017-02-08T02:00:53.274Z"),
					Note:           opt.Something("<p>I came. I saw. I conquered!</p>"),
					URL:            opt.Something("https://example.com/@joeblow"),
					URI:            opt.Something("https://example.com/users/joeblow"),
					Avatar:         opt.Something("https://files.example.com/avatar/joeblow.png"),
					AvatarStatic:   opt.Something("https://files.example.com/avatar-static/joeblow.png"),
					Header:         opt.Something("https://files.example.com/header/joeblow.png"),
					HeaderStatic:   opt.Something("https://files.example.com/header-static/joeblow/png"),
					FollowersCount: opt.Something(jsonint.Int64(12345)),
					FollowingCount: opt.Something(jsonint.Int64(210)),
					StatusesCount:  opt.Something(jsonint.Int64(9876543210)),
					LastStatusAt:   opt.Something("2023-09-27T00:31:59Z"),
					NoIndex:        opt.Something(false),
					//Emojis       []CustomEmoji               `json:"emojis"`
					//Roles        []Role                      `json:"roles"`
					//Fields       []Field                     `json:"fields"`
				},
			},
			Expected :
			`{`+
				`"id":"0918273645"`+
				`,`+
				`"action_taken":true`+
				`,`+
				`"action_taken_at":"2024-02-27T00:31:59Z"`+
				`,`+
				`"category":"spam"`+
				`,`+
				`"comment":"this f-er spamed me"`+
				`,`+
				`"forwarded":false`+
				`,`+
				`"created_at":"2023-09-27T00:31:59Z"`+
				`,`+
				`"status_ids":["7823598731","7916234158","5218091242"]`+
				`,`+
				`"rule_ids":["23","1","13","42"]`+
				`,`+
				`"target_account":{`+
					`"id":"12345"`+
					`,`+
					`"username":"joeblow"`+
					`,`+
					`"acct":"joeblow@example.com"`+
					`,`+
					`"display_name":"Joe Blow :-)"`+
					`,`+
					`"locked":false`+
					`,`+
					`"bot":false`+
					`,`+
					`"discoverable":true`+
					`,`+
					`"created_at":"2017-02-08T02:00:53.274Z"`+
					`,`+
					`"note":"\u003cp\u003eI came. I saw. I conquered!\u003c/p\u003e"`+
					`,`+
					`"url":"https://example.com/@joeblow"`+
					`,`+
					`"uri":"https://example.com/users/joeblow"`+
					`,`+
					`"avatar":"https://files.example.com/avatar/joeblow.png"`+
					`,`+
					`"avatar_static":"https://files.example.com/avatar-static/joeblow.png"`+
					`,`+
					`"header":"https://files.example.com/header/joeblow.png"`+
					`,`+
					`"header_static":"https://files.example.com/header-static/joeblow/png"`+
					`,`+
					`"followers_count":12345`+
					`,`+
					`"following_count":210`+
					`,`+
					`"statuses_count":9876543210`+
					`,`+
					`"last_status_at":"2023-09-27T00:31:59Z"`+
					`,`+
					`"noindex":false`+
					`,`+
					`"emojis":null`+
					`,`+
					`"roles":null`+
					`,`+
					`"fields":null`+
				`}`+
			`}`,
		},



		{
			Report: ent.Report{
				ID:            opt.Something("0918273645"),
				ActionTaken:   opt.Something(true),
				ActionTakenAt: nul.Something("2024-02-27T00:31:59Z"),
				Category:      opt.Something("spam"),
				Comment:       opt.Something("this f-er spamed me"),
				Forwarded:     opt.Something(false),
				CreatedAt:     opt.Something("2023-09-27T00:31:59Z"),
				StatusIDs:     nul.Something(jsonstr.CompileStrings("7823598731", "7916234158", "5218091242")),
				RuleIDs:       nul.Something(jsonstr.CompileStrings("23","1","13","42")),
				TargetAccount: ent.Account{
					ID:             opt.Something("12345"),
					UserName:       opt.Something("joeblow"),
					Acct:           opt.Something("joeblow@example.com"),
					DisplayName:    opt.Something("Joe Blow :-) :batman:"),
					Locked:         opt.Something(false),
					Bot:            opt.Something(false),
					Discoverable:   opt.Something(true),
					CreatedAt:      opt.Something("2017-02-08T02:00:53.274Z"),
					Note:           opt.Something("<p>I came. I saw. I conquered!</p>"),
					URL:            opt.Something("https://example.com/@joeblow"),
					URI:            opt.Something("https://example.com/users/joeblow"),
					Avatar:         opt.Something("https://files.example.com/avatar/joeblow.png"),
					AvatarStatic:   opt.Something("https://files.example.com/avatar-static/joeblow.png"),
					Header:         opt.Something("https://files.example.com/header/joeblow.png"),
					HeaderStatic:   opt.Something("https://files.example.com/header-static/joeblow/png"),
					FollowersCount: opt.Something(jsonint.Int64(12345)),
					FollowingCount: opt.Something(jsonint.Int64(210)),
					StatusesCount:  opt.Something(jsonint.Int64(9876543210)),
					LastStatusAt:   opt.Something("2023-09-27T00:31:59Z"),
					NoIndex:        opt.Something(true),
					Emojis:       []ent.CustomEmoji{
						ent.CustomEmoji{
							ShortCode:       opt.Something("batman"),
							URL:             opt.Something("https://files.example.com/emoji/batman.png"),
							StaticURL:       opt.Something("https://files.example.com/emoji-static/batman.png"),
							VisibleInPicker: opt.Something(true),
						},
						ent.CustomEmoji{
							ShortCode:       opt.Something("spiderman"),
							URL:             opt.Something("https://files.example.com/emoji/spiderman.png"),
							StaticURL:       opt.Something("https://files.example.com/emoji-static/spiderman.png"),
							VisibleInPicker: opt.Something(true),
							Category:        opt.Something("spiders"),
						},
					},
					//Roles        []Role                      `json:"roles"`
					//Fields       []Field                     `json:"fields"`
				},
			},
			Expected :
			`{`+
				`"id":"0918273645"`+
				`,`+
				`"action_taken":true`+
				`,`+
				`"action_taken_at":"2024-02-27T00:31:59Z"`+
				`,`+
				`"category":"spam"`+
				`,`+
				`"comment":"this f-er spamed me"`+
				`,`+
				`"forwarded":false`+
				`,`+
				`"created_at":"2023-09-27T00:31:59Z"`+
				`,`+
				`"status_ids":["7823598731","7916234158","5218091242"]`+
				`,`+
				`"rule_ids":["23","1","13","42"]`+
				`,`+
				`"target_account":{`+
					`"id":"12345"`+
					`,`+
					`"username":"joeblow"`+
					`,`+
					`"acct":"joeblow@example.com"`+
					`,`+
					`"display_name":"Joe Blow :-) :batman:"`+
					`,`+
					`"locked":false`+
					`,`+
					`"bot":false`+
					`,`+
					`"discoverable":true`+
					`,`+
					`"created_at":"2017-02-08T02:00:53.274Z"`+
					`,`+
					`"note":"\u003cp\u003eI came. I saw. I conquered!\u003c/p\u003e"`+
					`,`+
					`"url":"https://example.com/@joeblow"`+
					`,`+
					`"uri":"https://example.com/users/joeblow"`+
					`,`+
					`"avatar":"https://files.example.com/avatar/joeblow.png"`+
					`,`+
					`"avatar_static":"https://files.example.com/avatar-static/joeblow.png"`+
					`,`+
					`"header":"https://files.example.com/header/joeblow.png"`+
					`,`+
					`"header_static":"https://files.example.com/header-static/joeblow/png"`+
					`,`+
					`"followers_count":12345`+
					`,`+
					`"following_count":210`+
					`,`+
					`"statuses_count":9876543210`+
					`,`+
					`"last_status_at":"2023-09-27T00:31:59Z"`+
					`,`+
					`"noindex":true`+
					`,`+
					`"emojis":[`+
						`{`+
							`"shortcode":"batman"`+
							`,`+
							`"static_url":"https://files.example.com/emoji-static/batman.png"`+
							`,`+
							`"url":"https://files.example.com/emoji/batman.png"`+
							`,`+
							`"visible_in_picker":true`+
						`}`+
						`,`+
						`{`+
							`"category":"spiders"`+
							`,`+
							`"shortcode":"spiderman"`+
							`,`+
							`"static_url":"https://files.example.com/emoji-static/spiderman.png"`+
							`,`+
							`"url":"https://files.example.com/emoji/spiderman.png"`+
							`,`+
							`"visible_in_picker":true`+
						`}`+
					`]`+
					`,`+
					`"roles":null`+
					`,`+
					`"fields":null`+
				`}`+
			`}`,
		},



		{
			Report: ent.Report{
				ID:            opt.Something("0918273645"),
				ActionTaken:   opt.Something(true),
				ActionTakenAt: nul.Something("2024-02-27T00:31:59Z"),
				Category:      opt.Something("spam"),
				Comment:       opt.Something("this f-er spamed me"),
				Forwarded:     opt.Something(false),
				CreatedAt:     opt.Something("2023-09-27T00:31:59Z"),
				StatusIDs:     nul.Something(jsonstr.CompileStrings("7823598731", "7916234158", "5218091242")),
				RuleIDs:       nul.Something(jsonstr.CompileStrings("23","1","13","42")),
				TargetAccount: ent.Account{
					ID:             opt.Something("12345"),
					UserName:       opt.Something("joeblow"),
					Acct:           opt.Something("joeblow@example.com"),
					DisplayName:    opt.Something("Joe Blow :-) :batman:"),
					Locked:         opt.Something(false),
					Bot:            opt.Something(false),
					Discoverable:   opt.Something(true),
					CreatedAt:      opt.Something("2017-02-08T02:00:53.274Z"),
					Note:           opt.Something("<p>I came. I saw. I conquered!</p>"),
					URL:            opt.Something("https://example.com/@joeblow"),
					URI:            opt.Something("https://example.com/users/joeblow"),
					Avatar:         opt.Something("https://files.example.com/avatar/joeblow.png"),
					AvatarStatic:   opt.Something("https://files.example.com/avatar-static/joeblow.png"),
					Header:         opt.Something("https://files.example.com/header/joeblow.png"),
					HeaderStatic:   opt.Something("https://files.example.com/header-static/joeblow/png"),
					FollowersCount: opt.Something(jsonint.Int64(12345)),
					FollowingCount: opt.Something(jsonint.Int64(210)),
					StatusesCount:  opt.Something(jsonint.Int64(9876543210)),
					LastStatusAt:   opt.Something("2023-09-27T00:31:59Z"),
					NoIndex:        opt.Something(true),
					Emojis:       []ent.CustomEmoji{
						ent.CustomEmoji{
							ShortCode:       opt.Something("batman"),
							URL:             opt.Something("https://files.example.com/emoji/batman.png"),
							StaticURL:       opt.Something("https://files.example.com/emoji-static/batman.png"),
							VisibleInPicker: opt.Something(true),
						},
						ent.CustomEmoji{
							ShortCode:       opt.Something("spiderman"),
							URL:             opt.Something("https://files.example.com/emoji/spiderman.png"),
							StaticURL:       opt.Something("https://files.example.com/emoji-static/spiderman.png"),
							VisibleInPicker: opt.Something(true),
							Category:        opt.Something("spiders"),
						},
						ent.CustomEmoji{
							ShortCode:       opt.Something("busooka"),
							URL:             opt.Something("https://files.example.com/emoji/dads-name-for-his-sons-toy.png"),
							StaticURL:       opt.Something("https://files.example.com/emoji-static/dads-name-for-his-sons-toy.png"),
							VisibleInPicker: opt.Something(false),
							Category:        opt.Something("not-transformers"),
						},
					},
					Roles:        []ent.Role{
						ent.Role{
							ID:          opt.Something(jsonint.Int64(17)),
							Name :       opt.Something("editor"),
							Permissions: opt.Something(jsonint.Int64(8)),
							Highlighted: opt.Something(false),
						},
						ent.Role{
							ID:          opt.Something(jsonint.Int64(23)),
							Name:        opt.Something("super-fan"),
							Color:       opt.Something("#dd2211"),
							Permissions: opt.Something(jsonint.Int64(255)),
							Highlighted: opt.Something(true),
						},
					},
					//Fields       []Field                     `json:"fields"`
				},
			},
			Expected :
			`{`+
				`"id":"0918273645"`+
				`,`+
				`"action_taken":true`+
				`,`+
				`"action_taken_at":"2024-02-27T00:31:59Z"`+
				`,`+
				`"category":"spam"`+
				`,`+
				`"comment":"this f-er spamed me"`+
				`,`+
				`"forwarded":false`+
				`,`+
				`"created_at":"2023-09-27T00:31:59Z"`+
				`,`+
				`"status_ids":["7823598731","7916234158","5218091242"]`+
				`,`+
				`"rule_ids":["23","1","13","42"]`+
				`,`+
				`"target_account":{`+
					`"id":"12345"`+
					`,`+
					`"username":"joeblow"`+
					`,`+
					`"acct":"joeblow@example.com"`+
					`,`+
					`"display_name":"Joe Blow :-) :batman:"`+
					`,`+
					`"locked":false`+
					`,`+
					`"bot":false`+
					`,`+
					`"discoverable":true`+
					`,`+
					`"created_at":"2017-02-08T02:00:53.274Z"`+
					`,`+
					`"note":"\u003cp\u003eI came. I saw. I conquered!\u003c/p\u003e"`+
					`,`+
					`"url":"https://example.com/@joeblow"`+
					`,`+
					`"uri":"https://example.com/users/joeblow"`+
					`,`+
					`"avatar":"https://files.example.com/avatar/joeblow.png"`+
					`,`+
					`"avatar_static":"https://files.example.com/avatar-static/joeblow.png"`+
					`,`+
					`"header":"https://files.example.com/header/joeblow.png"`+
					`,`+
					`"header_static":"https://files.example.com/header-static/joeblow/png"`+
					`,`+
					`"followers_count":12345`+
					`,`+
					`"following_count":210`+
					`,`+
					`"statuses_count":9876543210`+
					`,`+
					`"last_status_at":"2023-09-27T00:31:59Z"`+
					`,`+
					`"noindex":true`+
					`,`+
					`"emojis":[`+
						`{`+
							`"shortcode":"batman"`+
							`,`+
							`"static_url":"https://files.example.com/emoji-static/batman.png"`+
							`,`+
							`"url":"https://files.example.com/emoji/batman.png"`+
							`,`+
							`"visible_in_picker":true`+
						`}`+
						`,`+
						`{`+
							`"category":"spiders"`+
							`,`+
							`"shortcode":"spiderman"`+
							`,`+
							`"static_url":"https://files.example.com/emoji-static/spiderman.png"`+
							`,`+
							`"url":"https://files.example.com/emoji/spiderman.png"`+
							`,`+
							`"visible_in_picker":true`+
						`}`+
						`,`+
						`{`+
							`"category":"not-transformers"`+
							`,`+
							`"shortcode":"busooka"`+
							`,`+
							`"static_url":"https://files.example.com/emoji-static/dads-name-for-his-sons-toy.png"`+
							`,`+
							`"url":"https://files.example.com/emoji/dads-name-for-his-sons-toy.png"`+
							`,`+
							`"visible_in_picker":false`+
						`}`+
					`]`+
					`,`+
					`"roles":[`+
						`{`+
							`"id":17`+
							`,`+
							`"name":"editor"`+
							`,`+
							`"color":""`+
							`,`+
							`"permissions":8`+
							`,`+
							`"highlighted":false`+
						`}`+
						`,`+
						`{`+
							`"id":23`+
							`,`+
							`"name":"super-fan"`+
							`,`+
							`"color":"#dd2211"`+
							`,`+
							`"permissions":255`+
							`,`+
							`"highlighted":true`+
						`}`+
					`]`+
					`,`+
					`"fields":null`+
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
