package main

import (
	"strings"
)

// // mysql 支持以下数据类型：
// // - 数值类型
// //     - 整数型
// //     - 浮点型
// // - 日期和时间类型
// // - 字符类型
// // - 空间类型
// // - json 数据类型
// type MySQLDataType struct {
// 	Name     string // 小写类型名: int/integer/mediumint/smallint/tinyint/char/varchar/text
// 	Unsigned bool   // 是否无符号数值
// 	ZeroFill bool   // 是否按0左填充
// 	M        uint8  // 整型 width 位数/浮点型数值个数
// 	D        uint8  // 浮点型 小数位数
// }

func getDataType(s string) string {
	mysqlDataType := s
	index := strings.Index(s, "(")
	if index != -1 {
		mysqlDataType = s[:index]
	}
	dataType := mysqlDataType2Golang[strings.ToLower(mysqlDataType)]
	return dataType
}

// TODO 完善类型
// mysql 字段类型和golang类型对应
var mysqlDataType2Golang = map[string]string{
	// numeric
	// "bit": // TODO
	"tinyint":          "int8",
	"bool":             "bool",
	"boolean":          "bool",
	"smallint":         "int16",
	"mediumint":        "int",
	"int":              "int",
	"integer":          "int",
	"bigint":           "int64",
	"decimal":          "float64",
	"dec":              "float64",
	"numeric":          "float64",
	"fixed":            "float64",
	"float":            "float64",
	"double":           "float64",
	"double precision": "float64",

	// date and time
	"date":      "time.Time",
	"datetime":  "time.Time",
	"timestamp": "time.Time",
	"time":      "time.Time",
	"year":      "time.Time",

	// string
	"char":       "string",
	"varchar":    "string",
	"binary":     "string",
	"varbinary":  "string",
	"tinyblob":   "string",
	"tinytext":   "string",
	"blob":       "string",
	"text":       "string",
	"mediumblob": "string",
	"mediumtext": "string",
	"longblob":   "string",
	"longtext":   "string",
	// "enum" // TODO
	// "set" // TODO
}
