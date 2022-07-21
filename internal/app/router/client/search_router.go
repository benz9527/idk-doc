// @Author Ben.Zheng
// @DateTime 2022/7/20 7:08 PM

package client

import "github.com/gofiber/fiber/v2"

const CLIENT_SEARCH_PRE = "/search"

func SearchGroups(app *fiber.App) *fiber.App {

	searchByTagGroup := app.Group(CLIENT_SEARCH_PRE + "/tags")
	{
		searchByTagGroup.Get("/:tagName")
	}

	searchByTitleGroup := app.Group(CLIENT_SEARCH_PRE + "/title")
	{
		searchByTitleGroup.Get("/:keywords")
	}

	searchByContentGroup := app.Group(CLIENT_SEARCH_PRE + "/content")
	{
		searchByContentGroup.Get("/:passage")
	}

	searchByTimestampGroup := app.Group(CLIENT_SEARCH_PRE + "/ts")
	{
		searchByTimestampGroup.Get("")
	}

	return app
}
