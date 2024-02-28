package main

import (
	"bytes"
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
	"github.com/matjam/maxbbs/internal/mecca"
	"github.com/matjam/maxbbs/internal/system"
	"github.com/mattn/go-isatty"
)

func main() {
	slog.SetDefault(slog.New(
		tint.NewHandler(os.Stdout, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
			NoColor:    !isatty.IsTerminal(os.Stdout.Fd()),
		}),
	))

	bbs := system.NewBBS()

	slog.Info("running test template")

	template := `This is just a string you know. System Name: [sys_name] [sysop_name]`

	var in bytes.Buffer
	var out bytes.Buffer

	interpreter := mecca.NewInterpreter(bbs)
	session := interpreter.NewSession(&in, &out)
	err := session.Compile("test_template", template)
	if err != nil {
		panic(err)
	}

	err = session.Exec("test_template")
	if err != nil {
		panic(err)
	}

	slog.Info("output", out.String())
}
