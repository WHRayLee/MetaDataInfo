package ast

import "SQLParser/token"

type Node interface {
	// TokenLiteral() will be used only for debugging
	// and testing.  returns
	//the literal value of the token it’s associated with.
	TokenLiteral() string
}

// The AST we are going to construct consists solely of Nodes that are connected to
// each other - it’s a tree after all.
type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// This Program node is going to be the root node of every AST our parser produces
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

//---------------------------------------------------------------------------
type LetStatement struct {
	Token token.Token // token.LET
	Name  *Identifier // Name to hold the identifier of the binding
	Value Expression  //  Value for the expression that produces the value.
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// we’ll use Identifier here to represent the name
// in a variable binding and later reuse it, to represent an identifier as part of or as a complete
// expression.
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
