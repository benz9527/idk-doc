// @Author Ben.Zheng
// @DateTime 2022/8/16 17:24

package test

import (
	"context"
	"os"
	"path/filepath"
	"strings"
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

				abs, err := filepath.Abs(".")
				asserter.Nil(err)
				content, err := os.ReadFile(filepath.Join(abs, "sqls", "V1__dev_only_for_ws_init.sql"))
				asserter.Nil(err)
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
					asserter.Nil(tx.Commit().Error)
				}

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
