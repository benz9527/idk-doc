// @Author Ben.Zheng
// @DateTime 2022/8/4 9:33

package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/pprof"
)

func Pprof() fiber.Handler {
	return pprof.New(pprof.Config{})
}
