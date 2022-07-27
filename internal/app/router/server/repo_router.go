// @Author Ben.Zheng
// @DateTime 2022/7/26 15:54

package server

import "github.com/gofiber/fiber/v2"

const SERVER_ADM_MGMT_REPOSITORIES_PRE = "/adm/mgmt/repositories"

func RepoGroups(app *fiber.App) *fiber.App {

	repoGroup := app.Group(SERVER_ADM_MGMT_REPOSITORIES_PRE)
	// Repository management.
	{
		// Get all repositories.
		repoGroup.Get("")
		/*
			TODO(Ben) It should response the default main page info.
		*/
		repoGroup.Get("/:repoId")
		// Create a new repository.
		// TODO(Ben) A new repository should create an unique root catalog at the same time.
		repoGroup.Post("/new")
		// Remove an exists repository.
		repoGroup.Delete("/:repoId")
		// Update metadata of an exists repository.
		repoGroup.Put("/:repoId")
	}

	// Catalog management.
	{
		// Create a new catalog under a repository.
		repoGroup.Post("/:repoId/catalog")
		// Only a catalog, without any document instance associates to, could be removed.
		repoGroup.Delete("/:repoId/catalog/:catalogId")
		// Update a catalog metadata, such as the name, location level.
		// TODO(Ben) A catalog could be moved to others under same repository as sub catalog.
		repoGroup.Put("/:repoId/catalog/:catalogId")
		/*
			TODO(Ben) The title of doc, code, issue and so on will be
			treated as a catalog.
		*/
		repoGroup.Get("/:repoId/catalogs")
	}

	return app
}
