package mstdn_test

import (
	"testing"

	"encoding/json"

	"sourcecode.social/reiver/go-mstdn"
)

func TestField_UnmarshalJSON(t *testing.T) {

	tests := []struct{
		JSON string
		Expected mstdn.Field
	}{
		{
			JSON: `{"name":"Location","value":"Metro Vancouver"}`,
			Expected: mstdn.FieldNameValue("Location", "Metro Vancouver"),
		},
		{
			JSON: `{"name":"Location","value":"Metro Vancouver", "verified_at":null}`,
			Expected: mstdn.FieldNameValue("Location", "Metro Vancouver"),
		},
		{
			JSON: `{"name":"Location","value":"Metro Vancouver", "verified_at":"2023-09-22T22:45:22Z"}`,
			Expected: mstdn.FieldVerifiedNameValue("2023-09-22T22:45:22Z", "Location", "Metro Vancouver"),
		},
	}

	for testNumber, test := range tests {

		var actual mstdn.Field

		err := json.Unmarshal([]byte(test.JSON), &actual)
		if nil != err {
			t.Errorf("For test #%d,did not expect an error but actually got one." , testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("JSON: %q", test.JSON)
			t.Logf("EXPECTED %#v", test.Expected)
			continue
		}

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			t.Logf("JSON: %q", test.JSON)
			continue
		}
	}
}
