// @Author Ben.Zheng
// @DateTime 2022/7/18 2:19 PM

package po

// References:
// https://www.sqlite.org/foreignkeys.html

// CatalogCore
// Makes the persistence object keep simple without belongs to or others reference schemas.
type CatalogCore struct {
	NanoIdFullMode
	WorkspaceId     string `gorm:"column:ws_id;type:varchar(21);index:idx_ws_catalog_mixed;<-;"`
	GoBackCatalogId string `gorm:"column:go_back_id;type:varchar(21);index:idx_ws_catalog_mixed;<-;"`
	Name            string `gorm:"column:name;type:nvarchar(32);uniqueIndex:idk_catalog_name;<-;"`
}

func (c CatalogCore) IsPoCore() bool {
	return true
}

// Catalog
// Unable to add self primary key as self table other column's foreign key in gorm.
// Self-associate table creation with SQL:
// CREATE TABLE idk_catalog (
//
//	id varchar(21) PRIMARY KEY,
//	created_at INTEGER,
//	updated_at INTEGER,
//	deleted_at DATETIME,
//	ws_id varchar(21) CONSTRAINT fk_idk_catalog_workspace REFERENCES idk_workspace,
//	go_back_id varchar(21) CONSTRAINT fk_id_catalog_catalog REFERENCES idk_catalog,
//	name nvarchar(32)
//
// )
type Catalog[T CatalogCore] struct {
	Core      T          `gorm:"embedded;"`
	Workspace *Workspace `gorm:"foreignKey:WorkspaceId;"`
	// Catalog   *T         `gorm:"foreignKey:GoBackCatalogId;"` // Only an exists and foreign table could create fk constraint.
}

func (c Catalog[T]) TableName() string {
	return "idk_catalogs"
}

func (c Catalog[T]) GetCore() T {
	return c.Core
}
