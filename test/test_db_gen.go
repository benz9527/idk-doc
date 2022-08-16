// @Author Ben.Zheng
// @DateTime 2022/8/16 17:29

package test

import (
	"os"
	"path/filepath"

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
