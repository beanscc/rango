package gormutil

import (
	"context"

	"github.com/jinzhu/gorm"
)

// Conn gorm.DB
type Conn struct {
	master *DB
	slave  *DB
}

// NewConn return new Conn
func NewConn(mc, sc *Config) (*Conn, error) {
	m, err := NewDB(mc)
	if err != nil {
		return nil, err
	}

	s, err := NewDB(sc)
	if err != nil {
		return nil, err
	}

	return &Conn{
		master: m,
		slave:  s,
	}, nil
}

// SetContextFunc 设置 gorm.DB 的 WithContext 方法
func (c *Conn) SetContextFunc(fn func(ctx context.Context, db *gorm.DB)) {
	c.master.SetContextFunc(fn)
	c.slave.SetContextFunc(fn)
}

// Master 返回 master 连接
func (c *Conn) Master() *DB {
	return c.master
}

// Slave 返回 slave 连接
func (c *Conn) Slave() *DB {
	return c.slave
}
