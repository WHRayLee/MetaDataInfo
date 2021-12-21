package model

type TableBasicInfo struct {
	dBName         string
	schemaName     string
	tableName      string
	tableShortName string
}

func (t *TableBasicInfo) DBName() string {
	return t.dBName
}

func (t *TableBasicInfo) SchemaName() string {
	return t.schemaName
}

func (t *TableBasicInfo) TableName() string {
	return t.tableName
}

func (t *TableBasicInfo) TableShortName() string {
	return t.tableShortName
}

func (t *TableBasicInfo) SetTableShortName(tableShortName string) {
	t.tableShortName = tableShortName
}

func (t *TableBasicInfo) SetTableName(tableName string) {
	t.tableName = tableName
}

func (t *TableBasicInfo) SetSchemaName(schemaName string) {
	t.schemaName = schemaName
}

func (t *TableBasicInfo) SetDBName(dBName string) {
	t.dBName = dBName
}

func NewTableBasicInfo() *TableBasicInfo {
	return &TableBasicInfo{}
}
