// @Author Ben.Zheng
// @DateTime 2022/7/19 9:50 AM

package middleware

// Cross-Origin Resource Sharing
// https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/benz9527/idk-doc/lib/collections"
)

func CORS() fiber.Handler {
	headers, _ := getAllowHeaders()
	exposed, _ := getExposeHeaders()
	methods, _ := getAllowMethods()
	return cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     headers,
		AllowMethods:     methods,
		AllowCredentials: true,
		ExposeHeaders:    exposed,
		Next:             nil,
	})
}

func getAllowHeaders(headers ...string) (string, int) {
	defaultHeaders := []string{
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
	defaultHeaders = additionalAppend(defaultHeaders, headers...)
	return strings.Join(defaultHeaders, ", "), len(defaultHeaders)
}

func getExposeHeaders(headers ...string) (string, int) {
	defaultHeaders := []string{
		"Content-Length",
		"Access-Control-Allow-Origin",
		"Access-Control-Allow-Headers",
		"Content-Type",
	}
	defaultHeaders = additionalAppend(defaultHeaders, headers...)
	return strings.Join(defaultHeaders, ", "), len(defaultHeaders)
}

// getAllowMethods
func getAllowMethods(methods ...string) (string, int) {
	defaultMethods := []string{
		"GET", "POST", "DELETE", "PATCH", "PUT",
	}
	defaultMethods = additionalAppend(defaultMethods, methods...)
	return strings.Join(defaultMethods, ", "), len(defaultMethods)
}

func additionalAppend(original []string, others ...string) []string {
	// Currently, we don't check if a string is available and correct for caller.
	if len(others) > 0 {
		for _, h := range others {
			if strings.Contains(h, ",") {
				unhandledList := strings.Split(h, ",")
				for _, item := range unhandledList {
					res := strings.Trim(item, " ")
					if len(res) > 0 {
						original = append(original, res)
					}
				}
			} else if strings.Contains(h, " ") {
				res := strings.Trim(h, " ")
				if len(res) > 0 {
					original = append(original, res)
				}
			}
		}
	}
	// After do distinct operation, the final result is disordered.
	original = collections.SliceStream[string](original).Distinct().ToSlice()
	return original
}
