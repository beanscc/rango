package main

import (
	"context"

	"github.com/Sirupsen/logrus"
	"github.com/beanscc/rango/database/mysql/gormutil"
	"github.com/beanscc/rango/database/mysql/gormutil/examples"
	"github.com/jinzhu/gorm"
)

func main() {
	getAppByIDWithContextLogger()
}

func getAppByIDWithContextLogger() {
	// new logger
	logger := gormutil.NewLogrusEntry()
	logrusEntry := logger.Writer.WithFields(logrus.Fields{
		"x-uid":        "userID-123",
		"x-request-id": "xxx-xxx-xxx-xxx",
		"x-commit-id":  "xxxxx",
	})

	logger.Writer = logrusEntry

	// 将 logger 存入 ctx
	ctx := context.WithValue(context.Background(), gormutil.LogrusCtxKey, logger)

	var resp examples.App
	model := examples.Instance()

	// 给 slave 设置 ctxLoggerKey
	//model.SlaveWithLoggerCtxKey(gormutil.LogrusCtxKey)  // 不设置 ctxLoggerKey 则无法使用存在ctx 中的logger对象，即不能使用预先定义的logger以及预设在logger上的日志信息

	// 下面的 query 日志中不会有上面logrusEntry 对象中设置的信息（ctx 中设置了 ctxLoggerKey，但没有使用 SlaveWithLoggerCtxKey 设置 ctxLoggerKey, SlaveWithLoggerFromContext 不知道从 ctx 中哪个key中获取 logger对象）
	err := model.SlaveWithLoggerFromContext(ctx).Where("id=?", 1).First(&resp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logrus.Errorf("query first failed. err=%v", err)
		return
	}

	logrus.Infof("query resp=%v", resp)

	var resp2 examples.App

	// 给 slave 设置 ctxLoggerKey
	model.SlaveWithLoggerCtxKey(gormutil.LogrusCtxKey) // 不设置 ctxLoggerKey 则无法使用存在ctx 中的logger对象，即不能使用预先定义的logger以及预设在logger上的日志信息

	// 下面的 query 日志中会有上面 logrusEntry 对象中设置的信息（因为 SlaveWithLoggerCtxKey 设置了 ctxLoggerKey）
	err = model.SlaveWithLoggerFromContext(ctx).Where("id=?", 1).First(&resp2).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logrus.Errorf("after SlaveWithLoggerCtxKey, SlaveWithLoggerFromContext query first failed. err=%v", err)
		return
	}

	logrus.Infof("query2 resp2=%v", resp2)

	var resp3 examples.App

	// 给 slave 设置 ctxLoggerKey
	//model.SlaveWithLoggerCtxKey(gormutil.LogrusCtxKey) // 不设置 ctxLoggerKey 则无法使用存在ctx 中的logger对象，即不能使用预先定义的logger以及预设在logger上的日志信息

	// 下面的 query 日志中同样不会有 logrusEntry 对象中设置的信息（因为所有的 with 方法都是对原有对象的一个 clone 操作，类似 withContext，然后重置 相关的变量）
	err = model.Slave().Where("id=?", 1).First(&resp3).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logrus.Errorf("after SlaveWithLoggerCtxKey, SlaveWithLoggerFromContext query first failed. err=%v", err)
		return
	}

	logrus.Infof("query3 resp3=%v", resp3)
}
