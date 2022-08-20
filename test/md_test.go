// @Author Ben.Zheng
// @DateTime 2022/8/18 15:40

package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"gorm.io/gorm"

	"github.com/benz9527/idk-doc/internal/pkg/model/po"
)

func Test_init_md_in_db(t *testing.T) {
	asserter := assert.New(t)
	var opts []fx.Option

	opts = append(opts, fx.Provide(genDevTestSQLiteDB))
	opts = append(opts, fx.Invoke(func(dbClient *gorm.DB, lifecycle fx.Lifecycle) {
		lifecycle.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {

				asserter.NotNil(dbClient)
				err := dbClient.AutoMigrate(
					&po.MarkdownRenderTemplate{},
					&po.MarkdownTemplate{},
					&po.Markdown[po.MarkdownCore]{},
				)
				asserter.Nil(err)

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
