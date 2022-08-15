// @Author Ben.Zheng
// @DateTime 2022/8/15 9:04

package biz

import (
	"go.uber.org/zap"

	"github.com/benz9527/idk-doc/internal/pkg/intf"
)

type StoreBiz struct {
	logger *zap.SugaredLogger
}

func NewStoreBiz(logger *zap.SugaredLogger) intf.BizPrototype[StoreBiz] {
	return func() StoreBiz {
		return StoreBiz{
			logger: logger,
		}
	}
}

func (s StoreBiz) BizError() error {
	return nil
}
