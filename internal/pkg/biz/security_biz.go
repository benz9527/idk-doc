// @Author Ben.Zheng
// @DateTime 2022/8/4 10:11

package biz

import (
	"go.uber.org/zap"

	"github.com/benz9527/idk-doc/internal/pkg/intf"
)

type SecurityBiz struct {
	logger *zap.SugaredLogger
}

func NewSecurityBiz(logger *zap.SugaredLogger) intf.BizPrototype[SecurityBiz] {
	return func() SecurityBiz {
		return SecurityBiz{
			logger: logger,
		}
	}
}

func (s SecurityBiz) BizError() error {
	return nil
}
