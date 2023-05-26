package mecca_test

import (
	"reflect"
	"testing"

	"github.com/matjam/maxbbs/internal/mecca"
)

func TestPlainString(t *testing.T) {
	input := `Just a plain string`
	got, err := mecca.Tokenize(input)
	if err != nil {
		t.Error(err)
	}

	expected := []mecca.Token{
		{Type: mecca.STRING, Value: `Just a plain string`},
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected %#v got %#v", expected, got)
	}
}

func TestEscapedBracket(t *testing.T) {
	input := `Want to check your mail [[Y,n]?`
	got, err := mecca.Tokenize(input)
	if err != nil {
		t.Error(err)
	}

	expected := []mecca.Token{
		{Type: mecca.STRING, Value: `Want to check your mail [Y,n]?`},
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected %#v got %#v", expected, got)
	}
}

func TestTokenization(t *testing.T) {
	input := `Just a [save][white]plain [load]string`
	got, err := mecca.Tokenize(input)
	if err != nil {
		t.Error(err)
	}

	expected := []mecca.Token{
		{Type: mecca.STRING, Value: "Just a "},
		{Type: mecca.SAVE, Value: "save"},
		{Type: mecca.WHITE, Value: "white"},
		{Type: mecca.STRING, Value: "plain "},
		{Type: mecca.LOAD, Value: "load"},
		{Type: mecca.STRING, Value: "string"},
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %#v expected %#v", got, expected)
	}
}

func TestCompoundTokenization(t *testing.T) {
	input := `Just a [save white]plain [load]string`
	got, err := mecca.Tokenize(input)
	if err != nil {
		t.Error(err)
	}

	expected := []mecca.Token{
		{Type: mecca.STRING, Value: "Just a "},
		{Type: mecca.SAVE, Value: "save"},
		{Type: mecca.WHITE, Value: "white"},
		{Type: mecca.STRING, Value: "plain "},
		{Type: mecca.LOAD, Value: "load"},
		{Type: mecca.STRING, Value: "string"},
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %#v expected %#v", got, expected)
	}
}

func TestTokenizationComments(t *testing.T) {
	input := `Just a [save white][comment ignore this]plain [load]string`
	got, err := mecca.Tokenize(input)
	if err != nil {
		t.Error(err)
	}

	expected := []mecca.Token{
		{Type: mecca.STRING, Value: "Just a "},
		{Type: mecca.SAVE, Value: "save"},
		{Type: mecca.WHITE, Value: "white"},
		{Type: mecca.COMMENT, Value: "ignore this"},
		{Type: mecca.STRING, Value: "plain "},
		{Type: mecca.LOAD, Value: "load"},
		{Type: mecca.STRING, Value: "string"},
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %#v expected %#v", got, expected)
	}
}
