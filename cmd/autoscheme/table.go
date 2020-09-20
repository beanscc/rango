package main

import (
	"fmt"
	"strings"
)

func buildTableStruct(packageName string, table string) (string, error) {
	tableInfo, err := getTableInfo(table)
	if err != nil {
		return "", err
	}
	columns, err := getTableFullColumns(table)
	if err != nil {
		return "", err
	}

	fields := make([]Field, 0, len(columns))
	for _, v := range columns {
		dataType := getDataType(v.Type)
		if strings.ToUpper(v.Null) == "YES" {
			dataType = "*" + dataType
		}

		// 将 默认值放入 comment
		comment := v.Comment
		if v.Default != "" {
			comment += fmt.Sprintf("(default: %s)", v.Default)
		}

		fields = append(fields, Field{
			Name: snake2Camel(v.Field),
			Type: dataType,
			Tags: []string{
				v.GormTag(),
				fmt.Sprintf(`json:"%s"`, v.Field),
			},
			Comment: comment,
		})
	}

	g := tableSchemeGenerator{
		PackageName:  packageName,
		TableName:    table,
		TableComment: tableInfo.Comment,
		Fields:       fields,
	}

	return g.String(), nil
}

type Name struct {
	Name string `gorm:"column:name"`
}

func getNames(ns []Name) []string {
	names := make([]string, 0, len(ns))
	for _, v := range ns {
		names = append(names, v.Name)
	}
	return names
}

func getDatabaseName() (string, error) {
	var names []Name
	err := conn().
		Raw(`select database() as name`).
		Scan(&names).Error
	return names[0].Name, err
}

func getTables() ([]string, error) {
	var tables []Name
	err := conn().
		Raw(`select table_name as name from information_schema.tables where table_schema=?`, dbName()).
		Scan(&tables).Error
	return getNames(tables), err
}

type Column struct {
	Field     string `gorm:"column:Field"`     // 字段名
	Type      string `gorm:"column:Type"`      // 字段类型: int(11)/int(10) unsigned/datetime/varchar(520)
	Collation string `gorm:"column:Collation"` // 字符集：utf8mb4_general_ci
	Null      string `gorm:"column:Null"`      // 是否允许 NULL
	Key       string `gorm:"column:Key"`       // Key：PRI MUL
	Default   string `gorm:"column:Default"`   // 默认值：NULL/0/CURRENT_TIMESTAMP/...
	Extra     string `gorm:"column:Extra"`     // 扩展信息：auto_increment/on update CURRENT_TIMESTAMP
	Comment   string `gorm:"column:Comment"`   // 注释
}

func (c *Column) GormTag() string {
	tags := []string{
		"column:" + c.Field,
	}

	if c.Key == "PRI" {
		tags = append(tags, "primaryKey")
	}

	if strings.Contains(c.Extra, "auto_increment") {
		tags = append(tags, "autoIncrement")
	}

	if c.Null == "NO" {
		tags = append(tags, "NOT NULL")
		if c.Default != "" {
			tags = append(tags, "default:"+c.Default)
		}
	}

	return fmt.Sprintf(`gorm:"%s"`, strings.Join(tags, ";"))
}

// 获取表的所有列字段信息
func getTableFullColumns(table string) ([]Column, error) {
	var out []Column
	err := conn().Raw(fmt.Sprintf("show full columns from `%s`", table)).Scan(&out).Error
	return out, err
}

type Table struct {
	Comment string `gorm:"column:table_comment"`
}

func getTableInfo(table string) (*Table, error) {
	var out Table
	err := conn().Raw(`select table_comment from information_schema.tables where table_schema = ? and table_name = ?`, dbName(), table).Scan(&out).Error
	return &out, err
}
