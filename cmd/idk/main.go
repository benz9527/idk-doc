// @Author Ben.Zheng
// @DateTime 2022/7/19 9:31 AM

package main

import (
	"context"
	"fmt"

	"github.com/benz9527/idk-doc/internal/pkg/ioc"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

func main() {

	// TODO(Ben) Should add boot command line part.

	ioc.Options = append(ioc.Options, fx.Provide(func() *fiber.App {
		return fiber.New()
	}))

	ioc.Options = append(ioc.Options, fx.Invoke(func(srv *fiber.App, v *viper.Viper, lifecycle fx.Lifecycle) {
		lifecycle.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				go func() {
					if err := srv.Listen(":8166"); err != nil {
						_ = fmt.Errorf("failed to listen and serve from server: %v", err)
					}
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return srv.Shutdown()
			},
		})
	}))

	app := fx.New(
		ioc.Options...,
	)

	app.Run()
}
