package main

import "fmt"

// TODO :

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

// const (
// 	ILLEGAL tokenCat = iota
// 	EOF

// 	IDENT

// 	// types
// 	BYTE
// 	INT16
// 	INT32
// 	INT64
// 	FLOAT64
// 	BOOL
// 	STR

// 	// literals
// 	BYTELIT
// 	INT16LIT
// 	INT32LIT
// 	INT64LIT
// 	FLOAT64LIT
// 	BOOLLIT
// 	STRLIT

// 	ASSIGN  // =
// 	FASSIGN // <<=
// 	PLUS    // +
// 	MIN     // -
// 	DIV     // /
// 	MUL     // *
// 	REM     // %

// 	EQ   // ==
// 	NE   // !=
// 	GT   // >
// 	GTE  // >=
// 	LT   // <
// 	LTE  // <=
// 	AND  // &&
// 	OR   // ||
// 	INCR // +=
// 	DECR // -=

// 	DOT       // .
// 	COMMA     // ,
// 	SCOLON    // ;
// 	LPAREN    // (
// 	RPAREN    // )
// 	LBRACE    // {
// 	RBRACE    // }
// 	LSBRACKET // [
// 	RSBRACKET // ]

// 	// keywords
// 	MAIN
// 	FN
// 	RET
// 	IF
// 	ELSE
// 	FOR
// 	BRK
// 	NXT
// 	XOR
// 	CLASS
// )

type Token struct {
	cat tokenCat
	ln  int
	col int
	str string
	// ival int64
	// fval float64
	// sval string
	// bval bool
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
			t = newToken(NE, l.row, l.col)
		}
	case '=':
		if l.peekCh() == '=' {
			l.readCh()
			t = newToken(EQ, l.row, l.col)
		} else {
			t = newToken(ASSIGN, l.row, l.col)
		}
	case '+':
		if l.peekCh() == '=' {
			l.readCh()
			t = newToken(INCR, l.row, l.col)
		} else {
			t = newToken(PLUS, l.row, l.col)
		}
	case '-':
		if l.peekCh() == '=' {
			l.readCh()
			t = newToken(DECR, l.row, l.col)
		} else {
			t = newToken(MIN, l.row, l.col)
		}
	case '*':
		t = newToken(MUL, l.row, l.col)
	case '/':
		t = newToken(DIV, l.row, l.col)
	case '%':
		t = newToken(REM, l.row, l.col)
	case '<':
		if l.peekCh() == '=' {
			l.readCh()
			t = newToken(LTE, l.row, l.col)

		} else if l.peekCh() == '<' {
			l.readCh()
			if l.peekCh() == '=' {
				l.readCh()
				t = newToken(FASSIGN, l.row, l.col)
			}
		} else {
			t = newToken(LT, l.row, l.col)
		}
	case '>':
		if l.peekCh() == '=' {
			l.readCh()
			t = newToken(GTE, l.row, l.col)
		} else {
			t = newToken(GT, l.row, l.col)
		}
	case '&':
		if l.peekCh() == '&' {
			l.readCh()
			t = newToken(AND, l.row, l.col)
		}
	case '|':
		if l.peekCh() == '|' {
			l.readCh()
			t = newToken(OR, l.row, l.col)
		}
	case ',':
		t = newToken(COMMA, l.row, l.col)
	case '.':
		t = newToken(ILLEGAL, l.row, l.col)
		for !isWSpace(l.peekCh()) {
			l.readCh()
		}
	case ';':
		t = newToken(SCOLON, l.row, l.col)
	case '(':
		t = newToken(LPAREN, l.row, l.col)
	case ')':
		t = newToken(RPAREN, l.row, l.col)
	case '{':
		t = newToken(LBRACE, l.row, l.col)
	case '}':
		t = newToken(RBRACE, l.row, l.col)
	case '[':
		t = newToken(LSBRACKET, l.row, l.col)
	case ']':
		t = newToken(RSBRACKET, l.row, l.col)
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
		t = newTokenWithVal(STRLIT, l.row, l.col, val)
	case '\n':
		l.readCh()
	case 0:
		t = newToken(EOF, l.row, l.col)
	default:
		if isAlpha(l.ch) {
			idn := l.getIdn()
			tCat, ok := LookupType(idn)
			if ok {
				t = newToken(tCat, l.row, l.col)
			} else {
				if idn == "true" || idn == "false" {
					tCat = BOOLLIT
				}
				lval.Str = idn
				t = newTokenWithVal(tCat, l.row, l.col, idn)
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
				lval.Str = ns
				t = newTokenWithVal(INT64LIT, l.row, l.col, ns)
			} else {
				lval.Str = ns
				t = newTokenWithVal(FLOAT64LIT, l.row, l.col, ns)
				if dotsC > 1 {
					t = newToken(ILLEGAL, l.row, l.col)
				}
			}
		} else {
			t = newToken(ILLEGAL, l.row, l.col)
		}
	}
	l.lastT = t
	l.readCh()

	return int(t.cat)
}

func (l *Lexer) NextToken() Token {
	var t Token
	l.skipWS()

	switch l.ch {
	case '@':
		if l.peekCh() == '@' {
			l.readCh()
			for !isNLine(l.ch) && !l.atEOF() {
				l.readCh()
			}
			return l.NextToken()
		}
	case '!':
		if l.peekCh() == '=' {
			l.readCh()
			t = newToken(NE, l.row, l.col)
		}
	case '=':
		if l.peekCh() == '=' {
			l.readCh()
			t = newToken(EQ, l.row, l.col)
		} else {
			t = newToken(ASSIGN, l.row, l.col)
		}
	case '+':
		if l.peekCh() == '=' {
			l.readCh()
			t = newToken(INCR, l.row, l.col)
		} else {
			t = newToken(PLUS, l.row, l.col)
		}
	case '-':
		if l.peekCh() == '=' {
			l.readCh()
			t = newToken(DECR, l.row, l.col)
		} else {
			t = newToken(MIN, l.row, l.col)
		}
	case '*':
		t = newToken(MUL, l.row, l.col)
	case '/':
		t = newToken(DIV, l.row, l.col)
	case '%':
		t = newToken(REM, l.row, l.col)
	case '<':
		if l.peekCh() == '=' {
			l.readCh()
			t = newToken(LTE, l.row, l.col)

		} else if l.peekCh() == '<' {
			l.readCh()
			if l.peekCh() == '=' {
				l.readCh()
				t = newToken(FASSIGN, l.row, l.col)
			}
		} else {
			t = newToken(LT, l.row, l.col)
		}
	case '>':
		if l.peekCh() == '=' {
			l.readCh()
			t = newToken(GTE, l.row, l.col)
		} else {
			t = newToken(GT, l.row, l.col)
		}
	case '&':
		if l.peekCh() == '&' {
			l.readCh()
			t = newToken(AND, l.row, l.col)
		}
	case '|':
		if l.peekCh() == '|' {
			l.readCh()
			t = newToken(OR, l.row, l.col)
		}
	case ',':
		t = newToken(COMMA, l.row, l.col)
	case '.':
		t = newToken(ILLEGAL, l.row, l.col)
		for !isWSpace(l.peekCh()) {
			l.readCh()
		}
	case ';':
		t = newToken(SCOLON, l.row, l.col)
	case '(':
		t = newToken(LPAREN, l.row, l.col)
	case ')':
		t = newToken(RPAREN, l.row, l.col)
	case '{':
		t = newToken(LBRACE, l.row, l.col)
	case '}':
		t = newToken(RBRACE, l.row, l.col)
	case '[':
		t = newToken(LSBRACKET, l.row, l.col)
	case ']':
		t = newToken(RSBRACKET, l.row, l.col)
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
		t = newTokenWithVal(STRLIT, l.row, l.col, val)
	case '\n':
		l.readCh()
	case 0:
		t = newToken(EOF, l.row, l.col)
	default:
		if isAlpha(l.ch) {
			idn := l.getIdn()
			tCat, ok := LookupType(idn)
			if ok {
				t = newToken(tCat, l.row, l.col)
			} else {
				if idn == "true" || idn == "false" {
					tCat = BOOLLIT
				}
				t = newTokenWithVal(tCat, l.row, l.col, idn)
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
				t = newTokenWithVal(INT64LIT, l.row, l.col, ns)
			} else {
				t = newTokenWithVal(FLOAT64LIT, l.row, l.col, ns)
				if dotsC > 1 {
					t = newToken(ILLEGAL, l.row, l.col)
				}
			}
		} else {
			t = newToken(ILLEGAL, l.row, l.col)
		}
	}
	l.lastT = t
	l.readCh()

	return t
}

func newToken(cat tokenCat, row, col int) Token {
	return Token{
		cat: cat,
		ln:  row + 1,
		col: col + 1,
	}
}

func newTokenWithVal(cat tokenCat, row, col int, val string) Token {
	var t Token

	// switch cat {
	// case INT64LIT:
	// 	n, _ := strconv.ParseInt(val, 10, 64)
	// 	t.ival = n
	// case FLOAT64LIT:
	// 	n, _ := strconv.ParseFloat(val, 64)
	// 	t.fval = n
	// case STRLIT:
	// 	t.sval = val
	// case BOOLLIT:
	// 	if val == "true" {
	// 		t.bval = true
	// 	}
	// default:
	// 	t.sval = val
	// }

	t.str = val
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
	for isWSpace(l.ch) || isNLine(l.ch) {
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
