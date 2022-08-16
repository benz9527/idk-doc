// @Author Ben.Zheng
// @DateTime 2022/8/16 15:31

package po

import (
	"gorm.io/gorm"
)

type BaseMetaNumericId struct {
	Id uint `gorm:"column:id;primary;autoIncrement;<-:create;"`
}

// BaseMetaStringId
// IDK-Doc uses the nano-id, so it can be set the len as fixed size 21.
// https://github.com/jaevor/go-nanoid
type BaseMetaStringId struct {
	Id string `gorm:"column:id;primary;size:21;<-:create;"`
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
