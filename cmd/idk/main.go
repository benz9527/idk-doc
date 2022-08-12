// @Author Ben.Zheng
// @DateTime 2022/7/19 9:31 AM

package main

import (
	"os"

	"go.uber.org/fx"

	"github.com/benz9527/idk-doc/internal/app/cli"
	"github.com/benz9527/idk-doc/internal/ioc"
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

	ioc.Init(opts.FilePath)

	app := fx.New(
		ioc.Options...,
	)

	app.Run()
}
