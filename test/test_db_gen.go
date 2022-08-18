// @Author Ben.Zheng
// @DateTime 2022/8/16 17:29

package test

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
	"gorm.io/gorm"

	"github.com/benz9527/idk-doc/internal/pkg/consts"
	"github.com/benz9527/idk-doc/internal/pkg/db"
	"github.com/benz9527/idk-doc/internal/pkg/file"
)

func genDevTestSQLiteDB() *gorm.DB {
	v := viper.New()
	v.Set("db.type", db.DB_SQLITE3)
	v.Set("db.name", filepath.Join(os.TempDir(), "idk_test.db"))
	v.Set("db.init.create_db", consts.DB_CREATION_ALWAYS)
	v.Set("db.additional.max_idle_conns", 10)
	v.Set("db.additional.max_open_conns", 32)
	v.Set("db.additional.max_live_time_per_conn", 60)
	reader := file.NewSimpleReader(v)
	return db.NewDatabaseClient(reader)
}

func callSQLFiles(dbClient *gorm.DB, files ...string) error {
	abs, err := filepath.Abs(".")
	if err != nil {
		return err
	}

	if len(files) <= 0 {
		return nil
	}

	for _, f := range files {
		content, err := os.ReadFile(filepath.Join(abs, "sqls", f))
		if err != nil {
			return err
		}
		tx := dbClient.Begin()
		hasRollback := false
		for _, sql := range strings.Split(string(content), ";") {
			if err = tx.Exec(sql + ";").Error; err != nil {
				tx.Rollback()
				hasRollback = true
				break
			}
		}
		if !hasRollback {
			if err = tx.Commit().Error; err != nil {
				return err
			}
		}
	}
	return nil
}