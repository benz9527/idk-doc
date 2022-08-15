// @Author Ben.Zheng
// @DateTime 2022/8/15 11:00

package biz

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

// go test -run none -bench Benchmark_new_store_biz -benchmem
func Benchmark_new_store_biz(b *testing.B) {
	asserter := assert.New(b)

	dev, err := zap.NewDevelopment()
	asserter.Nil(err)

	proto := NewStoreBiz(dev.Sugar())
	for i := 0; i < b.N; i++ {
		svc := proto()
		_ = svc.BizError()
	}
}
