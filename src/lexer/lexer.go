package lexer

import (
	"log"

	"github.com/claranceliberi/monkey-interpreter/src/token"
)

// struct that supports reading source codew
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

// reading one character at a time from input (source code)
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

// think of this as a method that return next character that `readChar` is going to pick
// or as a next character following current character in `l.ch`
func (l *Lexer) peekChar() byte {

	if l.readPosition > len(l.input){
		return 0
	}else{
		return l.input[l.readPosition]
	}
}

// skips whitespaces,tabs or new lines found in source code.
func (l *Lexer) skipWhiteSpace() {

	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}



// return the token type of the current char in lexer
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			l.readChar()
			tok.Literal = "=="
			tok.Type = token.EQ
		}else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch) 
	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			tok.Literal = "!="
			tok.Type = token.NOT_EQ

		}else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch) 
	case '*':
		tok = newToken(token.ASTERISK, l.ch) 
	case '<':
		tok = newToken(token.LT, l.ch) 
	case '>':
		tok = newToken(token.GT, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPARENTHESIS, l.ch)
	case ')':
		tok = newToken(token.RPARENTHESIS, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
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
			log.Println(tok.Literal)
			tok.Type = token.LookupIdentifier(tok.Literal)
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

// return the whole identifier from source code. this might be a keyword or variable name
// ex. `let`,`fn`,`variable_name`, `foo`, `bar`
func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

// read the whole number from source code
// TODO: for now only integers are supported, in future should support all types of numbers
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

// helps to form Token struct
func newToken(tokenType token.TokenType, literal byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(literal)}
}
