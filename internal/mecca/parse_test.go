package mecca_test

import (
	"bytes"
	"testing"

	"github.com/matjam/maxbbs/internal/mecca"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/unicode"
)

func TestPlainStringParse(t *testing.T) {
	input := `This is a plain string`
	expected := []byte(input) // should be a straight conversion with no change

	parser := mecca.NewParser(unicode.UTF8.NewEncoder())
	got, err := parser.Parse(input)
	if err != nil {
		t.Error(err)
	}
	if bytes.Compare(got, expected) != 0 {
		t.Errorf("got %v expected %v", got, expected)
	}
}

func TestStringToCodePage437(t *testing.T) {
	input := `╣❤️`
	expected := []byte{185, 26, 26}

	parser := mecca.NewParser(charmap.CodePage437.NewEncoder())
	got, err := parser.Parse(input)
	if err != nil {
		t.Error(err)
	}
	if bytes.Compare(got, expected) != 0 {
		t.Errorf("got %v expected %v", got, expected)
	}
}
