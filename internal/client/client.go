package client

import (
	"fmt"

	"github.com/matjam/telnet"
)

// clientState is what "page"  a client is current at in the BBS.
// For example, they start in STATE_CONNECT and then immediately
// go to STATE_LOGIN for the login page.
//
// There is a corresponding default mecca file for each clientState.

type clientState int

const (
	STATE_CONNECT clientState = iota
	STATE_LOGIN
	STATE_MAIN_MENU
)

// all paths are relative to the cfg
var stateMeccaMap = map[clientState]string{
	STATE_CONNECT:   "misc/logo.mec",
	STATE_LOGIN:     "misc/login.mec",
	STATE_MAIN_MENU: "misc/main_menu.mec",
}

type TelnetClient struct {
	W     uint16 // width
	H     uint16 // height
	Color bool   // color capable
	ANSI  bool   // ansi capable

	conn *telnet.Connection

	state clientState
}

type Option = func(c *TelnetClient)

func WithTerminalSize(W uint16, H uint16) Option {
	return func(c *TelnetClient) {
		c.W = W
		c.H = H
	}
}

func WithColor(colorEnabled bool) Option {
	return func(c *TelnetClient) {
		c.Color = colorEnabled
	}
}

func NewTelnetClient(conn *telnet.Connection, opts ...Option) *TelnetClient {
	tc := &TelnetClient{
		conn: conn,
	}

	for _, opt := range opts {
		opt(tc)
	}

	return tc
}

// Start the state machine
func (client *TelnetClient) Start() {

}

func (client *TelnetClient) Write(b []byte) (n int, err error) {
	return client.conn.Write(b)
}

func (client *TelnetClient) Writef(format string, args ...any) (n int, err error) {
	return client.conn.Write([]byte(fmt.Sprintf(format, args...)))
}

func (client *TelnetClient) Read(b []byte) (n int, err error) {
	return client.conn.Read(b)
}
