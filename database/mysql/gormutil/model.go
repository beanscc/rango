package gormutil

import (
	"context"
	"errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Model model struct
type Model struct {
	// master for write
	master *gormDB
	// slave for read
	slave *gormDB
}

// Config gorm db config
type Config struct {
	// master db config
	Master *DBConfig
	// slave db config
	Slave *DBConfig
}

// NewModel return new model
func NewModel(cfg *Config) (*Model, error) {
	if cfg == nil {
		return nil, errors.New("nil Config")
	}

	if cfg.Master == nil {
		return nil, errors.New("nil Config.Master")
	}

	if cfg.Slave == nil {
		return nil, errors.New("nil Config.Slave")
	}

	masterConn, err := newGormDB(cfg.Master)
	if err != nil {
		return nil, err
	}

	slaveConn, err := newGormDB(cfg.Slave)
	if err != nil {
		return nil, err
	}

	return &Model{
		master: masterConn,
		slave:  slaveConn,
	}, nil
}

// Master return new master conn without search conditions
func (m *Model) Master() *gorm.DB {
	return m.master.DB()
}

// Slave return new slave conn without search conditions
func (m *Model) Slave() *gorm.DB {
	return m.slave.DB()
}

// MasterWithContext return new master conn with log
func (m *Model) MasterWithLogger(log logger) *gorm.DB {
	return m.master.withLogger(log)
}

// SlaveWithLogger return new slave conn with log
func (m *Model) SlaveWithLogger(log logger) *gorm.DB {
	return m.slave.withLogger(log)
}

// MasterWithLoggerFromContext return new master conn with ctx's logger
func (m *Model) MasterWithLoggerFromContext(ctx context.Context) *gorm.DB {
	return m.master.withLoggerFormContext(ctx)
}

// SlaveWithLoggerFromContext return new slave conn with ctx's logger
func (m *Model) SlaveWithLoggerFromContext(ctx context.Context) *gorm.DB {
	return m.slave.withLoggerFormContext(ctx)
}

// MasterWithLoggerCtxKey set master's loggerCtxKey
func (m *Model) MasterWithLoggerCtxKey(key interface{}) {
	m.master = m.master.withLoggerCtxKey(key)
}

// SlaveWithLoggerCtxKey set slave's loggerCtxKey
func (m *Model) SlaveWithLoggerCtxKey(key interface{}) {
	m.slave = m.slave.withLoggerCtxKey(key)
}
