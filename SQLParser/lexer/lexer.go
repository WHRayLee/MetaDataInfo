package lexer

import "SQLParser/token"

type Lexer struct {
	input string
	// two “pointers” pointing into our input string is the fact that we will need to be
	// able to “peek” further into the input and look after the current character to see
	// what comes up next.

	// position points to the character in the input that corresponds to the ch byte
	position int // current position in input(points to current char)

	// readPosition always points to the “next” character in the input
	readPosition int  // current reading position in input ( after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.ReadChar()
	return l
}

// The purpose of readChar is to give us the next character and advance our position in the input
// string.
func (l *Lexer) ReadChar() {
	// If that’s
	// the case it sets l.ch to 0, which is the ASCII code for the "NUL" character and signifies either
	// “we haven’t read anything yet” or “end of file” for us.
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		// if we haven’t reached the end of
		// input yet it sets l.ch to the next character by accessing l.input[l.readPosition].
		l.ch = l.input[l.readPosition]
	}
	// After that l.position is updated to the just used l.readPosition and l.readPosition
	// is incremented by one.
	l.position = l.readPosition
	// l.readPosition always points to the next position where we’re going
	// to read from next and l.position always points to the position where we last read
	l.readPosition++
}

// The difficulty of parsing different languages often comes down to how far you have to peek
// ahead (or look backwards!) in the source code to make sense of it
func (l *Lexer) peerChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhiteSpace()
	switch l.ch {
	case '=':
		if l.peerChar() == '=' {
			ch := l.ch
			l.ReadChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peerChar() == '=' {
			ch := l.ch
			l.ReadChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
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
			tok.Type = token.LoopupIndent(tok.Literal)
			return tok
		} else if isDigital(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.ReadChar()
	return tok
}

func isDigital(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.ReadChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigital(l.ch) {
		l.ReadChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	// As you can see, in our case it contains the check ch == '_', which means that
	// we’ll treat _ as a letter and allow it in identifiers and keywords.
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.ReadChar()
	}
}
