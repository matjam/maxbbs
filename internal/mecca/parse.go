// package mecca implements a Maximus BBS compatible MECCA parser.
//
// MECCA files are plain UTF8 text and can contain special MECCA
// tokens which are parsed and translated by the MECCA parser into
// ANSI sequences in whatever character set is desired.
package mecca

import (
	"bytes"

	"github.com/gookit/slog"
	"golang.org/x/text/encoding"
)

// Parser contains all of the information that the MECCA parser needs
// to parse a MECCA file.
type Parser struct {
	encoder *encoding.Encoder
}

// use charmap.CodePage437.NewEncoder() to encode to CodePage437
func NewParser(encoder *encoding.Encoder) Parser {
	return Parser{
		encoder: encoding.ReplaceUnsupported(encoder),
	}
}

func (p *Parser) write(output *bytes.Buffer, s string) error {
	b, err := p.encoder.Bytes([]byte(s))
	if err != nil {
		return err
	}
	_, err = output.Write(b)
	return err
}

func (p *Parser) Parse(input string) ([]byte, error) {
	var output bytes.Buffer

	tokens, err := Tokenize(input)
	if err != nil {
		slog.Errorf("error parsing MECCA: %v", err.Error())
		return []byte{}, err
	}

	for _, t := range tokens {
		var err error
		switch t.Type {
		case STRING:
			err = p.write(&output, t.Value)
		}

		if err != nil {
			slog.Errorf("error writing %v: %v", t.Value, err)
		}
	}

	return output.Bytes(), nil
}
