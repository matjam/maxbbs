package main

import (
	"bytes"

	"github.com/gookit/slog"
	"github.com/matjam/maxbbs/internal/mecca"
)

func main() {
	slog.Configure(func(logger *slog.SugaredLogger) {
		f := logger.Formatter.(*slog.TextFormatter)
		f.EnableColor = true
	})

	slog.Info("running test template")

	template := `This is just a string you know. System Name: [sys_name] [sysop_name]`

	var in bytes.Buffer
	var out bytes.Buffer

	interpreter := mecca.NewInterpreter()
	session := interpreter.NewSession(&in, &out)
	err := session.Compile("test_template", template)
	if err != nil {
		slog.Panicf(err.Error())
	}

	err = session.Exec("test_template")
	if err != nil {
		slog.Panicf(err.Error())
	}

}
