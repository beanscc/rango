package gormutil

import (
	"context"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

var stdModel *model

type model struct {
	// master for write
	master *gorm.DB
	// slave for read
	slave *gorm.DB
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
func NewModel(master, slave *DBConfig) (*model, error) {
	return newModel(master, slave)
}

// newModel return new model with DBConfig
func newModel(master, slave *DBConfig) (*model, error) {
	masterConn, err := newDB(master)
	if err != nil {
		return nil, err
	}

	slaveConn, err := newDB(slave)
	if err != nil {
		return nil, err
	}

	return &model{master: masterConn, slave: slaveConn}, nil
}

// newDB return new gorm DB
func newDB(cfg *DBConfig) (*gorm.DB, error) {
	conn, err := gorm.Open("mysql", cfg.DSN)
	if err != nil {
		return nil, err
	}

	if cfg.Conn != nil {
		// set conn attr
		conn.DB().SetMaxIdleConns(cfg.Conn.MaxIdle)
		conn.DB().SetMaxOpenConns(cfg.Conn.MaxOpen)
		conn.DB().SetConnMaxLifetime(cfg.Conn.MaxLifeTime)
	}

	// unScoped && set logger
	conn = clone(conn)

	return conn, nil
}

// Master return model.master
func (m *model) Master() *gorm.DB {
	return m.master
}

// Slave return model.Slave
func (m *model) Slave() *gorm.DB {
	return m.slave
}

// MasterWithContext return new model.master with log
func (m *model) MasterWithLogger(log *logrus.Entry) *gorm.DB {
	nm := m.clone()
	nm.master.Debug().SetLogger(Logger{log})

	return m.master
}

// MasterWithContext return new model.master with ctx logger
func (m *model) MasterWithContext(ctx context.Context) *gorm.DB {
	logEntry := FromContext(ctx)
	return m.MasterWithLogger(logEntry)
}

// SlaveWithLogger return new gorm.DB with log
func (m *model) SlaveWithLogger(log *logrus.Entry) *gorm.DB {
	nm := m.clone()
	nm.slave.Debug().SetLogger(Logger{log})

	return nm.slave
}

// SlaveWithContext return new model.slave with ctx logger
func (m *model) SlaveWithContext(ctx context.Context) *gorm.DB {
	logEntry := FromContext(ctx)
	return m.SlaveWithLogger(logEntry)
}

// clone return a new model with it's copy db
func (m *model) clone() *model {
	nm := new(model)
	nm.master = clone(m.master)
	nm.slave = clone(m.slave)

	return nm
}

// clone return a new db connection without search conditions
func clone(db *gorm.DB) *gorm.DB {
	newDB := db.New().Unscoped()
	newDB.SetLogger(Logger{Writer: NewLogrusEntry()})

	return newDB
}
