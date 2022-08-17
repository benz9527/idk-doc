// @Author Ben.Zheng
// @DateTime 2022/7/18 2:03 PM

package po

// https://www.sqlite.org/datatype3.html

type Workspace struct {
	NanoIdFullMode
	Name  string `gorm:"column:name;uniqueIndex:idx_ws_name;type:varchar(32);<-"` // Used for request URL as by primary key.
	Icon  string `gorm:"column:icon;type:varchar(512);<-"`                        // ICON URL, displayed with fixed size in WebUI.
	Intro string `gorm:"column:intro;type:varchar(512);<-"`                       // Describes the workspace works for sth.
}

// TableName
// Provide static table name.
func (w Workspace) TableName() string {
	return "idk_workspace"
}
