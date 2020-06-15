package log_test

import (
	"context"
	"os"

	"github.com/beanscc/rango/log"
)

func Example_Basic() {
	// 使用全局 logger log.Std()
	log.Std().Info("test global logger") // output：{"level":"info","time":"2020-06-15T15:30:44.726+0800","file":"log/example_test.go:11","msg":"test global logger"}

	log.Std().Infof("test %s", "global logger") // output: {"level":"info","time":"2020-06-15T15:30:44.726+0800","file":"log/example_test.go:14","msg":"test global logger"}
}

func Example_SetLogger() {
	// in init process
	// ...

	// init logger
	// 重置 logger
	logger := log.NewZapLogger(log.DebugLevel)
	log.SetLogger(logger)

	log.Std().Info("log use customer NewZapLogger") // output: {"level":"info","time":"2020-06-15T15:39:57.303+0800","file":"log/example_test.go:26","msg":"log use customer NewZapLogger"}

	// init logger with file writer
	logger = log.NewZapLogger(log.LevelFromStr("debug"), os.Stdout, log.NewZapCoreFileWriteSyncer("/tmp/log.log", 100))
	log.SetLogger(logger)

	log.Std().Info("log output to os.Stdout and file") // output: {"level":"info","time":"2020-06-15T15:39:57.303+0800","file":"log/example_test.go:32","msg":"log output to os.Stdout and file"}

	// store logger to ctx
	ctx := log.NewContext(context.Background(), log.Std())

	// init logger with global filed
	logger2 := log.NewZapLogger(log.DebugLevel)
	logger2 = logger2.With("namespace", "project-group").With("project", "project-name").With("x-request-id", "xxxx-xxxx-xxxx")
	log.SetLogger(logger2)
	log.Std().Info("log with namespace and project") // output: {"level":"info","time":"2020-06-15T15:39:57.303+0800","file":"log/example_test.go:41","msg":"log with namespace and project","namespace":"project-group","project":"project-name","x-request-id":"xxxx-xxxx-xxxx"}

	log.FromContext(ctx).Info("log from ctx") // output: {"level":"info","time":"2020-06-15T15:39:57.303+0800","file":"log/example_test.go:43","msg":"log from ctx"}
}

func Example_FromContext() {
	/*
		// Eg: gin middleware
		func SetLogger() gin.HandlerFunc {
			return func(c *gin.Context) {
				key := `x-request-id`
				requestID := stringutil.UUID()
				c.Writer.Header().Set(key, requestID)
				c.Set(log.LoggerCtxKey(), log.Logger().With(key, requestID))
				c.Next()
			}
		}
		g := gin.New()
		g.Use(
			middleware.SetLogger(),
			middleware.Recovery(),
		)

		...

		func Handle(c *gin.Context) {
			var req HandleReq
			if err := c.BindJSON(&req); err != nil {
				log.FromContext(c).Errorf("Handle bind param failed. err:%v",err)
				c.JSON(http.OK, g.H{"code": 40000, "msg": "invalid param"})
			}

			....
		}
	*/

	// 在请求初始化时/开始前，可将 logger 对象保存到请求上下文中
	ctx := log.NewContext(context.Background(), log.Std().With("x-request-id", "xxxx-xxxx-xxxx"))
	// ... handle req

	// 在没有 ctx 的位置，继续使用全局 Std() logger
	log.Std().Info("no x-request-id field") // output: {"level":"info","time":"2020-06-15T15:57:10.464+0800","file":"log/example_test.go:82","msg":"no x-request-id field"}

	// 在后面的函数中，可通过 ctx 取出 logger，继续使用
	log.FromContext(ctx).Info("logger from ctx") // output: {"level":"info","time":"2020-06-15T15:57:10.464+0800","file":"log/example_test.go:85","msg":"logger from ctx","x-request-id":"xxxx-xxxx-xxxx"}
}
