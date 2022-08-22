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
	return "idk_code_languages"
}

type CodeTabCore struct {
	BaseMetaNumericId
	BaseMetaCreatedAt
	SubjectId  string `gorm:"column:subj_id;type:varchar(21);<-;"`
	LangTypeId string `gorm:"column:lang_type_id;type:int;<-;"`
	Name       string `gorm:"column:name;type:nvarchar(32);index:idx_code_tab_name;<-;"`
}

type CodeTab[T CodeTabCore] struct {
	Core T                               `gorm:"embedded;"`
	Lang *CodeLang                       `gorm:"foreignKey:LangTypeId;"`
	Map  *SubjectIdMap[SubjectIdMapCore] `gorm:"foreignKey:SubjectId;"`
}

func (c CodeTab[T]) TableName() string {
	return "idk_code_tabs"
}

func (c CodeTab[T]) GetCore() T {
	return c.Core
}

type CodeCore struct {
	AutoIncIdFullMode
	BaseVersionInfo
	TabId   int    `gorm:"column:tab_id;type:int;<-;"`
	Content string `gorm:"column:content;type:text;<-;"`
}

type Code[T CodeCore] struct {
	Core T                     `gorm:"embedded;"`
	Tab  *CodeTab[CodeTabCore] `gorm:"foreignKey:TabId;"`
}

func (c Code[T]) TableName() string {
	return "idk_codes"
}

func (c Code[T]) GetCore() T {
	return c.Core
}
