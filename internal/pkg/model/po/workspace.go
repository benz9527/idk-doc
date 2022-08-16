// @Author Ben.Zheng
// @DateTime 2022/7/18 2:03 PM

package po

type Workspace struct {
	NanoIdFullMode
	Name  string `gorm:"column:name;size:32;<-"`   // Used for request URL as by primary key.
	Icon  string `gorm:"column:icon;<-"`           // ICON URL, displayed with fixed size.
	Intro string `gorm:"column:intro;size:512;<-"` // Describes the workspace works for sth.
}

// TableName
// Provide static table name.
func (w Workspace) TableName() string {
	return "idk_workspace"
}
