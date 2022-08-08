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
)

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
				wd, err := cast.ToStringE(v.Get(APP_ROOT_WORKING_DIR))
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

	app := fx.New(
		Options...,
	)

	app.Run()
}
