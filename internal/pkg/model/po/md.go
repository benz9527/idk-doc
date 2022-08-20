// @Author Ben.Zheng
// @DateTime 2022/7/18 2:15 PM

package po

// MarkdownRenderTemplate
// Works for Web dynamic CSS render.
type MarkdownRenderTemplate struct {
	NanoIdFullMode
	Name    string `gorm:"column:name;type:nvarchar(32);uniqueIndex:idx_md_render_tmpl_name;<-;"`
	Content string `gorm:"column:content;type:text;<-;"`
}

func (m MarkdownRenderTemplate) TableName() string {
	return "idk_md_render_tmpl"
}

// MarkdownTemplate
// Quickly create markdown file. Default empty template should
// provide also.
type MarkdownTemplate struct {
	NanoIdFullMode
	Name    string `gorm:"column:name;type:nvarchar(32);uniqueIndex:idx_md_tmpl_name;<-;"`
	Content string `gorm:"column:content;type:text;<-;"`
}

func (m MarkdownTemplate) TableName() string {
	return "idk_md_tmpl"
}

// MarkdownCore
// Markdown file could stash temporary changes into OSS.
// After apply stash temp changes as official version content
// will save into database.
type MarkdownCore struct {
	AutoIncIdFullMode
	FileIdMapKey string `gorm:"column:id_map_key;type:varchar(21);<-;"`
	Title        string `gorm:"column:title;type:nvarchar(64);index:idx_md_name_ver;<-;"`
	Content      string `gorm:"column:content;type:text;<-;"`
	Version      string `gorm:"column:version;type:char(13);index:idx_md_name_ver;<-;"` // VYYYY.MMDD.00
}

type Markdown[T MarkdownCore] struct {
	Core T                         `gorm:"embedded;"`
	Map  *FileIdMap[FileIdMapCore] `gorm:"foreignKey:FileIdMapKey;"`
}

func (m Markdown[T]) TableName() string {
	return "idk_md"
}

func (m Markdown[T]) GetCore() T {
	return m.Core
}
