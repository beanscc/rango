package log

import (
	"context"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var once sync.Once
var logger *zap.SugaredLogger

// Init 初始化 logger 对象
// 若未指定或指定 level 不存在，则 level = info
// 若未设置 writers 则 输出到 标准输出
// eg: log.Init("group-name", "project-name", "debug", os.Stdout, log.NewFileWriteSyncer("/tmp/group-project.log", 100))
func Init(namespace, project, level string, writers ...zapcore.WriteSyncer) {
	once.Do(func() {
		logger = newLogger(namespace, project, level, writers...)
	})
}

func newLogger(namespace, project, level string, writers ...zapcore.WriteSyncer) *zap.SugaredLogger {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "log",
		CallerKey:      "file",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 默认到 标准输出
	if len(writers) == 0 {
		writers = []zapcore.WriteSyncer{
			os.Stdout,
		}
	}

	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zapLevel(level))
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(writers...),
		atomicLevel,
	)
	caller := zap.AddCaller()
	development := zap.Development()
	return zap.New(core, caller, development, zap.AddCallerSkip(0)).Sugar().With("namespace", namespace, "project", project)
}

var zapLevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

// 未知 level 则返回 InfoLevel
func zapLevel(level string) zapcore.Level {
	if l, ok := zapLevelMap[level]; ok {
		return l
	}
	return zap.InfoLevel
}

// Logger 获取全局logger 对象；需要先初始化
func Logger() *zap.SugaredLogger {
	if logger == nil {
		panic("nil logger")
	}

	return logger
}

// NewContext return new context with a *zap.SugaredLogger inside
// 若将 log 保存在 gin.Context 中：
//      // c = *gin.Context
//      c.Set(log.ContextKey(), logger)
func NewContext(ctx context.Context, log *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, ContextKey(), log)
}

// FromContext return log from ctx
func FromContext(ctx context.Context) *zap.SugaredLogger {
	if ctx == nil {
		panic("nil ctx")
	}

	val := ctx.Value(ContextKey())
	if v, ok := val.(*zap.SugaredLogger); ok {
		return v
	}

	return Logger()
}

type ctxKey string

const loggerCtxKey ctxKey = "Ctx-Key-Logger"

// ContextKey return
func ContextKey() ctxKey {
	return loggerCtxKey
}

// NewFileWriteSyncer 文件输出器
func NewFileWriteSyncer(file string, size int) zapcore.WriteSyncer {
	return zapcore.AddSync(&lumberjack.Logger{
		Filename: file,
		MaxSize:  size,
	})
}
