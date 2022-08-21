// @Author Ben.Zheng
// @DateTime 2022/8/17 18:44

package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"gorm.io/gorm"

	"github.com/benz9527/idk-doc/internal/pkg/consts"
	"github.com/benz9527/idk-doc/internal/pkg/model/po"
)

func Test_init_subj_id_map_in_db(t *testing.T) {
	asserter := assert.New(t)
	var opts []fx.Option

	opts = append(opts, fx.Provide(genDevTestSQLiteDB))
	opts = append(opts, fx.Invoke(func(dbClient *gorm.DB, lifecycle fx.Lifecycle) {
		lifecycle.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {

				asserter.NotNil(dbClient)
				err := dbClient.AutoMigrate(
					&po.Workspace{},
					&po.Catalog[po.CatalogCore]{},
					&po.SubjectIdMap[po.SubjectIdMapCore]{},
				)
				asserter.Nil(err)

				asserter.Nil(callSQLFiles(dbClient,
					"V1__dev_only_for_ws_init.sql",
					"V2__dev_catalog_with_ws_init.sql",
					"V3__dev_for_subj_init.sql",
				))

				var idMapList []po.SubjectIdMap[po.SubjectIdMapCore]
				err = dbClient.Where("file_type = ?", consts.FILE_TYPE_MD).
					Find(&idMapList).Error
				asserter.Nil(err)
				asserter.Equal(2, len(idMapList))

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
