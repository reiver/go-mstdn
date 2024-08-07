package ent_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-json"
	"github.com/reiver/go-mstdn/ent"
)

func TestField_MarshalJSON(t *testing.T) {

	tests := []struct{
		Value ent.Field
		Expected string
	}{
		{
			Value: ent.FieldNameValue("", ""),
			Expected:
				"{"                                   +"\n"+
				"\t"+  `"name":`        +`""`   +","  +"\n"+
				"\t"+  `"value":`       +`""`   +","  +"\n"+
				"\t"+  `"verified_at":` +`null` +""   +"\n"+
				"}"                                   +"\n",
		},



		{
			Value: ent.FieldNameValue("NAME", ""),
			Expected:
				"{"                                     +"\n"+
				"\t"+  `"name":`        +`"NAME"` +","  +"\n"+
				"\t"+  `"value":`       +`""`     +","  +"\n"+
				"\t"+  `"verified_at":` +`null`   +""   +"\n"+
				"}"                                     +"\n",
		},
		{
			Value: ent.FieldNameValue("", "VALUE"),
			Expected:
				"{"                                      +"\n"+
				"\t"+  `"name":`        +`""`      +","  +"\n"+
				"\t"+  `"value":`       +`"VALUE"` +","  +"\n"+
				"\t"+  `"verified_at":` +`null`    +""   +"\n"+
				"}"                                      +"\n",
		},
		{
			Value: ent.FieldNameValue("NAME", "VALUE"),
			Expected:
				"{"                                      +"\n"+
				"\t"+  `"name":`        +`"NAME"`  +","  +"\n"+
				"\t"+  `"value":`       +`"VALUE"` +","  +"\n"+
				"\t"+  `"verified_at":` +`null`    +""   +"\n"+
				"}"                                      +"\n",
		},



		{
			Value: ent.FieldNameValue("age", "47"),
			Expected:
				"{"                                    +"\n"+
				"\t"+  `"name":`        +`"age"` +","  +"\n"+
				"\t"+  `"value":`       +`"47"`  +","  +"\n"+
				"\t"+  `"verified_at":` +`null`  +""   +"\n"+
				"}"                                    +"\n",
		},



		{
			Value: ent.FieldVerifiedNameValue("2023-09-22T22:45:22Z", "NAME", ""),
			Expected:
				"{"                                                     +"\n"+
				"\t"+  `"name":`        +`"NAME"`                 +","  +"\n"+
				"\t"+  `"value":`       +`""`                     +","  +"\n"+
				"\t"+  `"verified_at":` +`"2023-09-22T22:45:22Z"` +""   +"\n"+
				"}"                                                     +"\n",
		},
		{
			Value: ent.FieldVerifiedNameValue("2023-09-22T22:45:22Z", "", "VALUE"),
			Expected:
				"{"                                                     +"\n"+
				"\t"+  `"name":`        +`""`                     +","  +"\n"+
				"\t"+  `"value":`       +`"VALUE"`                +","  +"\n"+
				"\t"+  `"verified_at":` +`"2023-09-22T22:45:22Z"` +""   +"\n"+
				"}"                                                     +"\n",
		},
		{
			Value: ent.FieldVerifiedNameValue("2023-09-22T22:45:22Z", "NAME", "VALUE"),
			Expected:
				"{"                                                     +"\n"+
				"\t"+  `"name":`        +`"NAME"`                 +","  +"\n"+
				"\t"+  `"value":`       +`"VALUE"`                +","  +"\n"+
				"\t"+  `"verified_at":` +`"2023-09-22T22:45:22Z"` +""   +"\n"+
				"}"                                                     +"\n",
		},



		{
			Value: ent.FieldVerifiedNameValue("2023-09-22T22:45:22Z", "age", "47"),
			Expected:
				"{"                                                    +"\n"+
				"\t"+  `"name":`        +`"age"`                  +"," +"\n"+
				"\t"+  `"value":`       +`"47"`                   +"," +"\n"+
				"\t"+  `"verified_at":` +`"2023-09-22T22:45:22Z"` +""  +"\n"+
				"}"                                                    +"\n",
		},
	}

	for testNumber, test := range tests {

		actualBytes, err := json.Marshal(test.Value)
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
