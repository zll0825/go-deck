package gorm

import (
	"gorm.io/driver/mysql"
	gogorm "gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"sync"
	"time"
)

// A Client wrap *gorm.DB with best practices for development
type Client struct {
	*gogorm.DB

	mux       sync.RWMutex
	config    *Config
	traceOnce sync.Once
}

// NewWithLogger creates mysql client with config and logger given.
func New(config *Config) (client *Client, err error) {

	mycfg, err := config.NewMycfg()
	if err != nil {
		return
	}

	loggerConfig := logger.Config{
		SlowThreshold: config.SlowThreshold,
		Colorful:      false,
		LogLevel:      logger.Silent,
	}

	switch config.LogLevel {
	case "silent":
		loggerConfig.LogLevel = logger.Silent
	case "error":
		loggerConfig.LogLevel = logger.Error
	case "warn":
		loggerConfig.LogLevel = logger.Warn
	case "info":
		loggerConfig.LogLevel = logger.Info
	default:
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		loggerConfig,
	)
	db, err := gogorm.Open(mysql.Open(mycfg.FormatDSN()), &gogorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return
	}

	sqlDB, _ := db.DB()

	if config.MaxOpenConns > 0 {
		sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	}
	if config.MaxIdleConns > 0 {
		sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	}
	if config.MaxLifeConns > 0 {
		sqlDB.SetConnMaxLifetime(time.Duration(config.MaxLifeConns) * time.Second)
	}

	err = sqlDB.Ping()
	if err != nil {
		return
	}

	config.mycfg = mycfg

	client = &Client{
		DB:     db,
		config: config,
	}
	return
}

// Select switches to a new database of dbname given by creating a new gorm instance.
func (c *Client) Select(dbname string) (client *Client, err error) {
	c.mux.RLock()
	if c.config.IsEqualDB(dbname) {
		c.mux.RUnlock()

		return c, nil
	}

	config, err := c.config.NewWithDB(dbname)
	if err != nil {
		c.mux.RUnlock()
		return
	}

	name := config.Name()

	// first, try loading a client from default manager
	client, err = DefaultMgr.NewClientWithLogger(name)
	if err == nil {
		c.mux.RUnlock()

		return client, nil
	}

	c.mux.RUnlock()

	// second, register new client for default manager
	c.mux.Lock()
	defer c.mux.Unlock()

	DefaultMgr.Add(name, config)

	return DefaultMgr.NewClientWithLogger(name)
}
