package log

import (
	"os"
	"sync"
)

var (
	_globalMutex  sync.RWMutex
	_globalLogger = NewZapLogger(DebugLevel, os.Stdout)
)

// Std() 返回全局 logger
// 使用 debug level，输出到标准输出 os.Stdout
func Std() Logger {
	_globalMutex.RLock()
	l := _globalLogger
	_globalMutex.RUnlock()
	return l
}

// SetLogger SetLogger 替换全局 logger
func SetLogger(logger Logger) {
	_globalMutex.Lock()
	_globalLogger = logger
	_globalMutex.Unlock()
}

// WithNamespaceAndProject
func WithNamespaceAndProject(logger Logger, namespace, project string) Logger {
	return logger.With("namespace", namespace).
		With("project", project)
}

// InitProject 初始化日志log级别，同时添加 namespace 和 project 全局字段，日志输出到标准输出
func InitProject(namespace, project string, level Level) {
	SetLogger(WithNamespaceAndProject(NewZapLogger(level, os.Stdout), namespace, project))
}
