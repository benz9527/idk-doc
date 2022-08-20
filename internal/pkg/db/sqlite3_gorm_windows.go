// @Author Ben.Zheng
// @DateTime 2022/8/4 17:36

package db

// References:
// https://github.com/glebarez/go-sqlite
// https://github.com/glebarez/sqlite (for gorm)
// https://www.sqlite.org/pragma.html#pragma_journal_mode
// https://www.sqlite.org/pragma.html#pragma_busy_timeout
// Using gorm official sqlite driver requires dev/prod env
// installed with gcc(linux)/mingw(win), otherwise it will
// run with error.

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"github.com/benz9527/idk-doc/internal/pkg/consts"
	"github.com/benz9527/idk-doc/internal/pkg/intf"
)

type sqlite3 struct {
	client *gorm.DB
	dbPath string
}

func newSQLite3DBClient(cfgReader intf.IConfigurationReader) intf.IDBInitializer {
	var (
		rwd, dbPath, execEnv string
		completedDBPath      string
		err                  error
		initializer          = &sqlite3{}
	)

	if dbPath, err = cfgReader.GetString("db.name"); err != nil {
		// TODO(Ben) Need more details.
		panic(err)
	}

	if execEnv, err = cfgReader.GetString("app.env"); err != nil || len(execEnv) == 0 {
		execEnv = consts.APP_RUNTIME_ENV_DEV
	}

	if rwd, err = cfgReader.GetString(consts.APP_ROOT_WORKING_DIR); err != nil && execEnv == consts.APP_RUNTIME_ENV_PROD {
		panic(err)
	}

	completedDBPath, dbPath = getDBPathByEnv(execEnv, dbPath, rwd)
	_, err = os.Stat(completedDBPath)
	initializer.dbPath = dbPath
	cond, _ := cfgReader.GetString("db.init.create_db")
	status, err := initializer.ShouldCreateDB(cond, os.IsNotExist(err))
	if err != nil {
		panic(err)
	}

	dbClient, err := gorm.Open(sqlite.Dialector{
		DriverName: sqlite.DriverName,
		DSN:        completedDBPath + "?_pragma=foreign_keys(1)&_pragma=busy_timeout(5000)&_pragma=journal_mode(WAL)",
	}, &gorm.Config{})
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

// InitSchema
// SQLite3 root schema sqlite_master will be created automatically.
//
// CREATE TABLE sqlite_master
// (
//
//	type TEXT,
//	name TEXT,
//	tbl_name TEXT,
//	rootpage INT,
//	sql TEXT,
//
// )
//
// This is function prepared for migration schema initialize for
// the database first creation for recreation.
func (s *sqlite3) InitSchema(status consts.DBInitStatus) error {
	if status != consts.RECREATED {
		return nil
	}

	tx := s.GetDBClient().Begin()
	// TODO(Ben) Migration schemas.
	return tx.Commit().Error
}

func getDBPathByEnv(env, dbPathFromYaml, rwd string) (completedDBPath string, convertedDBPath string) {
	if env == consts.APP_RUNTIME_ENV_DEV && !filepath.IsAbs(dbPathFromYaml) {
		convertedDBPath = filepath.Join(os.TempDir(), "idk_test.db")
		completedDBPath = convertedDBPath
		return
	}

	// "dev" with abs db path or "prod" handles parts.
	convertedDBPath = strings.ReplaceAll(dbPathFromYaml, "/", "\\")

	if match, err := regexp.MatchString(`^(\.\.\\?)+`, convertedDBPath); match && err == nil {
		panic(fmt.Errorf("not support multiple upper dir relative path [%s]", convertedDBPath))
	}

	if strings.HasPrefix(convertedDBPath, ".\\") {
		completedDBPath = rwd + convertedDBPath[2:] // Remove the "./"
	} else {
		completedDBPath = convertedDBPath
	}

	// Only support the suffix extension named db.
	if _, err := regexp.MatchString(`^[A-Za-z]:\\.*\\.*\.db$`, completedDBPath); err != nil {
		panic(fmt.Errorf("[%s] isn't a real path style format of windows, error %v", completedDBPath, err))
	}

	return completedDBPath, convertedDBPath
}
