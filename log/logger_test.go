package log_test

import (
	"context"
	"os"
	"testing"

	"github.com/beanscc/rango/log"
)

func TestLogger(t *testing.T) {
	log.Init("group-name", "project-name", "info", os.Stdout, log.NewFileWriteSyncer("test.log", 100))

	log.Logger().Infof("info log test. field_1=%v", "val-1")
	// 输出示例：{"level":"info","time":"2020-05-11T01:02:23.452+0800","file":"log/logger_test.go:14","msg":"info log test. field_1=val-1","namespace":"group-name","project":"project-name"}

	// x-request-id
	ctx := log.NewContext(context.Background(), log.Logger().With("x-request-id", "xxxx-xxxx-xxxx-xxxx"))

	log.FromContext(ctx).Info("test x-request-id")
	// output: {"level":"info","time":"2020-05-11T01:02:23.453+0800","file":"log/logger_test.go:20","msg":"test x-request-id","namespace":"group-name","project":"project-name","x-request-id":"xxxx-xxxx-xxxx-xxxx"}
}
