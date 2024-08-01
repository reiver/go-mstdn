package ent_test

import (
	"testing"

	"encoding/json"

	"sourcecode.social/reiver/go-jsonint"
	"github.com/reiver/go-opt"

	"github.com/reiver/go-mstdn/ent"
)

func TestRole_MarshalJSON(t *testing.T) {

	tests := []struct{
		JSON string
		Expected ent.Role
	}{
		{
			JSON:
				`{`+
					`"id":12345,`+
					`"name":"Master",`+
					`"color":"#1a2b3c",`+
					`"permissions":256,`+
					`"highlighted":true`+
				`}`,
			Expected: ent.Role{
				ID:          opt.Something(jsonint.Int64(12345)),
				Name:        opt.Something("Master"),
				Color:       opt.Something("#1a2b3c"),
				Permissions: opt.Something(jsonint.Int64(256)),
				Highlighted: opt.Something(true),
			},
		},
		{
			JSON:
				`{`+
					`"id":71727374,`+
					`"name":"Student",`+
					`"color":"#76cd54",`+
					`"permissions":664,`+
					`"highlighted":false`+
				`}`,
			Expected: ent.Role{
				ID:          opt.Something(jsonint.Int64(71727374)),
				Name:        opt.Something("Student"),
				Color:       opt.Something("#76cd54"),
				Permissions: opt.Something(jsonint.Int64(664)),
				Highlighted: opt.Something(false),
			},
		},



		{
			JSON:
				`{`+
					`"id":87,`+
					`"name":"QA Specialist",`+
					`"color":"",`+
					`"permissions":218,`+
					`"highlighted":true`+
				`}`,
			Expected: ent.Role{
				ID:          opt.Something(jsonint.Int64(87)),
				Name:        opt.Something("QA Specialist"),
				Permissions: opt.Something(jsonint.Int64(218)),
				Highlighted: opt.Something(true),
			},
		},
		{
			JSON:
				`{`+
					`"id":87,`+
					`"name":"QA Specialist",`+
					`"permissions":218,`+
					`"highlighted":true`+
				`}`,
			Expected: ent.Role{
				ID:          opt.Something(jsonint.Int64(87)),
				Name:        opt.Something("QA Specialist"),
				Permissions: opt.Something(jsonint.Int64(218)),
				Highlighted: opt.Something(true),
			},
		},
	}

	for testNumber, test := range tests {

		var actual ent.Role

		err := json.Unmarshal([]byte(test.JSON), &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("JSON:\n%s", test.JSON)
			t.Logf("EXPECTED:\n%#v", test.Expected)
			continue
		}

		{
			var expected ent.Role = test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
				t.Logf("EXPECTED:\n%#v", expected)
				t.Logf("ACTUAL:\n%#v", actual)
				t.Logf("JSON:\n%s", test.JSON)
				continue
			}
		}
	}
}
