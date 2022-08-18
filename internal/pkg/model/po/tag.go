// @Author Ben.Zheng
// @DateTime 2022/7/18 2:20 PM

package po

// Tag
// @field Type Example "00-000" first 2 char present the idk-doc assets type,
// the second 4 char present the sub source type.
type Tag struct {
	AutoIncIdFullMode
	Label string `gorm:"column:label;type:nvarchar(32);index:idx_tag_label_type_mixed;<-;"`
	Type  string `gorm:"column:type;type:char(6);index:idx_tag_label_type_mixed;<-;"`
}

func (t Tag) TableName() string {
	return "idk_tag"
}

// TagMap
// Without the foreign key definition directly.
// Actually, idk will use the table of "idk_tag" with other assets tables'
// primary key to compose the data.
type TagMap struct {
	BaseMetaNumericId
	AssetsId string `gorm:"column:assets_id;type:varchar(21);index:idx_tag_map_mixed;<-;"`
	TagId    uint   `gorm:"column:tag_id;type:int;index:idx_tag_map_mixed;<-;"`
}

func (t TagMap) TableName() string {
	return "idk_tag_map"
}
