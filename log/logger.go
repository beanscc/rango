package log

import (
	"context"
)

// Logger log 接口
// - 包含了一组不同级别的 log 方法
// - With 方法可向 logger 对象中添加 key/val 健值对，将输出到日志中
type Logger interface {
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})

	Info(v ...interface{})
	Infof(format string, v ...interface{})

	// 消息等级高于 Info，但不必担心，不是严重错误，
	Warn(v ...interface{})
	Warnf(format string, v ...interface{})

	// 消息等级很高，一般程序正常运行的话，理应没有此类错误
	Error(v ...interface{})
	Errorf(format string, v ...interface{})

	// logs a message, then panics
	Panic(v ...interface{})
	Panicf(format string, v ...interface{})

	// logs a message, then call os.Exit(1)
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})

	// With 用于向 logger 对象中设置 key/val 对，该健值对将输出到日志
	With(key string, val interface{}) Logger
}

// NewContext 将 Logger 对象保存到 ctx 上下文中
func NewContext(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, ContextKey(), logger)
}

// FromContext 从 ctx 上下文中获取 Logger，若不存在，则返回全局默认 std.Logger()
func FromContext(ctx context.Context) Logger {
	if ctx == nil {
		panic("log: nil ctx")
	}

	val := ctx.Value(ContextKey())
	if v, ok := val.(Logger); ok {
		return v
	}

	return Std()
}

// ContextKey 返回将 Logger 存入 context 上下文时，对应的 key
func ContextKey() string {
	return "github.com/beanscc/rango/log::context-key"
}
