package lexer

import (
	"fmt"
	"log"

	"github.com/claranceliberi/monkey-interpreter/src/token"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to the current char)
	readPosition int  // current reading position in input (after current car)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	log.Println("in new")
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// TODO: for now only ASCII is supported in lexer, add support for Unicode and UTF-8
func (l *Lexer) readChar() {

	fmt.Printf("char : %v", l.ch)
	fmt.Println()

	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) skipWhiteSpace() {

	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}

/* return the token type of the current char in lexer
 */
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

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
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			log.Printf(" is Letter %v", tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.ch) {
		l.readNumber()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {

	position := l.position

	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func newToken(tokenType token.TokenType, literal byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(literal)}
}
