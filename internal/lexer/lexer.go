package lexer

import (
	"C-Compiler/internal/token"
)

type Lexer struct {
	source  string
	pos     int
	readPos int
	ch      byte
}

func New(code string) *Lexer {
	l := &Lexer{source: code}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPos >= len(l.source) {
		l.ch = 0
	} else {
		l.ch = l.source[l.readPos]
	}

	l.pos = l.readPos

	l.readPos++
}

func (l *Lexer) readIdentifier() string {
	pos := l.pos
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.source[pos:l.pos]
}

func (l *Lexer) readNumber() string {
	pos := l.pos
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.source[pos:l.pos]
}
func (l *Lexer) skipSpaces() {
	for l.ch == '\n' || l.ch == '\t' || l.ch == ' ' {
		l.readChar()

	}
}
func (l *Lexer) peekChar() byte {
	if l.readPos >= len(l.source) {
		return 0
	} else {
		return l.source[l.readPos]
	}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func newTokenLiteral(tokenType token.TokenType, ch string) token.Token {
	return token.Token{Type: tokenType, Literal: ch}
}

func (l *Lexer) NextToken() token.Token {

	l.skipSpaces()

	var tok token.Token
	switch l.ch {

	case '=':
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: "=="} // If token is == returning Token EQ, with literal "=="
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '&':
		tok = newToken(token.AMPERSAND, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		if l.peekChar() == '/' {
			l.readChar()
			literal := l.readSingleLineComment()
			tok = newTokenLiteral(token.COMMENT, literal)
		} else if l.peekChar() == '*' {
			l.readChar()
			literal := l.readMultiLineComment()
			tok = newTokenLiteral(token.COMMENT, literal)
		} else {
			tok = newToken(token.SLASH, l.ch)
		}

	case '>':
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.Token{Type: token.GTE, Literal: ">="}
		} else {
			tok = newToken(token.RIGHT, l.ch)
		}

	case '<':
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.Token{Type: token.LTE, Literal: "<="}
		} else {
			tok = newToken(token.LEFT, l.ch)
		}
	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: "!="}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case 0:

		tok.Literal = ""
		tok.Type = token.EOF

	default:

		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookForIdent(tok.Literal)

			//check if data type is Pointer

			if tok.Type != token.IDENT && l.ch == '*' {

				tmp, ok := token.LookForDataTypers(tok.Type)
				if ok {
					l.readChar()
					tok.Type = tmp
					tok.Literal += "*"
				}
			}
			return tok
		} else {
			if isDigit(l.ch) {
				tok.Type = token.INT
				tok.Literal = l.readNumber()
				return tok
			} else {

				tok = newToken(token.ILLEGAL, l.ch)
			}
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) readSingleLineComment() string {
	l.readChar() // delete extra /
	pos := l.pos
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}

	return l.source[pos:l.pos]
}

func (l *Lexer) readMultiLineComment() string {
	l.readChar() //delete extra *
	pos := l.pos
	for {
		if l.ch == 0 {

			break
		}
		if l.ch == '*' && l.peekChar() == '/' {
			l.readChar()
			l.readChar()
			break
		}
		l.readChar()
	}

	return l.source[pos : l.pos-2] // l pos -2 to not count */
}

func isLetter(b byte) bool {
	if (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || b == '_' {
		return true
	} else {
		return false
	}
}

func isDigit(b byte) bool {
	return (b >= '0' && b <= '9')
}
