package mstdn

import (
	"testing"

	"strings"
)

func TestErrorJSON(t *testing.T) {

	tests := []struct{
		StatusCode int
		Message []interface{}
		Expected string
	}{
		{
			StatusCode: 400,
			Message: []interface{}{},
			Expected: `{"error":"Bad Request"}`,
		},
		{
			StatusCode: 401,
			Message: []interface{}{},
			Expected: `{"error":"Unauthorized"}`,
		},
		{
			StatusCode: 403,
			Message: []interface{}{},
			Expected: `{"error":"Forbidden"}`,
		},
		{
			StatusCode: 404,
			Message: []interface{}{},
			Expected: `{"error":"Not Found"}`,
		},
		{
			StatusCode: 422,
			Message: []interface{}{},
			Expected: `{"error":"Unprocessable Entity"}`,
		},
		{
			StatusCode: 429,
			Message: []interface{}{},
			Expected: `{"error":"Too Many Requests"}`,
		},
		{
			StatusCode: 500,
			Message: []interface{}{},
			Expected: `{"error":"Internal Server Error"}`,
		},
		{
			StatusCode: 503,
			Message: []interface{}{},
			Expected: `{"error":"Service Unavailable"}`,
		},



		{
			StatusCode: 400,
			Message: []interface{}{"very bad request"},
			Expected: `{"error":"very bad request"}`,
		},
		{
			StatusCode: 401,
			Message: []interface{}{"get out of here"},
			Expected: `{"error":"get out of here"}`,
		},
		{
			StatusCode: 403,
			Message: []interface{}{"no!!!!"},
			Expected: `{"error":"no!!!!"}`,
		},
		{
			StatusCode: 404,
			Message: []interface{}{"huh?"},
			Expected: `{"error":"huh?"}`,
		},
		{
			StatusCode: 422,
			Message: []interface{}{"WTF?"},
			Expected: `{"error":"WTF?"}`,
		},
		{
			StatusCode: 429,
			Message: []interface{}{"f-off"},
			Expected: `{"error":"f-off"}`,
		},
		{
			StatusCode: 500,
			Message: []interface{}{"crap!"},
			Expected: `{"error":"crap!"}`,
		},
		{
			StatusCode: 503,
			Message: []interface{}{"nope"},
			Expected: `{"error":"nope"}`,
		},



		{
			StatusCode: 400,
			Message: []interface{}{"very bad request", "you have been naughty"},
			Expected: `{"error":"very bad request","error_description":"you have been naughty"}`,
		},
		{
			StatusCode: 401,
			Message: []interface{}{"get out of here", "access denied!!! — you can keep on knocking but you can't come in"},
			Expected: `{"error":"get out of here","error_description":"access denied!!! — you can keep on knocking but you can't come in"}`,
		},
		{
			StatusCode: 403,
			Message: []interface{}{"no!!!!", "go away"},
			Expected: `{"error":"no!!!!","error_description":"go away"}`,
		},
		{
			StatusCode: 404,
			Message: []interface{}{"huh?", "what are you talking about"},
			Expected: `{"error":"huh?","error_description":"what are you talking about"}`,
		},
		{
			StatusCode: 422,
			Message: []interface{}{"WTF?", "i don't understand the words coming out of your mouth"},
			Expected: `{"error":"WTF?","error_description":"i don't understand the words coming out of your mouth"}`,
		},
		{
			StatusCode: 429,
			Message: []interface{}{"f-off", "slow it down, back it up"},
			Expected: `{"error":"f-off","error_description":"slow it down, back it up"}`,
		},
		{
			StatusCode: 500,
			Message: []interface{}{"crap!", "i f-ed up"},
			Expected: `{"error":"crap!","error_description":"i f-ed up"}`,
		},
		{
			StatusCode: 503,
			Message: []interface{}{"nope", "it not workie"},
			Expected: `{"error":"nope","error_description":"it not workie"}`,
		},



		{
			StatusCode: 400,
			Message: []interface{}{"very bad request", "you have been naughty", " :-)"},
			Expected: `{"error":"very bad request","error_description":"you have been naughty :-)"}`,
		},
		{
			StatusCode: 401,
			Message: []interface{}{"get out of here", "access denied!!! — you can keep on knocking but you can't come in", "."},
			Expected: `{"error":"get out of here","error_description":"access denied!!! — you can keep on knocking but you can't come in."}`,
		},
		{
			StatusCode: 403,
			Message: []interface{}{"no!!!!", "go away", "!"},
			Expected: `{"error":"no!!!!","error_description":"go away!"}`,
		},
		{
			StatusCode: 404,
			Message: []interface{}{"huh?", "what are you talking about", "?"},
			Expected: `{"error":"huh?","error_description":"what are you talking about?"}`,
		},
		{
			StatusCode: 422,
			Message: []interface{}{"WTF?", "i don't understand the words coming out of your mouth", "."},
			Expected: `{"error":"WTF?","error_description":"i don't understand the words coming out of your mouth."}`,
		},
		{
			StatusCode: 429,
			Message: []interface{}{"f-off", "slow it down, back it up", "."},
			Expected: `{"error":"f-off","error_description":"slow it down, back it up."}`,
		},
		{
			StatusCode: 500,
			Message: []interface{}{"crap!", "i f-ed up", " :-("},
			Expected: `{"error":"crap!","error_description":"i f-ed up :-("}`,
		},
		{
			StatusCode: 503,
			Message: []interface{}{"nope", "it not workie", "."},
			Expected: `{"error":"nope","error_description":"it not workie."}`,
		},
	}

	for testNumber, test := range tests {

		var actualBuffer strings.Builder

		errorJSON(&actualBuffer, test.StatusCode, test.Message...)

		actual := actualBuffer.String()
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED:\n%s", expected)
			t.Logf("ACTUAL:\n%s", actual)
			continue
		}
	}
}
