package gorm_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/beanscc/rango/database"
	"github.com/beanscc/rango/database/gorm"
	"github.com/beanscc/rango/log"
	_ "github.com/go-sql-driver/mysql"
	gorm2 "github.com/jinzhu/gorm"
)

func TestSlaveWithGormDefaultLogger(t *testing.T) {
	conn := getConn()

	res := struct {
		Name string
	}{}

	// with gorm default logger
	err := conn.Slave().Raw("select database() as name").Scan(&res).Error
	if err != nil {
		t.Errorf("[TestSlaveWithGormDefaultLogger] gorm exec raw sql failed. err:%v", err)
		return
	}
	t.Logf("[TestSlaveWithGormDefaultLogger] db name:%v", res.Name)
}

func TestSlaveWithZapLogger(t *testing.T) {
	conn := getConn()

	// 使用 zap 自定义的 logger
	log.Init("group", "test", "debug", os.Stdout, log.NewFileWriteSyncer("gorm.log", 10))
	ctx := log.NewContext(context.Background(), log.Logger().With("x-request-id", "xxxx-xxx-xxx"))

	// set LogWriter
	conn.SetLogWriter(func(ctx context.Context) gorm.LogWriter {
		return log.FromContext(ctx)
	})

	// 设置 withContextFunc
	conn.WithContextFunc(func(ctx context.Context, db *gorm2.DB) {
		db = db.Unscoped()
		log.Logger().Debug("call withContext")
	})

	res := struct {
		Name string
	}{}
	err := conn.SlaveWithContext(ctx).Raw(`select database() as "name"`).Scan(&res).Error
	if err != nil {
		t.Errorf("[TestSlaveWithZapLogger] gorm exec raw sql failed. err:%v", err)
		return
	}

	t.Logf("[TestSlaveWithZapLogger] db name:%v", res.Name)
}

func getConn() *gorm.Conn {
	masterConf := gorm.Config{
		Config: database.Config{
			DriverType:  "mysql",
			DSN:         "root:P4m@bpet@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local",
			MaxIdle:     16,
			MaxOpen:     64,
			MaxLifeTime: 10 * time.Minute,
		},
		Debug:    true,
		Unscoped: true,
	}

	slaverConf := gorm.Config{
		Config: database.Config{
			DriverType:  "mysql",
			DSN:         "root:P4m@bpet@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local",
			MaxIdle:     16,
			MaxOpen:     64,
			MaxLifeTime: 10 * time.Minute,
		},
		Debug:    true,
		Unscoped: true,
	}

	conn, err := gorm.NewConn(&masterConf, &slaverConf)
	if err != nil {
		panic(err)
	}
	return conn
}
