package gorm

import (
	"context"

	"github.com/beanscc/rango/database"
	"github.com/jinzhu/gorm"
)

// Config
// 根据 Driver 不同，在引用侧自行按 driver 类型，导入驱动注册包；如mysql驱动，则 import _ "github.com/go-sql-driver/mysql"
type Config struct {
	database.Config
	Debug    bool
	Unscoped bool
}

// Conn gorm.DB conn
type Conn struct {
	master          *gorm.DB
	slave           *gorm.DB
	logWriter       func(ctx context.Context) LogWriter
	withContextFunc func(ctx context.Context, db *gorm.DB)
}

// NewConn return new Conn
func NewConn(master, slave *Config) (*Conn, error) {
	mConn, err := newGormDB(master)
	if err != nil {
		return nil, err
	}

	sConn, err := newGormDB(slave)
	if err != nil {
		return nil, err
	}
	return &Conn{
		master: mConn,
		slave:  sConn,
	}, nil
}

// Master 返回 master 连接
func (c *Conn) Master() *gorm.DB {
	return c.master
}

// Slave 返回 slave 连接
func (c *Conn) Slave() *gorm.DB {
	return c.slave
}

// SetLogWriter 设置 LogWriter
func (c *Conn) SetLogWriter(fn func(ctx context.Context) LogWriter) {
	c.logWriter = fn
}

func (c *Conn) WithContextFunc(fn func(ctx context.Context, db *gorm.DB)) {
	c.withContextFunc = fn
}

// MasterWithContext 返回 master 的新连接
func (c *Conn) MasterWithContext(ctx context.Context) *gorm.DB {
	n := c.Master().New()
	c.withContext(ctx, n)
	return n
}

// MasterWithContext 返回 slave 的新连接
func (c *Conn) SlaveWithContext(ctx context.Context) *gorm.DB {
	n := c.Slave().New()
	c.withContext(ctx, n)
	return n
}

func (c *Conn) withContext(ctx context.Context, db *gorm.DB) {
	if c.logWriter != nil {
		logger := c.logWriter(ctx)
		db.SetLogger(NewLogger(logger))
	}

	if c.withContextFunc != nil {
		c.withContextFunc(ctx, db)
	}
}

// newGormDB return new gorm.DB
func newGormDB(c *Config) (*gorm.DB, error) {
	conn, err := gorm.Open(c.DriverType, c.DSN)
	if err != nil {
		return nil, err
	}

	if err := conn.DB().Ping(); err != nil {
		return nil, err
	}

	// set conn attr
	conn.DB().SetMaxIdleConns(c.MaxIdle)
	conn.DB().SetMaxOpenConns(c.MaxOpen)
	conn.DB().SetConnMaxLifetime(c.MaxLifeTime)

	// Debug mode
	if c.Debug {
		conn = conn.Debug()
	}

	if c.Unscoped {
		conn = conn.Unscoped()
	}

	return conn, nil
}
