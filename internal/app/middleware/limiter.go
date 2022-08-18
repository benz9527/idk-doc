// @Author Ben.Zheng
// @DateTime 2022/8/4 9:15

package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func Limiter() fiber.Handler {
	return limiter.New(limiter.Config{
		// Skip white list.
		Next: func(ctx *fiber.Ctx) bool {
			return ctx.IP() == ""
		},
		Max: 20,
		// Request saved a period of time in the mem.
		Expiration: 30 * time.Second,
		KeyGenerator: func(ctx *fiber.Ctx) string {
			return ctx.Get("X-FORWARD-FOR")
		},
		LimitReached: func(ctx *fiber.Ctx) error {
			return ctx.SendFile("./toofasr.html")
		},
		// Calculation:
		// weight_of_prev_win = prev_win_amount_req * (when_new_win / expiration)
		// rate = weight_of_prev_win + cur_win_amount_req
		LimiterMiddleware: limiter.SlidingWindow{},
		Storage:           nil,
	})
}
