// @Author Ben.Zheng
// @DateTime 2022/7/26 15:50

package server

import "github.com/gofiber/fiber/v2"

const SERVER_ADM_MGMT_CODES_PRE = "/adm/mgmt/codes"

func CodeGroups(app *fiber.App) *fiber.App {

	codeGroup := app.Group(SERVER_ADM_MGMT_CODES_PRE)
	{
		// Create a new and empty code document resource.
		// TODO(Ben) It should allocate some metadata, including repo id, catalog id.
		// TODO(Ben) A default code tab should be created when a code document resource is creating.
		codeGroup.Post("/new")
		// Fetch all metadata of a code document, including timestamp info, tags, tabs and so on.
		codeGroup.Get("/:codeId")
		// Logical deletion, just let client could not access to this and present it.
		codeGroup.Delete("/:codeId")
	}

	// Tags management
	{
		// Add new empty tags or attach old tags to this code resource.
		codeGroup.Post("/:codeId/tags")
		// Remove tags' associations from this code resource.
		// TODO(Ben) Pass tags by query parameter, and separated by comma.
		codeGroup.Delete("/:codeId/tags")
	}

	// Tabs management
	{
		// Create a new code tab.
		// TODO(Ben) It have to select a code rendering template.
		codeGroup.Post("/:codeId/tabs/new")
		// Stash content to code tab.
		codeGroup.Put("/:codeId/tabs/:tabId")
		// Save the content to code permanently.
		codeGroup.Post("/:codeId/tabs/:tabId/commit")
		// Fetch the code content in the tab.
		codeGroup.Get("/:codeId/tabs/:tabId")
		// Logical deletion, marking this as not accessible.
		codeGroup.Delete("/:codeId/tabs/:tabId")
		codeGroup.Delete("/:codeId/tabs/:tabId/stash")
		// Run the code.
		// TODO(Ben) If the code could be marked as runnable, we can run this by virtual environment.
		// This is a dangerous feature.
		codeGroup.Get("/:codeId/tabs/:tabId/run")
	}

	return app
}

const SERVER_ADM_MGMT_CODES_LANG_PRE = "/adm/mgmt/codes/lang"

// CodeLangAndRenderingTmplGroups
// A lang will have multiple code rendering templates.
func CodeLangAndRenderingTmplGroups(app *fiber.App) *fiber.App {

	langGroup := app.Group(SERVER_ADM_MGMT_CODES_LANG_PRE)
	{
		langGroup.Post("/:lang/new")
		// Only update the lang name.
		langGroup.Patch("/:langId")
		// Only an empty lang could be deleted.
		// i.e. without any code doc and rendering template associated to.
		langGroup.Delete("/:langId")
	}

	// Template management
	// It should have a default or empty template for template not found situation.
	tmplGroup := app.Group(SERVER_ADM_MGMT_CODES_LANG_PRE + "/tmpl")
	{
		// Create a new and empty code rendering template to a lang.
		tmplGroup.Post("/:langId/new")
		tmplGroup.Put("/:tmplId")
		tmplGroup.Delete("/:tmplId")
	}

	return app
}
