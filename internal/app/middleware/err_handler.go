// @Author Ben.Zheng
// @DateTime 2022/7/19 10:07 AM

package middleware

import "github.com/gofiber/fiber/v2"

func DefaultErrHandleware() func(ctx *fiber.Ctx, err error) error {
	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}
		ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
		return ctx.Status(code).SendString(err.Error())
	}
}
