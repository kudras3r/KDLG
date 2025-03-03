package lexer

import (
	"testing"
)

func TestNewLexer(t *testing.T) {
	input := "func main() {}"
	lexer := NewLexer(input)

	if lexer == nil {
		t.Fatalf("Expected a new lexer instance, got nil")
	}
}

func TestLexingCats(t *testing.T) {
	input := ("i32 i64 f32 f64 boo str  \n" +
		"= == + += - -= / /= * *= % %= \n" +
		"!= < > <= >= \n" +
		"&& || & | \n" +
		"++ -- \n" +
		". : , ( ) { } [ ]\n" +
		"main func return if else for break next x0r class var\n" +
		"\\ ;")
	lexer := NewLexer(input)

	var tokens []Token
	var token Token

	for token.kind != SCOLON {
		token = lexer.Lex()
		tokens = append(tokens, token)
	}

	validCats := []TokenCat{INT32, INT64, FLOAT32, FLOAT64, BOOL, STR,
		ASSIGN, EQ, PLUS, ADDOP, MIN, MINOP, DIV, DIVOP, MUL, MULOP, REM, REMOP,
		NE, LT, GT, LTE, GTE,
		LAND, LOR, BAND, BOR,
		INCR, DECR,
		DOT, COLON, COMMA, LPAREN, RPAREN, LBRACE, RBRACE, LSBRACKET, RSBRACKET,
		MAIN, FN, RET, IF, ELSE, FOR, BRK, NXT, XOR, CLASS, VAR,
		UNK,
	}

	i := 0
	for i < len(validCats) {
		if tokens[i].kind != validCats[i] {
			t.Errorf("invalid token at pos %d token: %d %s expected: %d", i, token.kind, token.val, validCats[i])
		}
		i++
	}
}

func TestLexingVals(t *testing.T) {
	input := (`12 12.1 true "strstr" ; `)
	lexer := NewLexer(input)

	var tokens []Token
	var token Token

	for token.kind != SCOLON {
		token = lexer.Lex()
		tokens = append(tokens, token)
	}

	validTokens := []Token{
		newTokenWithVal(INT64LIT, "12"),
		newTokenWithVal(FLOAT64LIT, "12.1"),
		newTokenWithVal(BOOLLIT, "true"),
		newTokenWithVal(STRLIT, "strstr"),
	}

	i := 0
	for i < len(validTokens) {
		if tokens[i].kind != validTokens[i].kind || tokens[i].val != validTokens[i].val {
			t.Errorf("invalid token cat or val at pos %d token: %d %s expected: %d %s",
				i, token.kind, token.val,
				validTokens[i].kind, validTokens[i].val)
		}

		i++
	}
}

func TestWhitespaceHandling(t *testing.T) {
	input := "   func   main()   {}   "
	lexer := NewLexer(input)

	token := lexer.Lex()

	if token.kind != FN {
		t.Errorf("Expected token type %d after whitespace, got %d", FN, token.kind)
	}
}
