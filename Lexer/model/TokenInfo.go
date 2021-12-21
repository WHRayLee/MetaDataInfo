package model

import "Lexer/common"

type TokenInfo struct {
	tokenType    common.TokenType
	tokenLiteral string
}

func (t *TokenInfo) TokenLiteral() string {
	return t.tokenLiteral
}

func (t *TokenInfo) SetTokenLiteral(tokenLiteral string) {
	t.tokenLiteral = tokenLiteral
}

func (t *TokenInfo) TokenType() common.TokenType {
	return t.tokenType
}

func (t *TokenInfo) SetTokenType(tokenType common.TokenType) {
	t.tokenType = tokenType
}

func NewTokenInfo(tokenType common.TokenType, tokenLiteral string) *TokenInfo {
	return &TokenInfo{tokenType: tokenType, tokenLiteral: tokenLiteral}
}
