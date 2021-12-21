package yacc

import (
	"fmt"
	"testing"
)

func TestLexer(t *testing.T) {
	sql := "insert into cvcxv.cvcxv.cxvcxv(c,b,c,cvxcvcxvxc) select a.cola" +
		", count(b.*) as tagtest from db.schema.table tab " +
		"CROSS JOIN db.schema.taba on a = b and c = d " +
		"CROSS JOIN dbb.schemab.tabab on e = h and g = h " +
		"inner join dbd.tabless on addd = cvxv and xcvcx = cvxcv " +
		" where c = d AND AAA = BBBB AND AJ = BH GROUP BY a,b,c,d"
	lexer := NewLexer(sql)
	yyParse(lexer)
	fmt.Println(lexer.(*Lexer).result.InsertColumnList())
	fmt.Println(lexer.(*Lexer).result.InsertTableInfo())
	fmt.Println(lexer.(*Lexer).result.SelectColumnList())
	fmt.Println(lexer.(*Lexer).result.SelectTableInfo())
	fmt.Println(lexer.(*Lexer).result.JoinList())
	fmt.Println(lexer.(*Lexer).result.WhereConditionList())
	fmt.Println(lexer.(*Lexer).result.GroupByList())
	for _, err := range lexer.(*Lexer).errs {
		fmt.Println(err.Error())
	}
}
