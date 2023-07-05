package mecca

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

var ErrUnknownToken = errors.New("ErrUnknownToken: unknown token")
var ErrNotEnoughArguments = errors.New("ErrNotEnoughArguments: not enough arguments to token")

// tokenize is a multiparse lexer/tokenizer that takes an input stream of utf-8 strings and
// emits an array of parsed tokens ready for the MECCA interpreter to execute.
func tokenize(input io.Reader) ([]Token, error) {
	var err error

	lexTokens, err := lex(input)
	if err != nil {
		return []Token{}, err
	}

	tokens := make([]Token, 0)
	cur := 0

	getArgs := func(argCount int) []string {
		args := make([]string, 0)
		if cur+argCount >= len(lexTokens) {
			err = fmt.Errorf("%w for token %v", ErrNotEnoughArguments, lexTokens[cur].Type)
			return []string{}
		}

		for i := 0; i < argCount; i++ {
			cur = cur + 1
			if lexTokens[cur].Type != L_TOKEN {
				err = fmt.Errorf("%w for token %v", ErrNotEnoughArguments, lexTokens[cur].Type)
				return []string{}
			}
			args = append(args, lexTokens[cur].Value)
		}

		return args
	}

	var l lexToken
	// now we take the tokens emitted by lex() and transform them into actual tokens. Tokens
	// can be either strings or commands, and some commands have arguments.
	for {
		if cur >= len(lexTokens) {
			break
		}

		l = lexTokens[cur]

		switch l.Type {
		case L_STRING:
			// Strings just come across straight, we don't touch them
			tokens = append(tokens, Token{Type: T_STRING, Value: l.Value})
		case L_TOKEN:
			switch strings.ToUpper(l.Value) {
			case "BLACK":
				tokens = append(tokens, Token{Type: T_BLACK})
			case "BLUE":
				tokens = append(tokens, Token{Type: T_BLUE})
			case "GREEN":
				tokens = append(tokens, Token{Type: T_GREEN})
			case "CYAN":
				tokens = append(tokens, Token{Type: T_CYAN})
			case "RED":
				tokens = append(tokens, Token{Type: T_RED})
			case "MAGENTA":
				tokens = append(tokens, Token{Type: T_MAGENTA})
			case "BROWN":
				tokens = append(tokens, Token{Type: T_BROWN})
			case "GRAY":
				tokens = append(tokens, Token{Type: T_GRAY})
			case "DARKGRAY":
				tokens = append(tokens, Token{Type: T_DARKGRAY})
			case "LIGHTBLUE":
				tokens = append(tokens, Token{Type: T_LIGHTBLUE})
			case "LIGHTGREEN":
				tokens = append(tokens, Token{Type: T_LIGHTGREEN})
			case "LIGHTCYAN":
				tokens = append(tokens, Token{Type: T_LIGHTCYAN})
			case "LIGHTRED":
				tokens = append(tokens, Token{Type: T_LIGHTRED})
			case "LIGHTMAGENTA":
				tokens = append(tokens, Token{Type: T_LIGHTMAGENTA})
			case "YELLOW":
				tokens = append(tokens, Token{Type: T_YELLOW})
			case "WHITE":
				tokens = append(tokens, Token{Type: T_WHITE})
			case "BG":
				tokens = append(tokens, Token{Type: T_BG})
			case "ON":
				tokens = append(tokens, Token{Type: T_ON})
			case "BLINK":
				tokens = append(tokens, Token{Type: T_BLINK})
			case "BRIGHT":
				tokens = append(tokens, Token{Type: T_BRIGHT})
			case "DIM":
				tokens = append(tokens, Token{Type: T_DIM})
			case "FG":
				tokens = append(tokens, Token{Type: T_FG, Args: getArgs(1)})
			case "LOAD":
				tokens = append(tokens, Token{Type: T_LOAD})
			case "SAVE":
				tokens = append(tokens, Token{Type: T_SAVE})
			case "STEADY":
				tokens = append(tokens, Token{Type: T_STEADY})
			// Cursor Control and Video Tokens
			case "BELL":
				tokens = append(tokens, Token{Type: T_BELL})
			case "BS":
				tokens = append(tokens, Token{Type: T_BS})
			case "CLEOL":
				tokens = append(tokens, Token{Type: T_CLEOL})
			case "CLEOS":
				tokens = append(tokens, Token{Type: T_CLEOS})
			case "CLS":
				tokens = append(tokens, Token{Type: T_CLS})
			case "CR":
				tokens = append(tokens, Token{Type: T_CR})
			case "DOWN":
				tokens = append(tokens, Token{Type: T_DOWN})
			case "LEFT":
				tokens = append(tokens, Token{Type: T_LEFT})
			case "LF":
				tokens = append(tokens, Token{Type: T_LF})
			case "LOCATE":
				tokens = append(tokens, Token{Type: T_LOCATE, Args: getArgs(2)})
			case "TAB":
				tokens = append(tokens, Token{Type: T_TAB})
			case "RIGHT":
				tokens = append(tokens, Token{Type: T_RIGHT})
			case "SYSOPBELL":
				tokens = append(tokens, Token{Type: T_SYSOPBELL})
			case "UP":
				tokens = append(tokens, Token{Type: T_UP})
				// Informational Tokens
			case "ALIST_FILE":
				tokens = append(tokens, Token{Type: T_ALIST_FILE})
			case "ALIST_MSG":
				tokens = append(tokens, Token{Type: T_ALIST_MSG})
			case "CITY":
				tokens = append(tokens, Token{Type: T_CITY})
			case "DATE":
				tokens = append(tokens, Token{Type: T_DATE})
			case "DL":
				tokens = append(tokens, Token{Type: T_DL})
			case "EXPIRY_DATE":
				tokens = append(tokens, Token{Type: T_EXPIRY_DATE})
			case "EXPIRY_TIME":
				tokens = append(tokens, Token{Type: T_EXPIRY_TIME})
			case "FILE_CAREA":
				tokens = append(tokens, Token{Type: T_FILE_CAREA})
			case "FILE_CNAME":
				tokens = append(tokens, Token{Type: T_FILE_CNAME})
			case "FILE_DAREA":
				tokens = append(tokens, Token{Type: T_FILE_DAREA})
			case "FILE_SAREA":
				tokens = append(tokens, Token{Type: T_FILE_SAREA})
			case "FNAME":
				tokens = append(tokens, Token{Type: T_FNAME})
			case "FIRST":
				tokens = append(tokens, Token{Type: T_FIRST})
			case "IP":
				tokens = append(tokens, Token{Type: T_IP})
			case "LASTCALL":
				tokens = append(tokens, Token{Type: T_LASTCALL})
			case "LASTUSER":
				tokens = append(tokens, Token{Type: T_LASTUSER})
			case "LENGTH":
				tokens = append(tokens, Token{Type: T_LENGTH})
			case "MINUTES":
				tokens = append(tokens, Token{Type: T_MINUTES})
			case "MSG_CAREA":
				tokens = append(tokens, Token{Type: T_MSG_CAREA})
			case "MSG_CMSG":
				tokens = append(tokens, Token{Type: T_MSG_CMSG})
			case "MSG_CNAME":
				tokens = append(tokens, Token{Type: T_MSG_CNAME})
			case "MSG_DAREA":
				tokens = append(tokens, Token{Type: T_MSG_DAREA})
			case "MSG_HMSG":
				tokens = append(tokens, Token{Type: T_MSG_HMSG})
			case "MSG_NUMMSG":
				tokens = append(tokens, Token{Type: T_MSG_NUMMSG})
			case "MSG_SAREA":
				tokens = append(tokens, Token{Type: T_MSG_SAREA})
			case "NETBALANCE":
				tokens = append(tokens, Token{Type: T_NETBALANCE})
			case "NETCREDIT":
				tokens = append(tokens, Token{Type: T_NETCREDIT})
			case "NETDEBIT":
				tokens = append(tokens, Token{Type: T_NETDEBIT})
			case "NETDL":
				tokens = append(tokens, Token{Type: T_NETDL})
			case "NODE_NUM":
				tokens = append(tokens, Token{Type: T_NODE_NUM})
			case "PHONE":
				tokens = append(tokens, Token{Type: T_PHONE})
			case "RATIO":
				tokens = append(tokens, Token{Type: T_RATIO})
			case "REALNAME":
				tokens = append(tokens, Token{Type: T_REALNAME})
			case "REMAIN":
				tokens = append(tokens, Token{Type: T_REMAIN})
			case "RESPONSE":
				tokens = append(tokens, Token{Type: T_RESPONSE})
			case "SYSCALL":
				tokens = append(tokens, Token{Type: T_SYSCALL})
			case "SYS_NAME":
				tokens = append(tokens, Token{Type: T_SYS_NAME})
			case "SYSOP_NAME":
				tokens = append(tokens, Token{Type: T_SYSOP_NAME})
			case "TIME":
				tokens = append(tokens, Token{Type: T_TIME})
			case "TIMEOFF":
				tokens = append(tokens, Token{Type: T_TIMEOFF})
			case "UL":
				tokens = append(tokens, Token{Type: T_UL})
			case "USER":
				tokens = append(tokens, Token{Type: T_USER})
			case "USERCALL":
				tokens = append(tokens, Token{Type: T_USERCALL})
				// Questionnaire Tokens
			case "ANSOPT":
				tokens = append(tokens, Token{Type: T_ANSOPT})
			case "ANSREQ":
				tokens = append(tokens, Token{Type: T_ANSREQ})
			case "CHOICE":
				tokens = append(tokens, Token{Type: T_CHOICE})
			case "LEAVE_COMMENT":
				tokens = append(tokens, Token{Type: T_LEAVE_COMMENT})
			case "MENU":
				tokens = append(tokens, Token{Type: T_MENU})
			case "OPEN":
				tokens = append(tokens, Token{Type: T_OPEN})
			case "POST":
				tokens = append(tokens, Token{Type: T_POST})
			case "READLN":
				tokens = append(tokens, Token{Type: T_READLN})
			case "SOPEN":
				tokens = append(tokens, Token{Type: T_SOPEN})
			case "STORE":
				tokens = append(tokens, Token{Type: T_STORE})
			case "WRITE":
				tokens = append(tokens, Token{Type: T_WRITE})
				// Privilege Level Controls
			case "ACS":
				tokens = append(tokens, Token{Type: T_ACS, Args: getArgs(1)})
			case "ACCESS":
				tokens = append(tokens, Token{Type: T_ACCESS, Args: getArgs(1)})
			case "ACSFILE":
				tokens = append(tokens, Token{Type: T_ACSFILE, Args: getArgs(1)})
			case "ACCESSFILE":
				tokens = append(tokens, Token{Type: T_ACCESSFILE, Args: getArgs(1)})
			case "PRIV_ABBREV":
				tokens = append(tokens, Token{Type: T_PRIV_ABBREV})
			case "PRIV_DESC":
				tokens = append(tokens, Token{Type: T_PRIV_DESC})
			case "PRIV_DOWN":
				tokens = append(tokens, Token{Type: T_PRIV_DOWN})
			case "PRIV_LEVEL":
				tokens = append(tokens, Token{Type: T_PRIV_LEVEL})
			case "PRIV_UP":
				tokens = append(tokens, Token{Type: T_PRIV_UP})
			case "SETPRIV":
				tokens = append(tokens, Token{Type: T_SETPRIV, Args: getArgs(1)})
				// Lock and Key control
			case "IFKEY":
				tokens = append(tokens, Token{Type: T_IFKEY})
			case "NOTKEY":
				tokens = append(tokens, Token{Type: T_NOTKEY})
			case "KEYON":
				tokens = append(tokens, Token{Type: T_KEYON})
			case "KEYOFF":
				tokens = append(tokens, Token{Type: T_KEYOFF})
				// Conditional and Flow Control Tokens
			case "B1200":
				tokens = append(tokens, Token{Type: T_B1200})
			case "B2400":
				tokens = append(tokens, Token{Type: T_B2400})
			case "B9600":
				tokens = append(tokens, Token{Type: T_B9600})
			case "COL80":
				tokens = append(tokens, Token{Type: T_COL80})
			case "COLOR":
				tokens = append(tokens, Token{Type: T_COLOR})
			case "COLOUR":
				tokens = append(tokens, Token{Type: T_COLOUR})
			case "ENDCOLOR":
				tokens = append(tokens, Token{Type: T_ENDCOLOR})
			case "ENDCOLOUR":
				tokens = append(tokens, Token{Type: T_ENDCOLOUR})
			case "ENDRIP":
				tokens = append(tokens, Token{Type: T_ENDRIP})
			case "EXPERT":
				tokens = append(tokens, Token{Type: T_EXPERT})
			case "EXIT":
				tokens = append(tokens, Token{Type: T_EXIT})
			case "FILENEW":
				tokens = append(tokens, Token{Type: T_FILENEW})
			case "GOTO":
				tokens = append(tokens, Token{Type: T_GOTO, Args: getArgs(1)})
			case "HOTKEYS":
				tokens = append(tokens, Token{Type: T_HOTKEYS})
			case "IFENTERED":
				tokens = append(tokens, Token{Type: T_IFENTERED})
			case "IFEXIST":
				tokens = append(tokens, Token{Type: T_IFEXIST})
			case "IFFSE":
				tokens = append(tokens, Token{Type: T_IFFSE})
			case "IFFSR":
				tokens = append(tokens, Token{Type: T_IFFSR})
			case "IFLANG":
				tokens = append(tokens, Token{Type: T_IFLANG})
			case "IFTASK":
				tokens = append(tokens, Token{Type: T_IFTASK})
			case "IFTIME":
				tokens = append(tokens, Token{Type: T_IFTIME})
			case "INCITY":
				tokens = append(tokens, Token{Type: T_INCITY})
			case "ISLOCAL":
				tokens = append(tokens, Token{Type: T_ISLOCAL})
			case "ISREMOTE":
				tokens = append(tokens, Token{Type: T_ISREMOTE})
			case "JUMP":
				tokens = append(tokens, Token{Type: T_JUMP})
			case "LABEL":
				tokens = append(tokens, Token{Type: T_LABEL, Args: getArgs(1)})
			case "MAXED":
				tokens = append(tokens, Token{Type: T_MAXED})
			case "MSG_ATTR":
				tokens = append(tokens, Token{Type: T_MSG_ATTR, Args: getArgs(1)})
			case "MSG_CONF":
				tokens = append(tokens, Token{Type: T_MSG_CONF})
			case "MSG_ECHO":
				tokens = append(tokens, Token{Type: T_MSG_ECHO})
			case "MSG_FILEATTACH":
				tokens = append(tokens, Token{Type: T_MSG_FILEATTACH})
			case "MSG_LOCAL":
				tokens = append(tokens, Token{Type: T_MSG_LOCAL})
			case "MSG_MATRIX":
				tokens = append(tokens, Token{Type: T_MSG_MATRIX})
			case "MSG_NEXT":
				tokens = append(tokens, Token{Type: T_MSG_NEXT})
			case "MSG_NOMSGS":
				tokens = append(tokens, Token{Type: T_MSG_NOMSGS})
			case "MSG_NONEW":
				tokens = append(tokens, Token{Type: T_MSG_NONEW})
			case "MSG_NOREAD":
				tokens = append(tokens, Token{Type: T_MSG_NOREAD})
			case "MSG_PRIOR":
				tokens = append(tokens, Token{Type: T_MSG_PRIOR})
			case "NO_KEYPRESS":
				tokens = append(tokens, Token{Type: T_NO_KEYPRESS})
			case "NOCOLOR":
				tokens = append(tokens, Token{Type: T_NOCOLOR})
			case "NOCOLOUR":
				tokens = append(tokens, Token{Type: T_NOCOLOUR})
			case "NORIP":
				tokens = append(tokens, Token{Type: T_NORIP})
			case "NOSTACKED":
				tokens = append(tokens, Token{Type: T_NOSTACKED})
			case "NOTONTODAY":
				tokens = append(tokens, Token{Type: T_NOTONTODAY})
			case "NOVICE":
				tokens = append(tokens, Token{Type: T_NOVICE})
			case "PERMANENT":
				tokens = append(tokens, Token{Type: T_PERMANENT})
			case "REGULAR":
				tokens = append(tokens, Token{Type: T_REGULAR})
			case "RIP":
				tokens = append(tokens, Token{Type: T_RIP})
			case "RIPHASFILE":
				tokens = append(tokens, Token{Type: T_RIPHASFILE})
			case "TAGGED":
				tokens = append(tokens, Token{Type: T_TAGGED})
			case "TOP":
				tokens = append(tokens, Token{Type: T_TOP})
				// Multinode Tokens
			case "APB":
				tokens = append(tokens, Token{Type: T_APB})
			case "CHAT_AVAIL":
				tokens = append(tokens, Token{Type: T_CHAT_AVAIL})
			case "CHAT_NOTAVAIL":
				tokens = append(tokens, Token{Type: T_CHAT_NOTAVAIL})
			case "WHO_IS_ON":
				tokens = append(tokens, Token{Type: T_WHO_IS_ON})
				// RIPscrip Graphics
			case "RIPDISPLAY":
				tokens = append(tokens, Token{Type: T_RIPDISPLAY})
			case "RIPPATH":
				tokens = append(tokens, Token{Type: T_RIPPATH})
			case "RIPSEND":
				tokens = append(tokens, Token{Type: T_RIPSEND})
			case "TEXTSIZE":
				tokens = append(tokens, Token{Type: T_TEXTSIZE, Args: getArgs(2)})
				// Miscellanous Tokens
			case "CKOFF":
				tokens = append(tokens, Token{Type: T_CKOFF})
			case "CKON":
				tokens = append(tokens, Token{Type: T_CKON})
			case "CLEAR_STACKED":
				tokens = append(tokens, Token{Type: T_CLEAR_STACKED})
			case "COMMENT":
				// comments eat everything until the ] so we can assume that
				// you can eat everything until you see an L_STRING
				for {
					if lexTokens[cur+1].Type == L_STRING {
						break
					}
					cur = cur + 1
				}
			case "COPY":
				tokens = append(tokens, Token{Type: T_COPY, Args: getArgs(1)})
			case "DELETE":
				tokens = append(tokens, Token{Type: T_DELETE})
			case "DISPLAY":
				tokens = append(tokens, Token{Type: T_DISPLAY})
			case "DOS":
				tokens = append(tokens, Token{Type: T_DOS})
			case "ENTER":
				tokens = append(tokens, Token{Type: T_ENTER})
			case "HANGUP":
				tokens = append(tokens, Token{Type: T_HANGUP})
			case "IBMCHARS":
				tokens = append(tokens, Token{Type: T_IBMCHARS})
			case "INCLUDE":
				tokens = append(tokens, Token{Type: T_INCLUDE, Args: getArgs(1)})
			case "KEY_POKE":
				tokens = append(tokens, Token{Type: T_KEY_POKE})
			case "LANGUAGE":
				tokens = append(tokens, Token{Type: T_LANGUAGE})
			case "LINK":
				tokens = append(tokens, Token{Type: T_LINK})
			case "LOG":
				tokens = append(tokens, Token{Type: T_LOG})
			case "MENU_CMD":
				tokens = append(tokens, Token{Type: T_MENU_CMD, Args: getArgs(1)})
			case "MENUPATH":
				tokens = append(tokens, Token{Type: T_MENUPATH})
			case "MEX":
				tokens = append(tokens, Token{Type: T_MEX})
			case "MORE":
				tokens = append(tokens, Token{Type: T_MORE})
			case "MOREOFF":
				tokens = append(tokens, Token{Type: T_MOREOFF})
			case "MOREON":
				tokens = append(tokens, Token{Type: T_MOREON})
			case "MSG_CHECKMAIL":
				tokens = append(tokens, Token{Type: T_MSG_CHECKMAIL})
			case "NEWFILES":
				tokens = append(tokens, Token{Type: T_NEWFILES})
			case "ONEXIT":
				tokens = append(tokens, Token{Type: T_ONEXIT})
			case "PAUSE":
				tokens = append(tokens, Token{Type: T_PAUSE})
			case "QUIT":
				tokens = append(tokens, Token{Type: T_QUIT})
			case "QUOTE":
				tokens = append(tokens, Token{Type: T_QUOTE})
			case "REPEAT":
				tokens = append(tokens, Token{Type: T_REPEAT})
			case "REPEATSEQ":
				// should be exactly two tokens after this; one should be
				// a L_STRING and one should be an L_TOKEN that consists
				// of a simple number.
				length := getArgs(1)
				if err != nil ||
					cur+2 >= len(lexTokens) ||
					lexTokens[cur+1].Type != L_STRING ||
					lexTokens[cur+2].Type != L_TOKEN {
					return []Token{}, fmt.Errorf("%w: error processing token %v", err, strings.ToUpper(l.Value))
				}
				args := []string{
					length[0],
					lexTokens[cur+1].Value,
					lexTokens[cur+2].Value,
				}
				cur = cur + 2
				tokens = append(tokens, Token{Type: T_REPEATSEQ, Args: args})
			case "TAG_READ":
				tokens = append(tokens, Token{Type: T_TAG_READ})
			case "TAG_WRITE":
				tokens = append(tokens, Token{Type: T_TAG_WRITE})
			case "TUNE":
				tokens = append(tokens, Token{Type: T_TUNE})
			case "XTERN_DOS":
				tokens = append(tokens, Token{Type: T_XTERN_DOS})
			case "XTERN_ERLVL":
				tokens = append(tokens, Token{Type: T_XTERN_ERLVL})
			case "XTERN_RUN":
				tokens = append(tokens, Token{Type: T_XTERN_RUN})
			default:
				if strings.HasPrefix(l.Value, "/") {
					tokens = append(tokens, Token{Type: T_LABEL, Value: l.Value[1:]})
				} else {
					return []Token{}, fmt.Errorf("%w: %v", ErrUnknownToken, strings.ToUpper(l.Value))
				}
			}
		}

		if err != nil {
			return []Token{}, fmt.Errorf("%w: error processing token %v", err, strings.ToUpper(l.Value))
		}

		cur = cur + 1
	}

	return tokens, nil
}
