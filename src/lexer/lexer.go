package lexer

import "github.com/claranceliberi/monkey-interpreter/src/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to the current char)
	readPosition int  // current reading position in input (after current car)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// TODO: for now only ASCII is supported in lexer, add support for Unicode and UTF-8
func (l *Lexer) readChar() {

	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

/* return the token type of the current char in lexer
 */
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPARENTHESIS, l.ch)
	case ')':
		tok = newToken(token.RPARENTHESIS, l.ch) 
	case ',':
		tok = newToken(token.COMMA, l.ch) 
	case '+':
		tok = newToken(token.PLUS, l.ch) 
	case '{':
		tok = newToken(token.LBRACE, l.ch)
		case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF

	}

	l.readChar()
	return tok;
}

func newToken(tokenType token.TokenType, literal byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(literal)}
}
