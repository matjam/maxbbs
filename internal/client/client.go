package client

import (
	"fmt"

	"github.com/matjam/maxbbs/internal/mecca"
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

const (
	defaultWidth        = 80
	defaultHeight       = 24
	defaultColorEnabled = false
	defaultAnsiEnabled  = false
)

type TelnetClient struct {
	width        uint16 // width
	height       uint16 // height
	colorEnabled bool   // color capable
	ansiEnabled  bool   // ansi capable

	conn *telnet.Connection
	mec  mecca.Interpreter

	state clientState
}

type Option func(*TelnetClient)

func WithSize(width uint16, height uint16) Option {
	return func(c *TelnetClient) {
		c.width = width
		c.height = height
	}
}

func WithColor(colorEnabled bool) Option {
	return func(c *TelnetClient) {
		c.colorEnabled = colorEnabled
	}
}

func WithANSI(ansiEnabled bool) Option {
	return func(c *TelnetClient) {
		c.ansiEnabled = ansiEnabled
	}
}

func NewTelnetClient(conn *telnet.Connection, bbs mecca.BBS, opts ...Option) *TelnetClient {
	tc := &TelnetClient{
		conn:         conn,
		mec:          mecca.NewInterpreter(bbs),
		width:        defaultWidth,
		height:       defaultHeight,
		colorEnabled: defaultColorEnabled,
		ansiEnabled:  defaultAnsiEnabled,
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
