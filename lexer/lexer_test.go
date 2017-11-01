package lexer

import (
	"fmt"
	"github.com/smazumder05/sabini/token"
	"testing"
)

func TestNextToken(t *testing.T) {

	input := "=+(){},;"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LBRACE, "("},
		{token.RBRACE, ")"},
		{token.LPAREN, "{"},
		{token.RPAREN, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lex := New(input)

	fmt.Printf("Current input: %s", input)
	for i, tt := range tests {
		tok := lex.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - token type is wrong, expected= %q, got = %q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - Literal is wrong, expected= %q, got = %q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
