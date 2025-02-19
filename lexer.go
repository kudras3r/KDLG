package main

import (
	"strconv"
)

// TODO :
// check for errors
// add scolon replace option !

type tokenCat int

var valid4Scolon = map[int]bool{
	int(IDENT):     true,
	int(INT16LIT):  true,
	int(INT32LIT):  true,
	int(INT64LIT):  true,
	int(STRLIT):    true,
	int(BOOLLIT):   true,
	int(BRK):       true,
	int(RET):       true,
	int(DECR):      true,
	int(INCR):      true,
	int(RBRACE):    true,
	int(RPAREN):    true,
	int(RSBRACKET): true,
}

var idnMap = map[string]tokenCat{
	"func":  FN,
	"ret":   RET,
	"for":   FOR,
	"brk":   BRK,
	"next":  NXT,
	"if":    IF,
	"main":  MAIN,
	"else":  ELSE,
	"x0r":   XOR,
	"class": CLASS,

	"i16": INT16,
	"i32": INT32,
	"i64": INT64,
	"f64": FLOAT64,
	"boo": BOOL,
	"str": STR,
}

func LookupType(idn string) (tokenCat, bool) {
	if iCat, ok := idnMap[idn]; ok {
		return iCat, true
	}
	return IDENT, false

}

const (
	ILLEGAL tokenCat = iota
	EOF

	IDENT

	// types
	BYTE
	INT16
	INT32
	INT64
	FLOAT64
	BOOL
	STR

	// literals
	BYTELIT
	INT16LIT
	INT32LIT
	INT64LIT
	FLOAT64LIT
	BOOLLIT
	STRLIT

	ASSIGN  // =
	FASSIGN // <<=
	PLUS    // +
	MIN     // -
	DIV     // /
	MUL     // *
	REM     // %

	EQ   // ==
	NE   // !=
	GT   // >
	GTE  // >=
	LT   // <
	LTE  // <=
	AND  // &&
	OR   // ||
	INCR // +=
	DECR // -=

	DOT       // .
	COMMA     // ,
	SCOLON    // ;
	LPAREN    // (
	RPAREN    // )
	LBRACE    // {
	RBRACE    // }
	LSBRACKET // [
	RSBRACKET // ]

	// keywords
	MAIN
	FN
	RET
	IF
	ELSE
	FOR
	BRK
	NXT
	XOR
	CLASS
)

type Token struct {
	cat  tokenCat
	ln   int
	ival int64
	fval float64
	sval string
	bval bool
}

type Lexer struct {
	src     string
	srcLen  int
	currPos int
	readPos int
	ch      byte
	row     int
	col     int
	lastT   Token
}

func New(input string) *Lexer {
	l := &Lexer{
		src:    input,
		srcLen: len(input),
	}
	l.readCh()
	return l
}

func (l *Lexer) NextToken() Token {
	var t Token
	l.skipWS()
	switch l.ch {
	case '!':
		if l.peekCh() == '=' {
			l.readCh()
			t = newToken(NE, l.row)
		}
	case '=':
		if l.peekCh() == '=' {
			l.readCh()
			nx := l.peekCh()
			if nx == ' ' || isAlpha(nx) || isNum(nx) {
				t = newToken(EQ, l.row)
			} else {
				t = newToken(ILLEGAL, l.row)
			}
			l.readCh()
		} else {
			t = newToken(ASSIGN, l.row)
		}
	case '+':
		if l.peekCh() == '=' {
			l.readCh()
			t = newToken(INCR, l.row)
		} else {
			t = newToken(PLUS, l.row)
		}
	case '-':
		if l.peekCh() == '=' {
			l.readCh()
			t = newToken(DECR, l.row)
		} else {
			t = newToken(MIN, l.row)
		}
	case '*':
		t = newToken(MUL, l.row)
	case '/':
		t = newToken(DIV, l.row)
	case '%':
		t = newToken(REM, l.row)
	case '<':
		if l.peekCh() == '=' {
			l.readCh()
			t = newToken(LTE, l.row)

		} else if l.peekCh() == '<' {
			l.readCh()
			if l.peekCh() == '=' {
				t = newToken(FASSIGN, l.row)
			}
		} else {
			t = newToken(LT, l.row)
		}
	case '>':
		if l.peekCh() == '=' {
			l.readCh()
			t = newToken(GTE, l.row)
		} else {
			t = newToken(GT, l.row)
		}
	case '&':
		if l.peekCh() == '&' {
			l.readCh()
			t = newToken(AND, l.row)
		}
	case '|':
		if l.peekCh() == '|' {
			l.readCh()
			t = newToken(OR, l.row)
		}
	case ',':
		t = newToken(COMMA, l.row)
	case '.':
		t = newToken(ILLEGAL, l.row)
		for !isWSpace(l.peekCh()) {
			l.readCh()
		}
	// case ';':
	// 	t = newToken(SCOLON, l.row)
	case '(':
		t = newToken(LPAREN, l.row)
	case ')':
		t = newToken(RPAREN, l.row)
	case '{':
		t = newToken(LBRACE, l.row)
	case '}':
		t = newToken(RBRACE, l.row)
	case '[':
		t = newToken(LSBRACKET, l.row)
	case ']':
		t = newToken(RSBRACKET, l.row)
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
		t = newTokenWithVal(STRLIT, l.row, val)
	case '\n':
		ltc := l.lastT.cat
		if valid4Scolon[int(ltc)] {
			t = newToken(SCOLON, l.row)
		}

	case 0:
		t = newToken(EOF, l.row)
	default:
		if isAlpha(l.ch) {
			idn := l.getIdn()
			tCat, ok := LookupType(idn)
			if ok {
				t = newToken(tCat, l.row)
			} else {
				if idn == "true" || idn == "false" {
					tCat = BOOLLIT
				}
				t = newTokenWithVal(tCat, l.row, idn)
			}
			return t
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
				t = newTokenWithVal(INT64LIT, l.row, ns)
			} else {
				t = newTokenWithVal(FLOAT64LIT, l.row, ns)
				if dotsC > 1 {
					t = newToken(ILLEGAL, l.row)
				}
			}
		} else {
			t = newToken(ILLEGAL, l.row)
		}
	}
	l.lastT = t
	l.readCh()
	return t
}

func newToken(cat tokenCat, row int) Token {
	return Token{
		cat: cat,
		ln:  row + 1,
	}
}

func newTokenWithVal(cat tokenCat, row int, val string) Token {
	var t Token

	switch cat {
	case INT64LIT:
		n, _ := strconv.ParseInt(val, 10, 64)
		t.ival = n
	case FLOAT64LIT:
		n, _ := strconv.ParseFloat(val, 64)
		t.fval = n
	case STRLIT:
		t.sval = val
	case BOOLLIT:
		if val == "true" {
			t.bval = true
		}
	default:
		t.sval = val
	}

	t.cat = cat
	t.ln = row + 1

	return t
}

func (l *Lexer) getNum() string {
	f := ""
	for isNum(l.ch) || l.ch == '.' {
		f += string(l.ch)
		l.readCh()
	}
	return f
}

func (l *Lexer) getIdn() string {
	idn := ""
	for isAlpha(l.ch) || isNum(l.ch) {
		idn += string(l.ch)
		l.readCh()
	}
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
	for isWSpace(l.ch) {
		l.readCh()
	}
}

// func (l *Lexer) skipComm() {
// 	for !isNLine(l.ch) {
// 		l.readCh()
// 	}
// }

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
