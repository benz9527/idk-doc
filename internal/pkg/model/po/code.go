// @Author Ben.Zheng
// @DateTime 8/20/22 5:45 PM

package po

// CodeLang
// Default lang is plain text.
type CodeLang struct {
	AutoIncIdFullMode
	Name string `gorm:"column:name;type:nvarchar(16);uniqueIndex:idx_code_lang;<-;"`
}

func (c CodeLang) TableName() string {
	return "idk_code_lang"
}

type CodeTabCore struct {
	BaseMetaNumericId
	CodeSpaceId string `gorm:"column:code_space_id;type:varchar(21);<-;"`
	LangTypeId  string `gorm:"column:lang_type_id;type:int;<-;"`
	Title       string `gorm:"column:title;type:nvarchar(32);index:idx_code_tab_title;<-;"`
}

type CodeTab[T CodeTabCore] struct {
	Core T                         `gorm:"embedded;"`
	Lang *CodeLang                 `gorm:"foreignKey:LangTypeId;"`
	Map  *FileIdMap[FileIdMapCore] `gorm:"foreignKey:CodeSpaceId;"`
}

func (c CodeTab[T]) TableName() string {
	return "idk_code_tabs"
}

func (c CodeTab[T]) GetCore() T {
	return c.Core
}

type CodeCore struct {
	AutoIncIdFullMode
	CodeSpaceId string `gorm:"column:code_space_id;type:varchar(21);<-;"`
	CodeTabId   int    `gorm:"column:code_tab_id;type:int;<-;"`
	Content     string `gorm:"column:content;type:text;<-;"`
	Version     string `gorm:"column:version;type:char(13);index:idx_md_name_ver;<-;"` // VYYYY.MMDD.00
}

type Code[T CodeCore] struct {
	Core T                         `gorm:"embedded;"`
	Map  *FileIdMap[FileIdMapCore] `gorm:"foreignKey:CodeSpaceId;"`
	Tab  *CodeTab[CodeTabCore]     `gorm:"foreignKey:CodeTabId;"`
}

func (c Code[T]) TableName() string {
	return "idk_codes"
}

func (c Code[T]) GetCore() T {
	return c.Core
}
