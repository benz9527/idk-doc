// @Author Ben.Zheng
// @DateTime 2022/8/15 9:13

package ioc

import (
	"go.uber.org/fx"

	"github.com/benz9527/idk-doc/internal/pkg/biz"
)

func businesses() []fx.Option {
	opts := make([]fx.Option, 0, 128)
	opts = append(opts, fx.Provide(fx.Annotate(
		biz.NewSecurityBiz,
		fx.ResultTags(`name:"securityBizProto"`),
	)))
	opts = append(opts, fx.Provide(fx.Annotate(
		biz.NewStoreBiz,
		fx.ResultTags(`name:"storeBizProto"`),
	)))
	return opts
}
