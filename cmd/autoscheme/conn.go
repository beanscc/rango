package main

import (
	"time"

	"github.com/beanscc/rango/database"
	"github.com/beanscc/rango/database/gormutil"
)

var globalDBConn *gormutil.Conn
var globalDBName string

func initDB(dsn string) error {
	m := gormutil.Config{
		Config: database.Config{
			DriverType:  "mysql",
			DSN:         dsn,
			MaxIdle:     10,
			MaxOpen:     2,
			MaxLifeTime: 10 * time.Minute,
		},
		Debug:    false,
		Unscoped: true,
	}

	db, err := gormutil.NewConn(&m, &m)
	if err != nil {
		return err
	}

	globalDBConn = db

	// 获取当前连接的 db 名
	globalDBName, err = getDatabaseName()
	return err
}

func conn() *gormutil.Conn {
	return globalDBConn
}

func dbName() string {
	return globalDBName
}
