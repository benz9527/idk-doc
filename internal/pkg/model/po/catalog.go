// @Author Ben.Zheng
// @DateTime 2022/7/18 2:19 PM

package po

// CatalogCore
// Makes the persistence object keep simple without belongs to or others reference schemas.
type CatalogCore struct {
	NanoIdFullMode
	WorkspaceId     string `gorm:"column:ws_id;type:varchar(21);index;<-;"`
	GoBackCatalogId string `gorm:"column:go_back_id;type:varchar(21);<-;"`
	Name            string `gorm:"column:name;type:nvarchar(32);uniqueIndex:idk_catalog_name;<-;"`
}

func (c CatalogCore) IsPoCore() bool {
	return true
}

type Catalog[T CatalogCore] struct {
	Core      T          `gorm:"embedded;"`
	Workspace *Workspace `gorm:"foreignKey:WorkspaceId;"`
	Catalog   *T         `gorm:"foreignKey:GoBackCatalogId;"`
}

func (c Catalog[T]) TableName() string {
	return "idk_catalog"
}

func (c Catalog[T]) GetCore() T {
	return c.Core
}
