package examples

import (
	"time"

	"github.com/beanscc/rango/database/mysql/gormutil"
)

type App struct {
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

func (a App) TableName() string {
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
