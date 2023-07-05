package mecca

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

var ErrUnexpectedEndOfInput = errors.New("ErrUnexpectedEndOfInput: unexpected end of input")
var ErrTokenStartInsideTokenBlock = errors.New("ErrTokenStartInsideTokenBlock: start of a token inside an existing token block is illegal")

type lexTokenType int

const (
	L_STRING lexTokenType = iota
	L_TOKEN
)

func (t lexTokenType) String() string {
	switch t {
	case L_STRING:
		return "L_STRING"
	case L_TOKEN:
		return "L_TOKEN"
	}

	return "UNKNOWN"
}

type lexToken struct {
	Type  lexTokenType
	Value string
}

func (t lexToken) String() string {
	if t.Type == L_STRING || t.Type == L_TOKEN {
		return fmt.Sprintf(`%v(%v)`, t.Type, t.Value)
	}

	return fmt.Sprintf(`%v`, t.Type)
}

type lexer struct {
	tokens   []lexToken
	row      uint
	col      uint
	position uint
	input    *bufio.Reader
	c        rune
}

func (t *lexer) append(token lexToken) {
	if (token.Type == L_STRING || token.Type == L_TOKEN) && len(token.Value) == 0 {
		return
	}

	t.tokens = append(t.tokens, token)
}

// lex takes the input stream of bytes that we take to be a UTF-8 stream
// and converts it to a set of tokens or strings. Anything contained within
// [] is considered a "token" by the lexer, anything else is a string.
// The tokenizer then later takes this string/token output and further
// tokenizes it to MECCA strings and commands.
func lex(input io.Reader) ([]lexToken, error) {
	t := lexer{
		tokens: make([]lexToken, 0, 10),
		input:  bufio.NewReader(input),
	}

	var value strings.Builder
	inToken := false

	for {
		c, n, _ := t.input.ReadRune()
		if c == '\n' {
			t.row = t.row + 1
			t.col = 0
		} else {
			t.col = t.col + 1
		}

		if n > 0 {
			switch c {
			case '[':
				// check if this is the start of a [[ which is a [ literal
				c, n, err := t.input.ReadRune()
				if n > 0 && c == '[' {
					value.WriteRune('[')
					t.col = t.col + 1
					continue
				}
				if err := t.input.UnreadRune(); err != nil {
					return t.tokens, err
				}
				if err != nil {
					return t.tokens, nil
				}

				if inToken {
					// starting a token inside a token isn't going to work
					return []lexToken{}, fmt.Errorf("at row %v col %v: %w", t.row, t.col, ErrTokenStartInsideTokenBlock)
				}

				t.append(lexToken{Type: L_STRING, Value: value.String()})
				value.Reset()
				inToken = true
			case ']':
				if inToken {
					t.append(lexToken{Type: L_TOKEN, Value: value.String()})
					value.Reset()
					inToken = false
				} else {
					value.WriteRune(c)
				}
			case ' ':
				if inToken {
					t.append(lexToken{Type: L_TOKEN, Value: value.String()})
					value.Reset()
				} else {
					value.WriteRune(c)
				}
			default:
				value.WriteRune(c)
			}
		} else {
			if inToken {
				return []lexToken{}, fmt.Errorf("%w when processing token block", ErrUnexpectedEndOfInput)
			}
			t.append(lexToken{Type: L_STRING, Value: value.String()})
			break
		}
	}

	return t.tokens, nil
}
