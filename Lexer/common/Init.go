package common

func init() {
	GlobalKeyWords["ILLEGAL"] = ILLEGAL
	GlobalKeyWords["EOF"] = EOF
	// Identifiers + literals
	GlobalKeyWords["INT"] = INT
	GlobalKeyWords["IDENT"] = IDENT
	// Operators
	GlobalKeyWords["!"] = BANG
	GlobalKeyWords["-"] = MINUS
	GlobalKeyWords["+"] = PLUS
	GlobalKeyWords["!="] = NOT_EQ
	GlobalKeyWords["="] = EQ
	GlobalKeyWords["<"] = GT
	GlobalKeyWords[">"] = LT
	GlobalKeyWords["/"] = SLASH
	GlobalKeyWords["*"] = ASTERISK
	// Delimiters
	GlobalKeyWords["}"] = RBRACE
	GlobalKeyWords["{"] = LBRACE
	GlobalKeyWords[")"] = RPAREN
	GlobalKeyWords["("] = LPAREN
	GlobalKeyWords[";"] = SEMICOLON
	GlobalKeyWords[","] = COMMA
}
