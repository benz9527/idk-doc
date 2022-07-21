// @Author Ben.Zheng
// @DateTime 2022/7/20 7:12 PM

package client

import "github.com/gofiber/fiber/v2"

const CLIENT_REPOSITORIES_PRE = "/repositories"

func RepoGroups(app *fiber.App) *fiber.App {

	repoGroup := app.Group(CLIENT_REPOSITORIES_PRE)
	{
		/*
			TODO(Ben) It should response the default main page info.
		*/
		repoGroup.Get("/:repoId")
		/*
			TODO(Ben) The title of doc, code, issue and so on will be
			treated as a catalog.
		*/
		repoGroup.Get("/:repoId/catalogs")
	}

	return app
}
