// @Author Ben.Zheng
// @DateTime 2022/8/8 22:14

package db

import (
	"github.com/benz9527/idk-doc/internal/pkg/intf"

	"gorm.io/gorm"
)

func newMySQLDBClient(cfgReader intf.IConfigurationReader) *gorm.DB {
	panic("implement in future")
}
