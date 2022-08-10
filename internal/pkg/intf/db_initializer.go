// @Author Ben.Zheng
// @DateTime 2022/8/9 20:17

package intf

import (
	"github.com/benz9527/idk-doc/internal/pkg/consts"

	"gorm.io/gorm"
)

type IDBInitializer interface {
	GetDBClient() *gorm.DB
	ShouldCreateDB(condition string, notPresent bool) (consts.DBInitStatus, error)
	InitSchema(status consts.DBInitStatus) error
}
