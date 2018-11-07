package gormutil

import (
	"context"
	"time"

	"github.com/jinzhu/gorm"
)

// gormDB gormDB
type gormDB struct {
	// db gorm.DB
	db *gorm.DB
	// debug gorm debug mode
	debug bool
	// unscoped (建议true，防止表结构中有 deleted_at 字段时，gorm 加该查询条件，不容易发现) return all record including deleted record, refer Soft Delete https://jinzhu.github.io/gorm/crud.html#soft-delete
	unscoped bool
	// log logger entry
	log logger
	// ctx context
	ctx context.Context
	// ctxLoggerKey logger 对象保存在 ctx 中的 key
	ctxLoggerKey interface{}
}

// DBConfig db 配置 参数
type DBConfig struct {
	// db dsn
	DSN string
	// db conn config
	Conn *ConnConfig
	// logger entry
	Log logger
	// debug gorm debug mode
	Debug bool
	// unscoped (建议true，防止表结构中有 deleted_at 字段时，gorm 加该查询条件，不容易发现) return all record including deleted record, refer Soft Delete https://jinzhu.github.io/gorm/crud.html#soft-delete
	Unscoped bool
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

// DefaultLogger default custom logger
var DefaultLogger = NewLogrusEntry()

// newGormDB return new gorm DB
func newGormDB(cfg *DBConfig) (*gormDB, error) {
	conn, err := gorm.Open("mysql", cfg.DSN)
	if err != nil {
		return nil, err
	}

	if err := conn.DB().Ping(); err != nil {
		return nil, err
	}

	connCfg := cfg.Conn
	if cfg.Conn == nil {
		connCfg = DefaultConnConfig
	}

	// set conn attr
	conn.DB().SetMaxIdleConns(connCfg.MaxIdle)
	conn.DB().SetMaxOpenConns(connCfg.MaxOpen)
	conn.DB().SetConnMaxLifetime(connCfg.MaxLifeTime)

	// Debug mode
	if cfg.Debug {
		conn = conn.Debug()
	}

	if cfg.Unscoped {
		conn = conn.Unscoped()
	}

	logEntry := cfg.Log
	if cfg.Log == nil {
		logEntry = DefaultLogger
	}
	// set gorm logger
	conn.SetLogger(logEntry)

	return &gormDB{
		db:       conn,
		debug:    cfg.Debug,
		unscoped: cfg.Unscoped,
		log:      logEntry,
		ctx:      context.Background(),
	}, nil
}

// DB return new gorm.DB without search conditions
func (g *gormDB) DB() *gorm.DB {
	// new will clean search conditions
	db := g.db.New()

	if g.unscoped {
		db = db.Unscoped()
	}

	return db
}

// Context return context
func (g *gormDB) Context() context.Context {
	if g.ctx == nil {
		return context.Background()
	}

	return g.ctx
}

// WithContext return new gormDB with ctx
func (g *gormDB) WithContext(ctx context.Context) *gormDB {
	if ctx == nil {
		panic("nil context")
	}

	ng := new(gormDB)
	*ng = *g
	ng.ctx = ctx

	return ng
}

// setLogger set gormDB log
func (g *gormDB) setLogger(log logger) {
	g.log = log
}

// withLogger set gorm.DB with new logger, without change gormDB's log
func (g *gormDB) withLogger(log logger) *gorm.DB {
	db := g.DB()
	db.SetLogger(log)
	return db
}

// withLoggerCtxKey 设置上下文中存储 loogger 对象的 key
// note: 若后面要使用 MasterLoggerFormContext/SlaveLoggerFromContext 方法，则需在前面调用此方法设置 logger 对象在上下文中存储的 key
func (g *gormDB) withLoggerCtxKey(key interface{}) *gormDB {
	if key == nil {
		panic("nil ctxLoggerKey")
	}

	ng := new(gormDB)
	*ng = *g
	ng.ctxLoggerKey = key

	return ng
}

// withLoggerFormContext 使用上下文中的logger对象作为本次连接的 logger
func (g *gormDB) withLoggerFormContext(ctx context.Context) *gorm.DB {
	logEntry := LoggerFromContext(ctx, g.ctxLoggerKey, g.log)
	return g.withLogger(logEntry)
}
