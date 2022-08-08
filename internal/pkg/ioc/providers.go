// @Author Ben.Zheng
// @DateTime 2022/8/5 14:27

package ioc

import (
	"os"
	"sync"

	"github.com/spf13/viper"
	"go.uber.org/fx"
)

var (
	Options = make([]fx.Option, 0, 256)
	once    = sync.Once{}
)

const (
	APP_ROOT_WORKING_DIR = "APP_ROOT_WORKING_DIR"
	EMPTY_DIR            = ""
)

func init() {
	once.Do(func() {
		// Viper dep
		Options = append(Options, fx.Provide(func() *viper.Viper {
			// Make viper as idk-doc application global const variable storage.
			v := viper.New()
			if wd, err := os.Getwd(); err == nil {
				// At main entrypoint, it will be got the real root working directory of applicaiton.
				v.Set(APP_ROOT_WORKING_DIR, wd)
			} else {
				// Set as empty works for unit tests.
				v.Set(APP_ROOT_WORKING_DIR, EMPTY_DIR)
			}
			return v
		}))
		// TODO(Ben) Gorm SQLite3 dep
		//Options = append(Options, fx.Provide(db.NewSQLite3DBClient))
	})
}
