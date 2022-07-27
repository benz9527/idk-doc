// @Author Ben.Zheng
// @DateTime 2022/7/26 17:13

package server

import "github.com/gofiber/fiber/v2"

const SERVER_ADM_MGMT_DOCS_PRE = "/adm/mgmt/docs"

func DocGroups(app *fiber.App) *fiber.App {

	docGroup := app.Group(SERVER_ADM_MGMT_DOCS_PRE)
	{
		// Create a new and empty document.
		docGroup.Post("/new")
		// Fetch a document content.
		docGroup.Get("/:docId")
		docGroup.Delete("/:docId")
		// As temporarily stash edit content.
		docGroup.Put("/:docId")
		// Really save into database. Then we should remove before stash logs of this doc.
		docGroup.Post("/:docId/commit")
		docGroup.Delete("/:docId/stash")
		/*
			TODO(Ben) Metadata including doc created timestamp, update timestamp, font size,
			doc rendering template, catalog enabled, access enabled, contents page enabled.
		*/
		docGroup.Get("/:docId/metadata")
		// TODO(Ben) Summary including doc title, abstract, tags, thumbnail and other metadata.
		docGroup.Get("/:docId/summary")
		/*
			TODO(Ben) catalogs will list sub-doc under current doc. Current doc will play as
			a contents page also. It means current doc has another page to show catalog table.
		*/
		docGroup.Get("/:docId/catalogs")
	}

	// Tags management
	{
		// Add new tags or attach old tags to this doc resource.
		docGroup.Post("/:docId/tags")
		// Remove tags' associations from this doc resource.
		// TODO(Ben) Pass tags by query parameter, and separated by comma.
		docGroup.Delete("/:docId/tags")
	}

	return app
}

const SERVER_ADM_MGMT_DOCS_TMPL_PRE = "/adm/mgmt/docs/tmpl"

// DocTmplGroups
// Doc template is global enable or belong to an exists repository.
// The application should have a root default and empty template.
func DocTmplGroups(app *fiber.App) *fiber.App {

	tmplGroup := app.Group(SERVER_ADM_MGMT_DOCS_TMPL_PRE)
	{
		tmplGroup.Post("/new")
		tmplGroup.Get("/:tmplId")
		// Mark/Unmark a document template is global or belong to an exists repository.
		tmplGroup.Put("/:tmplId/mark")
		// Template could be deleted really.
		tmplGroup.Delete("/:tmplId")
	}

	return app
}
