// @Author Ben.Zheng
// @DateTime 2022/7/18 2:15 PM

package po

type MarkdownMeta struct {
	BaseMetaStringId
}

type Markdown struct {
	BaseMetaStringId
	Title   string `gorm:"column:title;<-"`
	Content string `gorm:"column:content;<-"`
}

type MarkdownTemplate struct {
	ContentTemplate string
}
