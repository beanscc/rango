package gormutil

import (
	"context"

	"github.com/beanscc/rango/database"
	"github.com/beanscc/rango/log"
	"github.com/jinzhu/gorm"
)

// Config
// 根据 Driver 不同，在引用侧自行按 driver 类型，导入驱动注册包；如mysql驱动，则 import _ "github.com/go-sql-driver/mysql"
type Config struct {
	database.Config
	Debug    bool
	Unscoped bool
}

// DB gorm.DB
type DB struct {
	*gorm.DB
	contextFunc func(ctx context.Context, db *gorm.DB)
}

// NewDB new DB
func NewDB(c *Config) (*DB, error) {
	db, err := newGormDB(c)
	if err != nil {
		return nil, err
	}

	return &DB{
		DB:          db,
		contextFunc: nil,
	}, nil
}

// WithContext return a new *gorm.DB with context
func (d *DB) WithContext(ctx context.Context) *gorm.DB {
	n := d.DB.New()
	if d.contextFunc != nil {
		d.contextFunc(ctx, n)
	}
	return n
}

// SetContextFunc 设置 gorm.DB 的 WithContext 方法
func (d *DB) SetContextFunc(fn func(ctx context.Context, db *gorm.DB)) {
	d.contextFunc = fn
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

func DefaultContextFunc() func(ctx context.Context, db *gorm.DB) {
	return func(ctx context.Context, db *gorm.DB) {
		db.SetLogger(NewLogger(log.FromContext(ctx)))
	}
}
