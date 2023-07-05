package main

import (
	"time"

	"github.com/gookit/slog"
	"github.com/matjam/maxbbs/internal/client"
	"github.com/matjam/maxbbs/internal/system"
	"github.com/matjam/telnet"
	"github.com/matjam/telnet/options"
)

func main() {
	bbs := system.NewBBS()

	slog.Configure(func(logger *slog.SugaredLogger) {
		f := logger.Formatter.(*slog.TextFormatter)
		f.EnableColor = true
	})

	slog.Infof("starting server %v", bbs.SysName())

	telnetHandler := func(c *telnet.Connection) {
		slog.Infof("Connection received: %s", c.RemoteAddr())

		time.Sleep(time.Millisecond)
		nh := c.OptionHandlers[telnet.TeloptNAWS].(*options.NAWSHandler)
		slog.Infof("Client width: %d, height: %d", nh.Width, nh.Height)

		conn := client.NewTelnetClient(c)

		slog.Printf("Goodbye %s!", c.RemoteAddr())
	}

	svr := telnet.NewServer("127.0.0.1:9999", telnet.HandleFunc(telnetHandler), options.NAWSOption)
	svr.ListenAndServe()
}
