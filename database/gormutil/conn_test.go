package gormutil_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/beanscc/rango/database"
	"github.com/beanscc/rango/database/gormutil"
	"github.com/beanscc/rango/log"
	_ "github.com/go-sql-driver/mysql"
)

type testRes struct {
	Name string
}

func TestSlaveWithGormDefaultLogger(t *testing.T) {
	conn := testGetConn()
	var res testRes
	// with gorm default logger
	err := conn.Slave().
		Raw("select database() as name").
		Scan(&res).
		Error
	if err != nil {
		t.Errorf("[TestSlaveWithGormDefaultLogger] gorm exec raw sql failed. err:%v", err)
		return
	}
	/*
		查询SQL日志输出：
		(/Users/yan/work/github.com/beanscc/rango/database/gormutil/conn_test.go:25)
		[2020-06-10 01:35:08]  [1.35ms]   select database() as name
		[1 rows affected or returned ]
	*/

	t.Logf("[TestSlaveWithGormDefaultLogger] db name:%v", res.Name)
}

func TestSlaveWithZapLogger(t *testing.T) {
	// 使用 zap 自定义的 logger
	log.Init("group", "test", "debug", os.Stdout)
	ctx := log.NewContext(context.Background(), log.Logger().With("x-request-id", "xxxx-xxx-xxx"))

	conn := testGetConn()
	conn.SetContextFunc(gormutil.DefaultContextFunc())

	var res testRes
	err := conn.Slave().
		WithContext(ctx).
		Raw(`select database() as "name"`).
		Scan(&res).
		Error
	if err != nil {
		t.Errorf("[TestSlaveWithZapLogger] gorm exec raw sql failed. err:%v", err)
		return
	}

	// 查询SQL日志输出：
	// {"level":"info","time":"2020-06-10T01:33:18.043+0800","file":"gormutil/logger.go:166","msg":"[gorm-time]:2020-06-10T01:33:18.043104+08:00; [gorm-file]:/Users/yan/work/github.com/beanscc/rango/database/gormutil/conn_test.go:46; [gorm-latency]:4.77ms; [gorm-sql]:select database() as \"name\"; [gorm-rows]:1 rows affected or returned","namespace":"group","project":"test","x-request-id":"xxxx-xxx-xxx"}

	t.Logf("[TestSlaveWithZapLogger] db name:%v", res.Name)
}

func testGetConn() *gormutil.Conn {
	mc := gormutil.Config{
		Config: database.Config{
			DriverType:  "mysql",
			DSN:         `root:P4m@bpet@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local`,
			MaxIdle:     16,
			MaxOpen:     64,
			MaxLifeTime: 10 * time.Minute,
		},
		Debug:    true,
		Unscoped: true,
	}
	sc := mc
	conn, err := gormutil.NewConn(&mc, &sc)
	if err != nil {
		panic(err)
	}
	return conn
}
