package lexer

import "github.com/agp745/Interpreter-Go/token"

type Lexer struct {
	input string
	position int // current position in input (points to current char)
	readPosition int // current reading position in input (after current char)
	char byte // current char under examination

}

func New(input string) *Lexer {
	lex := &Lexer{input: input}
	lex.readChar()
	return lex
}

//Gives the next character and advances position in the input string.
func (lex *Lexer) readChar() {
	if lex.readPosition >= len(lex.input) {
		lex.char = 0
	} else {
		lex.char = lex.input[lex.readPosition]
	}

	lex.position = lex.readPosition
	lex.readPosition += 1
}

func (lex *Lexer) NextToken() token.Token {
	var tok token.Token

	lex.skipWhitespace()

	switch lex.char {
	case '=':
		if lex.peekChar() == '=' {
			char := lex.char
			lex.readChar()
			literal := string(char) + string(lex.char)
			tok = token.Token{ Type: token.EQ, Literal: literal }
		} else {
			tok = newToken(token.ASSIGN, lex.char)
		}
	case '+':
		tok = newToken(token.PLUS, lex.char)
	case '-':
		tok = newToken(token.MINUS, lex.char)
	case '*':
		tok = newToken(token.ASTERISK, lex.char)
	case '/':
		tok = newToken(token.SLASH, lex.char)
	case '!':
		if lex.peekChar() == '=' {
			char := lex.char
			lex.readChar()
			literal := string(char) + string(lex.char)
			tok = token.Token{ Type: token.NOT_EQ, Literal: literal }
		} else {
			tok = newToken(token.BANG, lex.char)
		}
	case '<':
		tok = newToken(token.LT, lex.char)
	case '>':
		tok = newToken(token.GT, lex.char)
	case ';':
		tok = newToken(token.SEMICOLON, lex.char)
	case ',':
		tok = newToken(token.COMMA, lex.char)
	case '(':
		tok = newToken(token.LPAREN, lex.char)
	case ')':
		tok = newToken(token.RPAREN, lex.char)
	case '{':
		tok = newToken(token.LBRACE, lex.char)
	case '}':
		tok = newToken(token.RBRACE, lex.char)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(lex.char) {
			tok.Literal = lex.readIdentifier()
			tok.Type = token.LookupIdent((tok.Literal))
			return tok
		} else if isDigit(lex.char){
			tok.Type = token.INT
			tok.Literal = lex.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, lex.char)
		}
	}

	lex.readChar()
	return tok
}


func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type: tokenType,
		Literal: string(ch),
	}
}

func (lex *Lexer) readIdentifier() string {
	postion := lex.position
	for isLetter(lex.char) {
		lex.readChar()
	}
	return lex.input[postion:lex.position]
}

func (lex *Lexer) skipWhitespace() {
	for lex.char == ' ' || lex.char == '\t' || lex.char == '\n' || lex.char == '\r' {
		lex.readChar()
	}
}

func (lex *Lexer) readNumber() string {
	position := lex.position
	for isDigit(lex.char) {
		lex.readChar()
	}
	return lex.input[position:lex.position]
}

func (lex *Lexer) peekChar() byte {
	if (lex.readPosition >= len(lex.input)) {
		return 0
	} else {
		return lex.input[lex.readPosition]
	}
}

func isLetter(char byte) bool {
	return 'a' <= char && char >= 'z' || 'A' <= char && char >= 'Z' || char == '_'
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}