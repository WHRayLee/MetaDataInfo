package yacc

import (
	"Lexer/common"
	"Lexer/model"
	"errors"
	"fmt"
	"strings"
)

type Lexer struct {
	input        []byte
	position     int
	readPosition int
	ch           byte
	errs         []error
	result       model.ParseResult
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) NextToken() model.TokenInfo {
	l.skipWhiteSpace()
	var tok model.TokenInfo
	switch l.ch {
	case '+', '-', '/', '{', '}', '(', ')', '.', ',':
		tok.SetTokenLiteral(string(l.ch))
		if tokenTypeInfo, ok := common.GlobalKeyWords[string(l.ch)]; ok {
			tok.SetTokenType(tokenTypeInfo)
		} else {
			tok.SetTokenType(common.ILLEGAL)
		}
		l.readChar()
	case '=':
		tok.SetTokenLiteral("eq")
		if tokenTypeInfo, ok := common.GlobalKeyWords[string(l.ch)]; ok {
			tok.SetTokenType(tokenTypeInfo)
		} else {
			tok.SetTokenType(common.ILLEGAL)
		}
		l.readChar()
	case '*':
		tok.SetTokenLiteral("ASTERISK")
		if tokenTypeInfo, ok := common.GlobalKeyWords[string(l.ch)]; ok {
			tok.SetTokenType(tokenTypeInfo)
		} else {
			tok.SetTokenType(common.ILLEGAL)
		}
		l.readChar()
	case '!':
		peekCh := l.peekChar()
		switch peekCh {
		case '=':
			currentCh := l.ch
			l.readChar()
			tok.SetTokenLiteral(string(currentCh) + string(l.ch))
			tok.SetTokenLiteral(string(common.GlobalKeyWords[string(currentCh)+string(l.ch)]))
		default:
			tok.SetTokenLiteral(string(l.ch))
			tok.SetTokenType(common.GlobalKeyWords[string(l.ch)])
		}
	case 0:
		tok.SetTokenLiteral("")
		tok.SetTokenType(common.EOF)
	default:
		if l.isLetter() {
			identifier := l.readIdentifier()
			identifier = strings.ToLower(identifier)
			tok.SetTokenLiteral(identifier)
			tok.SetTokenType(common.IDENT)
		} else if l.isDigital() {
			number := l.readNumber()
			tok.SetTokenLiteral(number)
			// todo 区分整数和浮点数
			tok.SetTokenType(common.INT)
		} else {
			tok.SetTokenType(common.ILLEGAL)
			tok.SetTokenLiteral(string(l.ch))
		}
	}

	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for l.isLetter() {
		l.readChar()
	}
	return string(l.input[position:l.position])
}

func (l *Lexer) readNumber() string {
	position := l.position
	for l.isFloat() {
		l.readChar()
	}
	return string(l.input[position:l.position])
}

func (l *Lexer) isLetter() bool {
	return 'a' <= l.ch && l.ch <= 'z' || 'A' <= l.ch && l.ch <= 'Z' || l.ch == '_'
}

func (l *Lexer) isDigital() bool {
	return '0' <= l.ch && l.ch <= '9'
}

func (l *Lexer) isFloat() bool {
	return '0' <= l.ch && l.ch <= '9' && l.ch == '.'
}

func (l *Lexer) Lex(lval *yySymType) int {
	// 跳过空白字符
	l.skipWhiteSpace()
	// 获取下一个token
	token := l.NextToken()
	//
	fmt.Println("token:", token)
	tokenToLower := strings.ToLower(token.TokenLiteral())
	switch tokenToLower {
	case "":
		return 0
	case "*", ";", ",", "(", ".", ")", "'", "=":
		lval.str = tokenToLower
		return int([]byte(tokenToLower)[0])
	case "select":
		lval.str = tokenToLower
		return SELECT
	case "insert":
		lval.str = tokenToLower
		return INSERT
	case "into":
		lval.str = tokenToLower
		return INTO
	case "from":
		lval.str = tokenToLower
		return FROM
	case "as":
		lval.str = tokenToLower
		return AS
	case "right":
		lval.str = tokenToLower
		return RIGHT
	case "inner":
		lval.str = tokenToLower
		return INNER
	case "cross":
		lval.str = tokenToLower
		return CROSS
	case "natural":
		lval.str = tokenToLower
		return NATURAL
	case "left":
		lval.str = tokenToLower
		return LEFT
	case "outer":
		lval.str = tokenToLower
		return OUTER
	case "join":
		lval.str = tokenToLower
		return JOIN
	case "on":
		lval.str = tokenToLower
		return ON
	case "and":
		lval.str = tokenToLower
		return AND
	case "eq":
		lval.str = tokenToLower
		return eq
	case "asterisk":
		lval.str = tokenToLower
		return ASTERISK
	case "where":
		lval.str = tokenToLower
		return WHERE
	case "group":
		lval.str = tokenToLower
		return GROUP
	case "by":
		lval.str = tokenToLower
		return BY
	case "sum":
		lval.str = tokenToLower
		return SUM
	case "max":
		lval.str = tokenToLower
		return MAX
	case "min":
		lval.str = tokenToLower
		return MIN
	case "count":
		lval.str = tokenToLower
		return COUNT
	case "avg":
		lval.str = tokenToLower
		return AVG
	default:
		lval.str = tokenToLower
		return Identifier
	}
}

func (l *Lexer) Error(s string) {
	err := errors.New(s)
	l.errs = append(l.errs, err)
}

func NewLexer(input string) ILexer {
	l := &Lexer{
		input: []byte(input),
	}
	l.readChar()
	return l
}
