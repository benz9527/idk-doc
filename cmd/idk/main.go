// @Author Ben.Zheng
// @DateTime 2022/7/19 9:31 AM

package main

import (
	"context"
	"fmt"
	"github.com/benz9527/idk-doc/internal/ioc"
	"os"

	"github.com/benz9527/idk-doc/internal/app/cli"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func main() {
	IDK(os.Args[1:])
}

func IDK(args []string) {
	opts := cli.NewBootOptions()
	if err := opts.Parse(args); err != nil {
		// Application exit with error.
		os.Exit(1)
	}

	ioc.Options = append(ioc.Options, fx.Provide(func() *fiber.App {
		return fiber.New()
	}))

	ioc.Options = append(ioc.Options, fx.Invoke(func(srv *fiber.App, lifecycle fx.Lifecycle) {
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

	ioc.Init(opts.FilePath)

	app := fx.New(
		ioc.Options...,
	)

	app.Run()
}
