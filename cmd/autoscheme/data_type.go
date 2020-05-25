package main

import (
	"strings"
)

func getDataType(s string) string {
	types := strings.Split(s, " ")
	typeName := types[0]
	index := strings.Index(typeName, "(")
	if index != -1 {
		typeName = s[:index]
	}

	dataType, ok := mysqlDataType2Golang[strings.ToLower(typeName)]
	if !ok {
		return "unknown"
	}

	// go 只有整数型有 unsigned 属性
	switch dataType {
	case "int8", "int16", "int32", "int", "int64":
		// numeric unsigned/zerofill
		if strings.Contains(s, "unsigned") {
			dataType = "u" + dataType
		}
	}

	return dataType
}

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
	"enum":       "string",
	"set":        "string",
}
