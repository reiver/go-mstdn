package ent

import (
	"testing"

	"bytes"
	"encoding/json"

	"github.com/reiver/go-opt"
)

func TestCustomEmoji_MarshalJSON(t *testing.T) {

	tests := []struct{
		Value CustomEmoji
		Expected string
	}{
		{
			Value: CustomEmoji{
				ShortCode:       opt.Something[string](""),
				URL:             opt.Something[string](""),
				StaticURL:       opt.Something[string](""),
				VisibleInPicker: opt.Something[bool](false),
			},
			Expected:
				"{"                                            +"\n"+
				"\t"+  `"shortcode":`         +`""`      +","  +"\n"+
				"\t"+  `"url":`               +`""`      +","  +"\n"+
				"\t"+  `"static_url":`        +`""`      +","  +"\n"+
				"\t"+  `"visible_in_picker":` +`false`   +""   +"\n"+
				"}"                                            +"\n",
		},



		{
			Value: CustomEmoji{
				ShortCode:       opt.Something[string]("bananas"),
				URL:             opt.Something[string](""),
				StaticURL:       opt.Something[string](""),
				VisibleInPicker: opt.Something[bool](false),
			},
			Expected:
				"{"                                                   +"\n"+
				"\t"+  `"shortcode":`         +`"bananas"`      +","  +"\n"+
				"\t"+  `"url":`               +`""`             +","  +"\n"+
				"\t"+  `"static_url":`        +`""`             +","  +"\n"+
				"\t"+  `"visible_in_picker":` +`false`          +""   +"\n"+
				"}"                                                   +"\n",
		},
		{
			Value: CustomEmoji{
				ShortCode:       opt.Something[string](""),
				URL:             opt.Something[string]("https://example.com/emoji/cracra"),
				StaticURL:       opt.Something[string](""),
				VisibleInPicker: opt.Something[bool](false),
			},
			Expected:
				"{"                                                                            +"\n"+
				"\t"+  `"shortcode":`         +`""`                                      +","  +"\n"+
				"\t"+  `"url":`               +`"https://example.com/emoji/cracra"`      +","  +"\n"+
				"\t"+  `"static_url":`        +`""`                                      +","  +"\n"+
				"\t"+  `"visible_in_picker":` +`false`                                   +""   +"\n"+
				"}"                                                                            +"\n",
		},
		{
			Value: CustomEmoji{
				ShortCode:       opt.Something[string](""),
				URL:             opt.Something[string](""),
				StaticURL:       opt.Something[string]("https://static.example.com/img/emoji/cracra.png"),
				VisibleInPicker: opt.Something[bool](false),
			},
			Expected:
				"{"                                                                                           +"\n"+
				"\t"+  `"shortcode":`         +`""`                                                     +","  +"\n"+
				"\t"+  `"url":`               +`""`                                                     +","  +"\n"+
				"\t"+  `"static_url":`        +`"https://static.example.com/img/emoji/cracra.png"`      +","  +"\n"+
				"\t"+  `"visible_in_picker":` +`false`                                                  +""   +"\n"+
				"}"                                                                                     +"\n",
		},
		{
			Value: CustomEmoji{
				ShortCode:       opt.Something[string](""),
				URL:             opt.Something[string](""),
				StaticURL:       opt.Something[string](""),
				VisibleInPicker: opt.Something[bool](true),
			},
			Expected:
				"{"                                           +"\n"+
				"\t"+  `"shortcode":`         +`""`     +","  +"\n"+
				"\t"+  `"url":`               +`""`     +","  +"\n"+
				"\t"+  `"static_url":`        +`""`     +","  +"\n"+
				"\t"+  `"visible_in_picker":` +`true`   +""   +"\n"+
				"}"                                           +"\n",
		},
	}

	for testNumber, test := range tests {

		actualBytes, err := test.Value.MarshalJSON()
		if nil != err {
			t.Errorf("For test #%d, did not expect to get an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: %#v", test.Value)
			continue
		}

		actual := string(actualBytes)

		var expected string
		{
			var buffer bytes.Buffer

			err := json.Compact(&buffer, []byte(test.Expected))
			if nil != err {
				panic(err)
			}

			expected = buffer.String()
		}

		if expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED:\n%s", expected)
			t.Logf("ACTUAL:\n%s", actual)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("VALUE: %#v", test.Value)
			continue
		}

	}
}
