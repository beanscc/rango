package gormutil

import (
	"context"
	"time"

	"github.com/jinzhu/gorm"
)

// model model struct
type model struct {
	// master for write
	master *gorm.DB
	// slave for read
	slave *gorm.DB
	// logger
	logger logger
	// ctxLoggerKey logger 对象保存在 ctx 中的 key
	ctxLoggerKey interface{}
	// debug
	debug bool
	// unscoped (建议true，防止表结构中有 deleted_at 字段时，gorm 加该查询条件，不容易发现) return all record including deleted record, refer Soft Delete https://jinzhu.github.io/gorm/crud.html#soft-delete
	unscoped bool
}

// DBConfig db 配置 参数
type DBConfig struct {
	// db dsn
	DSN string
	// db conn config
	Conn *ConnConfig
}

// ConnConfig db 连接参数
type ConnConfig struct {
	MaxIdle     int
	MaxOpen     int
	MaxLifeTime time.Duration
}

// DefaultConnConfig 默认连接参数
var DefaultConnConfig = &ConnConfig{
	MaxIdle:     16,
	MaxOpen:     64,
	MaxLifeTime: 10 * time.Minute,
}

// NewModel return
func NewModel(master, slave *DBConfig, log logger, debug, unscoped bool) (*model, error) {
	masterConn, err := newGormDB(master, debug, unscoped)
	if err != nil {
		return nil, err
	}
	masterConn.SetLogger(log)

	slaveConn, err := newGormDB(slave, debug, unscoped)
	if err != nil {
		return nil, err
	}
	slaveConn.SetLogger(log)

	return &model{
		master:   masterConn,
		slave:    slaveConn,
		logger:   log,
		debug:    debug,
		unscoped: unscoped,
	}, nil
}

// newGormDB return new gorm DB
func newGormDB(cfg *DBConfig, debug, unscoped bool) (*gorm.DB, error) {
	conn, err := gorm.Open("mysql", cfg.DSN)
	if err != nil {
		return nil, err
	}

	var connCfg *ConnConfig
	if cfg.Conn == nil {
		connCfg = DefaultConnConfig
	} else {
		connCfg = cfg.Conn
	}

	// set conn attr
	conn.DB().SetMaxIdleConns(connCfg.MaxIdle)
	conn.DB().SetMaxOpenConns(connCfg.MaxOpen)
	conn.DB().SetConnMaxLifetime(connCfg.MaxLifeTime)

	if debug {
		conn = conn.Debug()
	}

	if unscoped {
		conn = conn.Unscoped()
	}

	return conn, nil
}

// SetLoggerCtxKey 设置上下文中存储 loogger 对象的 key
// note: 若后面要使用 MasterLoggerFormContext/SlaveLoggerFromContext 方法，则需在前面调用此方法设置 logger 对象在上下文中存储的 key
func (m *model) SetLoggerCtxKey(key interface{}) {
	if key == nil {
		panic("logger key in context can not be nil")
	}

	m.ctxLoggerKey = key
}

// Master return model.master without search conditions
func (m *model) Master() *gorm.DB {
	return m.clone().master
}

// Slave return model.Slave without search conditions
func (m *model) Slave() *gorm.DB {
	return m.clone().slave
}

// MasterWithContext return new model.master with log
func (m *model) MasterWithLogger(log logger) *gorm.DB {
	nm := m.clone()
	nm.master.SetLogger(log)

	return m.master
}

// MasterLoggerFromContext return new model.master with ctx logger
func (m *model) MasterLoggerFromContext(ctx context.Context) *gorm.DB {
	logEntry := LoggerFromContext(ctx, m.ctxLoggerKey, m.logger)
	return m.MasterWithLogger(logEntry)
}

// SlaveWithLogger return new gorm.DB with log
func (m *model) SlaveWithLogger(log logger) *gorm.DB {
	nm := m.clone()
	nm.slave.SetLogger(log)

	return nm.slave
}

// SlaveLoggerFromContext return new model.slave with ctx logger
func (m *model) SlaveLoggerFromContext(ctx context.Context) *gorm.DB {
	logEntry := LoggerFromContext(ctx, m.ctxLoggerKey, m.logger)
	return m.SlaveWithLogger(logEntry)
}

// clone return a new model with it's new db connection without search conditions
func (m *model) clone() *model {
	nm := new(model)
	// clone 其他参数
	*nm = *m
	// clean search conditions
	nm.master = m.master.New()
	if nm.unscoped {
		nm.master = nm.master.Unscoped()
	}

	nm.slave = m.slave.New()
	if nm.unscoped {
		nm.slave = nm.slave.Unscoped()
	}

	return nm
}
