package main

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"strings"
)

type Field struct {
	Name    string   // 字段名
	Type    string   // 字段类型
	Tags    []string // tag 标签
	Comment string   // 字段注释
}

func (f Field) String() string {
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

	return strings.Join(cs, " ")
}

type StructGenerator struct {
	Name    string  // 结构体名称
	Comment string  // 结构体注释
	Fields  []Field // 结构体字段
}

func (g StructGenerator) Format() (string, error) {
	src, err := format.Source([]byte(g.String()))
	if err != nil {
		// Should never happen, but can arise when developing this code.
		// The user can compile the output to see the error.
		// log.Printf("warning: internal error: invalid Go generated: %s", err)
		// log.Printf("warning: compile the package to analyze the error")
		return g.String(), err
	}
	return string(src), nil
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
	fields := make([]string, 0, len(g.Fields))
	for _, v := range g.Fields {
		fields = append(fields, v.String())
	}
	return fields
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

// tableSchemeGenerator 表结构生成器
type tableSchemeGenerator struct {
	PackageName  string  // package name
	TableName    string  // table_name
	TableComment string  // 表注释
	Fields       []Field // 结构体字段
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
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("package %s\n", g.PackageName))

	imports := g.getImports()
	if len(imports) > 0 {
		buf.WriteString("\nimport (\n\t")
		buf.WriteString(strings.Join(imports, "\n\t"))
		buf.WriteString("\n)\n")
	}

	sg := StructGenerator{
		Name:    g.getStructName(),
		Comment: g.getTableComment(),
		Fields:  g.Fields,
	}

	buf.WriteString(sg.String()) // struct
	buf.WriteString(fmt.Sprintf(`
// TableName table name
func (t %s) TableName() string {
	return %q
}`, sg.Name, g.TableName,
	))

	src, err := format.Source(buf.Bytes())
	if err != nil {
		log.Printf("warning: internal error: invalid generated: %s", err)
		log.Printf("warning: compile the package to analyze the error")
		return buf.String()
	}

	return string(src)
}
