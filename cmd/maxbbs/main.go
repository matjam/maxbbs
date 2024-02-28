package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
	"github.com/matjam/maxbbs/internal/client"
	"github.com/matjam/maxbbs/internal/config"
	"github.com/matjam/maxbbs/internal/system"
	"github.com/matjam/telnet"
	"github.com/matjam/telnet/options"
	"github.com/mattn/go-isatty"
)

func main() {
	// This is share between all the sessions, and provides access to messages, forums, etc.
	bbs := system.NewBBS()

	slog.SetDefault(slog.New(
		tint.NewHandler(os.Stdout, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
			NoColor:    !isatty.IsTerminal(os.Stdout.Fd()),
		}),
	))

	slog.Info("starting server " + bbs.SystemName())

	telnetHandler := func(c *telnet.Connection) {
		slog.Info("Connection received", "remote", c.RemoteAddr())

		time.Sleep(time.Millisecond)
		nh := c.OptionHandlers[telnet.TeloptNAWS].(*options.NAWSHandler)
		slog.Info("Client terminal settings", "width", nh.Width, "height", nh.Height)

		conn := client.NewTelnetClient(c, client.BBSOption(bbs))
		conn.Start()

		slog.Info("Disconnected", "remote", c.RemoteAddr())
	}

	listenAddress := fmt.Sprintf("%v:%v", config.Get().Server.Telnet.Host, config.Get().Server.Telnet.Port)

	svr := telnet.NewServer(
		listenAddress,
		telnet.HandleFunc(telnetHandler),
		options.NAWSOption,
	)
	slog.Info("Telnet server listening", "address", listenAddress)

	if err := svr.ListenAndServe(); err != nil {
		slog.Error("Unable to start telnet server", "error", err.Error())
	}
}
