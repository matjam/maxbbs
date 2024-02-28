package mecca

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"reflect"
	"strings"
	"sync"
)

type cachedTemplate struct {
	template     string
	tokens       []Token
	labelIndexes map[string]int
	pc           int // program counter

}

type Interpreter struct {
	templateCache sync.Map // stores a map of name:cachedTemplate
	bbs           BBS
	in            io.Reader // attached input from terminal
	out           io.Writer // attached output to terminal
}

type Option func(session *Interpreter)

// NewInterpreter initializes a new interpreter
func NewInterpreter(bbs BBS, input io.Reader, output io.Writer, opts ...Option) *Interpreter {
	i := Interpreter{
		bbs: bbs,
		in:  input,
		out: output,
	}

	for _, opt := range opts {
		opt(&i)
	}

	return &i
}

// Compile takes the given MECCA template string and compiles it to MECCA
// Tokens. Each name/template pair should be unique and consistent. If you
// provide a new template for a given name, the cached compiled template
// will be replaced.
func (s *Interpreter) Compile(name string, template string) error {
	t := strings.NewReader(template)
	tokens, err := tokenize(t)
	if err != nil {
		return err
	}

	labels := make(map[string]int)
	// fetch all the labels and store their indexes
	for i, token := range tokens {
		if token.Type == T_LABEL {
			labels[token.Value] = i
		}
	}

	compiled := cachedTemplate{
		template:     template,
		tokens:       tokens,
		labelIndexes: labels,
		pc:           0,
	}

	s.templateCache.Store(name, compiled)

	return nil
}

var ErrNotCompiled = errors.New("ErrNotCompiled: template not compiled")
var ErrInvalidTemplateCache = errors.New("ErrInvalidTemplateCache: templateCache contains something other than cachedTemplate")

// Exec will execute the named template, which must be already compiled. Exec
// may block if the template requires input, and MECCA templates may execute
// for an unlimited amount of time as you can jump to different parts of the
// same file, or another file, etc.
func (s InterpreterSession) Exec(name string) error {
	t, ok := s.templateCache.Load(name)
	if !ok {
		return fmt.Errorf("%w: %v not in cache", ErrNotCompiled, name)
	}
	template, ok := t.(cachedTemplate)
	if !ok {
		return fmt.Errorf("%w: %v not correct type", ErrInvalidTemplateCache, reflect.TypeOf(t))
	}

	// and here the magic happens
	for {
		if template.pc >= len(template.tokens) {
			break
		}

		token := template.tokens[template.pc]
		handlerFunc, ok := handlerTable[token.Type]
		if ok {
			err := handlerFunc(s, token)
			if err != nil {
				slog.Error("template token error", "template", name, "token", token.Type, "error", err.Error())
			}
		} else {
			slog.Warn("template token unimplemented", "template", name, "token", token.Type)
		}

		template.pc++
	}

	return nil
}

type tokenHandlerFunc func(InterpreterSession, Token) error

var handlerTable = map[TokenType]tokenHandlerFunc{
	T_STRING:     handleString,
	T_SYS_NAME:   handleSysName,
	T_SYSOP_NAME: handleSysopName,
}

func handleString(s InterpreterSession, t Token) error {
	_, err := s.out.Write([]byte(t.Value))
	return err
}

func handleSysName(s InterpreterSession, t Token) error {
	_, err := s.out.Write([]byte(s.bbs.SystemName()))
	return err
}

func handleSysopName(s InterpreterSession, t Token) error {
	_, err := s.out.Write([]byte(s.bbs.SysopName()))
	return err
}
