package gormutil

import (
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

// logger 在上下文中存储的key
type LoggerCtxKey string

const (
	LogrusCtxKey = LoggerCtxKey("logrus")
)

// NewLogrusEntry new logrus.Entry with json formatter
func NewLogrusEntry() *logrus.Entry {
	log := logrus.New()
	log.Formatter = &logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
	}

	logEntry := logrus.NewEntry(log)

	return logEntry
}
