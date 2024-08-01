package admn_test

import (
	"testing"

	"encoding/json"

	"github.com/reiver/go-opt"

	"github.com/reiver/go-mstdn/ent/admn"
)

func TestDimensionData_MarshalJSON(t *testing.T) {

	tests := []struct{
		DimensionData admn.DimensionData
		Expected string
	}{
		{
			DimensionData: admn.DimensionData{
				Key:      opt.Something("fa"),
				HumanKey: opt.Something("Persian"),
				Value:    opt.Something("7"),
			},
			Expected: `{`+
				`"key":"fa"`+
				`,`+
				`"human_key":"Persian"`+
				`,`+
				`"value":"7"`+
			`}`,
		},



		{
			DimensionData: admn.DimensionData{
				Key:        opt.Something("drive"),
				HumanKey:   opt.Something("Drive Space"),
				Value:      opt.Something("2097152"),
				Unit:       opt.Something("bytes"),
			},
			Expected: `{`+
				`"key":"drive"`+
				`,`+
				`"human_key":"Drive Space"`+
				`,`+
				`"value":"2097152"`+
				`,`+
				`"unit":"bytes"`+
			`}`,
		},
		{
			DimensionData: admn.DimensionData{
				Key:        opt.Something("drive"),
				HumanKey:   opt.Something("Drive Space"),
				Value:      opt.Something("2097152"),
				Unit:       opt.Something("bytes"),
				HumanValue: opt.Something("2,097,152 bytes"),
			},
			Expected: `{`+
				`"key":"drive"`+
				`,`+
				`"human_key":"Drive Space"`+
				`,`+
				`"value":"2097152"`+
				`,`+
				`"unit":"bytes"`+
				`,`+
				`"human_value":"2,097,152 bytes"`+
			`}`,
		},



		{
			DimensionData: admn.DimensionData{
				Key:        opt.Something("primes-founds"),
				HumanKey:   opt.Something("Number of Prime-Numbers Found"),
				Value:      opt.Something("8723389322"),
				HumanValue: opt.Something("8,723,389,322 Prime-Numbers"),
			},
			Expected: `{`+
				`"key":"primes-founds"`+
				`,`+
				`"human_key":"Number of Prime-Numbers Found"`+
				`,`+
				`"value":"8723389322"`+
				`,`+
				`"human_value":"8,723,389,322 Prime-Numbers"`+
			`}`,
		},
	}

	for testNumber, test := range tests {

		actualBytes, err := json.Marshal(test.DimensionData)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("DIMENSION-DATA: %#v", test.DimensionData)
			continue
		}

		{
			actual := string(actualBytes)
			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
				t.Logf("EXPECTED:\n%s", expected)
				t.Logf("ACTUAL:\n%s", actual)
			t.Logf("DIMENSION-DATA: %#v", test.DimensionData)
				continue
			}
		}
	}
}
