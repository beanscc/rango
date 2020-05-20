package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/beanscc/rango/utils/stringutil"
)

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
	err := conn().Slave().
		Raw(`select database() as name`).
		Scan(&names).Error
	return names[0].Name, err
}

func getTables() ([]string, error) {
	var tables []Name
	err := conn().Slave().
		Raw(`select table_name as name from information_schema.tables where table_schema=?`, dbName()).
		Scan(&tables).Error
	return getNames(tables), err
}

type Column struct {
	Field     string `gorm:"column:Field"`     // 字段名
	Type      string `gorm:"column:Type"`      // 字段类型: int(11) datetime varchar(520)
	Collation string `gorm:"column:Collation"` // 字符集：utf8mb4_general_ci
	Null      string `gorm:"column:Null"`      // 是否允许 NULL
	Key       string `gorm:"column:Key"`       // Key：PRI MUL
	Default   string `gorm:"column:Default"`   // 默认值：NULL/0/CURRENT_TIMESTAMP/...
	Extra     string `gorm:"column:Extra"`     // 扩展信息：auto_increment/on update CURRENT_TIMESTAMP
	Comment   string `gorm:"column:Comment"`   // 注释
}

// 获取表的所有列字段信息
func getTableFullColumns(table string) ([]Column, error) {
	var out []Column
	err := conn().Slave().Raw(`show full columns from ` + table).Scan(&out).Error
	return out, err
}

type Table struct {
	Comment string `gorm:"column:table_comment"`
}

func getTableInfo(table string) (*Table, error) {
	var out Table
	err := conn().Slave().Raw(`select table_comment from information_schema.tables where table_schema = ? and table_name = ?`, dbName(), table).Scan(&out).Error
	return &out, err
}

func buildTableStructBuffer(packageName string, table string) (string, error) {
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
		fields = append(fields, Field{
			Name: stringutil.Snake2Camel(v.Field, true),
			Type: dataType,
			Tags: []string{
				fmt.Sprintf(`gorm:"column:%s"`, v.Field),
				fmt.Sprintf(`json:"%s"`, v.Field),
			},
			Comment: v.Comment,
		})
	}

	g := tableStructGenerator{
		PackageName:  packageName,
		TableName:    table,
		TableComment: tableInfo.Comment,
		Fields:       fields,
	}

	return g.String(), nil
}

func getDataType(s string) string {
	mysqlDataType := s
	index := strings.Index(s, "(")
	if index != -1 {
		mysqlDataType = s[:index]
	}
	dataType := mysqlDataType2Golang[strings.ToLower(mysqlDataType)]
	return dataType
}

type tableStructGenerator struct {
	PackageName  string
	TableName    string // table_name
	TableComment string // 表注释
	Fields       []Field
}

type Field struct {
	Name    string   // 字段名
	Type    string   // 字段类型
	Tags    []string // tag 标签
	Comment string   // 字段注释
}

func (f Field) Columns() []string {
	cs := []string{
		f.Name, // name
		f.Type, // type
	}

	if len(f.Tags) > 0 {
		cs = append(cs, fmt.Sprintf("`%s`", strings.Join(f.Tags, " "))) // tag
	}

	if f.Comment != "" {
		cs = append(cs, fmt.Sprintf("// %s", f.Comment)) // comment
	}

	return cs
}

func (g tableStructGenerator) formatFields() []string {
	fieldsColumns := make([][]string, 0, len(g.Fields))
	for _, v := range g.Fields {
		fieldsColumns = append(fieldsColumns, v.Columns())
	}

	// 计算每列字符串显示所需的最大字节数  map[int]int, 然后按各列的最长单位对齐
	columnsMaxDisplayLen := calColumnsMaxDisplayLen(fieldsColumns)
	rows := make([]string, 0, len(fieldsColumns))
	for _, row := range fieldsColumns {
		rowStr := formatRow(row, columnsMaxDisplayLen)
		rows = append(rows, rowStr)
	}

	return rows
}

func (g tableStructGenerator) getImports() []string {
	var imports []string
	for _, v := range g.Fields {
		if v.Type == "time.Time" {
			imports = []string{
				`"time"`,
			}
			break
		}
	}

	return imports
}

func (g tableStructGenerator) String() string {
	camelTableName := stringutil.Snake2Camel(g.TableName, true)
	fields := g.formatFields()
	imports := g.getImports()

	tableStruct := fmt.Sprintf(tableStructFormat,
		camelTableName,
		g.TableComment,
		camelTableName,
		strings.Join(fields, "\n\t"),
		camelTableName,
		g.TableName,
	)

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("package %s\n", g.PackageName))
	// import
	if len(imports) > 0 {
		buf.WriteString("\nimport (\n\t")
		buf.WriteString(strings.Join(imports, "\n\t"))
		buf.WriteString("\n)\n")
	}
	buf.WriteString(tableStruct)

	return buf.String()
}

const tableStructFormat = `
// %s %s
type %s struct {
	%s
}

// TableName table name
func (t %s) TableName() string {
	return %q
}
`

// TODO 完善类型
// mysql 字段类型和golang类型对应
var mysqlDataType2Golang = map[string]string{
	// int
	"tinyint": "int8",
	"int":     "int",
	"bigint":  "int64",
	"decimal": "float64",

	// string
	"char":    "string",
	"varchar": "string",
	"text":    "string",

	// bool

	// time
	"timestamp": "time.Time",
	"datetime":  "time.Time",
}

// ============== column 对齐 ================
func formatRow(columns []string, columnsMaxDisplayLen map[int]int) string {
	columnData := make([]string, 0, len(columns))
	for index, column := range columns {
		charLen := utf8.RuneCountInString(column)
		displayLen := calStrDisplayLen(column)
		// 若该字符串等宽显示所需的字节数小于该列等宽显示所需的最大字节数，
		// 则给其 fmt 的时的字符长度 = 原字符长度 + （maxDisplayLen - displayLen）
		if displayLen < columnsMaxDisplayLen[index] {
			charLen += columnsMaxDisplayLen[index] - displayLen
		}

		if index == len(columns)-1 { // 最后一列，不用补齐
			columnData = append(columnData, column)
		} else {
			columnData = append(columnData, fmt.Sprintf("%-*s", charLen, column))
		}
	}

	return strings.Join(columnData, " ")
}

// calColumnsMaxDisplayLen 计算多行数据每列数据显示的最大字节长度
func calColumnsMaxDisplayLen(rows [][]string) map[int]int {
	columnsMaxDisplayLen := make(map[int]int, 0)
	for _, row := range rows {
		for index, column := range row {
			cdLen := calStrDisplayLen(column)
			if _, ok := columnsMaxDisplayLen[index]; !ok {
				columnsMaxDisplayLen[index] = cdLen
			} else {
				if columnsMaxDisplayLen[index] < cdLen {
					columnsMaxDisplayLen[index] = cdLen
				}
			}
		}
	}

	return columnsMaxDisplayLen
}

// calStrDisplayLen 计算字符串等宽显示需要的字节长度
func calStrDisplayLen(s string) int {
	tmp, l := s, 0
	for {
		if len(tmp) == 0 {
			break
		}
		_, i := utf8.DecodeRuneInString(tmp)
		if charDisplayLen, ok := runeDisplayLenMap[i]; ok {
			l += charDisplayLen
		} else {
			l += i
		}

		tmp = tmp[i:]
	}

	return l
}

// runeDisplayLenMap utf8字符等宽显示所需的字节长度
var runeDisplayLenMap = map[int]int{
	1: 1, // 如 英文字母
	2: 1, // 如 埃塞俄比亚语
	3: 2, // 如 中文、日文
	4: 2, // 如 表情：😂
	5: 2,
	6: 2,
}
