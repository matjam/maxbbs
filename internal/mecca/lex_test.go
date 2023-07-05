package mecca

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

func TestLexer(t *testing.T) {
	type testCase struct {
		line     int
		input    string
		expected []lexToken
		err      error
	}

	cases := []testCase{
		{
			getTestCaseLine(),
			`[white]Leave a message to [sysop_name] [[Y,n]? [gray ansreq menu]yn|`,
			[]lexToken{
				{Type: L_TOKEN, Value: "white"},
				{Type: L_STRING, Value: "Leave a message to "},
				{Type: L_TOKEN, Value: "sysop_name"},
				{Type: L_STRING, Value: " [Y,n]? "},
				{Type: L_TOKEN, Value: "gray"},
				{Type: L_TOKEN, Value: "ansreq"},
				{Type: L_TOKEN, Value: "menu"},
				{Type: L_STRING, Value: "yn|"},
			},
			nil,
		},
		{
			getTestCaseLine(),
			`Some text [white]Leave a message to [sysop_name] [[Y,n]? [gray ansreq menu]yn|`,
			[]lexToken{
				{Type: L_STRING, Value: "Some text "},
				{Type: L_TOKEN, Value: "white"},
				{Type: L_STRING, Value: "Leave a message to "},
				{Type: L_TOKEN, Value: "sysop_name"},
				{Type: L_STRING, Value: " [Y,n]? "},
				{Type: L_TOKEN, Value: "gray"},
				{Type: L_TOKEN, Value: "ansreq"},
				{Type: L_TOKEN, Value: "menu"},
				{Type: L_STRING, Value: "yn|"},
			},
			nil,
		},
		{
			getTestCaseLine(),
			`text [   token] text [token    ] text [   token   ] text`,
			[]lexToken{
				{Type: L_STRING, Value: "text "},
				{Type: L_TOKEN, Value: "token"},
				{Type: L_STRING, Value: " text "},
				{Type: L_TOKEN, Value: "token"},
				{Type: L_STRING, Value: " text "},
				{Type: L_TOKEN, Value: "token"},
				{Type: L_STRING, Value: " text"},
			},
			nil,
		},
		{
			getTestCaseLine(),
			`text [   token`,
			[]lexToken{},
			ErrUnexpectedEndOfInput,
		},
		{
			getTestCaseLine(),
			`text [   token [ invalid ] text`,
			[]lexToken{},
			ErrTokenStartInsideTokenBlock,
		},
	}

	for _, c := range cases {
		r := strings.NewReader(c.input)
		got, err := lex(r)

		if c.err != nil {
			if err == nil || !errors.Is(err, c.err) {
				t.Errorf("\n\n    line: %v\nexpected error: %v\n     got error: %v\n\n", c.line, c.err, err)
			}
		} else {
			if err != nil {
				t.Errorf("\n\n    line: %v\nexpected error: %v\n     got error: %v\n\n", c.line, c.err, err)
			}

			if !reflect.DeepEqual(got, c.expected) {
				t.Errorf("\n\n    line: %v\nexpected: %v\n     got: %v\n\n", c.line, c.expected, got)
			}
		}
	}
}
