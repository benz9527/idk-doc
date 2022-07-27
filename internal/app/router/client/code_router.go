// @Author Ben.Zheng
// @DateTime 2022/7/20 7:10 PM

package client

import "github.com/gofiber/fiber/v2"

const CLIENT_CODES_PRE = "/codes"

func CodeGroups(app *fiber.App) *fiber.App {

	codeGroup := app.Group(CLIENT_CODES_PRE)
	{
		codeGroup.Get("/:codeId")
		codeGroup.Get("/:codeId/metadata")
		codeGroup.Get("/:codeId/tabs")
		codeGroup.Get("/:codeId/tabs/:tabId")
	}

	return app
}
