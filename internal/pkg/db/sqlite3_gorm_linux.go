// @Author Ben.Zheng
// @DateTime 2022/8/4 17:36

package db

import (
	"github.com/benz9527/idk-doc/internal/pkg/consts"
	"os"
	"strings"

	"github.com/benz9527/idk-doc/internal/pkg/intf"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func newSQLite3DBClient(cfgReader intf.IConfigurationReader) *gorm.DB {
	dbPath, err := cfgReader.GetString("db.name")
	if err != nil {
		// TODO(Ben) Need more details.
		panic(err)
	}

	var completedDbPath string
	if strings.HasPrefix(dbPath, "./") {
		rwd, err := cfgReader.GetString(consts.APP_ROOT_WORKING_DIR)
		if err != nil {
			panic(err)
		}
		completedDbPath = rwd + dbPath[2:] // Remove the "./"
	} else {
		completedDbPath = dbPath
	}

	_, err = os.Stat(completedDbPath)
	if os.IsNotExist(err) {
		// TODO(Ben) Need more details.
		panic(err)
	}

	dbClient, err := gorm.Open(sqlite.Dialector{
		DriverName: sqlite.DriverName,
		DSN:        completedDbPath,
	})
	if err != nil {
		panic(err)
	}
	return dbClient
}
