package ent_test

import (
	"testing"

	"github.com/reiver/go-json"
	"github.com/reiver/go-jsonpp"
	"github.com/reiver/go-jsonstr"
	"github.com/reiver/go-nul"
	"github.com/reiver/go-opt"

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
				TargetAccount: opt.Something(demoAccount1()),
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
				`"target_account":`+toJSON(demoAccount1())+
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
				TargetAccount: opt.Something(demoAccount1()),
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
				`"target_account":`+toJSON(demoAccount1())+
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
				TargetAccount: opt.Something(demoAccount1()),
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
				`"target_account":`+toJSON(demoAccount1())+
			`}`,
		},
	}

	for testNumber, test := range tests {

		actualBytes, err := json.Marshal(&test.Report)
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
