package mecca_test

import (
	"bytes"
	"testing"

	"github.com/gookit/slog"
	"github.com/matjam/maxbbs/internal/mecca"
)

func TestInterpreter(t *testing.T) {
	slog.Configure(func(logger *slog.SugaredLogger) {
		f := logger.Formatter.(*slog.TextFormatter)
		f.EnableColor = true
	})

	template := `This is just a string you know. System Name: [sys_name] [sysop_name]`

	var in bytes.Buffer
	var out bytes.Buffer

	interpreter := mecca.NewInterpreter()
	session := interpreter.NewSession(&in, &out)
	err := session.Compile("test_template", template)
	if err != nil {
		t.Fatalf(err.Error())
	}

	err = session.Exec("test_template")
	if err != nil {
		t.Fatalf(err.Error())
	}

	t.Fatalf(out.String())
}
