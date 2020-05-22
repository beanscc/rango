package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

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

// runeDisplayLenMap utf8字符等宽显示所需的字节长度
var runeDisplayLenMap = map[int]int{
	1: 1, // 如 英文字母
	2: 1, // 如 埃塞俄比亚语
	3: 2, // 如 中文、日文
	4: 2, // 如 表情：😂
	5: 2,
	6: 2,
}

type StructGenerator struct {
	Name    string  // 结构体名称
	Comment string  // 结构体注释
	Fields  []Field // 结构体字段
}

func (g StructGenerator) String() string {
	fields := g.formatFields()

	return fmt.Sprintf(`// %s %s
type %s struct {
	%s
}`,
		g.Name,
		g.Comment,
		g.Name,
		strings.Join(fields, "\n\t"),
	)
}

// formatFields 格式化对齐所有字段
func (g StructGenerator) formatFields() []string {
	fieldsColumns := make([][]string, 0, len(g.Fields))
	for _, v := range g.Fields {
		fieldsColumns = append(fieldsColumns, v.Columns())
	}

	// 计算每列字符串显示所需的最大字节数  map[int]int, 然后按各列的最长单位对齐
	columnsMaxDisplayLen := g.calColumnsMaxDisplayLen(fieldsColumns)
	rows := make([]string, 0, len(fieldsColumns))
	for _, fieldColumns := range fieldsColumns {
		rowStr := g.formatField(fieldColumns, columnsMaxDisplayLen)
		rows = append(rows, rowStr)
	}

	return rows
}

// formatField 按各列对齐所需的最大长度，格式化字段输出
func (g StructGenerator) formatField(columns []string, columnsMaxDisplayLen map[int]int) string {
	columnData := make([]string, 0, len(columns))
	for index, column := range columns {
		charLen := utf8.RuneCountInString(column)
		displayLen := g.calStrDisplayLen(column)
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
func (g StructGenerator) calColumnsMaxDisplayLen(rows [][]string) map[int]int {
	columnsMaxDisplayLen := make(map[int]int, 0)
	for _, row := range rows {
		for index, column := range row {
			cdLen := g.calStrDisplayLen(column)
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
func (g StructGenerator) calStrDisplayLen(s string) int {
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

// 常用大写缩写
var commonInitialisms = map[string]bool{
	"ACL":   true,
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SQL":   true,
	"SSH":   true,
	"TCP":   true,
	"TLS":   true,
	"TTL":   true,
	"UDP":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
	"XMPP":  true,
	"XSRF":  true,
	"XSS":   true,
	"MOQ":   true,
	"OS":    true,
	"WIFI":  true,
}

// 下划线格式转大写驼峰
// under_score -> UnderScore
func snake2Camel(s string) string {
	ss := strings.Split(s, "_")
	for i, v := range ss {
		if u := strings.ToUpper(v); commonInitialisms[u] {
			ss[i] = u
		} else {
			ss[i] = strings.Title(v)
		}
	}
	return strings.Join(ss, "")
}

// =============== struct generator end ===============

// tableSchemeGenerator 表结构生成器
type tableSchemeGenerator struct {
	PackageName  string // package name
	TableName    string // table_name
	TableComment string // 表注释
	Fields       []Field
}

func (g tableSchemeGenerator) getImports() []string {
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

func (g tableSchemeGenerator) getTableComment() string {
	tableComment := g.TableComment
	if tableComment == "" {
		tableComment = "..."
	}

	return tableComment
}

func (g tableSchemeGenerator) getStructName() string {
	return snake2Camel(g.TableName)
}

func (g tableSchemeGenerator) String() string {
	imports := g.getImports()
	structName := g.getStructName()

	sg := StructGenerator{
		Name:    structName,
		Comment: g.getTableComment(),
		Fields:  g.Fields,
	}

	ts := fmt.Sprintf(`
%s

// TableName table name
func (t %s) TableName() string {
	return %q
}`,
		sg.String(),
		structName,
		g.TableName,
	)

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("package %s\n", g.PackageName))
	if len(imports) > 0 {
		buf.WriteString("\nimport (\n\t")
		buf.WriteString(strings.Join(imports, "\n\t"))
		buf.WriteString("\n)\n")
	}
	buf.WriteString(ts)

	return buf.String()
}
