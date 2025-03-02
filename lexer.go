package main

import "fmt"

// TODO :
// auto scol

var idnMap = map[string]Token{
	"func":        FN,
	"ret":         RET,
	"for":         FOR,
	"brk":         BRK,
	"next":        NXT,
	"if":          IF,
	"main":        MAIN,
	"else":        ELSE,
	"x0r":         XOR,
	"class":       CLASS,
	"constructor": CONSTRUCT,

	"chr": CHR,
	"i16": INT16,
	"i32": INT32,
	"i64": INT64,
	"f32": FLOAT32,
	"f64": FLOAT64,
	"boo": BOOL,
	"str": STR,
}

func LookupType(idn string) (Token, bool) {
	if iCat, ok := idnMap[idn]; ok {
		return iCat, true
	}
	return IDENT, false
}

const (
	EOF = iota
	ILLEGAL

	// IDENT

	//types
	// CHR
	// INT16
	// INT32
	// INT64
	// FLOAT64
	// BOOL
	// STR

	// literals
	BYTELIT
	// INT16LIT
	// INT32LIT
	// // INT64LIT
	// FLOAT64LIT
	// BOOLLIT
	// STRLIT

	// ASSIGN  // =
	// FASSIGN // <<=
	// PLUS    // +
	// MIN // -
	// DIV     // /
	// MUL     // *
	// REM // %

	// EQ   // ==
	// NE   // !=
	// GT   // >
	// GTE  // >=
	// LT   // <
	// LTE  // <=
	// AND  // &&
	// OR   // ||
	// INCR // +=
	// DECR // -=

	// DOT   // .
	// COMMA // ,
	// SCOLON    // ;
	// LPAREN    // (
	// RPAREN    // )
	// LBRACE    // {
	// RBRACE    // }
	// LSBRACKET // [
	// RSBRACKET // ]

	// keywords
	MAIN
	// FN
	// RET
	// IF
	// ELSE
	// FOR
	// BRK
	NXT
	XOR
	// CLASS
)

type Token int

type Lexer struct {
	src     string
	srcLen  int
	currPos int
	readPos int
	ch      byte
	row     int
	col     int
}

func New(input string) *Lexer {
	l := &Lexer{
		src:    input,
		srcLen: len(input),
	}
	l.readCh()
	return l
}

func (l *Lexer) Error(s string) {
	fmt.Println(s)
}

func (l *Lexer) Lex(lval *yySymType) int {
	var t Token
	l.skipWS()

	switch l.ch {
	case '@':
		if l.peekCh() == '@' {
			l.readCh()
			for !isNLine(l.ch) && !l.atEOF() {
				l.readCh()
			}
			return l.Lex(lval)
		}
	case '!':
		if l.peekCh() == '=' {
			l.readCh()
			t = NE
		} else {
			t = NOT
		}
	case '=':
		if l.peekCh() == '=' {
			l.readCh()
			t = EQ
		} else {
			t = ASSIGN
		}
	case '+':
		if l.peekCh() == '=' {
			l.readCh()
			t = INCR
		} else {
			t = PLUS
		}
	case '-':
		if l.peekCh() == '=' {
			l.readCh()
			t = DECR
		} else {
			t = MIN
		}
	case '*':
		t = MUL
	case '/':
		t = DIV
	case '%':
		t = REM
	case '<':
		if l.peekCh() == '=' {
			l.readCh()
			t = LTE

		} else if l.peekCh() == '<' {
			l.readCh()
			if l.peekCh() == '=' {
				l.readCh()
				t = FASSIGN
			}
		} else {
			t = LT
		}
	case '>':
		if l.peekCh() == '=' {
			l.readCh()
			t = GTE
		} else {
			t = GT
		}
	case '&':
		if l.peekCh() == '&' {
			l.readCh()
			t = AND
		}
	case '|':
		if l.peekCh() == '|' {
			l.readCh()
			t = OR
		}
	case ',':
		t = COMMA
	case '.':
		t = ILLEGAL
		for !isWSpace(l.peekCh()) {
			l.readCh()
		}
	case ';':
		t = SCOLON
	case '(':
		t = LPAREN
	case ')':
		t = RPAREN
	case '{':
		t = LBRACE
	case '}':
		t = RBRACE
	case '[':
		t = LSBRACKET
	case ']':
		t = RSBRACKET
	case '#':
		startPos := l.currPos + 1
		for {
			l.readCh()
			if l.ch == '#' {
				l.readCh()
				break
			}
			if l.ch == 0 {
				break
			}
		}
		val := l.src[startPos : l.currPos-1]
		t = STRLIT
		lval.S = val
	case '\n':
		l.readCh()
	case 0:
		t = 0
	default:
		if isAlpha(l.ch) {
			idn := l.getIdn()
			tCat, ok := LookupType(idn)
			if ok {
				t = tCat
			} else {
				if idn == "true" || idn == "false" {
					tCat = BOOLLIT
				}
				lval.S = idn
				t = tCat
			}
		} else if isNum(l.ch) {
			ns := l.getNum()
			isInt, dotsC := true, 0
			for _, r := range ns {
				if r == '.' {
					isInt = false
					dotsC++
				}
			}
			if isInt {
				lval.S = ns
				t = INT64LIT
			} else {
				lval.S = ns
				t = FLOAT64LIT
				if dotsC > 1 {
					t = EOF
				}
			}
		} else {
			t = EOF
		}
	}
	l.readCh()

	return int(t)
}

func (l *Lexer) getNum() string {
	f := ""
	for isNum(l.ch) || l.ch == '.' {
		f += string(l.ch)
		l.readCh()
	}
	l.readPos--
	return f
}

func (l *Lexer) getIdn() string {
	idn := ""
	for isAlpha(l.ch) || isNum(l.ch) {
		idn += string(l.ch)
		l.readCh()
	}
	l.readPos--
	return idn
}

func (l *Lexer) readCh() {
	if l.atEOF() {
		l.ch = 0
		return
	}

	if isNLine(l.ch) {
		l.row++
	}

	l.ch = l.src[l.readPos]
	l.currPos = l.readPos
	l.readPos++
	l.col++
}

func (l *Lexer) peekCh() byte {
	if l.readPos > l.srcLen {
		return 0
	}
	return l.src[l.readPos]
}

func (l *Lexer) skipWS() {
	for isWSpace(l.ch) || isNLine(l.ch) {
		l.readCh()
	}
}

func isNum(c byte) bool {
	return c >= '0' && c <= '9'
}

func isAlpha(c byte) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z'
}

func (l *Lexer) atEOF() bool {
	return l.readPos >= l.srcLen
}

func isWSpace(c byte) bool {
	return c == ' '
}

func isNLine(c byte) bool {
	return c == '\n'
}
