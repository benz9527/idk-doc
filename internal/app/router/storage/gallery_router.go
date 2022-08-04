// @Author Ben.Zheng
// @DateTime 2022/7/21 10:51 AM

package storage

import "github.com/gofiber/fiber/v2"

const STORAGE_GALLERY_PRE = "/storage/gallery"

func GalleryGroups(app *fiber.App) *fiber.App {

	galleryGroup := app.Group(STORAGE_GALLERY_PRE)
	{
		// In order to support non-static image like
		// GIF. We should provide the type of image.
		galleryGroup.Get("/:imageId.:imageType")
	}

	// MinIO
	galleryAdmGroup := app.Group("/adm" + STORAGE_GALLERY_PRE)
	{
		/*
			TODO(Ben) Uploaded image should be compressed and marked with watermark.
			An image should be associated to a doc instance.
		*/
		galleryAdmGroup.Post("")
		/*
			TODO(Ben) Only the metadata of image could be updated.
		*/
		galleryAdmGroup.Put("/:imageId")
		/*
			TODO(Ben) Logical deletion and do further compressing.
		*/
		galleryAdmGroup.Delete("/:imageId")
	}

	return app
}
