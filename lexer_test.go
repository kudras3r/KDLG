package main

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

func TestLexingValidInput(t *testing.T) {
	input := ("asd chr i16 i32 i64 f32 f64 boo str \n" +
		"12 12.1 true #asd#  \n" +
		"= <<= + - / * %\n" +
		"== != < > <= >= \n" +
		"&& || \n" +
		"+= -=\n" +
		". , ( ) { } [ ]\n" +
		"main func ret if else for brk nxt x0r class\n" +
		"\\" +
		";")
	lexer := NewLexer(input)

	var lval yySymType
	var tokens []int
	var token int

	for token != SCOLON {
		token = lexer.Lex(&lval)
		tokens = append(tokens, token)
	}

	validTokens := []int{IDENT, CHR, INT16, INT32, INT64, FLOAT32, FLOAT64, BOOL, STR,
		INT64LIT, FLOAT64LIT, BOOLLIT, STRLIT,
		ASSIGN, FASSIGN, PLUS, MIN, DIV, MUL, REM,
		EQ, NE, LT, GT, LTE, GTE,
		AND, OR,
		INCR, DECR,
		DOT, COMMA, LPAREN, RPAREN, LBRACE, RBRACE, LSBRACKET, RSBRACKET,
		MAIN, FN, RET, IF, ELSE, FOR, BRK, NXT, XOR, CLASS,
		UNK,
		SCOLON}

	for i := 0; i < 5; i++ {
		if tokens[i] != validTokens[i] {
			t.Errorf("invalid token at pos %d", i)
		}
	}
}

func TestWhitespaceHandling(t *testing.T) {
	input := "   func   main()   {}   "
	lexer := NewLexer(input)

	var lval yySymType
	token := lexer.Lex(&lval)

	if token != FN {
		t.Errorf("Expected token type %d after whitespace, got %d", FN, token)
	}
}
