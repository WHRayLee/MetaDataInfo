package model

type ColumnInfo struct {
	shortTableName  string
	columnName      string
	columnAliasName string
	columnComment   string
}

func (c *ColumnInfo) SetColumnComment(columnComment string) {
	c.columnComment = columnComment
}

func (c *ColumnInfo) SetColumnAliasName(columnAliasName string) {
	c.columnAliasName = columnAliasName
}

func (c *ColumnInfo) SetColumnName(columnName string) {
	c.columnName = columnName
}

func (c *ColumnInfo) SetShortTableName(shortTableName string) {
	c.shortTableName = shortTableName
}

func NewColumnInfo() *ColumnInfo {
	return &ColumnInfo{}
}
