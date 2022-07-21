// @Author Ben.Zheng
// @DateTime 2022/7/19 9:31 AM

package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Listen(":8166")
}
