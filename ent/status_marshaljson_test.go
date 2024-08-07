package ent

import (
	"testing"

	"encoding/json"

	"github.com/reiver/go-nul"
	"github.com/reiver/go-opt"
)

func TestStatus_JSON(t *testing.T) {

	tests := []struct{
			Status Status
			Expected string
	}{
		{
			Status:Status{
				ID: opt.Something("123"),
			},
			Expected: `{"id":"123"}`,
		},



		{
			Status:Status{
				ID: opt.Something("123"),
				URI: nul.Something("https://example.com/status/123"),
			},
			Expected: `{"id":"123","uri":"https://example.com/status/123"}`,
		},
		{
			Status:Status{
				ID: opt.Something("123"),
				URL: nul.Something("https://example.com/@joeblow/status/123"),
			},
			Expected: `{"id":"123","url":"https://example.com/@joeblow/status/123"}`,
		},
		{
			Status:Status{
				ID: opt.Something("123"),
				CreatedAt: opt.Something("yesterday"),
			},
			Expected: `{"id":"123","created_at":"yesterday"}`,
		},
//@TODO: Account
		{
			Status:Status{
				ID: opt.Something("123"),
				Content: opt.Something("Hello world!"),
			},
			Expected: `{"id":"123","content":"Hello world!"}`,
		},
		{
			Status:Status{
				ID: opt.Something("123"),
				Visibility: opt.Something("public"),
			},
			Expected: `{"id":"123","visibility":"public"}`,
		},
		{
			Status:Status{
				ID: opt.Something("123"),
				Sensitive: opt.Something(true),
			},
			Expected: `{"id":"123","sensitive":true}`,
		},
		{
			Status:Status{
				ID: opt.Something("123"),
				SpoilerText: opt.Something("the rabbit did it"),
			},
			Expected: `{"id":"123","spoiler_text":"the rabbit did it"}`,
		},
//@TODO: MediaAttachments
//@TODO: Application
//@TODO: Mentions

		{
			Status:Status{
				ID: opt.Something("123"),
				Tags: []Tag{
					Tag{
						Name:opt.Something("akkoma"),
						URL:opt.Something("https://example.com/tags/akkoma"),
					},
					Tag{
						Name:opt.Something("fediverse"),
						URL:opt.Something("https://example.com/tags/fediverse"),
					},
				},
			},
			Expected: `{"id":"123","tags":[{"name":"akkoma","url":"https://example.com/tags/akkoma"},{"name":"fediverse","url":"https://example.com/tags/fediverse"}]}`,
		},

//@TODO: Emojis









		{
			Status:Status{
				ID: opt.Something("123"),
				Language: nul.Something("en"),
			},
			Expected: `{"id":"123","language":"en"}`,
		},
		{
			Status:Status{
				ID: opt.Something("123"),
				Text: nul.Something("bing bong bang"),
			},
			Expected: `{"id":"123","text":"bing bong bang"}`,
		},
		{
			Status:Status{
				ID: opt.Something("123"),
				EditedAt: nul.Something("today"),
			},
			Expected: `{"id":"123","edited_at":"today"}`,
		},
		{
			Status:Status{
				ID: opt.Something("123"),
				Favourited: opt.Something(true),
			},
			Expected: `{"id":"123","favourited":true}`,
		},
		{
			Status:Status{
				ID: opt.Something("123"),
				Reblogged: opt.Something(true),
			},
			Expected: `{"id":"123","reblogged":true}`,
		},
		{
			Status:Status{
				ID: opt.Something("123"),
				Muted: opt.Something(true),
			},
			Expected: `{"id":"123","muted":true}`,
		},
		{
			Status:Status{
				ID: opt.Something("123"),
				Bookmarked: opt.Something(true),
			},
			Expected: `{"id":"123","bookmarked":true}`,
		},
		{
			Status:Status{
				ID: opt.Something("123"),
				Pinned: opt.Something(true),
			},
			Expected: `{"id":"123","pinned":true}`,
		},
		{
			Status:Status{
				ID: opt.Something("123"),
				Filtered: json.RawMessage([]byte(`{"once":1,"twice":2,"thrice":3,"fource":4}`)),
			},
			Expected: `{"id":"123","filtered":{"once":1,"twice":2,"thrice":3,"fource":4}}`,
		},
	}

	for testNumber, test := range tests {

		actualBytes, err := json.Marshal(&test.Status)

		if nil != err {
			t.Errorf("For test #%d, did not expect to get an error but actually did.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			continue
		}

		actual := string(actualBytes)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual string-serialized form is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("EXPECTED:\n%s", expected)
			t.Logf("ACTUAL:\n%s", actual)
			t.Logf("STATUS:\n%#v", test.Status)
			continue
		}
	}
}
