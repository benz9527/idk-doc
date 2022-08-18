// @Author Ben.Zheng
// @DateTime 2022/8/8 21:48

package db

import (
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/benz9527/idk-doc/internal/pkg/intf"
)

const (
	DB_SQLITE3    = "SQLite3"
	DB_PostgreSQL = "PostgreSQL"
	DB_MySQL      = "MySQL"
)

// TODO(Ben) May have to add gorm logger.

type DBClientFactory func(reader intf.IConfigurationReader) intf.IDBInitializer

var (
	DBClients = map[string]DBClientFactory{
		DB_MySQL:      newMySQLDBClient,
		DB_SQLITE3:    newSQLite3DBClient,
		DB_PostgreSQL: newPostgreSQLDBClient,
	}
)

func NewDatabaseClient(cfgReader intf.IConfigurationReader, logger logger.Interface) *gorm.DB {
	dbType, err := cfgReader.GetString("db.type")
	if err != nil {
		panic(err)
	}

	fn, ok := DBClients[dbType]
	if !ok {
		panic(fmt.Errorf("unknown and unsupported database type [%s]", dbType))
	}

	return postOpenDBClient(fn(cfgReader).GetDBClient(), cfgReader, logger)
}

func postOpenDBClient(client *gorm.DB, cfgReader intf.IConfigurationReader, logger logger.Interface) *gorm.DB {
	// If query result is empty, we mask the gorm v2 will occur bug
	// issue: https://github.com/go-gorm/gorm/issues/3789
	_ = client.Callback().Query().Before("gorm:query").
		Register("disable_raise_record_not_found", func(db *gorm.DB) {
			db.Statement.RaiseErrorOnNotFound = false
		})
	// Replace default logger to show SQL info.
	client.Logger = logger

	db, err := client.DB()
	if err != nil {
		panic(fmt.Errorf("unable to set db client connection properties, %v", err))
	}

	maxOpen, err := cfgReader.GetInt64("db.additional.max_open_conns")
	if err != nil {
		maxOpen = 64
	}
	maxIdle, _ := cfgReader.GetInt64("db.additional.max_idle_conns")
	if err != nil {
		maxIdle = 16
	}
	// TODO(Ben) Connection optimized.
	if maxIdle >= maxOpen {
		maxOpen = 64
		maxIdle = 16
	}
	maxLiveTime, _ := cfgReader.GetInt64("db.additional.max_live_time_per_conn")
	if err != nil {
		maxLiveTime = 60 // second
	}

	db.SetMaxOpenConns(int(maxOpen))
	db.SetMaxIdleConns(int(maxIdle))
	db.SetConnMaxIdleTime(30 * time.Second)
	db.SetConnMaxLifetime(time.Duration(maxLiveTime) * time.Second)
	return client
}
