package mecca

import (
	"errors"
	"fmt"
	"io"
	"reflect"
	"strings"
	"sync"

	"github.com/gookit/slog"
)

type cachedTemplate struct {
	template     string
	tokens       []Token
	labelIndexes map[string]int
	pc           int // program counter

}

type Interpreter struct {
	templateCache sync.Map // stores a map of name:cachedTemplate
}

type InterpreterSession struct {
	*Interpreter
	in  io.Reader // attached input from terminal
	out io.Writer // attached output to terminal
}

// NewInterpreter initializes a new interpreter which can be shared among
// multiple terminal connections.
func NewInterpreter() Interpreter {
	return Interpreter{}
}

// NewSession will create an InterpreterSession which is specific to a
// given terminal connection.
func (i *Interpreter) NewSession(in io.Reader, out io.Writer) InterpreterSession {
	return InterpreterSession{i, in, out}
}

// Compile takes the given MECCA template string and compiles it to MECCA
// Tokens. Each name/template pair should be unique and consistent. If you
// provide a new template for a given name, the cached compiled template
// will be replaced.
func (s InterpreterSession) Compile(name string, template string) error {
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
				slog.Errorf("template %v token %v error: %v", name, token.Type, err.Error())
			}
		} else {
			slog.Warnf("template %v token %v unimplemented", name, token.Type)
		}

		template.pc++
	}

	return nil
}

type tokenHandlerFunc func(InterpreterSession, Token) error

var handlerTable = map[TokenType]tokenHandlerFunc{
	T_STRING:   handleString,
	T_SYS_NAME: handleSysName,
}

func handleString(s InterpreterSession, t Token) error {
	_, err := s.out.Write([]byte(t.Value))
	return err
}

func handleSysName(s InterpreterSession, t Token) error {
	_, err := s.out.Write([]byte("MaxBBS"))
	return err
}
