package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/logging"
	"github.com/matjam/mecca"
)

func handler(next ssh.Handler) ssh.Handler {
	return func(sess ssh.Session) {
		// Create a new interpreter with a session.
		interpreter := mecca.NewInterpreter(mecca.WithTemplateRoot("config/mec"), mecca.WithSession(sess))

		_, _, active := sess.Pty()
		if !active {
			next(sess)
			return
		}

		if err := interpreter.RenderTemplate("welcome.mec", map[string]any{
			"bbsversion": "v1.0.0",
			"bbsname":    "MaxBBS",
			"sysopname":  "Max",
		}); err != nil {
			slog.Error("RenderTemplate error", "err", err)
			next(sess)
			return
		}

		time.Sleep(5 * time.Second)
		next(sess)
	}
}

func main() {
	charmLogger := log.NewWithOptions(
		os.Stderr,
		log.Options{
			ReportTimestamp: true,
			TimeFormat:      time.Kitchen,
		},
	)
	slog.SetDefault(slog.New(charmLogger))

	port := 3456
	s, err := wish.NewServer(
		wish.WithAddress(fmt.Sprintf(":%d", port)),
		wish.WithHostKeyPath(".ssh/ssh_server_key"),
		wish.WithMiddleware(handler, logging.MiddlewareWithLogger(charmLogger)),
	)
	if err != nil {
		slog.Error("Failed to create server", "err", err)
		os.Exit(1)
	}
	slog.Info("SSH server listening on", "port", port)
	slog.Info(fmt.Sprintf("To connect from your local machine run: ssh localhost -p %d", port))
	if err := s.ListenAndServe(); err != nil {
		slog.Error("Server error", "err", err)
		os.Exit(1)
	}
}
