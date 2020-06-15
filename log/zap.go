package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type zapLogger struct {
	zap *zap.SugaredLogger
}

func (l *zapLogger) Debug(v ...interface{}) {
	l.zap.Debug(v...)
}

func (l *zapLogger) Debugf(format string, v ...interface{}) {
	l.zap.Debugf(format, v...)
}

func (l *zapLogger) Info(v ...interface{}) {
	l.zap.Info(v...)
}
func (l *zapLogger) Infof(format string, v ...interface{}) {
	l.zap.Infof(format, v...)
}

func (l *zapLogger) Warn(v ...interface{}) {
	l.zap.Warn(v...)
}
func (l *zapLogger) Warnf(format string, v ...interface{}) {
	l.zap.Warnf(format, v...)
}

func (l *zapLogger) Error(v ...interface{}) {
	l.zap.Error(v...)
}
func (l *zapLogger) Errorf(format string, v ...interface{}) {
	l.zap.Errorf(format, v...)
}

func (l *zapLogger) Panic(v ...interface{}) {
	l.zap.Panic(v...)
}
func (l *zapLogger) Panicf(format string, v ...interface{}) {
	l.zap.Panicf(format, v...)
}

func (l *zapLogger) Fatal(v ...interface{}) {
	l.zap.Fatal(v...)
}

func (l *zapLogger) Fatalf(format string, v ...interface{}) {
	l.zap.Fatalf(format, v...)
}

func (l *zapLogger) With(key string, val interface{}) Logger {
	s := l.zap.With(key, val)
	nl := new(zapLogger)
	nl.zap = s
	return nl
}

var zapLevel = map[Level]zapcore.Level{
	DebugLevel: zapcore.DebugLevel,
	InfoLevel:  zapcore.InfoLevel,
	WarnLevel:  zapcore.WarnLevel,
	ErrorLevel: zapcore.ErrorLevel,
	PanicLevel: zapcore.PanicLevel,
	FatalLevel: zapcore.FatalLevel,
}

func NewZapLogger(level Level, writers ...zapcore.WriteSyncer) Logger {
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

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(writers...),
		zap.NewAtomicLevelAt(zapLevel[level]),
	)
	return &zapLogger{zap: zap.New(core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	).Sugar()}
}

// NewZapCoreFileWriteSyncer 文件输出器
func NewZapCoreFileWriteSyncer(file string, size int) zapcore.WriteSyncer {
	return zapcore.AddSync(&lumberjack.Logger{
		Filename: file,
		MaxSize:  size,
	})
}
