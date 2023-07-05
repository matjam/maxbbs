package mecca

import "github.com/pborman/ansi"

type TokenType int

func (t TokenType) String() string {
	s, ok := TokenTypeMap[t]
	if !ok {
		return "UNKNOWN"
	}
	return s
}

type Token struct {
	Type  TokenType
	Value string
	Args  []string
}

// FGColors maps mecca colors to ANSI escape sequences
var FGColors = map[TokenType][]byte{
	T_BLACK:        []byte(ansi.BlackText),
	T_BLUE:         []byte(ansi.BlueText),
	T_GREEN:        []byte(ansi.GreenText),
	T_CYAN:         []byte(ansi.CyanText),
	T_RED:          []byte(ansi.RedText),
	T_MAGENTA:      []byte(ansi.MagentaText),
	T_BROWN:        []byte(ansi.YellowText),
	T_GRAY:         []byte(ansi.FaintWhiteText),
	T_DARKGRAY:     []byte(ansi.BoldBlackText),
	T_LIGHTBLUE:    []byte(ansi.BoldBlueText),
	T_LIGHTGREEN:   []byte(ansi.BoldGreenText),
	T_LIGHTCYAN:    []byte(ansi.BoldCyanText),
	T_LIGHTRED:     []byte(ansi.BoldRedText),
	T_LIGHTMAGENTA: []byte(ansi.BoldMagentaText),
	T_YELLOW:       []byte(ansi.BoldYellowText),
	T_WHITE:        []byte(ansi.WhiteText),
}

// BGColors maps mecca colors to ANSI escape sequences
var BGColors = map[TokenType][]byte{
	T_BLACK:   ansi.SGR.S().Format(ansi.BlackBackground),
	T_BLUE:    ansi.SGR.S().Format(ansi.BlueBackground),
	T_GREEN:   ansi.SGR.S().Format(ansi.GreenBackground),
	T_CYAN:    ansi.SGR.S().Format(ansi.CyanBackground),
	T_RED:     ansi.SGR.S().Format(ansi.RedBackground),
	T_MAGENTA: ansi.SGR.S().Format(ansi.MagentaBackground),
	T_BROWN:   ansi.SGR.S().Format(ansi.YellowBackground),
	T_GRAY:    ansi.SGR.S().Format(ansi.WhiteBackground),
}

const (
	T_STRING TokenType = iota
	T_UNIMPLEMENTED
	T_UNKNOWN
	// Color Tokens
	T_BLACK
	T_BLUE
	T_GREEN
	T_CYAN
	T_RED
	T_MAGENTA
	T_BROWN
	T_GRAY
	T_DARKGRAY
	T_LIGHTBLUE
	T_LIGHTGREEN
	T_LIGHTCYAN
	T_LIGHTRED
	T_LIGHTMAGENTA
	T_YELLOW
	T_WHITE
	T_BG
	T_ON
	T_BLINK
	T_BRIGHT
	T_DIM
	T_FG
	T_LOAD
	T_SAVE
	T_STEADY
	// Cursor Control and Video Tokens
	T_BELL
	T_BS
	T_CLEOL
	T_CLEOS
	T_CLS
	T_CR
	T_DOWN
	T_LEFT
	T_LF
	T_LOCATE
	T_TAB
	T_RIGHT
	T_SYSOPBELL
	T_UP
	// Informational Tokens
	T_ALIST_FILE
	T_ALIST_MSG
	T_CITY
	T_DATE
	T_DL
	T_EXPIRY_DATE
	T_EXPIRY_TIME
	T_FILE_CAREA
	T_FILE_CNAME
	T_FILE_DAREA
	T_FILE_SAREA
	T_FNAME
	T_FIRST
	T_IP
	T_LASTCALL
	T_LASTUSER
	T_LENGTH
	T_MINUTES
	T_MSG_CAREA
	T_MSG_CMSG
	T_MSG_CNAME
	T_MSG_DAREA
	T_MSG_HMSG
	T_MSG_NUMMSG
	T_MSG_SAREA
	T_NETBALANCE
	T_NETCREDIT
	T_NETDEBIT
	T_NETDL
	T_NODE_NUM
	T_PHONE
	T_RATIO
	T_REALNAME
	T_REMAIN
	T_RESPONSE
	T_SYSCALL
	T_SYS_NAME
	T_SYSOP_NAME
	T_TIME
	T_TIMEOFF
	T_UL
	T_USER
	T_USERCALL
	// Questionnaire Tokens
	T_ANSOPT
	T_ANSREQ
	T_CHOICE
	T_LEAVE_COMMENT
	T_MENU
	T_OPEN
	T_POST
	T_READLN
	T_SOPEN
	T_STORE
	T_WRITE
	// Privilege Level Controls
	T_ACS
	T_ACCESS
	T_ACSFILE
	T_ACCESSFILE
	T_PRIV_ABBREV
	T_PRIV_DESC
	T_PRIV_DOWN
	T_PRIV_LEVEL
	T_PRIV_UP
	T_SETPRIV
	// Lock and Key control
	T_IFKEY
	T_NOTKEY
	T_KEYON
	T_KEYOFF
	// Conditional and Flow Control Tokens
	T_B1200
	T_B2400
	T_B9600
	T_COL80
	T_COLOR
	T_COLOUR
	T_ENDCOLOR
	T_ENDCOLOUR
	T_ENDRIP
	T_EXPERT
	T_EXIT
	T_FILENEW
	T_GOTO
	T_HOTKEYS
	T_IFENTERED
	T_IFEXIST
	T_IFFSE
	T_IFFSR
	T_IFLANG
	T_IFTASK
	T_IFTIME
	T_INCITY
	T_ISLOCAL
	T_ISREMOTE
	T_JUMP
	T_LABEL
	T_MAXED
	T_MSG_ATTR
	T_MSG_CONF
	T_MSG_ECHO
	T_MSG_FILEATTACH
	T_MSG_LOCAL
	T_MSG_MATRIX
	T_MSG_NEXT
	T_MSG_NOMSGS
	T_MSG_NONEW
	T_MSG_NOREAD
	T_MSG_PRIOR
	T_NO_KEYPRESS
	T_NOCOLOR
	T_NOCOLOUR
	T_NORIP
	T_NOSTACKED
	T_NOTONTODAY
	T_NOVICE
	T_PERMANENT
	T_REGULAR
	T_RIP
	T_RIPHASFILE
	T_TAGGED
	T_TOP
	// Multinode Tokens
	T_APB
	T_CHAT_AVAIL
	T_CHAT_NOTAVAIL
	T_WHO_IS_ON
	// RIPscrip Graphics
	T_RIPDISPLAY
	T_RIPPATH
	T_RIPSEND
	T_TEXTSIZE
	// Miscellanous Tokens
	T_CKOFF
	T_CKON
	T_CLEAR_STACKED
	T_COMMENT
	T_COPY
	T_DELETE
	T_DISPLAY
	T_DOS
	T_ENTER
	T_HANGUP
	T_IBMCHARS
	T_INCLUDE
	T_KEY_POKE
	T_LANGUAGE
	T_LINK
	T_LOG
	T_MENU_CMD
	T_MENUPATH
	T_MEX
	T_MORE
	T_MOREOFF
	T_MOREON
	T_MSG_CHECKMAIL
	T_NEWFILES
	T_ONEXIT
	T_PAUSE
	T_QUIT
	T_QUOTE
	T_REPEAT
	T_REPEATSEQ
	T_TAG_READ
	T_TAG_WRITE
	T_TUNE
	T_XTERN_DOS
	T_XTERN_ERLVL
	T_XTERN_RUN
)

var TokenTypeMap = map[TokenType]string{
	T_STRING:        "T_STRING",
	T_UNIMPLEMENTED: "T_UNIMPLEMENTED",
	T_UNKNOWN:       "T_UNKNOWN",
	// Color Tokens
	T_BLACK:        "T_BLACK",
	T_BLUE:         "T_BLUE",
	T_GREEN:        "T_GREEN",
	T_CYAN:         "T_CYAN",
	T_RED:          "T_RED",
	T_MAGENTA:      "T_MAGENTA",
	T_BROWN:        "T_BROWN",
	T_GRAY:         "T_GRAY",
	T_DARKGRAY:     "T_DARKGRAY",
	T_LIGHTBLUE:    "T_LIGHTBLUE",
	T_LIGHTGREEN:   "T_LIGHTGREEN",
	T_LIGHTCYAN:    "T_LIGHTCYAN",
	T_LIGHTRED:     "T_LIGHTRED",
	T_LIGHTMAGENTA: "T_LIGHTMAGENTA",
	T_YELLOW:       "T_YELLOW",
	T_WHITE:        "T_WHITE",
	T_BG:           "T_BG",
	T_ON:           "T_ON",
	T_BLINK:        "T_BLINK",
	T_BRIGHT:       "T_BRIGHT",
	T_DIM:          "T_DIM",
	T_FG:           "T_FG",
	T_LOAD:         "T_LOAD",
	T_SAVE:         "T_SAVE",
	T_STEADY:       "T_STEADY",
	// Cursor Control and Video Tokens
	T_BELL:      "T_BELL",
	T_BS:        "T_BS",
	T_CLEOL:     "T_CLEOL",
	T_CLEOS:     "T_CLEOS",
	T_CLS:       "T_CLS",
	T_CR:        "T_CR",
	T_DOWN:      "T_DOWN",
	T_LEFT:      "T_LEFT",
	T_LF:        "T_LF",
	T_LOCATE:    "T_LOCATE",
	T_TAB:       "T_TAB",
	T_RIGHT:     "T_RIGHT",
	T_SYSOPBELL: "T_SYSOPBELL",
	T_UP:        "T_UP",
	// Informational Tokens
	T_ALIST_FILE:  "T_ALIST_FILE",
	T_ALIST_MSG:   "T_ALIST_MSG",
	T_CITY:        "T_CITY",
	T_DATE:        "T_DATE",
	T_DL:          "T_DL",
	T_EXPIRY_DATE: "T_EXPIRY_DATE",
	T_EXPIRY_TIME: "T_EXPIRY_TIME",
	T_FILE_CAREA:  "T_FILE_CAREA",
	T_FILE_CNAME:  "T_FILE_CNAME",
	T_FILE_DAREA:  "T_FILE_DAREA",
	T_FILE_SAREA:  "T_FILE_SAREA",
	T_FNAME:       "T_FNAME",
	T_FIRST:       "T_FIRST",
	T_IP:          "T_IP",
	T_LASTCALL:    "T_LASTCALL",
	T_LASTUSER:    "T_LASTUSER",
	T_LENGTH:      "T_LENGTH",
	T_MINUTES:     "T_MINUTES",
	T_MSG_CAREA:   "T_MSG_CAREA",
	T_MSG_CMSG:    "T_MSG_CMSG",
	T_MSG_CNAME:   "T_MSG_CNAME",
	T_MSG_DAREA:   "T_MSG_DAREA",
	T_MSG_HMSG:    "T_MSG_HMSG",
	T_MSG_NUMMSG:  "T_MSG_NUMMSG",
	T_MSG_SAREA:   "T_MSG_SAREA",
	T_NETBALANCE:  "T_NETBALANCE",
	T_NETCREDIT:   "T_NETCREDIT",
	T_NETDEBIT:    "T_NETDEBIT",
	T_NETDL:       "T_NETDL",
	T_NODE_NUM:    "T_NODE_NUM",
	T_PHONE:       "T_PHONE",
	T_RATIO:       "T_RATIO",
	T_REALNAME:    "T_REALNAME",
	T_REMAIN:      "T_REMAIN",
	T_RESPONSE:    "T_RESPONSE",
	T_SYSCALL:     "T_SYSCALL",
	T_SYS_NAME:    "T_SYS_NAME",
	T_SYSOP_NAME:  "T_SYSOP_NAME",
	T_TIME:        "T_TIME",
	T_TIMEOFF:     "T_TIMEOFF",
	T_UL:          "T_UL",
	T_USER:        "T_USER",
	T_USERCALL:    "T_USERCALL",
	// Questionnaire Tokens
	T_ANSOPT:        "T_ANSOPT",
	T_ANSREQ:        "T_ANSREQ",
	T_CHOICE:        "T_CHOICE",
	T_LEAVE_COMMENT: "T_LEAVE_COMMENT",
	T_MENU:          "T_MENU",
	T_OPEN:          "T_OPEN",
	T_POST:          "T_POST",
	T_READLN:        "T_READLN",
	T_SOPEN:         "T_SOPEN",
	T_STORE:         "T_STORE",
	T_WRITE:         "T_WRITE",
	// Privilege Level Controls
	T_ACS:         "T_ACS",
	T_ACCESS:      "T_ACCESS",
	T_ACSFILE:     "T_ACSFILE",
	T_ACCESSFILE:  "T_ACCESSFILE",
	T_PRIV_ABBREV: "T_PRIV_ABBREV",
	T_PRIV_DESC:   "T_PRIV_DESC",
	T_PRIV_DOWN:   "T_PRIV_DOWN",
	T_PRIV_LEVEL:  "T_PRIV_LEVEL",
	T_PRIV_UP:     "T_PRIV_UP",
	T_SETPRIV:     "T_SETPRIV",
	// Lock and Key control
	T_IFKEY:  "T_IFKEY",
	T_NOTKEY: "T_NOTKEY",
	T_KEYON:  "T_KEYON",
	T_KEYOFF: "T_KEYOFF",
	// Conditional and Flow Control Tokens
	T_B1200:          "T_B1200",
	T_B2400:          "T_B2400",
	T_B9600:          "T_B9600",
	T_COL80:          "T_COL80",
	T_COLOR:          "T_COLOR",
	T_COLOUR:         "T_COLOUR",
	T_ENDCOLOR:       "T_ENDCOLOR",
	T_ENDCOLOUR:      "T_ENDCOLOUR",
	T_ENDRIP:         "T_ENDRIP",
	T_EXPERT:         "T_EXPERT",
	T_EXIT:           "T_EXIT",
	T_FILENEW:        "T_FILENEW",
	T_GOTO:           "T_GOTO",
	T_HOTKEYS:        "T_HOTKEYS",
	T_IFENTERED:      "T_IFENTERED",
	T_IFEXIST:        "T_IFEXIST",
	T_IFFSE:          "T_IFFSE",
	T_IFFSR:          "T_IFFSR",
	T_IFLANG:         "T_IFLANG",
	T_IFTASK:         "T_IFTASK",
	T_IFTIME:         "T_IFTIME",
	T_INCITY:         "T_INCITY",
	T_ISLOCAL:        "T_ISLOCAL",
	T_ISREMOTE:       "T_ISREMOTE",
	T_JUMP:           "T_JUMP",
	T_LABEL:          "T_LABEL",
	T_MAXED:          "T_MAXED",
	T_MSG_ATTR:       "T_MSG_ATTR",
	T_MSG_CONF:       "T_MSG_CONF",
	T_MSG_ECHO:       "T_MSG_ECHO",
	T_MSG_FILEATTACH: "T_MSG_FILEATTACH",
	T_MSG_LOCAL:      "T_MSG_LOCAL",
	T_MSG_MATRIX:     "T_MSG_MATRIX",
	T_MSG_NEXT:       "T_MSG_NEXT",
	T_MSG_NOMSGS:     "T_MSG_NOMSGS",
	T_MSG_NONEW:      "T_MSG_NONEW",
	T_MSG_NOREAD:     "T_MSG_NOREAD",
	T_MSG_PRIOR:      "T_MSG_PRIOR",
	T_NO_KEYPRESS:    "T_NO_KEYPRESS",
	T_NOCOLOR:        "T_NOCOLOR",
	T_NOCOLOUR:       "T_NOCOLOUR",
	T_NORIP:          "T_NORIP",
	T_NOSTACKED:      "T_NOSTACKED",
	T_NOTONTODAY:     "T_NOTONTODAY",
	T_NOVICE:         "T_NOVICE",
	T_PERMANENT:      "T_PERMANENT",
	T_REGULAR:        "T_REGULAR",
	T_RIP:            "T_RIP",
	T_RIPHASFILE:     "T_RIPHASFILE",
	T_TAGGED:         "T_TAGGED",
	T_TOP:            "T_TOP",
	// Multinode Tokens
	T_APB:           "T_APB",
	T_CHAT_AVAIL:    "T_CHAT_AVAIL",
	T_CHAT_NOTAVAIL: "T_CHAT_NOTAVAIL",
	T_WHO_IS_ON:     "T_WHO_IS_ON",
	// RIPscrip Graphics
	T_RIPDISPLAY: "T_RIPDISPLAY",
	T_RIPPATH:    "T_RIPPATH",
	T_RIPSEND:    "T_RIPSEND",
	T_TEXTSIZE:   "T_TEXTSIZE",
	// Miscellanous Tokens
	T_CKOFF:         "T_CKOFF",
	T_CKON:          "T_CKON",
	T_CLEAR_STACKED: "T_CLEAR_STACKED",
	T_COMMENT:       "T_COMMENT",
	T_COPY:          "T_COPY",
	T_DELETE:        "T_DELETE",
	T_DISPLAY:       "T_DISPLAY",
	T_DOS:           "T_DOS",
	T_ENTER:         "T_ENTER",
	T_HANGUP:        "T_HANGUP",
	T_IBMCHARS:      "T_IBMCHARS",
	T_INCLUDE:       "T_INCLUDE",
	T_KEY_POKE:      "T_KEY_POKE",
	T_LANGUAGE:      "T_LANGUAGE",
	T_LINK:          "T_LINK",
	T_LOG:           "T_LOG",
	T_MENU_CMD:      "T_MENU_CMD",
	T_MENUPATH:      "T_MENUPATH",
	T_MEX:           "T_MEX",
	T_MORE:          "T_MORE",
	T_MOREOFF:       "T_MOREOFF",
	T_MOREON:        "T_MOREON",
	T_MSG_CHECKMAIL: "T_MSG_CHECKMAIL",
	T_NEWFILES:      "T_NEWFILES",
	T_ONEXIT:        "T_ONEXIT",
	T_PAUSE:         "T_PAUSE",
	T_QUIT:          "T_QUIT",
	T_QUOTE:         "T_QUOTE",
	T_REPEAT:        "T_REPEAT",
	T_REPEATSEQ:     "T_REPEATSEQ",
	T_TAG_READ:      "T_TAG_READ",
	T_TAG_WRITE:     "T_TAG_WRITE",
	T_TUNE:          "T_TUNE",
	T_XTERN_DOS:     "T_XTERN_DOS",
	T_XTERN_ERLVL:   "T_XTERN_ERLVL",
	T_XTERN_RUN:     "T_XTERN_RUN",
}
