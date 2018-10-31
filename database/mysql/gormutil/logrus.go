package gormutil

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

// Logger default logger
type Logger struct {
	Writer *logrus.Entry
}

// Print format & print log
func (logger Logger) Print(values ...interface{}) {
	fields := LogFormatter(values...)
	logEntry := logger.Writer.WithFields(fields)
	if fields["gorm-level"].(string) == LogLevelSQL {
		logEntry.Info("")
	} else {
		logEntry.Error("gorm err. See the gorm-msg field for details.")
	}
}

type ContextKey string

const (
	logrusCtxKey = ContextKey("logrus")
)

// FromContext 从 ctx 中获取 logrus 对象
func FromContext(ctx context.Context) *logrus.Entry {
	log := ctx.Value(logrusCtxKey)
	if log == nil {
		return NewLogrusEntry()
	}

	return log.(*logrus.Entry)
}

// ToContext 将 logrus 对象存入 ctx
func ToContext(ctx context.Context, log *logrus.Entry) context.Context {
	return context.WithValue(ctx, logrusCtxKey, log)
}

// NewLogrusEntry new logrus.Entry
func NewLogrusEntry() *logrus.Entry {
	log := logrus.New()
	log.Formatter = &logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
	}

	// logEntry := log.WithFields(logrus.Fields{
	// // "commit_id": util.CommitID(),
	// })
	logEntry := logrus.NewEntry(log)

	return logEntry
}
