package log_test

import (
	"context"
	log2 "log"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/beanscc/rango/log"
)

func Test_Std(t *testing.T) {
	log.Std().Debug("debug test")

	wg := sync.WaitGroup{}
	wg.Add(1)
	// 异步使用全局默认 logger
	l := log.Std()
	go func() {
		defer wg.Done()
		time.Sleep(3 * time.Second)
		l.Debug("setlogger 后，原 logger 是否有效")
	}()
	log.Std().Debugf("caller deep test:%s", "sss")

	// 设置新 logger
	logger := log.NewZapLogger(log.LevelFromStr("debug")).
		With("namespace", "g1").
		With("group", "p1")
	log.SetLogger(logger)
	log.Std().Debug("debug use logger")

	// set logger to ctx
	ctx := log.NewContext(context.Background(), log.Std())

	// 设置新 logger2, 同步输出到 /tmp/log.log 文件, 同时加 namespace 和 project 全局字段
	logger2 := log.NewZapLogger(log.LevelFromStr("debug"), os.Stdout, log.NewZapCoreFileWriteSyncer("/tmp/log.log", 100))
	log.SetLogger(log.WithNamespaceAndProject(logger2, "g2", "p2"))
	log.Std().Debug("debug use logger2")

	log.FromContext(ctx).Debug("log from ctx")

	wg.Wait()
}

func Test_A(t *testing.T) {
	// 在请求初始化时/开始前，可将 logger 对象保存到请求上下文中
	ctx := log.NewContext(context.Background(), log.Std().With("x-request-id", "xxxx-xxxx-xxxx"))
	// ... handle req

	// 在没有 ctx 的位置，继续使用全局 Std() logger
	log.Std().Info("no x-request-id field") // output: {"level":"info","time":"2020-06-15T15:57:10.464+0800","file":"log/logger_test.go:54","msg":"no x-request-id field"}

	// 在后面的函数中，可通过 ctx 取出 logger，继续使用
	log.FromContext(ctx).Info("logger from ctx") // output: {"level":"info","time":"2020-06-15T15:57:10.464+0800","file":"log/logger_test.go:57","msg":"logger from ctx","x-request-id":"xxxx-xxxx-xxxx"}
}

// go test -v -count=1 -bench=BenchmarkDebug1 -benchmem -run BenchmarkDebug1
func BenchmarkDebug1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go func() {
			log.Std().Debug("BenchmarkDebug1")
		}()
	}
}

// go test -v -count=1 -bench=BenchmarkDebug2 -benchmem -run BenchmarkDebug2
func BenchmarkDebug2(b *testing.B) {
	logger := log.NewZapLogger(log.LevelFromStr("debug"))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		go func() {
			logger.Debug("BenchmarkDebug2")
		}()
	}
	b.StopTimer()
}

// go test -v -count=1 -bench=BenchmarkDebug3 -benchmem -run BenchmarkDebug3
func BenchmarkDebug3(b *testing.B) {
	// logger := log.NewZapLogger(log.LevelFromStr("debug"))
	for i := 0; i < b.N; i++ {
		go func() {
			log2.Print("BenchmarkDebug3")
		}()
	}
}
