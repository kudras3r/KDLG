package main

type tokenCat int

const (
	ILLEGAL tokenCat = iota
	EOF

	IDENT
	INT
	// TODO int32... float...

	ASSIGN
	// TODO fast assign
	PLUS // +
	MIN  // -
	DIV  // /
	MUL  // *

	EQ  // ==
	NE  // !=
	GT  // >
	GTE // >=
	LT  // <
	LTE // <=

	COMMA  // ,
	SCOLON // ;
	LPAREN // (
	RPAREN // )
	LBRACE // {
	RBRACE // }

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
)

type token struct {
	cat    tokenCat
	lineno int
	lex    string
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
	return &Lexer{
		sourse:    input,
		sourseLen: len(input),
	}
}

func newToken(cat tokenCat) token {
	return token{
		cat: cat,
	}
}

func newTokenWithVal(cat tokenCat, lex string) token {
	return token{
		cat: cat,
		lex: lex,
	}
}

func (l *Lexer) getNum() string {
	// TODO float read
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
