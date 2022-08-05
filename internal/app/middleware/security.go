// @Author Ben.Zheng
// @DateTime 2022/8/4 9:59

package middleware

import "github.com/gofiber/fiber/v2"

func Security() fiber.Handler {
	// Request URL access authentication.
	return func(ctx *fiber.Ctx) error {
		return nil
	}
}
