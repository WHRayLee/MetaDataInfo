package model

type ParseResult struct {
	sourceSQLString    string
	insertTableInfo    TableBasicInfo
	insertColumnList   []ColumnInfo
	selectTableInfo    TableBasicInfo
	selectColumnList   []ColumnInfo
	joinList           []JoinInfo
	whereConditionList []ConditionInfo
	groupByList        []ColumnInfo
	orderByList        []ColumnInfo
}

func (p *ParseResult) SourceSQLString() string {
	return p.sourceSQLString
}

func (p *ParseResult) InsertTableInfo() TableBasicInfo {
	return p.insertTableInfo
}

func (p *ParseResult) InsertColumnList() []ColumnInfo {
	return p.insertColumnList
}

func (p *ParseResult) SelectTableInfo() TableBasicInfo {
	return p.selectTableInfo
}

func (p *ParseResult) SelectColumnList() []ColumnInfo {
	return p.selectColumnList
}

func (p *ParseResult) JoinList() []JoinInfo {
	return p.joinList
}

func (p *ParseResult) WhereConditionList() []ConditionInfo {
	return p.whereConditionList
}

func (p *ParseResult) GroupByList() []ColumnInfo {
	return p.groupByList
}

func (p *ParseResult) OrderByList() []ColumnInfo {
	return p.orderByList
}

func (p *ParseResult) SetSourceSQLString(sourceSQLString string) {
	p.sourceSQLString = sourceSQLString
}

func (p *ParseResult) SetInsertTableInfo(insertTableInfo TableBasicInfo) {
	p.insertTableInfo = insertTableInfo
}

func (p *ParseResult) SetInsertColumnList(insertColumnList []ColumnInfo) {
	p.insertColumnList = insertColumnList
}

func (p *ParseResult) SetSelectTableInfo(selectTableInfo TableBasicInfo) {
	p.selectTableInfo = selectTableInfo
}

func (p *ParseResult) SetSelectColumnList(selectColumnList []ColumnInfo) {
	p.selectColumnList = selectColumnList
}

func (p *ParseResult) SetWhereConditionList(whereConditionList []ConditionInfo) {
	p.whereConditionList = whereConditionList
}

func (p *ParseResult) SetGroupByList(groupByList []ColumnInfo) {
	p.groupByList = groupByList
}

func (p *ParseResult) SetOrderByList(orderByList []ColumnInfo) {
	p.orderByList = orderByList
}

func (p *ParseResult) SetJoinList(joinList []JoinInfo) {
	p.joinList = joinList
}

func NewParseResult() *ParseResult {
	return &ParseResult{}
}
