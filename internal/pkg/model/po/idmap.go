// @Author Ben.Zheng
// @DateTime 2022/8/17 17:12

package po

import "github.com/benz9527/idk-doc/internal/pkg/consts"

// FileIdMapCore
// All type of files have to relate with this intermedia table.
// Different file has diverse table definition.
type FileIdMapCore struct {
	BaseMetaStringId                 // File id.
	Type             consts.FileType `gorm:"column:file_type;type:tinyint(8);<-;"`
	WorkspaceId      string          `gorm:"column:ws_id;type:varchar(21);index:idx_file_ids;priority:10;<-;"`
	CatalogId        string          `gorm:"column:catalog_id;type:varchar(21);index:idx_file_ids;priority:11;<-;"`
}

type FileIdMap[T FileIdMapCore] struct {
	Core      T                     `gorm:"embedded;"`
	Workspace *Workspace            `gorm:"foreignKey:WorkspaceId;"`
	Catalog   *Catalog[CatalogCore] `gorm:"foreignKey:CatalogId;"`
}

func (f FileIdMap[T]) TableName() string {
	return "idk_file_id_map"
}

func (f FileIdMap[T]) GetCore() T {
	return f.Core
}
