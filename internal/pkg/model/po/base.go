// @Author Ben.Zheng
// @DateTime 2022/8/16 15:31

package po

import (
	"gorm.io/gorm"
)

type BaseMetaNumericId struct {
	Id uint `gorm:"column:id;primaryKey;autoIncrement;<-:create;"`
}

// BaseMetaStringId
// IDK-Doc uses the nano-id, so it can be set the len as fixed size 21.
// IDK-Doc only support the single replica.
// So the id just keep it as simple as possible. (2022.08.17)
// https://github.com/jaevor/go-nanoid
type BaseMetaStringId struct {
	Id string `gorm:"column:id;primaryKey;type:varchar(21);<-:create;"`
}

type BaseMetaCreatedAt struct {
	CreatedAt int64 `gorm:"column:created_at;autoCreateTime:milli;<-"`
}

type BaseMetaUpdatedAt struct {
	UpdatedAt int64 `gorm:"column:updated_at;autoUpdateTime:milli;<-"`
}

type BaseMetaDeletedAt struct {
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type NanoIdFullMode struct {
	BaseMetaStringId
	BaseMetaCreatedAt
	BaseMetaUpdatedAt
	BaseMetaDeletedAt
}

type AutoIncIdFullMode struct {
	BaseMetaNumericId
	BaseMetaCreatedAt
	BaseMetaUpdatedAt
	BaseMetaDeletedAt
}

type BaseVersionInfo struct {
	Version string `gorm:"column:version;type:char(13);index:idx_md_name_ver;<-;"` // VYYYY.MMDD.00
}
