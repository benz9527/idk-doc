// @Author Ben.Zheng
// @DateTime 2022/8/4 17:36

package db

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/benz9527/idk-doc/internal/pkg/consts"
	"github.com/benz9527/idk-doc/internal/pkg/intf"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type sqlite3 struct {
	client *gorm.DB
	dbPath string
}

func newSQLite3DBClient(cfgReader intf.IConfigurationReader) intf.IDBInitializer {
	dbPath, err := cfgReader.GetString("db.name")
	if err != nil {
		// TODO(Ben) Need more details.
		panic(err)
	}

	var completedDbPath string
	if match, err := regexp.MatchString(`^([\.\.]/?)+`, dbPath); match && err == nil {
		panic(fmt.Errorf("not support multiple upper dir relative path [%s]", dbPath))
	}

	if strings.HasPrefix(dbPath, "./") {
		rwd, err := cfgReader.GetString(consts.APP_ROOT_WORKING_DIR)
		if err != nil {
			panic(err)
		}
		completedDbPath = rwd + dbPath[2:] // Remove the "./"
	} else {
		completedDbPath = dbPath
	}

	initializer := &sqlite3{}

	_, err = os.Stat(completedDbPath)
	initializer.dbPath = dbPath
	cond, _ := cfgReader.GetString("db.init.create_db")
	status, err := initializer.ShouldCreateDB(cond, os.IsNotExist(err))
	if err != nil {
		panic(err)
	}

	dbClient, err := gorm.Open(sqlite.Dialector{
		DriverName: sqlite.DriverName,
		DSN:        completedDbPath,
	})
	if err != nil {
		panic(err)
	}
	initializer.client = dbClient

	if err = initializer.InitSchema(status); err != nil {
		panic(fmt.Errorf("init empty sqlite with sqlite_master schema with error, %v", err))
	}

	return initializer
}

func (s *sqlite3) GetDBClient() *gorm.DB {
	return s.client
}

func (s *sqlite3) ShouldCreateDB(condition string, notPresent bool) (consts.DBInitStatus, error) {

	var hasDeleted bool
	switch condition {
	case consts.DB_CREATION_ALWAYS:
		if !notPresent {
			if err := os.Remove(s.dbPath); err != nil {
				return consts.ONLY_REMOVED, err
			}
			hasDeleted = true
		}
		fallthrough
	case consts.DB_CREATION_NEVER,
		consts.DB_CREATION_IF_NOT_PRESENT:
		fallthrough
	default:
		// It is unacceptable for SQLite3 never to create DB file event if there isn't present.
		if !notPresent && !hasDeleted {
			return consts.NEVER_CHANGED, nil
		}

		dir := filepath.Dir(s.dbPath)
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			if err = os.Mkdir(dir, 0666); err != nil {
				return consts.RECREATE_WITH_ERR, fmt.Errorf("create sqlite3 db directory with error, %v", err)
			}
		}
		if _, err := os.Create(s.dbPath); err != nil {
			return consts.RECREATE_WITH_ERR, fmt.Errorf("create sqlite3 db file with error, %v", err)
		}
	}

	return consts.RECREATED, nil
}

func (s *sqlite3) InitSchema(status consts.DBInitStatus) error {
	if status != consts.RECREATED {
		return nil
	}

	tx := s.GetDBClient().Begin()
	if err := tx.Exec(`CREATE TABLE sqlite_master
(
    type TEXT,
    name TEXT,
    tbl_name TEXT,
    rootpage INT,
    sql TEXT,
)
`).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
