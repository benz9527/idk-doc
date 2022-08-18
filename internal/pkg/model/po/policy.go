// @Author Ben.Zheng
// @DateTime 2022/8/4 10:05

package po

// RBACPolicy
// Serv as casbin policies loading source.
type RBACPolicy struct {
	Ptype string `gorm:"column:ptype;type:nvarchar(32);index:idx_rbac_p_type;<-;"`
	V0    string `gorm:"column:v0;type:nvarchar(512);<-;"`
	V1    string `gorm:"column:v1;type:nvarchar(512);<-;"`
	V2    string `gorm:"column:v2;type:nvarchar(512);<-;"`
	V3    string `gorm:"column:v3;type:nvarchar(512);<-;"`
	V4    string `gorm:"column:v4;type:nvarchar(512);<-;"`
	V5    string `gorm:"column:v5;type:nvarchar(512);<-;"`
}

func (R RBACPolicy) TableName() string {
	return "idk_rbac_policies"
}
