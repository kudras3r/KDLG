package main

// TODO :
// add line field to token ?
// add types

type tokenCat int

var iMap = map[string]tokenCat{
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
}

func LookupType(idn string) (tokenCat, bool) {
	if iCat, ok := iMap[idn]; ok {
		return iCat, true
	}
	return IDENT, false

}

const (
	ILLEGAL tokenCat = iota
	EOF

	IDENT
	BYTE
	INT16
	INT32
	INT64
	FLOAT
	BOOL
	BOOLLIT
	STR

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
	cat tokenCat
	col int
	val string
}

type Lexer struct {
	sourse    string
	sourseLen int
	currPos   int
	readPos   int
	ch        byte
	row       int
	col       int
}

func New(input string) *Lexer {
	l := &Lexer{
		sourse:    input,
		sourseLen: len(input),
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
			t = newToken(NE, l.col)
		}
	case '=':
		if l.peekCh() == '=' {
			l.readCh()
			t = newToken(EQ, l.col)
		} else {
			t = newToken(ASSIGN, l.col)
		}
	case '+':
		if l.peekCh() == '=' {
			l.readCh()
			t = newToken(INCR, l.col)
		} else {
			t = newToken(PLUS, l.col)
		}
	case '-':
		if l.peekCh() == '=' {
			l.readCh()
			t = newToken(DECR, l.col)
		} else {
			t = newToken(MIN, l.col)
		}
	case '*':
		t = newToken(MUL, l.col)
	case '/':
		t = newToken(DIV, l.col)
	case '%':
		t = newToken(REM, l.col)
	case '<':
		if l.peekCh() == '=' {
			l.readCh()
			t = newToken(LTE, l.col)

		} else if l.peekCh() == '<' {
			l.readCh()
			if l.peekCh() == '=' {
				t = newToken(FASSIGN, l.col)
			}
		} else {
			t = newToken(LT, l.col)
		}
	case '>':
		if l.peekCh() == '=' {
			l.readCh()
			t = newToken(GTE, l.col)
		} else {
			t = newToken(GT, l.col)
		}
	case '&':
		if l.peekCh() == '&' {
			l.readCh()
			t = newToken(AND, l.col)
		}
	case '|':
		if l.peekCh() == '|' {
			l.readCh()
			t = newToken(OR, l.col)
		}
	case ',':
		t = newToken(COMMA, l.col)
	case '.':
		t = newToken(DOT, l.col)
	case ';':
		t = newToken(SCOLON, l.col)
	case '(':
		t = newToken(LPAREN, l.col)
	case ')':
		t = newToken(RPAREN, l.col)
	case '{':
		t = newToken(LBRACE, l.col)
	case '}':
		t = newToken(RBRACE, l.col)
	case '[':
		t = newToken(LSBRACKET, l.col)
	case ']':
		t = newToken(RSBRACKET, l.col)
	case 0:
		t = newToken(EOF, l.col)
	default:
		if isAlpha(l.ch) {
			idn := l.getIdn()
			tCat, ok := LookupType(idn)
			if ok {
				t = newToken(tCat, l.col)
			} else {
				t = newTokenWithVal(tCat, l.col, idn)
			}
			return t
		} else if isNum(l.ch) {
			// TODO : float ints...
			n := l.getInt()
			t = newTokenWithVal(INT64, l.col, n)
		} else {
			t = newToken(ILLEGAL, l.col)
		}

	}
	l.readCh()
	return t
}

func newToken(cat tokenCat, col int) Token {
	return Token{
		cat: cat,
		col: col,
	}
}

func newTokenWithVal(cat tokenCat, col int, val string) Token {
	return Token{
		cat: cat,
		col: col,
		val: val,
	}
}

func (l *Lexer) getFloat() string {
	f := ""
	for isNum(l.ch) || l.ch == '.' {
		f += string(l.ch)
	}
	return f
}

func (l *Lexer) getInt() string {
	n := ""
	for isNum(l.ch) {
		n += string(l.ch)
		l.readCh()
	}
	return n
}

func (l *Lexer) getIdn() string {
	idn := ""
	for isAlpha(l.ch) {
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

	l.ch = l.sourse[l.readPos]
	l.currPos = l.readPos
	l.readPos++
	l.col++
}

func (l *Lexer) peekCh() byte {
	if l.readPos > l.sourseLen {
		return 0
	}
	return l.sourse[l.readPos]
}

func (l *Lexer) skipWS() {
	for isWSpace(l.ch) {
		l.readCh()
	}
}

func (l *Lexer) skipComm() {
	for !isNLine(l.ch) {
		l.readCh()
	}
}

func isNum(c byte) bool {
	return c >= '0' && c <= '9'
}

func isAlpha(c byte) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z'
}

func isAlphaNum(c byte) bool {
	return isAlpha(c) || isNum(c)
}

func (l *Lexer) atEOF() bool {
	return l.readPos >= l.sourseLen
}

func isWSpace(c byte) bool {
	return c == ' '
}

func isNLine(c byte) bool {
	return c == '\n'
}
