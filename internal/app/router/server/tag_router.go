// @Author Ben.Zheng
// @DateTime 2022/7/26 17:14

package server

import "github.com/gofiber/fiber/v2"

const SERVER_ADM_MGMT_TAGS_PRE = "/adm/mgmt/tags"

func TagGroups(app *fiber.App) *fiber.App {

	tagGroup := app.Group(SERVER_ADM_MGMT_TAGS_PRE)
	{
		tagGroup.Get("")
		tagGroup.Delete("/:tagId")
	}

	return app
}
