// @Author Ben.Zheng
// @DateTime 2022/8/8 16:48

package ioc

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"gorm.io/gorm"

	"github.com/benz9527/idk-doc/internal/pkg/consts"
	"github.com/benz9527/idk-doc/internal/pkg/intf"
)

const UPPER_RELATIVE_DIR = "../../"

func Test_viper_provide_with_RWD_as_global(t *testing.T) {
	asserter := assert.New(t)

	// Auto shutdown
	Options = append(Options, fx.Invoke(func(shutdowner fx.Shutdowner) {
		time.Sleep(2 * time.Second)
		_ = shutdowner.Shutdown()
	}))

	Options = append(Options, fx.Invoke(func(v *viper.Viper, lifecycle fx.Lifecycle) {
		lifecycle.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				wd, err := cast.ToStringE(v.Get(consts.APP_ROOT_WORKING_DIR))
				asserter.Nil(err)
				expectedWd, err := filepath.Abs(".")
				asserter.Nil(err)
				asserter.Equal(expectedWd, wd)
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return nil
			},
		})
	}))

	Init(UPPER_RELATIVE_DIR + "conf/idk-boot.yaml")

	app := fx.New(
		Options...,
	)

	app.Run()
}

func Test_readers_read_from_upper_relative_dir(t *testing.T) {
	asserter := assert.New(t)

	// Auto shutdown
	Options = append(Options, fx.Invoke(func(shutdowner fx.Shutdowner) {
		time.Sleep(2 * time.Second)
		_ = shutdowner.Shutdown()
	}))

	Options = append(Options, fx.Invoke(func(cfgReader intf.IConfigurationReader, lifecycle fx.Lifecycle) {
		lifecycle.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				asserter.NotNil(cfgReader)
				typ, err := cfgReader.GetString("db.type")
				asserter.Nil(err)
				asserter.Equal("SQLite3", typ)
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return nil
			},
		})
	}))

	Init(UPPER_RELATIVE_DIR + "conf/idk-boot.yaml")

	app := fx.New(
		Options...,
	)

	app.Run()
}

func Test_gorm_run_with_sqlite3_init(t *testing.T) {
	asserter := assert.New(t)

	// Auto shutdown
	Options = append(Options, fx.Invoke(func(shutdowner fx.Shutdowner) {
		time.Sleep(3 * time.Second)
		_ = shutdowner.Shutdown()
	}))

	Options = append(Options, fx.Invoke(func(dbClient *gorm.DB, lifecycle fx.Lifecycle) {
		lifecycle.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				asserter.NotNil(dbClient)
				var expected int
				tx := dbClient.Begin()
				err := tx.Exec(`CREATE TABLE IF NOT EXISTS tbl_test (num INT)`).Error
				if err != nil {
					tx.Rollback()
					expected = 0
				} else {
					asserter.Nil(tx.Commit().Error)
					expected = 1
				}
				var count int
				dbClient.Raw(`SELECT count(1) FROM sqlite_master WHERE type IN('table')`).
					Scan(&count)
				t.Log(count)
				asserter.Equal(expected, count)
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return nil
			},
		})
	}))

	Init(UPPER_RELATIVE_DIR + "conf/idk-boot.yaml")

	app := fx.New(
		Options...,
	)

	app.Run()
}
