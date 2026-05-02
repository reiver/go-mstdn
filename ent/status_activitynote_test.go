package ent_test

import (
	"testing"

	"codeberg.org/reiver/go-activitypub"
	"github.com/reiver/go-nul"
	"github.com/reiver/go-opt"

	"github.com/reiver/go-mstdn/ent"
)

func TestStatus_ActivityNote(t *testing.T) {

	tests := []struct{
		Status ent.Status
		Expected activitypub.Note
	}{
		{
			Status: ent.Status{},
			Expected: activitypub.Note{},
		},



		{
			Status: ent.Status{
				Account: ent.Account{
					ID: opt.Something("https://example.com/users/joeblow"),
				},
			},
			Expected: func() activitypub.Note {
				var note activitypub.Note
				note.AttributedTo = append(note.AttributedTo, activitypub.ObjectID("https://example.com/users/joeblow"))
				return note
			}(),
		},



		{
			Status: ent.Status{
				Content: opt.Something("Hello world!"),
			},
			Expected: func() activitypub.Note {
				var note activitypub.Note
				note.Content = opt.Something("Hello world!")
				return note
			}(),
		},



		{
			Status: ent.Status{
				CreatedAt: nul.Something("2023-09-27T22:06:19Z"),
			},
			Expected: func() activitypub.Note {
				var note activitypub.Note
				note.Published = opt.Something("2023-09-27T22:06:19Z")
				return note
			}(),
		},



		{
			Status: ent.Status{
				URI: nul.Something("https://example.com/status/123"),
			},
			Expected: func() activitypub.Note {
				var note activitypub.Note
				note.ID = opt.Something("https://example.com/status/123")
				note.URL = opt.Something("https://example.com/status/123")
				return note
			}(),
		},
		{
			Status: ent.Status{
				URL: nul.Something("https://example.com/@joeblow/status/123"),
			},
			Expected: func() activitypub.Note {
				var note activitypub.Note
				note.ID = opt.Something("https://example.com/@joeblow/status/123")
				note.URL = opt.Something("https://example.com/@joeblow/status/123")
				return note
			}(),
		},
		{
			Status: ent.Status{
				URI: nul.Something("https://example.com/status/123"),
				URL: nul.Something("https://example.com/@joeblow/status/123"),
			},
			Expected: func() activitypub.Note {
				var note activitypub.Note
				note.ID = opt.Something("https://example.com/status/123")
				note.URL = opt.Something("https://example.com/status/123")
				return note
			}(),
		},



		{
			Status: ent.Status{
				SpoilerText: nul.Something("the rabbit did it"),
			},
			Expected: func() activitypub.Note {
				var note activitypub.Note
				note.Summary = nul.Something("the rabbit did it")
				return note
			}(),
		},
		{
			Status: ent.Status{
				SpoilerText: nul.Something(""),
			},
			Expected: func() activitypub.Note {
				var note activitypub.Note
				note.Summary = nul.Something("")
				return note
			}(),
		},



		{
			Status: ent.Status{
				Tags: []ent.Tag{
					{
						Name: opt.Something("akkoma"),
						URL:  opt.Something("https://example.com/tags/akkoma"),
					},
					{
						Name: opt.Something("fediverse"),
						URL:  opt.Something("https://example.com/tags/fediverse"),
					},
				},
			},
			Expected: func() activitypub.Note {
				var note activitypub.Note

				var hashtag1 activitypub.HashTag
				hashtag1.Name = opt.Something("akkoma")
				hashtag1.HRef = opt.Something("https://example.com/tags/akkoma")

				var hashtag2 activitypub.HashTag
				hashtag2.Name = opt.Something("fediverse")
				hashtag2.HRef = opt.Something("https://example.com/tags/fediverse")

				note.Tags = append(note.Tags, hashtag1, hashtag2)
				return note
			}(),
		},



		{
			Status: ent.Status{
				Account: ent.Account{
					ID: opt.Something("https://example.com/users/joeblow"),
				},
				Content:     opt.Something("<p>Hello world!</p>"),
				CreatedAt:   nul.Something("2023-09-27T22:06:19Z"),
				URI:         nul.Something("https://example.com/status/123"),
				URL:         nul.Something("https://example.com/@joeblow/status/123"),
				SpoilerText: nul.Something("the rabbit did it"),
				Tags: []ent.Tag{
					{
						Name: opt.Something("fediverse"),
						URL:  opt.Something("https://example.com/tags/fediverse"),
					},
				},
			},
			Expected: func() activitypub.Note {
				var note activitypub.Note
				note.AttributedTo = append(note.AttributedTo, activitypub.ObjectID("https://example.com/users/joeblow"))
				note.Content = opt.Something("<p>Hello world!</p>")
				note.Published = opt.Something("2023-09-27T22:06:19Z")
				note.ID = opt.Something("https://example.com/status/123")
				note.URL = opt.Something("https://example.com/status/123")
				note.Summary = nul.Something("the rabbit did it")

				var hashtag activitypub.HashTag
				hashtag.Name = opt.Something("fediverse")
				hashtag.HRef = opt.Something("https://example.com/tags/fediverse")
				note.Tags = append(note.Tags, hashtag)

				return note
			}(),
		},
	}

	for testNumber, test := range tests {

		var actual activitypub.Note

		err := test.Status.ActivityNote(&actual)

		if nil != err {
			t.Errorf("For test #%d, did not expect to get an error but actually did.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			continue
		}

		expected := test.Expected

		expectedJSON := toJSON(expected)
		actualJSON   := toJSON(actual)

		if expectedJSON != actualJSON {
			t.Errorf("For test #%d, the actual note is not what was expected.", testNumber)
			t.Logf("EXPECTED: %s", expectedJSON)
			t.Logf("ACTUAL:   %s", actualJSON)
			t.Logf("STATUS: %#v", test.Status)
			continue
		}
	}
}

func TestStatus_ActivityNote_nilNote(t *testing.T) {

	var status ent.Status

	err := status.ActivityNote(nil)

	if nil == err {
		t.Errorf("Expected an error when passing nil note but did not get one.")
		return
	}
}
