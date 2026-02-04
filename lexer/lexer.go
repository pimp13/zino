package lexer

import (
	"github.com/pimp13/zino/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func (l *Lexer) isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readString() string {
	l.readChar()
	pos := l.position
	for l.ch != '"' && l.ch != 0 {
		l.readChar()
	}
	str := l.input[pos:l.position]
	l.readChar()
	return str
}

func (l *Lexer) readIdentifier() string {
	pos := l.position
	for l.isLetter(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func (l *Lexer) readNumber() string {
	pos := l.position
	for l.isDigit(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = l.newToken(token.ASSIGN, l.ch)
	case '+':
		tok = l.newToken(token.PLUS, l.ch)
	case '-':
		tok = l.newToken(token.MINUS, l.ch)
	case '*':
		tok = l.newToken(token.MULTIPLY, l.ch)
	case '/':
		tok = l.newToken(token.DIVIDE, l.ch)
	case '(':
		tok = l.newToken(token.LEFT_PAREN, l.ch)
	case ')':
		tok = l.newToken(token.RIGHT_PAREN, l.ch)
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case '\n':
		tok = l.newToken(token.NEWLINE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if l.isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.IDENT
			return tok
		} else if l.isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.NUMBER
			return tok
		} else {
			tok = l.newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}
