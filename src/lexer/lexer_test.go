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
	input := `let five = 5;
	let ten = 10;
	
	let add = fn(x,y){
		x+y;
	};
	
	let result = add(five,ten);`

	tests := []ExpectedToken{
		{token.LET, "let"},
		{token.IDENTIFIER, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "add"},
		{token.FUNCTION, "fn"},
		{token.LPARENTHESIS, "("},
		{token.IDENTIFIER, "x"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "y"},
		{token.RPARENTHESIS, ")"},
		{token.LBRACE, "{"},
		{token.IDENTIFIER, "x"},
		{token.PLUS, "+"},
		{token.IDENTIFIER, "y"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "result"},
		{token.ASSIGN, "="},
		{token.IDENTIFIER, "add"},
		{token.LPARENTHESIS, "("},
		{token.IDENTIFIER, "five"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "10"},
		{token.RPARENTHESIS, ")"},
		{token.SEMICOLON, ";"},
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
