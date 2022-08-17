// @Author Ben.Zheng
// @DateTime 2022/8/17 10:25

package test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"gorm.io/gorm"

	"github.com/benz9527/idk-doc/internal/pkg/intf"
	"github.com/benz9527/idk-doc/internal/pkg/model/po"
)

func Test_icore_intf(t *testing.T) {
	asserter := assert.New(t)
	c1 := po.CatalogCore{
		NanoIdFullMode: po.NanoIdFullMode{
			BaseMetaStringId:  po.BaseMetaStringId{Id: "123"},
			BaseMetaCreatedAt: po.BaseMetaCreatedAt{CreatedAt: time.Now().UnixMilli()},
			BaseMetaUpdatedAt: po.BaseMetaUpdatedAt{UpdatedAt: time.Now().UnixMilli()},
			BaseMetaDeletedAt: po.BaseMetaDeletedAt{},
		},
		WorkspaceId:     "1234",
		GoBackCatalogId: "1234",
	}
	var catalog intf.ICore[po.CatalogCore] = po.Catalog[po.CatalogCore]{
		Core: c1,
	}
	c2 := catalog.GetCore()
	asserter.Equal(c1.BaseMetaStringId, c2.BaseMetaStringId)
	asserter.Equal(c1.BaseMetaCreatedAt, c2.BaseMetaCreatedAt)
	asserter.Equal(c1.BaseMetaUpdatedAt, c2.BaseMetaUpdatedAt)
	asserter.Equal(c1.WorkspaceId, c2.WorkspaceId)
	asserter.Equal(c1.GoBackCatalogId, c2.GoBackCatalogId)
}

func Test_init_catalog_in_db(t *testing.T) {
	asserter := assert.New(t)
	var opts []fx.Option

	opts = append(opts, fx.Provide(genDevTestSQLiteDB))
	opts = append(opts, fx.Invoke(func(dbClient *gorm.DB, lifecycle fx.Lifecycle) {
		lifecycle.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {

				asserter.NotNil(dbClient)
				err := dbClient.AutoMigrate(&po.Workspace{}, &po.Catalog[po.CatalogCore]{})
				asserter.Nil(err)

				asserter.Nil(callSQLFiles(dbClient,
					"V1__dev_only_for_ws_init.sql",
					"V2__dev_catalog_with_ws_init.sql",
				))

				var golangCatalogs []po.Catalog[po.CatalogCore]
				dbClient.Where(&po.Catalog[po.CatalogCore]{
					Core: po.CatalogCore{
						WorkspaceId: "9dnARHACYfwylzhwQARON",
					},
				}).Find(&golangCatalogs)
				asserter.Equal(3, len(golangCatalogs))
				asserter.Equal("Golang-catalog1", golangCatalogs[0].GetCore().Name)
				asserter.Equal("Golang-catalog2", golangCatalogs[1].GetCore().Name)
				asserter.Equal("Golang-catalog3", golangCatalogs[2].GetCore().Name)

				golangCatalogs = make([]po.Catalog[po.CatalogCore], 0, 4)
				dbClient.Where("ws_id = ? AND go_back_id = ?",
					"9dnARHACYfwylzhwQARON",
					"z1B1xNlU3RdF3X4TgdSGF").
					Find(&golangCatalogs)
				asserter.Equal(2, len(golangCatalogs))
				asserter.Equal("Golang-catalog2", golangCatalogs[0].GetCore().Name)
				asserter.Equal("Golang-catalog3", golangCatalogs[1].GetCore().Name)

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
