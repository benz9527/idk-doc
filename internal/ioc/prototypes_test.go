// @Author Ben.Zheng
// @DateTime 2022/8/15 9:47

package ioc

import (
	"testing"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/benz9527/idk-doc/internal/pkg/biz"
	"github.com/benz9527/idk-doc/internal/pkg/intf"
)

func Test_businesses(t *testing.T) {
	opts := make([]fx.Option, 0, 16)
	// Auto shutdown
	opts = append(opts, fx.Provide(func() (*zap.SugaredLogger, error) {
		dev, err := zap.NewDevelopment()
		if err != nil {
			return nil, err
		}
		return dev.Sugar(), nil
	}))
	opts = append(opts, businesses()...)
	opts = append(opts, fx.Invoke(func(shutdowner fx.Shutdowner) {
		time.Sleep(2 * time.Second)
		_ = shutdowner.Shutdown()
	}))
	opts = append(opts, fx.Invoke(func(deps struct {
		fx.In
		Sec intf.BizPrototype[biz.SecurityBiz] `name:"securityBizProto"`
	}) {
		svc := deps.Sec()
		t.Logf("sec svc 1 %p", &svc)
	}))
	opts = append(opts, fx.Invoke(func(deps struct {
		fx.In
		Sec intf.BizPrototype[biz.SecurityBiz] `name:"securityBizProto"`
	}) {
		svc := deps.Sec()
		t.Logf("sec svc 2 %p", &svc)
	}))
	opts = append(opts, fx.Invoke(func(deps struct {
		fx.In
		Store intf.BizPrototype[biz.StoreBiz] `name:"storeBizProto"`
	}) {
		svc := deps.Store()
		t.Logf("store svc %p", &svc)
	}))
	app := fx.New(opts...)
	app.Run()
}
