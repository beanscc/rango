package gormutil_test

import (
	"context"
	"testing"
	"time"

	"github.com/beanscc/rango/database/mysql/gormutil"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type app struct {
	ID        int64     `db:"id" gorm:"column:id"`
	Name      string    `db:"name" gorm:"column:name"`
	AppID     string    `db:"app_id" gorm:"column:app_id"`
	Secret    string    `db:"secret" gorm:"column:secret"`
	Sign      string    `db:"sign" gorm:"column:sign"`
	Status    bool      `db:"status" gorm:"column:status"`
	EndTime   int64     `db:"end_time" gorm:"column:end_time"`
	StartTime int64     `db:"start_time" gorm:"column:start_time"`
	Ctime     time.Time `db:"ctime" gorm:"column:ctime"`
	Utime     time.Time `db:"utime" gorm:"column:utime"`
	Operator  string    `db:"operator" gorm:"column:operator"`
}

func (a app) TableName() string {
	return "app"
}

/*

-- 创建测试表：

CREATE TABLE IF NOT EXISTS `app` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'APP 名称',
  `app_id` CHAR(40) NOT NULL DEFAULT '' COMMENT '应用APPID',
  `secret` VARCHAR(45) NOT NULL DEFAULT '' COMMENT '应用APP secret',
  `sign` VARCHAR(45) NOT NULL DEFAULT '' COMMENT '应用签名key',
  `start_time` BIGINT(20) NOT NULL DEFAULT 0 COMMENT '生效时间',
  `end_time` BIGINT(20) NOT NULL DEFAULT 0 COMMENT '结束时间',
  `status` INT NOT NULL DEFAULT 0 COMMENT '应用状态；0-停用；1-启用',
  `description` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '描述',
  `operator` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '操作人',
  `ctime` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `utime` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC),
  UNIQUE INDEX `app_id_UNIQUE` (`app_id` ASC),
  UNIQUE INDEX `name_UNIQUE` (`name` ASC))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COMMENT = '应用APP'

-- 添加测试数据

INSERT INTO `test`.`app`
(
`name`,
`app_id`,
`secret`,
`sign`,
`start_time`,
`end_time`,
`status`,
`description`,
`operator`)
VALUES
('t1', "app_id_1_dfsdfsdfsdfsddf", "secret_1_sfsdfsdfsdfsd", "sign_1_dsfsdfvsdghadfg", 14811110152, 1523772288, 0, "app_id_1_desc", "yx"),
('t2', "app_id_2_dfsdfsdfsdfsddf", "secret_2_sfsdfsdfsdfsd", "sign_2_dsfsdfvsdghadfg", 1482220152, 1523772288, 1, "app_id_2_desc", "yx"),
('t3', "app_id_3_dfsdfsdfsdfsddf", "secret_3_sfsdfsdfsdfsd", "sign_4_dsfsdfvsdghadfg", 1483330152, 1523772288, 0, "app_id_3_desc", "yx");
*/

var model *gormutil.Model

func getModelConfig() *gormutil.Config {
	dbCfg := &gormutil.DBConfig{
		DSN:  "root@/test?parseTime=true",
		Conn: gormutil.DefaultConnConfig,
		//Log:
		Debug:    true,
		Unscoped: true,
	}
	return &gormutil.Config{
		Master: dbCfg,
		Slave:  dbCfg,
	}
}
func Instance() *gormutil.Model {
	if model == nil {
		cfg := getModelConfig()
		var err error
		model, err = gormutil.NewModel(cfg)
		if err != nil {
			panic(err)
		}
	}

	return model
}

// go test -v -run Test_GetAPPByID
func Test_GetAPPByID(t *testing.T) {
	var resp app
	err := Instance().Slave().Where("id=?", 1).First(&resp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		t.Errorf("query first failed. err=%v", err)
		return
	}

	t.Logf("query resp=%v", resp)
}

func Test_GetAppByIDWithContextLogger(t *testing.T) {
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

	var resp app
	model := Instance()

	// 给 slave 设置 ctxLoggerKey
	//model.SlaveWithLoggerCtxKey(gormutil.LogrusCtxKey)  // 不设置 ctxLoggerKey 则无法使用存在ctx 中的logger对象，即不能使用预先定义的logger以及预设在logger上的日志信息

	// 下面的 query 日志中不会有上面logrusEntry 对象中设置的信息（ctx 中设置了 ctxLoggerKey，但没有使用 SlaveWithLoggerCtxKey 设置 ctxLoggerKey, SlaveWithLoggerFromContext 不知道从 ctx 中哪个key中获取 logger对象）
	err := model.SlaveWithLoggerFromContext(ctx).Where("id=?", 1).First(&resp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		t.Errorf("query first failed. err=%v", err)
		return
	}

	t.Logf("query resp=%v", resp)

	var resp2 app

	// 给 slave 设置 ctxLoggerKey
	model.SlaveWithLoggerCtxKey(gormutil.LogrusCtxKey) // 不设置 ctxLoggerKey 则无法使用存在ctx 中的logger对象，即不能使用预先定义的logger以及预设在logger上的日志信息

	// 下面的 query 日志中会有上面 logrusEntry 对象中设置的信息（因为 SlaveWithLoggerCtxKey 设置了 ctxLoggerKey）
	err = model.SlaveWithLoggerFromContext(ctx).Where("id=?", 1).First(&resp2).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		t.Errorf("after SlaveWithLoggerCtxKey, SlaveWithLoggerFromContext query first failed. err=%v", err)
		return
	}

	t.Logf("query2 resp2=%v", resp2)

	var resp3 app

	// 给 slave 设置 ctxLoggerKey
	//model.SlaveWithLoggerCtxKey(gormutil.LogrusCtxKey) // 不设置 ctxLoggerKey 则无法使用存在ctx 中的logger对象，即不能使用预先定义的logger以及预设在logger上的日志信息

	// 下面的 query 日志中同样不会有 logrusEntry 对象中设置的信息（因为所有的 with 方法都是对原有对象的一个 clone 操作，类似 withContext，然后重置 相关的变量）
	err = model.Slave().Where("id=?", 1).First(&resp3).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		t.Errorf("after SlaveWithLoggerCtxKey, SlaveWithLoggerFromContext query first failed. err=%v", err)
		return
	}

	t.Logf("query3 resp3=%v", resp3)
}
