// @Author Ben.Zheng
// @DateTime 2022/7/20 7:39 PM

package client

import "github.com/gofiber/fiber/v2"

const CLIENT_TAGS_PRE = "/tags"

func TagGroups(app *fiber.App) *fiber.App {

	tagGroup := app.Group(CLIENT_TAGS_PRE)
	{
		tagGroup.Get("")
	}

	return app
}
