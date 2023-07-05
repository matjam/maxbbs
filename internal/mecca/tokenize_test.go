package mecca

import (
	"errors"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

func getTestCaseLine() int {
	_, _, line, _ := runtime.Caller(1)

	return line
}

func TestTokenizer(t *testing.T) {
	type testCase struct {
		line     int
		input    string
		expected []Token
		err      error
	}

	cases := []testCase{
		{
			getTestCaseLine(),
			`[white]Leave a message to [sysop_name] [[Y,n]? [gray ansreq menu]yn|`,
			[]Token{
				{Type: T_WHITE, Value: "", Args: []string(nil)},
				{Type: T_STRING, Value: "Leave a message to ", Args: []string(nil)},
				{Type: T_SYSOP_NAME, Value: "", Args: []string(nil)},
				{Type: T_STRING, Value: " [Y,n]? ", Args: []string(nil)},
				{Type: T_GRAY, Value: "", Args: []string(nil)},
				{Type: T_ANSREQ, Value: "", Args: []string(nil)},
				{Type: T_MENU, Value: "", Args: []string(nil)},
				{Type: T_STRING, Value: "yn|", Args: []string(nil)},
			},
			nil,
		},
		{
			getTestCaseLine(),
			`Some text [white]Leave a message to [sysop_name] [[Y,n]? [gray ansreq menu]yn|`,
			[]Token{
				{Type: T_STRING, Value: "Some text ", Args: []string(nil)},
				{Type: T_WHITE, Value: "", Args: []string(nil)},
				{Type: T_STRING, Value: "Leave a message to ", Args: []string(nil)},
				{Type: T_SYSOP_NAME, Value: "", Args: []string(nil)},
				{Type: T_STRING, Value: " [Y,n]? ", Args: []string(nil)},
				{Type: T_GRAY, Value: "", Args: []string(nil)},
				{Type: T_ANSREQ, Value: "", Args: []string(nil)},
				{Type: T_MENU, Value: "", Args: []string(nil)},
				{Type: T_STRING, Value: "yn|", Args: []string(nil)},
			},
			nil,
		},
		{
			getTestCaseLine(),
			`text [   sysop_name] text [sysop_name    ] text [   sysop_name   ] text`,
			[]Token{
				{Type: T_STRING, Value: "text ", Args: []string(nil)},
				{Type: T_SYSOP_NAME, Value: "", Args: []string(nil)},
				{Type: T_STRING, Value: " text ", Args: []string(nil)},
				{Type: T_SYSOP_NAME, Value: "", Args: []string(nil)},
				{Type: T_STRING, Value: " text ", Args: []string(nil)},
				{Type: T_SYSOP_NAME, Value: "", Args: []string(nil)},
				{Type: T_STRING, Value: " text", Args: []string(nil)},
			},
			nil,
		},
		{
			getTestCaseLine(),
			`text [   token`,
			[]Token{},
			ErrUnexpectedEndOfInput,
		},
		{
			getTestCaseLine(),
			`text [   token [ invalid ] text`,
			[]Token{},
			ErrTokenStartInsideTokenBlock,
		},
		{
			getTestCaseLine(),
			`text [comment this should be completely ignored] text`,
			[]Token{
				{Type: T_STRING, Value: "text ", Args: []string(nil)},
				{Type: T_STRING, Value: " text", Args: []string(nil)},
			},
			nil,
		},
		{
			getTestCaseLine(),
			`text [invalid token] text`,
			[]Token{
				{Type: T_STRING, Value: "text ", Args: []string(nil)},
				{Type: T_STRING, Value: " text", Args: []string(nil)},
			},
			ErrUnknownToken,
		},
		{
			getTestCaseLine(),
			`text [fg black] text`,
			[]Token{
				{Type: T_STRING, Value: "text ", Args: []string(nil)},
				{Type: T_FG, Value: "", Args: []string{"black"}},
				{Type: T_STRING, Value: " text", Args: []string(nil)},
			},
			nil,
		},
		{
			getTestCaseLine(),
			`text [locate 12 13] text`,
			[]Token{
				{Type: T_STRING, Value: "text ", Args: []string(nil)},
				{Type: T_LOCATE, Value: "", Args: []string{"12", "13"}},
				{Type: T_STRING, Value: " text", Args: []string(nil)},
			},
			nil,
		},
		{
			getTestCaseLine(),
			`text [locate 12] text`,
			[]Token{},
			ErrNotEnoughArguments,
		},
		{
			getTestCaseLine(),
			`text [repeatseq 4]four[10] text`,
			[]Token{
				{Type: T_STRING, Value: "text ", Args: []string(nil)},
				{Type: T_REPEATSEQ, Value: "", Args: []string{"4", "four", "10"}},
				{Type: T_STRING, Value: " text", Args: []string(nil)},
			},
			nil,
		},
		{
			getTestCaseLine(),
			`text [fg white]text [goto jump]text [/jump]text`,
			[]Token{
				{Type: T_STRING, Value: "text ", Args: []string(nil)},
				{Type: T_FG, Value: "", Args: []string{"white"}},
				{Type: T_STRING, Value: "text ", Args: []string(nil)},
				{Type: T_GOTO, Value: "", Args: []string{"jump"}},
				{Type: T_STRING, Value: "text ", Args: []string(nil)},
				{Type: T_LABEL, Value: "jump", Args: []string(nil)},
				{Type: T_STRING, Value: "text", Args: []string(nil)},
			},
			nil,
		},
	}

	for _, c := range cases {
		r := strings.NewReader(c.input)
		got, err := tokenize(r)

		if c.err != nil {
			if err == nil || !errors.Is(err, c.err) {
				t.Errorf("\n\n    line: %v\nexpected error: %v\n     got error: %v\n\n", c.line, c.err, err)
			}
		} else {
			if err != nil {
				t.Errorf("\n\n    line: %v\nexpected error: %v\n     got error: %v\n\n", c.line, c.err, err)
			}

			if !reflect.DeepEqual(got, c.expected) {
				t.Errorf("\n\n    line: %v\nexpected: %#v\n     got: %#v\n\n", c.line, c.expected, got)
			}
		}
	}

}
