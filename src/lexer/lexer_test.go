package lexer

import (
	"testing"

	"github.com/claranceliberi/monkey-interpreter/src/token"
)

type ExpectedToken struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func TestNextToken(t *testing.T) {
	input := "=+(){},;"

	tests := []ExpectedToken{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPARENTHESIS, "("},
		{token.RPARENTHESIS, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	_lexer := New(input)

	for i, expectedToken := range tests {
		token := _lexer.NextToken()

		if token.Type != expectedToken.expectedType {
			t.Fatalf("tests [%d] - tokentype wrong. expected=%q, got=%q", i, expectedToken.expectedType, token.Type)
		}

		if token.Literal != expectedToken.expectedLiteral {
			t.Fatalf("tests [%d] - literal wrong. expected=%q, got=%q", i, expectedToken.expectedLiteral, token.Literal)
		}
	}

}
