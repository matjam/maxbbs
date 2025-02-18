package main

import (
	"fmt"
	"log"
	"time"

	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	lm "github.com/charmbracelet/wish/logging"
	"github.com/matjam/mecca"
)

// Handle SSH requests.
func handler(next ssh.Handler) ssh.Handler {
	return func(sess ssh.Session) {
		// Create a new renderer.
		interpreter := mecca.NewInterpreter(mecca.WithTemplateRoot("config/mec"), mecca.WithSession(sess))

		_, _, active := sess.Pty()
		if !active {
			next(sess)
			return
		}

		err := interpreter.RenderTemplate("welcome.mec", nil)
		if err != nil {
			log.Println(err)
			next(sess)
			return
		}

		time.Sleep(5 * time.Second)

		next(sess)
	}
}

func main() {
	port := 3456
	s, err := wish.NewServer(
		wish.WithAddress(fmt.Sprintf(":%d", port)),
		wish.WithHostKeyPath("ssh_example"),
		wish.WithMiddleware(handler, lm.Middleware()),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("SSH server listening on port %d", port)
	log.Printf("To connect from your local machine run: ssh localhost -p %d", port)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
