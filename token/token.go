package token;



type Token struct {
	Type string  // helps to distinguish INTEGER,STRING,RIGHT_BRACES or other supported token types
	Literal string // holds the token value, if INTEGER, will be any number, if RIGHT_BRACES literal will be '}'
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	// Identifiers + literals
	IDENTIFIER = "IDENTIFIER" // variable or function names ex. add,foobar,x,y
	INT = "INT" // any integer ex. 123445

	// Math operators
	ASSIGN = "="
	PLUS = "+"

	// Delimiters
	COMMA = ","
	SEMICOLON = ";"

	LPARENTHESIS="("
	RPARENTHESIS=")"
	LBRACE="{"
	RBRACE=""
	
	// Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
)