package yacc

import (
	"Lexer/model"
)

type ILexer interface {
	yyLexer
	NextToken() model.TokenInfo
	skipWhiteSpace()
	readChar()
	peekChar() byte
}
