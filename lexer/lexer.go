package lexer

// TODO :
// auto scol

const (
	EOF   = iota // 0
	UNK          // \ ...
	IDENT        // asd

	// types
	INT32   // i32
	INT64   // i64
	FLOAT32 // f32
	FLOAT64 // f64
	BOOL    // boo
	STR     // str

	// literals
	INT32LIT   // 12
	INT64LIT   // 12
	FLOAT32LIT // 12.1
	FLOAT64LIT // 12.1
	BOOLLIT    // true
	STRLIT     // "golang"

	ASSIGN // =
	EQ     // ==
	PLUS   // +
	ADDOP  // +=
	MIN    // -
	MINOP  // -=
	DIV    // /
	DIVOP  // /=
	MUL    // *
	MULOP  // *=
	REM    // %
	REMOP  // %=
	POW    // ^
	POWOP  // ^=

	LSH  // <<
	RSH  // >>
	GT   // >
	GTE  // >=
	LT   // <
	LTE  // <=
	LNOT // !
	NE   // !=
	LAND // &&
	LOR  // ||
	BAND // &
	BOR  // |
	INCR // ++
	DECR // --

	DOT       // .
	COMMA     // ,
	SCOLON    // ;
	COLON     // :
	LPAREN    // (
	RPAREN    // )
	LBRACE    // {
	RBRACE    // }
	LSBRACKET // [
	RSBRACKET // ]

	// keywords
	MAIN      // main
	FN        // func
	RET       // return
	IF        // if
	ELSE      // else
	FOR       // for
	ETERN     // etern
	BRK       // break
	NXT       // next
	XOR       // x0r
	CLASS     // class
	CONSTRUCT // constructor
	DESTRUCT  // destructor
	STRUCT    // struct
	VAR       // var
)

var idnMap = map[string]TokenCat{
	"func":        FN,
	"return":      RET,
	"for":         FOR,
	"break":       BRK,
	"next":        NXT,
	"if":          IF,
	"main":        MAIN,
	"else":        ELSE,
	"x0r":         XOR,
	"class":       CLASS,
	"constructor": CONSTRUCT,
	"destructor":  DESTRUCT,
	"etern":       ETERN,
	"true":        BOOLLIT,
	"false":       BOOLLIT,
	"struct":      STRUCT,
	"var":         VAR,

	"i32": INT32,
	"i64": INT64,
	"f32": FLOAT32,
	"f64": FLOAT64,
	"boo": BOOL,
	"str": STR,
}

func LookupСat(idn string) TokenCat {
	if iCat, ok := idnMap[idn]; ok {
		return iCat
	}
	return IDENT
}

type TokenCat int

type Token struct {
	kind TokenCat
	val  string
}

type Lexer struct {
	src     string
	srcLen  int
	currPos int
	readPos int
	ch      byte
	row     int
	col     int
}

func NewLexer(input string) *Lexer {
	l := &Lexer{
		src:    input,
		srcLen: len(input),
	}
	l.readCh()
	return l
}

func (l *Lexer) Lex() Token {
	var tok Token
	l.skipWS()

	switch l.ch {
	case '@':
		if l.peekCh() == '@' {
			l.readCh()
			for !isNLine(l.ch) && !l.atEOF() {
				l.readCh()
			}
			return l.Lex()
		}
	case '!':
		tok = l.checkOp(LNOT)
	case '=':
		tok = l.checkOp(ASSIGN)
	case '+':
		if l.peekCh() == '+' {
			l.readCh()
			tok = newToken(INCR)
		} else {
			tok = l.checkOp(PLUS)
		}
	case '-':
		if l.peekCh() == '-' {
			l.readCh()
			tok = newToken(DECR)
		} else {
			tok = l.checkOp(MIN)
		}
	case '*':
		tok = l.checkOp(MUL)
	case '/':
		tok = l.checkOp(DIV)
	case '%':
		tok = l.checkOp(REM)
	case '<':
		tok = l.checkOp(LT)
	case '>':
		tok = l.checkOp(GT)
	case '&':
		if l.peekCh() == '&' {
			l.readCh()
			tok = newToken(LAND)
		} else {
			tok = newToken(BAND)
		}
	case '|':
		if l.peekCh() == '|' {
			l.readCh()
			tok = newToken(LOR)
		} else {
			tok = newToken(BOR)
		}
	case ',':
		tok = newToken(COMMA)
	case '.':
		tok = newToken(DOT)
	case ';':
		tok = newToken(SCOLON)
	case ':':
		tok = newToken(COLON)
	case '(':
		tok = newToken(LPAREN)
	case ')':
		tok = newToken(RPAREN)
	case '{':
		tok = newToken(LBRACE)
	case '}':
		tok = newToken(RBRACE)
	case '[':
		tok = newToken(LSBRACKET)
	case ']':
		tok = newToken(RSBRACKET)
	case '"':
		startPos := l.currPos + 1
		for {
			l.readCh()
			if l.ch == '"' {
				l.readCh()
				break
			}
			if l.ch == 0 {
				break
			}
		}
		val := l.src[startPos : l.currPos-1]
		tok = newTokenWithVal(STRLIT, val)
		l.readPos--
		l.col--
	case '\n':
		l.col = 1
		l.readCh()
	case 0:
		tok = newToken(EOF)
	default:
		if isAlpha(l.ch) {
			idn := l.getIdn()
			tok = newTokenWithVal(LookupСat(idn), idn)
		} else if isNum(l.ch) {
			ns, isInt := l.getNum()
			if isInt {
				tok = newTokenWithVal(INT64LIT, ns)
			} else {
				tok = newTokenWithVal(FLOAT64LIT, ns)
			}
		} else {
			tok = newToken(UNK)
		}
	}
	l.readCh()

	return tok
}

func newToken(cat TokenCat) Token {
	return Token{
		kind: cat,
	}
}

func newTokenWithVal(cat TokenCat, val string) Token {
	return Token{
		kind: cat,
		val:  val,
	}
}

func (l *Lexer) checkOp(cat TokenCat) Token {
	var tok Token

	if l.peekCh() != '=' {
		tok = newToken(cat)
	} else {
		l.readCh()
		tok = newToken(cat + 1)
	}

	return tok
}

func (l *Lexer) getNum() (string, bool) {
	n := ""
	dotsC := 0
	for isNum(l.ch) || l.ch == '.' {
		if l.ch == '.' {
			dotsC++
		}
		if dotsC > 1 {
			break
		}
		n += string(l.ch)
		l.readCh()
	}
	l.readPos--
	l.col--
	if dotsC == 1 {
		return n, false
	}
	return n, true
}

func (l *Lexer) getIdn() string {
	idn := ""
	for isAlpha(l.ch) || isNum(l.ch) {
		idn += string(l.ch)
		l.readCh()
	}
	l.readPos--
	l.col--
	return idn
}

func (l *Lexer) readCh() {
	if l.atEOF() {
		l.ch = 0
		return
	}

	if isNLine(l.ch) {
		l.row++
		l.col = 0
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
