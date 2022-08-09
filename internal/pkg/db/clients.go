// @Author Ben.Zheng
// @DateTime 2022/8/8 21:48

package db

import (
	"fmt"

	"github.com/benz9527/idk-doc/internal/pkg/intf"

	"gorm.io/gorm"
)

const (
	DB_SQLITE3    = "SQLite3"
	DB_PostgreSQL = "PostgreSQL"
	DB_MySQL      = "MySQL"
)

// TODO(Ben) May have to add gorm logger.

type DBClientFactory func(reader intf.IConfigurationReader) *gorm.DB

var (
	DBClients = map[string]DBClientFactory{
		DB_MySQL:      newMySQLDBClient,
		DB_SQLITE3:    newSQLite3DBClient,
		DB_PostgreSQL: newPostgreSQLDBClient,
	}
)

func NewDatabaseClient(cfgReader intf.IConfigurationReader) *gorm.DB {
	dbType, err := cfgReader.GetString("db.type")
	if err != nil {
		panic(err)
	}

	fn, ok := DBClients[dbType]
	if !ok {
		panic(fmt.Errorf("unknown and unsupported database type [%s]", dbType))
	}

	return fn(cfgReader)
}
