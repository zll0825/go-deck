package service

import (
	"bytes"
	"gorm.io/gorm"
	"text/template"
)

type Generate struct {
	DB *gorm.DB

}

type TplData struct {
	StructName         string  `json:"structName"`
	TableName          string  `json:"tableName"`
	PackageName        string  `json:"packageName"`
	Abbreviation       string  `json:"abbreviation"`
	Description        string  `json:"description"`
	AutoCreateApiToSql bool    `json:"autoCreateApiToSql"`
	AutoMoveFile       bool    `json:"autoMoveFile"`
	Fields             []Field `json:"fields"`
}

type Field struct {
	FieldName       string `json:"fieldName"`
	FieldDesc       string `json:"fieldDesc"`
	FieldType       string `json:"fieldType"`
	FieldJson       string `json:"fieldJson"`
	DataType        string `json:"dataType"`
	DataTypeLong    string `json:"dataTypeLong"`
	Comment         string `json:"comment"`
	ColumnName      string `json:"columnName"`
	FieldSearchType string `json:"fieldSearchType"`
	DictType        string `json:"dictType"`
}

type DB struct {
	Database string `json:"database" gorm:"column:database"`

	Table []Table `gorm:"-"`
}

type Table struct {
	TableName string `json:"tableName"`

	Column []Column `gorm:"-"`
}

type Column struct {
	ColumnName    string `json:"columnName" gorm:"column:column_name"`
	DataType      string `json:"dataType" gorm:"column:data_type"`
	DataTypeLong  string `json:"dataTypeLong" gorm:"column:data_type_long"`
	ColumnComment string `json:"columnComment" gorm:"column:column_comment"`
}

func NewGenerate(DB *gorm.DB) *Generate {
	return &Generate{
		DB: DB,
	}
}

func (g *Generate)GetDBs() ([]DB, error)  {
	db := make([]DB, 0)
	sql := "SELECT SCHEMA_NAME AS `database` FROM INFORMATION_SCHEMA.SCHEMATA;"
	err := g.DB.Raw(sql).Scan(&db).Error
	if err != nil {
		return nil, err
	}

	return db, err
}

func (g *Generate)GetTables(dbName string) ([]Table, error) {
	table := make([]Table, 0)
	sql := "SELECT TABLE_NAME as table_name FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = ?"
	err := g.DB.Raw(sql, dbName).Scan(&table).Error
	if err != nil {
		return nil, err
	}

	return table, err
}

func (g *Generate) GetColumns(tableName, dbName string) ([]Column, error) {
	column := make([]Column, 0)
	sql := `
SELECT 
COLUMN_NAME column_name,
DATA_TYPE data_type,
CASE DATA_TYPE 
	WHEN 'longtext' THEN c.CHARACTER_MAXIMUM_LENGTH 
	WHEN 'varchar' THEN c.CHARACTER_MAXIMUM_LENGTH
	WHEN 'double' THEN CONCAT_WS( ',', c.NUMERIC_PRECISION, c.NUMERIC_SCALE ) 
	WHEN 'decimal' THEN CONCAT_WS( ',', c.NUMERIC_PRECISION, c.NUMERIC_SCALE ) 
	WHEN 'int' THEN c.NUMERIC_PRECISION 
	WHEN 'bigint' THEN c.NUMERIC_PRECISION ELSE '' END AS data_type_long,
COLUMN_COMMENT column_comment 
FROM INFORMATION_SCHEMA.COLUMNS c 
WHERE table_name = ? AND table_schema = ?`
	err := g.DB.Raw(sql, tableName, dbName).Scan(&column).Error
	if err != nil {
		return nil, err
	}

	return column, nil
}

func (g *Generate) GenerateFile()  {
}

func (g *Generate) BuildData()  {
	fields := make([]Field, 0)
	fields = append(fields, Field{
		FieldName:       "1",
		FieldDesc:       "2",
		FieldType:       "3",
		FieldJson:       "4",
		DataType:        "5",
		DataTypeLong:    "6",
		Comment:         "7",
		ColumnName:      "8",
		FieldSearchType: "9",
		DictType:        "10",
	})
	tplData := &TplData{
		StructName:         "a",
		TableName:          "b",
		PackageName:        "c",
		Abbreviation:       "d",
		Description:        "e",
		AutoCreateApiToSql: false,
		AutoMoveFile:       false,
		Fields:             fields,
	}

	tpl, err := g.LoadTemplate("/Users/zll/Develop/go/src/github.com/zll0825/go-deck/cmd/template/entity.go.tpl")
	if err != nil {
		return
	}

	var buf bytes.Buffer
	err = tpl.Execute(&buf, tplData)
	if err != nil {
		panic(err.Error())
	}

	println( buf.String())
}

func (g *Generate) LoadTemplate(path string) (*template.Template, error) {
	return template.ParseFiles(path)
}
