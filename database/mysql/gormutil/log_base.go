package gormutil

import (
	"context"
	"database/sql/driver"
	"fmt"
	"reflect"
	"regexp"
	"time"
	"unicode"

	"github.com/beanscc/utils/stringutil"
)

var (
	sqlRegexp                = regexp.MustCompile(`\?`)
	numericPlaceHolderRegexp = regexp.MustCompile(`\$\d+`)
)

// gorm 日志级别
const (
	LogLevelSQL = "sql"
	LogLevelErr = "error"
)

// gorm 查询日志字段key定义
const (
	LogFieldLevel   = "gorm-level"
	LogFieldTime    = "gorm-time"
	LogFieldFile    = "gorm-file"
	LogFieldLatency = "gorm-latency"
	LogFieldSQL     = "gorm-sql"
	LogFieldRows    = "gorm-rows"
	LogFieldMsg     = "gorm-msg"
)

// NowFunc returns current time, this function is exported in order to be able
// to give the flexibility to the developer to customize it according to their
// needs, e.g:
//    gorm.NowFunc = func() time.Time {
//      return time.Now().UTC()
//    }
var NowFunc = func() time.Time {
	return time.Now()
}

func isPrintable(s string) bool {
	for _, r := range s {
		if !unicode.IsPrint(r) {
			return false
		}
	}
	return true
}

// LogFormatter gorm log formatter
var LogFormatter = func(values ...interface{}) (fields map[string]interface{}) {
	if len(values) > 1 {
		var (
			sql             string
			formattedValues []string
			level           = values[0]
			currentTime     = NowFunc().Format(time.RFC3339Nano)
			source          = values[1]
		)

		fields = map[string]interface{}{
			LogFieldLevel: level,
			LogFieldTime:  currentTime,
			LogFieldFile:  source,
		}

		if level == LogLevelSQL {
			// duration
			fields[LogFieldLatency] = fmt.Sprintf("%.2fms", float64(values[2].(time.Duration).Nanoseconds()/1e4)/100.0)

			// sql
			for _, value := range values[4].([]interface{}) {
				indirectValue := reflect.Indirect(reflect.ValueOf(value))
				if indirectValue.IsValid() {
					value = indirectValue.Interface()
					if t, ok := value.(time.Time); ok {
						formattedValues = append(formattedValues, fmt.Sprintf("'%v'", t.Format("2006-01-02 15:04:05")))
					} else if b, ok := value.([]byte); ok {
						if str := string(b); isPrintable(str) {
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", str))
						} else {
							formattedValues = append(formattedValues, "'<binary>'")
						}
					} else if r, ok := value.(driver.Valuer); ok {
						if value, err := r.Value(); err == nil && value != nil {
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
						} else {
							formattedValues = append(formattedValues, "NULL")
						}
					} else {
						formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
					}
				} else {
					formattedValues = append(formattedValues, "NULL")
				}
			}

			// differentiate between $n placeholders or else treat like ?
			if numericPlaceHolderRegexp.MatchString(values[3].(string)) {
				sql = values[3].(string)
				for index, value := range formattedValues {
					placeholder := fmt.Sprintf(`\$%d([^\d]|$)`, index+1)
					sql = regexp.MustCompile(placeholder).ReplaceAllString(sql, value+"$1")
				}
			} else {
				formattedValuesLength := len(formattedValues)
				for index, value := range sqlRegexp.Split(values[3].(string), -1) {
					sql += value
					if index < formattedValuesLength {
						sql += formattedValues[index]
					}
				}
			}

			// 将sql中的空白统一替换成一个空格
			sql = stringutil.SimplifyWhitespace(sql)
			// 影响行数
			affectedRows := fmt.Sprintf("%v rows affected or returned", values[5])

			fields[LogFieldSQL] = sql
			fields[LogFieldRows] = affectedRows
		} else {
			fields[LogFieldMsg] = values[2:]
		}
	}

	return
}

// logger gorm log interface
type logger interface {
	Print(v ...interface{})
}

// LoggerFromContext 从 ctx 按指定 key 取出 log 对象，若为该对象实现了 logger 接口，则直接返回，否则，使用指定的默认 logger 接口实现对象
func LoggerFromContext(ctx context.Context, loggerKey interface{}, defaultLogger logger) logger {
	log := ctx.Value(loggerKey)
	if log == nil {
		return defaultLogger
	}

	if log, ok := log.(logger); ok {
		return log
	}

	return defaultLogger
}

// LoggerToContext 将实现 logger 接口的对象存入 ctx
func LoggerToContext(ctx context.Context, loggerKey interface{}, log logger) context.Context {
	return context.WithValue(ctx, loggerKey, log)
}
