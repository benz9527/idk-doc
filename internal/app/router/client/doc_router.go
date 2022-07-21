// @Author Ben.Zheng
// @DateTime 2022/7/20 7:09 PM

package client

import "github.com/gofiber/fiber/v2"

const CLIENT_DOCS_PRE = "/docs"

func DocGroups(app *fiber.App) *fiber.App {

	docGroup := app.Group(CLIENT_DOCS_PRE)
	{
		docGroup.Get("/:docId")
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
	return app
}
