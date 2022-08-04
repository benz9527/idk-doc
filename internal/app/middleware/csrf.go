// @Author Ben.Zheng
// @DateTime 2022/8/3 22:15

package middleware

// Cross-site request forgery
// https://developer.mozilla.org/en-US/docs/Glossary/CSRF
// https://tech.meituan.com/2018/10/11/fe-security-csrf.html

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
)

func CSRF() fiber.Handler {
	return csrf.New(csrf.Config{
		// Fetch from request:
		// - "header:<xxx>"
		// - "query:<xxx>"
		// - "param:<xxx>"
		// - “form:<xxx>”
		// - "cookie:<xxx>"
		KeyLookup:      "header:X-CSRF-TOKEN",
		CookieName:     "csrf_",
		CookieSameSite: "Strict",
		Expiration:     1 * time.Hour,
		KeyGenerator: func() string {
			return ""
		},
	})
}
