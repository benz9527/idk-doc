// @Author Ben.Zheng
// @DateTime 2022/7/21 10:51 AM

package storage

import "github.com/gofiber/fiber/v2"

const STORAGE_ATTACHEMTNS_PRE = "/storage/attachments"

func AttachmentGroups(app *fiber.App) *fiber.App {

	attachmentGroup := app.Group(STORAGE_ATTACHEMTNS_PRE)
	{
		attachmentGroup.Get("/:attachmentId")
	}

	// MongoDB
	attachmentAdmGroup := app.Get("/adm" + STORAGE_ATTACHEMTNS_PRE)
	{
		/*
			TODO(Ben) Attachment should be enabled compress processing.
			Attachment should be associated to a doc instance.
			And it should be set a size limitation.
		*/
		attachmentAdmGroup.Post("")
		/*
			TODO(Ben) Modification only limited to update attachment's metadata instead of binary.
		*/
		attachmentAdmGroup.Put("/:attachmentId")
		/*
			TODO(Ben) Logical deletion instead of physical one.
		*/
		attachmentAdmGroup.Delete("/:attachmentId")
	}

	return app
}
