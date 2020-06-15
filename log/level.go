package log

import "fmt"

// Level log level
type Level int8

// log level
const (
	DebugLevel Level = iota + 1
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	FatalLevel
)

// Level string
func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	case PanicLevel:
		return "panic"
	case FatalLevel:
		return "fatal"
	default:
		return fmt.Sprintf("Level(%d)", l)
	}
}

// LevelFromStr 从 Level 的 string 格式获取 Level，如是为定义的格式，则返回 error 级别
func LevelFromStr(str string) Level {
	switch str {
	case DebugLevel.String():
		return DebugLevel
	case InfoLevel.String():
		return InfoLevel
	case WarnLevel.String():
		return WarnLevel
	case ErrorLevel.String():
		return ErrorLevel
	case PanicLevel.String():
		return PanicLevel
	case FatalLevel.String():
		return FatalLevel
	default:
		return ErrorLevel
	}
}
