package main

import (
	"testing"
)

func TestLexerInitialization(t *testing.T) {
	input := "i64\nasd <<= *\n  123 123.3  true @@    asd \nkek "
	lexer := New(input)

	if lexer == nil {
		t.Fatalf("Expected lexer to be initialized, got nil")
	}
}

func TestTokenGeneration(t *testing.T) {
	input := "i64 <<= 123 1.2 @@ comment\nboo true"
	lexer := New(input)

	tests := []struct {
		expectedCat tokenCat
		expectedStr string
		expectedLn  int
	}{
		{INT64, "", 1},
		{FASSIGN, "", 1},
		{INT64LIT, "123", 1},
		{FLOAT64LIT, "1.2", 1},
		{BOOL, "", 2},
		{BOOLLIT, "true", 2},
		{EOF, "", 2},
	}

	for _, test := range tests {
		tok := lexer.NextToken()
		if tok.cat != test.expectedCat {
			t.Errorf("Expected token category %d, got %d", test.expectedCat, tok.cat)
		}
		if tok.str != test.expectedStr {
			t.Errorf("Expected token string '%s', got '%s'", test.expectedStr, tok.str)
		}
	}
}

func TestCommentHandling(t *testing.T) {
	input := "i64 <<= 123 @@ this is a comment\n"
	lexer := New(input)

	for {
		tok := lexer.NextToken()
		if tok.cat == EOF {
			break
		}
		if tok.cat == ILLEGAL {
			t.Errorf("Expected no illegal tokens, got one")
		}
	}
}

func TestEndOfInput(t *testing.T) {
	input := ""
	lexer := New(input)

	tok := lexer.NextToken()
	if tok.cat != EOF {
		t.Errorf("Expected EOF token, got %d", tok.cat)
	}
}
