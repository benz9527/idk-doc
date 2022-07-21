// @Author Ben.Zheng
// @DateTime 2022/7/21 9:53 AM

package client

import "github.com/gofiber/fiber/v2"

const CLIENT_ISSUES_PRE = "/issues"

func IssueGroups(app *fiber.App) *fiber.App {

	issueGroup := app.Group(CLIENT_ISSUES_PRE)
	{
		issueGroup.Get("/:issueId")
		issueGroup.Get(":/:issueId/metadata")
		issueGroup.Get("/:issueId/related")
	}
	return app
}
