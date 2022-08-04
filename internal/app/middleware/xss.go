// @Author Ben.Zheng
// @DateTime 2022/7/19 9:46 AM

package middleware

import "github.com/gofiber/fiber/v2"

func XSS() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ctx.Set("X-XSS-Protection", "1; mode=block")
		ctx.Set("X-Content-Type-Options", "nosniff")
		ctx.Set("X-Download-Options", "noopen")
		ctx.Set("Strict-Transport-Security", "max-age=5184000")
		ctx.Set("X-Frame-Options", "SAMEORIGIN")
		ctx.Set("X-DNS-Prefetch-Control", "off")

		return ctx.Next()
	}
}
