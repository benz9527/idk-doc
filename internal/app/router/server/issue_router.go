// @Author Ben.Zheng
// @DateTime 2022/7/26 17:13

package server

import "github.com/gofiber/fiber/v2"

const SERVER_ADM_MGMT_ISSUES_PRE = "/adm/mgmt/issues"

func IssueGroups(app *fiber.App) *fiber.App {

	issueGroup := app.Group(SERVER_ADM_MGMT_ISSUES_PRE)
	{
		issueGroup.Post("/new")
		// Stash the changes.
		issueGroup.Put("/:issueId")
		// Save as permanently.
		issueGroup.Post("/:issueId/commit")
		issueGroup.Get("/:issueId")
	}

	// Related issues management
	{
		// Query all related issues.
		issueGroup.Get("/:issueId/related")
		issueGroup.Post("/:issueId/related")
		// TODO(Ben) Pass tags by query parameter, and separated by comma.
		issueGroup.Delete("/:issueId/related")
	}

	// Tags management
	{
		// Add new tags or attach old tags to this issue resource.
		issueGroup.Post("/:issueId/tags")
		// Remove tags' associations from this issue resource.
		// TODO(Ben) Pass tags by query parameter, and separated by comma.
		issueGroup.Delete("/:issueId/tags")
	}

	return app
}
