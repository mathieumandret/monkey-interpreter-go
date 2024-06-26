package lexer

import "interpreter/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	currentChar  byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) atEnd() bool {
	return l.readPosition >= len(l.input)
}

func (l *Lexer) readChar() {
	if l.atEnd() {
		l.currentChar = 0
	} else {
		l.currentChar = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peek() byte {
	if l.atEnd() {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) parseTwoCharsToken(tokenType token.TokenType) token.Token {
	first := l.currentChar
	l.readChar()
	return token.Token{Type: tokenType, Literal: string(first) + string(l.currentChar)}
}

func (l *Lexer) NextToken() token.Token {
	l.skipWhitespace()
	var tok token.Token

	switch l.currentChar {
	case '=':
		if l.peek() == '=' {
			tok = l.parseTwoCharsToken(token.EQ)
		} else {
			tok = newToken(token.ASSIGN, l.currentChar)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.currentChar)
	case '(':
		tok = newToken(token.LPAREN, l.currentChar)
	case ')':
		tok = newToken(token.RPAREN, l.currentChar)
	case '+':
		tok = newToken(token.PLUS, l.currentChar)
	case '-':
		tok = newToken(token.MINUS, l.currentChar)
	case '!':
		if l.peek() == '=' {
			tok = l.parseTwoCharsToken(token.NOT_EQ)
		} else {
			tok = newToken(token.BANG, l.currentChar)
		}
	case '*':
		tok = newToken(token.ASTERISK, l.currentChar)
	case '/':
		tok = newToken(token.SLASH, l.currentChar)
	case '<':
		tok = newToken(token.LT, l.currentChar)
	case '>':
		tok = newToken(token.GT, l.currentChar)
	case ',':
		tok = newToken(token.COMMA, l.currentChar)
	case '{':
		tok = newToken(token.LBRACE, l.currentChar)
	case '}':
		tok = newToken(token.RBRACE, l.currentChar)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		// If the char is not a known symbol it might be the start
		// of an identifier
		if isLetter(l.currentChar) {
			// Literal defines the type, as this can be a keyword
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			// readIdentifier handles moving the current position cursor
			return tok
		} else if isDigit(l.currentChar) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			// If not, then it's an illegal character
			tok = newToken(token.ILLEGAL, l.currentChar)
		}
	}
	// Advance to the next position
	l.readChar()
	return tok
}

func (l *Lexer) readToken(checking_func func(byte) bool) string {
	startpos := l.position
	for checking_func(l.currentChar) {
		l.readChar()
	}
	return l.input[startpos:l.position]
}

func (l *Lexer) readIdentifier() string {
	return l.readToken(isLetter)
}

func (l *Lexer) readNumber() string {
	return l.readToken(isDigit)
}

func (l *Lexer) skipWhitespace() {
	for l.currentChar == ' ' || l.currentChar == '\t' || l.currentChar == '\n' || l.currentChar == '\r' {
		l.readChar()
	}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
