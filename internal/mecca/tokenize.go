package mecca

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/exp/slices"
)

var ErrEndOfInput = errors.New("end of input")
var ErrUnknownToken = errors.New("unknown token")
var ErrMissingParameters = errors.New("token requires parameters")

type Token struct {
	Type  TokenType
	Value string
}

type tokenizer struct {
	tokens   []Token
	row      uint
	col      uint
	position uint
	input    string
	c        rune
}

func (t *tokenizer) append(token Token) {
	if len(token.Value) == 0 {
		return
	}

	t.tokens = append(t.tokens, token)
}

func Tokenize(input string) ([]Token, error) {
	t := tokenizer{
		tokens: make([]Token, 0, 10),
		input:  input,
	}

	var value bytes.Buffer

	// start the state machine, at the start we're processing strings
	// so we are just writing a token that holds a string.
	for {
		// fetch the next character
		if c, err := t.next(); err == nil {
			if c == '[' {
				peek, err := t.peek()
				if peek == '[' || err != nil {
					// regardless of what happens, we write this [ character because
					// either we're doing [[ or we're doing [ and its the end of the
					// input.
					value.WriteRune(c)
					if err != nil {
						// must be the end of the input, so we write the final string value
						t.append(Token{Type: STRING, Value: value.String()})
						break
					}
				} else {
					// we're in token processing now
					tokens, err := t.tokenStart()
					if err != nil {
						if errors.Is(err, ErrEndOfInput) {
							t.tokens = append(t.tokens, tokens...)
							return t.tokens
						}
						return []Token{}, err
					}

					t.tokens = append(t.tokens, tokens...)
				}
			} else {
				value.WriteRune(c)
			}
		} else {
			// we're at the end of the input
			t.append(Token{Type: STRING, Value: value.String()})
			break
		}
	}

	return t.tokens, nil
}

// these are tokens that have parameters inside the square brackets, like [comment xxx yyy zzz]
var tokensWithTokenParameters = []TokenType{
	FG,
	LOCATE,
	ACS,
	ACCESS,
	ACSFILE,
	ACCESSFILE,
	SETPRIV,
	GOTO,
	LABEL,
	MSG_ATTR,
	TEXTSIZE,
	COMMENT,
	COPY,
	INCLUDE,
	MENU_CMD,
	REPEATSEQ,
}

// This function handles when we're "inside" a token block. It will keep adding
// tokens to its internal list until done, then return them. Its guaranteed to
// not return "empty" tokens.
func (t *tokenizer) tokenStart() ([]Token, error) {
	var value bytes.Buffer
	var tokens []Token

	for {
		if c, err := t.next(); err == nil {
			switch c {
			case ']':
				tokenType, err := parseToken(value.String())
				if err != nil {
					return tokens, err
				}

				// if this is a token that requires parameters, then we should error
				if slices.Contains(tokensWithTokenParameters, tokenType) {
					return tokens, fmt.Errorf("%w with token %v", ErrMissingParameters, strings.ToLower(value.String()))
				}

				tokens = append(tokens, Token{Type: tokenType, Value: value.String()})
				return tokens, nil
			case ' ':
				// end of the token but start of a new token
				tokenType, err := parseToken(value.String())
				if err != nil {
					return tokens, err
				}
				tokens = append(tokens, Token{Type: tokenType, Value: value.String()})
				value.Reset()

				if slices.Contains(tokensWithTokenParameters, tokenType) {

				}
			default:
				value.WriteRune(c)
			}
		} else {
			return tokens, err
		}
	}
}

func (t *tokenizer) readUntil(r rune) (string, error) {

}

// gets the next character
func (t *tokenizer) next() (rune, error) {
	if t.position >= uint(len(t.input)) {
		return ' ', ErrEndOfInput
	}

	t.c = []rune(t.input)[t.position]
	t.position = t.position + 1
	return t.c, nil
}

// peeks at the next character
func (t *tokenizer) peek() (rune, error) {
	if t.position >= uint(len(t.input)) {
		return ' ', ErrEndOfInput
	}

	return []rune(t.input)[t.position], nil
}

func parseToken(token string) (TokenType, error) {
	switch strings.ToLower(token) {
	case "black":
		return BLACK, nil
	case "blue":
		return BLUE, nil
	case "green":
		return GREEN, nil
	case "cyan":
		return CYAN, nil
	case "red":
		return RED, nil
	case "magenta":
		return MAGENTA, nil
	case "brown":
		return BROWN, nil
	case "gray":
		return GRAY, nil
	case "darkgray":
		return DARKGRAY, nil
	case "lightblue":
		return LIGHTBLUE, nil
	case "lightgren":
		return LIGHTGREEN, nil
	case "lightcyan":
		return LIGHTCYAN, nil
	case "lightred":
		return LIGHTRED, nil
	case "lightmagenta":
		return LIGHTMAGENTA, nil
	case "yellow":
		return YELLOW, nil
	case "white":
		return WHITE, nil
	case "bg":
		return BG, nil
	case "on":
		return ON, nil
	case "blink":
		return BLINK, nil
	case "bright":
		return BRIGHT, nil
	case "dim":
		return DIM, nil
	case "fg":
		return FG, nil
	case "load":
		return LOAD, nil
	case "save":
		return SAVE, nil
	case "steady":
		return STEADY, nil
	case "comment":
		return COMMENT, nil
	default:
		return EXIT, fmt.Errorf("%w: %v", ErrUnknownToken, token)
	}
}

// for _, c := range input {
// 	col = col + 1

// 	if c == '\n' {
// 		row = row + 1
// 		col = 0
// 	}

// 	switch current_context {
// 	case STRING:
// 		if c == '[' {
// 			current_context = TOKEN
// 		} else {
// 			valueString.WriteRune(c)
// 		}
// 	case TOKEN:
// 		if current_token == NONE {
// 			switch c {
// 			case '[':
// 				current_context = STRING
// 				valueString.WriteRune(c)
// 			case ']':
// 				current_context = STRING
// 				tokenType, err := parseToken(tokenString.String())
// 				if err != nil {
// 					return []Token{}, fmt.Errorf("error parsing row %v col %v: %w", row, col, err)
// 				}
// 				tokens.append(Token{
// 					Type:  tokenType,
// 					Value: tokenString.String(),
// 				})
// 				tokenString.Reset()
// 			case ' ': // space can be used to start another token, or some tokens have arguments
// 				tokenType, err := parseToken(tokenString.String())
// 				if err != nil {
// 					return []Token{}, fmt.Errorf("error parsing row %v col %v: %w", row, col, err)
// 				}
// 				tokens.append(Token{
// 					Type:  tokenType,
// 					Value: tokenString.String(),
// 				})
// 				tokenString.Reset()

// 				switch tokenType {
// 				case COMMENT:
// 					current_token = COMMENT
// 				}
// 			default:
// 				// we're actually starting a token, so we need to emit the string we
// 				// were building before
// 				if valueString.Len() > 0 {
// 					tokens.append(Token{
// 						Type:  STRING,
// 						Value: valueString.String(),
// 					})
// 					valueString.Reset()
// 				}

// 				tokenString.WriteRune(c)
// 			}
// 		} else {
// 			switch current_token {
// 			case COMMENT:
// 				// comment tokens ignore everything until the token is closed with ]
// 				switch c {
// 				case ']':
// 					current_context = STRING
// 					tokens.append(Token{
// 						Type:  COMMENT,
// 						Value: tokenString.String(), // might as well capture the comment
// 					})
// 					tokenString.Reset()
// 				default:
// 					tokenString.WriteRune(c)
// 				}
// 			}
// 		}
// 	}
// }
