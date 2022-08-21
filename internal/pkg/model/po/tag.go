// @Author Ben.Zheng
// @DateTime 2022/7/18 2:20 PM

package po

// Tag
// @field Type Example "00-000" first 2 char present the idk-doc assets type,
// the second 4 char present the sub source type.
type Tag struct {
	AutoIncIdFullMode
	Label string `gorm:"column:label;type:nvarchar(32);uniqueIndex:idx_tag_label_type_mixed;<-;"`
	Type  string `gorm:"column:type;type:char(6);uniqueIndex:idx_tag_label_type_mixed;<-;"`
}

func (t Tag) TableName() string {
	return "idk_tags"
}

// TagMapCore
// Without the foreign key definition directly.
// Actually, idk will use the table of "idk_tags" with other assets tables'
// primary key to compose the data.
type TagMapCore struct {
	BaseMetaNumericId
	BaseMetaCreatedAt
	BaseMetaDeletedAt
	AssetsId string `gorm:"column:assets_id;type:varchar(21);index:idx_tag_map_mixed;<-;"`
	TagId    uint   `gorm:"column:tag_id;type:int;index:idx_tag_map_mixed;<-;"`
}

type TagMap[T TagMapCore] struct {
	Core T    `gorm:"embedded;"`
	Tag  *Tag `gorm:"foreignKey:TagId;"`
}

func (t TagMap[T]) TableName() string {
	return "idk_tag_map"
}

func (t TagMap[T]) GetCore() T {
	return t.Core
}
