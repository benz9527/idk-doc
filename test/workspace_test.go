// @Author Ben.Zheng
// @DateTime 2022/8/16 17:24

package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"gorm.io/gorm"

	"github.com/benz9527/idk-doc/internal/pkg/model/po"
)

func Test_init_workspace_in_db(t *testing.T) {
	asserter := assert.New(t)
	var opts []fx.Option

	opts = append(opts, fx.Provide(genDevTestSQLiteDB))
	opts = append(opts, fx.Invoke(func(dbClient *gorm.DB, lifecycle fx.Lifecycle) {
		lifecycle.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {

				asserter.NotNil(dbClient)
				err := dbClient.AutoMigrate(&po.Workspace{})
				asserter.Nil(err)

				asserter.Nil(callSQLFiles(dbClient, "V1__dev_only_for_ws_init.sql"))

				var count int64
				expectedCount := int64(3)
				dbClient.Model(&po.Workspace{}).Count(&count)
				asserter.Equal(expectedCount, count)

				return nil
			},
			OnStop: func(ctx context.Context) error {
				return nil
			},
		})
	}))

	app := fx.New(opts...)
	defer app.Stop(context.Background())
	if err := app.Start(context.Background()); err != nil {
		t.Failed()
	}
}
