package parser

import (
	"SQLParser/ast"
	"SQLParser/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}
	test := []struct{
		expectedIndentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}
	for i, tt := range test {
		stmt := program.Statements[i]
		if !testLetStatements(t, stmt, tt.expectedIndentifier) {
			return
		}
	}
}

func testLetStatements(t *testing.T, stmt ast.Statement, indentifier string) bool {
	if stmt.TokenLiteral() != "let" {
		t.Errorf("stmt.TokenLiteral not 'let'. got=%q", stmt.TokenLiteral())
		return false
	}
	letStmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", letStmt)
		return false
	}
	if letStmt.Name.Value != indentifier {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", indentifier, letStmt.Name.Value)
		return false
	}
	if letStmt.Name.TokenLiteral() != indentifier {
		t.Errorf("stmt.Name not '%s'. got=%s", indentifier, letStmt.Name)
		return false
	}
	return true
}
