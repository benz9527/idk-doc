// @Author Ben.Zheng
// @DateTime 2022/8/5 14:27

package ioc

import (
	"os"
	"sync"

	"github.com/benz9527/idk-doc/internal/pkg/consts"
	"github.com/benz9527/idk-doc/internal/pkg/db"
	"github.com/benz9527/idk-doc/internal/pkg/file"
	"github.com/benz9527/idk-doc/internal/pkg/intf"

	"github.com/spf13/viper"
	"go.uber.org/fx"
)

var (
	Options = make([]fx.Option, 0, 256)
	once    = sync.Once{}
)

func Init(filepath string) {
	once.Do(func() {
		// Without any another dependencies. Just a simple constructor.
		// Viper dep
		Options = append(Options, fx.Provide(func() *viper.Viper {
			// Make viper as idk-doc application global const variable storage.
			v := viper.New()
			if wd, err := os.Getwd(); err == nil {
				// At main entrypoint, it will be got the real root working directory of applicaiton.
				v.Set(consts.APP_ROOT_WORKING_DIR, wd)
			} else {
				// Set as empty works for unit tests.
				v.Set(consts.APP_ROOT_WORKING_DIR, consts.EMPTY_DIR)
			}
			return v
		}))
		Options = append(Options, fx.Provide(func(viper *viper.Viper) intf.IConfigurationReader {
			return file.NewConfigurationReader(viper, filepath)
		}))

		// Contains dependencies reference. Uses invoke function to finish construction and DI.
		// TODO(Ben) Gorm dep
		Options = append(Options, fx.Provide(db.NewDatabaseClient))
	})

}
