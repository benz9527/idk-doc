// @Author Ben.Zheng
// @DateTime 2022/7/19 9:50 AM

package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CORSware() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     getAllowHeaders(),
		AllowMethods:     getAllowMethods(),
		AllowCredentials: true,
		ExposeHeaders:    getExposeHeaders(),
		Next:             nil,
	})
}

func getAllowHeaders() string {
	headers := []string{
		"Access-Control-Allow-Headers",
		"Authorization",
		"User-Agent",
		"Keep-Alive",
		"Content-Type",
		"X-Requested-With",
		"X-CSRF-Token",
		"AccessToken",
		"Token",
	}
	return strings.Join(headers, ", ")
}

func getExposeHeaders() string {
	headers := []string{
		"Content-Length",
		"Access-Control-Allow-Origin",
		"Access-Control-Allow-Headers",
		"Content-Type",
	}
	return strings.Join(headers, ", ")
}

func getAllowMethods() string {
	methods := []string{
		"GET", "POST", "DELETE", "PATCH", "PUT",
	}
	return strings.Join(methods, ", ")
}
