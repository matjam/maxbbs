package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gookit/slog"
	"github.com/matjam/telnet"
	"github.com/matjam/telnet/linereader"
	"github.com/matjam/telnet/options"
)

func main() {
	slog.Configure(func(logger *slog.SugaredLogger) {
		f := logger.Formatter.(*slog.TextFormatter)
		f.EnableColor = true
	})

	svr := telnet.NewServer("127.0.0.1:9999", telnet.HandleFunc(telnetHandler), options.NAWSOption)
	svr.ListenAndServe()
}

var mecTemplate = `

`

type TelnetClient struct {
	W     uint16 // width
	H     uint16 // height
	Color bool   // color capable
	ANSI  bool   // ansi capable

	*telnet.Connection
}

func (client TelnetClient) Writef(format string, args ...any) {
	client.Write([]byte(fmt.Sprintf(format, args...)))
}

func telnetHandler(c *telnet.Connection) {
	log.Printf("Connection received: %s", c.RemoteAddr())
	lr := linereader.New()
	go lr.ReadLines(c)

	time.Sleep(time.Millisecond)
	nh := c.OptionHandlers[telnet.TeloptNAWS].(*options.NAWSHandler)
	log.Printf("Client width: %d, height: %d", nh.Width, nh.Height)

	client := TelnetClient{
		Connection: c,
	}

	client.Writef("Welcome to MAXBBS v1.0.0\n\n")
	client.Writef("your terminal was detected as %vx%v, is that correct?\n", nh.Width, nh.Height)
	client.Writef("[Y/n]? ")

	for line := range lr.C {
		log.Printf("got %v\n", line)

	}

	log.Printf("Goodbye %s!", c.RemoteAddr())
}
