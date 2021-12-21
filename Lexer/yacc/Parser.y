%{
package yacc

import "Lexer/model"

func setResult(l yyLexer, v model.ParseResult) {
	l.(*Lexer).result = v
}
%}
%union {
	str string
	strs []string
	tpParseResult model.ParseResult
	tpTableBasicInfo model.TableBasicInfo
	tpColumnList []model.ColumnInfo
	tpJoinList []model.JoinInfo
	tpConditionList []model.ConditionInfo
	tpConditionInfo model.ConditionInfo
	tpJoinInfo model.JoinInfo
	tpColumnInfo model.ColumnInfo
}
%token<str> Identifier AS INSERT INTO FROM SELECT LEFT OUTER JOIN RIGHT INNER CROSS NATURAL ON AND eq WHERE GROUP BY SUM AVG COUNT MAX MIN ASTERISK
%type <str> AliasOpt joinTypeOpt leftCondition rightCondition  OnAndOpt
%type <tpColumnList> insertColumnListStmt insertColumnFiledList selectColumnListStmt groupByList orderByList ColumnListOpt groupByListStmt
%type <tpParseResult> InsertSelectStmt
%type <tpTableBasicInfo> insertTableInfoStmt basicTableInfoStmt selectTableInfoStmt
%type <tpJoinList> joinList
%type <tpConditionList> whereConditionList ConditionList whereConditionListInfo
%type <tpConditionInfo> ConditionStmtOpt whereConditionStmtOpt
%type <tpColumnInfo> groupByStmtOpt ColumnField
%type <tpJoinInfo> joinStmtOpt

%left '('
%right ')'

%start Start
%%
Start:
	InsertSelectStmt
	{
		setResult(yylex, $1)
	}
InsertSelectStmt: insertTableInfoStmt insertColumnListStmt selectColumnListStmt selectTableInfoStmt joinList whereConditionList groupByList orderByList
	{
		parseResult := model.NewParseResult()
		parseResult.SetInsertTableInfo($1)
		parseResult.SetInsertColumnList($2)
		parseResult.SetSelectTableInfo($4)
		parseResult.SetSelectColumnList($3)
		parseResult.SetJoinList($5)
		parseResult.SetWhereConditionList($6)
		parseResult.SetGroupByList($7)
		parseResult.SetOrderByList($8)
		$$ = *parseResult
	}
insertTableInfoStmt:
	{
		tableBasicInfo := model.NewTableBasicInfo()
		$$ = *tableBasicInfo
	}
|	INSERT INTO basicTableInfoStmt
	{
		$$ = $3
	}

selectTableInfoStmt:
	FROM basicTableInfoStmt
	{
		$$ = $2
	}
|
	{
		tableBasicInfo := model.NewTableBasicInfo()
                $$ = *tableBasicInfo
	}

basicTableInfoStmt:
	Identifier '.' Identifier '.' Identifier AliasOpt
	{
		tableBasicInfo := model.NewTableBasicInfo()
		tableBasicInfo.SetDBName($1)
		tableBasicInfo.SetSchemaName($3)
		tableBasicInfo.SetTableName($5)
		tableBasicInfo.SetTableShortName($6)
		$$ = *tableBasicInfo
	}
|	Identifier '.' Identifier AliasOpt
	{
		tableBasicInfo := model.NewTableBasicInfo()
		tableBasicInfo.SetSchemaName($1)
		tableBasicInfo.SetTableName($3)
		tableBasicInfo.SetTableShortName($4)
                $$ = *tableBasicInfo
	}
|	Identifier AliasOpt
	{
		tableBasicInfo := model.NewTableBasicInfo()
		tableBasicInfo.SetTableName($1)
		tableBasicInfo.SetTableShortName($2)
                $$ = *tableBasicInfo
	}

AliasOpt:
	{
		$$ = ""
	}
|	Identifier
	{
		$$ = $1
	}
|	AS Identifier
	{
		$$ = $2
	}
insertColumnListStmt:
	{
		columnInfo := model.NewColumnInfo()
		columnInfo.SetColumnName("*")
		$$ = []model.ColumnInfo{*columnInfo}
	}
|	'(' insertColumnFiledList ')'
	{
		$$ = $2
	}

insertColumnFiledList:
	Identifier
	{
		columnInfo := model.NewColumnInfo()
		columnInfo.SetColumnName($1)
		$$ = []model.ColumnInfo{*columnInfo}
	}
|	insertColumnFiledList ',' Identifier
	{
		columnInfo := model.NewColumnInfo()
		columnInfo.SetColumnName($3)
		$$ = append($1, *columnInfo)
	}
selectColumnListStmt:
	{
		columnInfo := model.NewColumnInfo()
		$$ = []model.ColumnInfo{*columnInfo}
	}
|	SELECT ColumnListOpt
	{
		$$ = $2
	}

ColumnListOpt:
	{
		columnInfo := model.NewColumnInfo()
		$$ = []model.ColumnInfo{*columnInfo}
	}
|
	ASTERISK
	{
		columnInfo := model.NewColumnInfo()
		columnInfo.SetColumnName("*")
		$$ = []model.ColumnInfo{*columnInfo}
	}
|
	ColumnField
	{
		$$ = []model.ColumnInfo{$1}
	}
|	ColumnListOpt ',' ColumnField
	{
                $$ = append($1, $3)
	}

ColumnField:
	Identifier AliasOpt
	{
		columnInfo := model.NewColumnInfo()
		columnInfo.SetColumnName($1)
		columnInfo.SetColumnAliasName($2)
		$$ = *columnInfo
	}
|	Identifier '.' Identifier AliasOpt
	{
		columnInfo := model.NewColumnInfo()
		columnInfo.SetShortTableName($1)
                columnInfo.SetColumnName($3)
                columnInfo.SetColumnAliasName($4)
                $$ = *columnInfo
	}
|	AggStmt '(' Identifier '.' Identifier ')' AliasOpt
	{
        	columnInfo := model.NewColumnInfo()
        	columnInfo.SetShortTableName($3)
                columnInfo.SetColumnName($5)
                columnInfo.SetColumnAliasName($7)
                $$ = *columnInfo
        }
|	AggStmt '(' Identifier '.' ASTERISK ')' AliasOpt
	{
        	columnInfo := model.NewColumnInfo()
        	columnInfo.SetShortTableName($3)
                columnInfo.SetColumnName($5)
                columnInfo.SetColumnAliasName($7)
                $$ = *columnInfo
        }
|	AggStmt '(' Identifier ')' AliasOpt
	{
        	columnInfo := model.NewColumnInfo()
                columnInfo.SetColumnName($3)
                columnInfo.SetColumnAliasName($5)
                $$ = *columnInfo
        }
|	AggStmt '(' ASTERISK ')' AliasOpt
	{
        	columnInfo := model.NewColumnInfo()
                columnInfo.SetColumnName($3)
                columnInfo.SetColumnAliasName($5)
                $$ = *columnInfo
        }
AggStmt:
	SUM
|	AVG
|	MAX
|	COUNT
|	MIN

joinList:
	joinList joinStmtOpt
	{
		$$ = append($1, $2)
	}
|	joinStmtOpt
	{
		$$ = []model.JoinInfo{$1}
	}

joinStmtOpt:
	joinTypeOpt basicTableInfoStmt ConditionList
        {
        	joinInfo := model.NewJoinInfo()
        	joinInfo.SetJoinType($1)
        	tableBasicInfo := model.NewTableBasicInfo()
        	tableBasicInfo = &$2
                joinInfo.SetJoinTableInfo(*tableBasicInfo)
                joinInfo.SetJoinConditionList($3)
   		$$ = *joinInfo
        }

joinTypeOpt:
	LEFT OUTER JOIN
	{
		$$ = "LeftOuterJoin"
	}
|	LEFT JOIN
	{
		$$ = "LeftJoin"
	}
|	RIGHT OUTER JOIN
	{
		$$ = "RightOuterJoin"
	}
|	RIGHT JOIN
	{
		$$ = "RightJoin"
	}
|	INNER JOIN
	{
		$$ = "InnerJoin"
	}
|	CROSS JOIN
	{
		$$ = "CrossJoin"
	}
|	NATURAL JOIN
	{
		$$ = "NaturalJoin"
	}

ConditionList:
	{
		conditionInfo := model.NewConditionInfo()
		conditionInfo.SetLeftColumnInfo("1")
		conditionInfo.SetRightColumnInfo("1")
		conditionInfo.SetSign("=")
		$$ = []model.ConditionInfo{*conditionInfo}
	}
|	ConditionList ConditionStmtOpt
	{
		$$ = append($1, $2)
	}

ConditionStmtOpt:
	OnAndOpt leftCondition eq rightCondition
	{
		conditionInfo := model.NewConditionInfo()
         	conditionInfo.SetLeftColumnInfo($2)
         	conditionInfo.SetRightColumnInfo($4)
         	conditionInfo.SetSign("=")
		$$ = *conditionInfo
	}

OnAndOpt:
	ON
|	AND



leftCondition:
	Identifier

rightCondition:
	Identifier

whereConditionList:
	{
		conditionInfo := model.NewConditionInfo()
                conditionInfo.SetLeftColumnInfo("1")
                conditionInfo.SetRightColumnInfo("1")
                conditionInfo.SetSign("=")
                $$ = []model.ConditionInfo{*conditionInfo}
	}
|	WHERE whereConditionListInfo
	{
		$$ = $2
	}
whereConditionListInfo:
	whereConditionStmtOpt
	{
		$$ = []model.ConditionInfo{$1}
	}
|	whereConditionListInfo whereConditionStmtOpt
	{
		$$ = append($1, $2)
	}

whereConditionStmtOpt:
	leftCondition eq rightCondition
	{
		conditionInfo := model.NewConditionInfo()
                conditionInfo.SetLeftColumnInfo($1)
                conditionInfo.SetRightColumnInfo($3)
                conditionInfo.SetSign("=")
                $$ = *conditionInfo
	}
|	AND leftCondition eq rightCondition
	{
		conditionInfo := model.NewConditionInfo()
                conditionInfo.SetLeftColumnInfo($2)
                conditionInfo.SetRightColumnInfo($4)
                conditionInfo.SetSign("=")
                $$ = *conditionInfo
	}


groupByList:
	{
		columnInfo := model.NewColumnInfo()
                $$ = []model.ColumnInfo{*columnInfo}
	}
|	GROUP BY groupByListStmt
	{
		$$ = $3
	}
groupByListStmt:
	groupByListStmt ',' groupByStmtOpt
	{
		$$ = append($1, $3)
	}
|	groupByStmtOpt
	{
		$$ = []model.ColumnInfo{$1}
	}
groupByStmtOpt:
	Identifier
	{
		columnInfo := model.NewColumnInfo()
		columnInfo.SetColumnName($1)
		$$ = *columnInfo
	}
orderByList:
	{
		columnInfo := model.NewColumnInfo()
		$$ = []model.ColumnInfo{*columnInfo}
	}
%%