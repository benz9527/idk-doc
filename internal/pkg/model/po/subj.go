// @Author Ben.Zheng
// @DateTime 2022/8/17 17:12

package po

import "github.com/benz9527/idk-doc/internal/pkg/consts"

// SubjectIdMapCore
// All type of files have to relate with this intermedia table.
// Different file has diverse table definition.
type SubjectIdMapCore struct {
	BaseMetaStringId                 // Subject id.
	Type             consts.FileType `gorm:"column:file_type;type:tinyint(8);<-;"`
	WorkspaceId      string          `gorm:"column:ws_id;type:varchar(21);index:idx_subj_ids;priority:10;<-;"`
	CatalogId        string          `gorm:"column:catalog_id;type:varchar(21);index:idx_subj_ids;priority:11;<-;"`
}

type SubjectIdMap[T SubjectIdMapCore] struct {
	Core      T                     `gorm:"embedded;"`
	Workspace *Workspace            `gorm:"foreignKey:WorkspaceId;"`
	Catalog   *Catalog[CatalogCore] `gorm:"foreignKey:CatalogId;"`
}

func (f SubjectIdMap[T]) TableName() string {
	return "idk_subj_id_map"
}

func (f SubjectIdMap[T]) GetCore() T {
	return f.Core
}
