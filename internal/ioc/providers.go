// @Author Ben.Zheng
// @DateTime 2022/8/5 14:27

package ioc

// Doc https://pkg.go.dev/go.uber.org/fx

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"

	"github.com/benz9527/idk-doc/internal/pkg/cache"
	"github.com/benz9527/idk-doc/internal/pkg/consts"
	"github.com/benz9527/idk-doc/internal/pkg/db"
	"github.com/benz9527/idk-doc/internal/pkg/file"
	"github.com/benz9527/idk-doc/internal/pkg/intf"
	"github.com/benz9527/idk-doc/internal/pkg/logger"
)

var (
	Options = make([]fx.Option, 0, 256)
	once    = sync.Once{}
)

func Init(filepath string) {
	once.Do(func() {
		// 1. Without any another dependencies. Just a simple constructor.
		Options = append(Options, fx.Provide(
			fx.Annotate(banner, fx.ResultTags(`name:"banner"`)),
		))
		Options = append(Options, fx.Provide(func() *fiber.App {
			return fiber.New()
		}))
		// Viper dep
		Options = append(Options, fx.Provide(func() *viper.Viper {
			// Make viper as idk-doc application global const variable storage.
			v := viper.New()
			if wd, err := os.Getwd(); err == nil {
				// At main entrypoint, it will be got the real root working directory of application.
				v.Set(consts.APP_ROOT_WORKING_DIR, wd)
			} else {
				// Set as empty works for unit tests.
				v.Set(consts.APP_ROOT_WORKING_DIR, consts.EMPTY_DIR)
			}
			return v
		}))

		// 2. Contains dependencies reference. Uses invoke function to finish construction and DI.
		Options = append(Options, fx.Provide(func(viper *viper.Viper) intf.IConfigurationReader {
			return file.NewConfigurationReader(viper, filepath)
		}))
		Options = append(Options, fx.Provide(logger.NewLogger))
		// Fx logger.
		Options = append(Options, fx.WithLogger(func(cfgReader intf.IConfigurationReader) fxevent.Logger {
			env, err := cfgReader.GetString("app.env")
			if err != nil || len(env) == 0 {
				env = consts.APP_RUNTIME_ENV_DEV
			}

			if env == consts.APP_RUNTIME_ENV_PROD {
				return fxevent.NopLogger
			}
			return &fxevent.ConsoleLogger{
				W: os.Stdout,
			}
		}))
		// Database.
		Options = append(Options, fx.Provide(db.NewDatabaseClient))
		Options = append(Options, fx.Provide(cache.NewRedisClient))

		// 3. Invocations.
		// Application
		Options = append(Options, fx.Invoke(func(deps struct {
			fx.In
			// Fields must be exported.
			Srv    *fiber.App
			Lc     fx.Lifecycle
			Logger *zap.SugaredLogger
			Banner string `name:"banner"`
		}) {
			deps.Lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					go func() {
						fmt.Print(deps.Banner)
						if err := deps.Srv.Listen(":8166"); err != nil {
							_ = fmt.Errorf("failed to listen and serve from server: %v", err)
						}
					}()
					return nil
				},
				OnStop: func(ctx context.Context) error {
					return deps.Srv.Shutdown()
				},
			})
		}))
	})

}

// https://www.bootschool.net/ascii (ansi-shadow)
func banner() string {
	return `
██╗██████╗ ██╗  ██╗     ██████╗  ██████╗  ██████╗
██║██╔══██╗██║ ██╔╝     ██╔══██╗██╔═══██╗██╔════╝
██║██║  ██║█████╔╝█████╗██║  ██║██║   ██║██║     
██║██║  ██║██╔═██╗╚════╝██║  ██║██║   ██║██║     
██║██████╔╝██║  ██╗     ██████╔╝╚██████╔╝╚██████╗
╚═╝╚═════╝ ╚═╝  ╚═╝     ╚═════╝  ╚═════╝  ╚═════╝
`
}
