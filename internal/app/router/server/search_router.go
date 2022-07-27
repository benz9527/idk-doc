// @Author Ben.Zheng
// @DateTime 2022/7/26 17:13

package server

import "github.com/gofiber/fiber/v2"

const SERVER_ADM_MGMT_SEARCH_PRE = "/adm/mgmt/search"

func SearchGroups(app *fiber.App) *fiber.App {

	searchByTagGroup := app.Group(SERVER_ADM_MGMT_SEARCH_PRE + "/tags")
	{
		searchByTagGroup.Get("/:tagName")
	}

	searchByTitleGroup := app.Group(SERVER_ADM_MGMT_SEARCH_PRE + "/title")
	{
		searchByTitleGroup.Get("/:keywords")
	}

	searchByContentGroup := app.Group(SERVER_ADM_MGMT_SEARCH_PRE + "/content")
	{
		searchByContentGroup.Get("/:passage")
	}

	searchByTimestampGroup := app.Group(SERVER_ADM_MGMT_SEARCH_PRE + "/ts")
	{
		searchByTimestampGroup.Get("")
	}

	return app
}
