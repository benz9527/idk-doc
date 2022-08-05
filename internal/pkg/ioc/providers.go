// @Author Ben.Zheng
// @DateTime 2022/8/5 14:27

package ioc

import (
	"github.com/benz9527/idk-doc/internal/pkg/db"

	"go.uber.org/fx"
)

var Options = make([]fx.Option, 0, 256)

func init() {
	Options = append(Options, fx.Provide(db.NewSQLite3DBClient))
}
