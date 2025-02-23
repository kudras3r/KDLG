// Code generated by goyacc -o parser.go parser1.y. DO NOT EDIT.

//line parser1.y:2
package main

import __yyfmt__ "fmt"

//line parser1.y:2

//line parser1.y:5
type yySymType struct {
	yys int
	Str string
	Int int
	Boo bool
	Flt float64
}

const (
	ILLEGAL tokenCat = iota
	EOF

	// IDENT

	// // types
	// BYTE
	// INT16
	// INT32
	// INT64
	// FLOAT64
	// BOOL
	// STR

	// // literals
	// BYTELIT
	// INT16LIT
	// INT32LIT
	// INT64LIT
	// FLOAT64LIT
	// BOOLLIT
	// STRLIT

	// ASSIGN  // =
	FASSIGN // <<=
	PLUS    // +
	MIN     // -
	DIV     // /
	MUL     // *
	REM     // %

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

	DOT       // .
	COMMA     // ,
	// SCOLON    // ;
	LPAREN    // (
	RPAREN    // )
	LBRACE    // {
	RBRACE    // }
	LSBRACKET // [
	RSBRACKET // ]

	// keywords
	// MAIN
	// FN
	// RET
	// IF
	// ELSE
	// FOR
	// BRK
	// NXT
	// XOR
	// CLASS
)

const BYTELIT = 57346
const INT16LIT = 57347
const INT32LIT = 57348
const INT64LIT = 57349
const FLOAT64LIT = 57350
const BOOLLIT = 57351
const STRLIT = 57352
const IDENT = 57353
const BYTE = 57354
const INT16 = 57355
const INT32 = 57356
const INT64 = 57357
const FLOAT64 = 57358
const BOOL = 57359
const STR = 57360
const MAIN = 57361
const FN = 57362
const RET = 57363
const IF = 57364
const ELSE = 57365
const FOR = 57366
const BRK = 57367
const NXT = 57368
const XOR = 57369
const CLASS = 57370
const ETERN = 57371
const INCR = 57372
const DECR = 57373
const EQ = 57374
const NE = 57375
const GT = 57376
const GTE = 57377
const LT = 57378
const LTE = 57379
const AND = 57380
const OR = 57381
const ASSIGN = 57382
const SCOLON = 57383

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"BYTELIT",
	"INT16LIT",
	"INT32LIT",
	"INT64LIT",
	"FLOAT64LIT",
	"BOOLLIT",
	"STRLIT",
	"IDENT",
	"BYTE",
	"INT16",
	"INT32",
	"INT64",
	"FLOAT64",
	"BOOL",
	"STR",
	"MAIN",
	"FN",
	"RET",
	"IF",
	"ELSE",
	"FOR",
	"BRK",
	"NXT",
	"XOR",
	"CLASS",
	"ETERN",
	"INCR",
	"DECR",
	"EQ",
	"NE",
	"GT",
	"GTE",
	"LT",
	"LTE",
	"AND",
	"OR",
	"ASSIGN",
	"SCOLON",
	"'.'",
	"','",
	"'['",
	"']'",
	"'('",
	"')'",
	"'{'",
	"'}'",
	"';'",
	"'-'",
	"'!'",
	"'*'",
	"'/'",
	"'%'",
	"'+'",
	"'<'",
	"'>'",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 23

var yyAct = [...]int8{
	22, 19, 17, 14, 18, 16, 11, 3, 4, 5,
	6, 7, 8, 9, 15, 20, 12, 10, 13, 2,
	1, 0, 21,
}

var yyPact = [...]int16{
	-5, -1000, 3, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-37, -1000, -1000, -39, -43, -1000, 4, -1000, 3, -45,
	-1000, -43, -1000,
}

var yyPgo = [...]int8{
	0, 20, 19, 18, 17, 16, 3, 16, 16, 16,
	16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
	16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
	16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
	16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
	16, 16,
}

var yyR1 = [...]int8{
	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	4, 4, 5, 3, 3, 6, 6, 7, 8, 8,
	11, 11, 10, 9, 12, 12, 13, 13, 14, 14,
	15, 17, 16, 16, 16, 16, 16, 16, 16, 16,
	16, 18, 25, 25, 25, 21, 22, 23, 23, 30,
	30, 31, 24, 32, 32, 32, 33, 33, 34, 34,
	35, 35, 19, 20, 36, 36, 36, 36, 37, 37,
	37, 37, 28, 40, 40, 39, 39, 38, 27, 27,
	27, 27, 41, 41, 42, 42, 42, 43, 43, 43,
	43, 44, 44, 44, 45, 45, 45, 45, 46, 46,
	47, 47, 47, 48, 48, 49, 49, 29, 29, 26,
	50, 50, 51, 51, 51,
}

var yyR2 = [...]int8{
	0, 3, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 3, 1, 3, 1, 3, 6, 1, 3,
	1, 3, 2, 3, 1, 0, 1, 2, 1, 1,
	2, 2, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 2, 1, 1, 1, 5, 7, 6, 8, 1,
	2, 2, 9, 1, 1, 0, 1, 0, 1, 0,
	1, 3, 2, 3, 1, 3, 1, 1, 1, 1,
	1, 1, 4, 1, 3, 1, 0, 3, 4, 4,
	6, 6, 1, 1, 2, 2, 1, 1, 3, 3,
	3, 1, 3, 3, 1, 1, 1, 1, 1, 3,
	1, 3, 3, 1, 3, 1, 3, 1, 1, 3,
	1, 1, 1, 1, 1,
}

var yyChk = [...]int16{
	-1000, -1, -2, 12, 13, 14, 15, 16, 17, 18,
	-4, 11, -5, -3, -6, 11, 42, 41, 43, 44,
	11, -6, 45,
}

var yyDef = [...]int8{
	0, -2, 0, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 0, 13, 15, 0, 1, 0, 0,
	12, 14, 16,
}

var yyTok1 = [...]int8{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 52, 3, 3, 3, 55, 3, 3,
	46, 47, 53, 56, 43, 51, 42, 54, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 50,
	57, 3, 58, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 44, 3, 45, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 48, 3, 49,
}

var yyTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
}

var yyTok3 = [...]int8{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 5
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := int(yyPact[state])
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && int(yyChk[int(yyAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || int(yyExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := int(yyExca[i])
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = int(yyTok1[0])
		goto out
	}
	if char < len(yyTok1) {
		token = int(yyTok1[char])
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = int(yyTok2[char-yyPrivate])
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = int(yyTok3[i+0])
		if token == char {
			token = int(yyTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(yyTok2[1]) /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = int(yyPact[yystate])
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = int(yyAct[yyn])
	if int(yyChk[yyn]) == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = int(yyDef[yystate])
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && int(yyExca[xi+1]) == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = int(yyExca[xi+0])
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = int(yyExca[xi+1])
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = int(yyPact[yyS[yyp].yys]) + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = int(yyAct[yyn]) /* simulate a shift of "error" */
					if int(yyChk[yystate]) == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= int(yyR2[yyn])
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = int(yyR1[yyn])
	yyg := int(yyPgo[yyn])
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = int(yyAct[yyg])
	} else {
		yystate = int(yyAct[yyj])
		if int(yyChk[yystate]) != -yyn {
			yystate = int(yyAct[yyg])
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	}
	goto yystack /* stack new state and value */
}
