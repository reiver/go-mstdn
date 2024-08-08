package ent_test

import (
	"testing"

	"github.com/reiver/go-json"
	"github.com/reiver/go-jsonpp"
	"github.com/reiver/go-nul"
	"github.com/reiver/go-opt"

	"github.com/reiver/go-mstdn/ent"
)

func TestApplication_MarshalJSON(t *testing.T) {

	tests := []struct{
		Application ent.Application
		Expected string
	}{
		{
			Application: ent.Application{},
			Expected: "{}",
		},



		{
			Application: ent.Application{
				Name:     opt.Something("acme app"),
				WebSite: nul.Null[string](),
				VapidKey: opt.Something("BHgNMADAUjgYgM4PZtHkY3yTQRYD-ibS_qrWYg2KPBRidocowKcOc-8YpyItumamkGph2bk8FuryT4-p3Eymwz8"),
			},
			Expected:
				`{`+
					`"name":"acme app"`+
					`,`+
					`"website":null`+
					`,`+
					`"vapid_key":"BHgNMADAUjgYgM4PZtHkY3yTQRYD-ibS_qrWYg2KPBRidocowKcOc-8YpyItumamkGph2bk8FuryT4-p3Eymwz8"`+
				`}`,
		},



		{
			Application: ent.Application{
				Name:     opt.Something("client fish"),
				WebSite:  nul.Something("http://example.fish/"),
				VapidKey: opt.Something("BLV6IwZiUgNnReINKtfgpt-zNCUF8jXTIsvA7Pa1-TTTLOEkeG-UtWVhKraRGgAcGUnrMBBzQPPFxTEao7L_Oz"),
			},
			Expected:
				`{`+
					`"name":"client fish"`+
					`,`+
					`"website":"http://example.fish/"`+
					`,`+
					`"vapid_key":"BLV6IwZiUgNnReINKtfgpt-zNCUF8jXTIsvA7Pa1-TTTLOEkeG-UtWVhKraRGgAcGUnrMBBzQPPFxTEao7L_Oz"`+
				`}`,
		},



		{
			Application: ent.Application{
				Name:     opt.Something("frontodon"),
				WebSite:  nul.Something("http://something.tld/"),
				VapidKey: opt.Something("BLCtzyFc2xtnAqWYW3m4M0v47Uym9mKd5yQMvY3FLWzdXN2vpjzF3iQ413fKwNOITkhJ6_tlvTZELL876uokpM4"),
				ClientID: opt.Something("22"),
			},
			Expected:
				`{`+
					`"name":"frontodon"`+
					`,`+
					`"website":"http://something.tld/"`+
					`,`+
					`"vapid_key":"BLCtzyFc2xtnAqWYW3m4M0v47Uym9mKd5yQMvY3FLWzdXN2vpjzF3iQ413fKwNOITkhJ6_tlvTZELL876uokpM4"`+
					`,`+
					`"client_id":"22"`+
				`}`,
		},



		{
			Application: ent.Application{
				Name:         opt.Something("super-gorilla"),
				WebSite:      nul.Something("http://example.gorilla/"),
				VapidKey:     opt.Something("BNi9UYyWxKVULR-FEaCWt3NuAHrlyyIz7zYRyIyLP0Q46ePHsiLbd8wHG3VXy-wTAgzFsRb5pGxJZDeX3FROwlE"),
				ClientID:     opt.Something("22"),
				ClientSecret: opt.Something("NwlOvca6TQjFutgobDM6voupU8kIqzN0h_oa1pEqBD4"),
			},
			Expected:
				`{`+
					`"name":"super-gorilla"`+
					`,`+
					`"website":"http://example.gorilla/"`+
					`,`+
					`"vapid_key":"BNi9UYyWxKVULR-FEaCWt3NuAHrlyyIz7zYRyIyLP0Q46ePHsiLbd8wHG3VXy-wTAgzFsRb5pGxJZDeX3FROwlE"`+
					`,`+
					`"client_id":"22"`+
					`,`+
					`"client_secret":"NwlOvca6TQjFutgobDM6voupU8kIqzN0h_oa1pEqBD4"`+
				`}`,
		},
	}

	for testNumber, test := range tests {

		actualBytes, err := json.Marshal(test.Application)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("EXPECTED:\n%s", test.Expected)
			t.Logf("EXPECTED:\n%s", jsonpp.SPrettyPrint([]byte(test.Expected)))
			t.Logf("APPLICATION: %#v", test.Application)
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
				t.Logf("APPLICATION: %#v", test.Application)
				continue
			}
		}
	}
}
