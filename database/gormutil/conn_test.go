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

type Res struct {
	Name string
}

func TestSlaveWithGormDefaultLogger(t *testing.T) {
	conn := getConn()
	var res Res
	// with gorm default logger
	err := conn.Slave().
		Raw("select database() as name").
		Scan(&res).
		Error
	if err != nil {
		t.Errorf("[TestSlaveWithGormDefaultLogger] gorm exec raw sql failed. err:%v", err)
		return
	}
	t.Logf("[TestSlaveWithGormDefaultLogger] db name:%v", res.Name)
}

func TestSlaveWithZapLogger(t *testing.T) {
	conn := getConn()
	conn.SetContextFunc(gormutil.DefaultContextFunc())
	// 使用 zap 自定义的 logger
	log.Init("group", "test", "debug", os.Stdout)
	ctx := log.NewContext(context.Background(), log.Logger().With("x-request-id", "xxxx-xxx-xxx"))

	var res Res
	err := conn.Slave().
		WithContext(ctx).
		Raw(`select database() as "name"`).
		Scan(&res).
		Error
	if err != nil {
		t.Errorf("[TestSlaveWithZapLogger] gorm exec raw sql failed. err:%v", err)
		return
	}

	t.Logf("[TestSlaveWithZapLogger] db name:%v", res.Name)
}

func getConn() *gormutil.Conn {
	driverType := "mysql"
	dsn := `root:P4m@bpet@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local`
	mc := gormutil.Config{
		Config: database.Config{
			DriverType:  driverType,
			DSN:         dsn,
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
