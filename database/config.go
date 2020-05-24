package database

import "time"

// Config db 配置 参数
type Config struct {
	DriverType  string        `json:"driver_type" yaml:"driver_type"`   // database driver type name; eg: mysql
	DSN         string        `json:"dsn" yaml:"dsn"`                   // database 连接字符串；eg: user:password@tcp(127.0.0.1:3306)/database?charset=utf8mb4&parseTime=True&loc=Local
	MaxIdle     int           `json:"max_idle" yaml:"max_idle"`         // the maximum number of connections in the idle connection pool
	MaxOpen     int           `json:"max_open" yaml:"max_open"`         // the maximum number of open connections to the database
	MaxLifeTime time.Duration `json:"max_lifetime" yaml:"max_lifetime"` // the maximum amount of time a connection may be reused
}
