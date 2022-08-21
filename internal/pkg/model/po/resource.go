// @Author Ben.Zheng
// @DateTime 2022/8/18 13:37

package po

type ResourceGroup struct {
	NanoIdFullMode
	Name string `gorm:"column:name;type:nvarchar(32);uniqueIndex:idx_img_grp_name;<-;"`
}

func (i ResourceGroup) TableName() string {
	return "idk_res_groups"
}

// ResourceCore
// @field OssURL Display the resource local save path or OSS
// save URL.
// @field Extension For image, png, svg, jpeg and gif will be
// included. For attachment, doc, execl, ppt, pdf, zip, 7z,
// tar.gz and so on are acceptable.
// @field Checksum Used to check idk-doc resource integrity.
// Potential algo including SHA-256(64 characters), SHA-512
// (128 characters), MD5 and so on.
type ResourceCore struct {
	NanoIdFullMode
	OssURL    string `gorm:"column:url;type:varchar(256);<-;"`
	Extension string `gorm:"column:ext;type:varchar(8);<-;"`
	GroupId   string `gorm:"column:res_group_id;type:varchar(21);index;<-;"`
	Size      int64  `gorm:"column:size;type:bigint;<-;"`
	Checksum  string `gorm:"column:checksum;type:varchar(128);<-;"`
}

type ImageCore struct {
	ResourceCore
	Watermark string `gorm:"column:watermark;type:nvarchar(32);<-;"`
}

type Image[T ImageCore] struct {
	Core  T              `gorm:"embedded;"`
	Group *ResourceGroup `gorm:"foreignKey:GroupId;"`
}

func (i Image[T]) TableName() string {
	return "idk_res_images"
}

func (i Image[T]) GetCore() T {
	return i.Core
}

type AttachmentCore struct {
	ResourceCore
	Name string `gorm:"column:name;type:nvarchar(32);index;<-;"`
}

type Attachment[T AttachmentCore] struct {
	Core  T              `gorm:"embedded;"`
	Group *ResourceGroup `gorm:"foreignKey:GroupId;"`
}

func (a Attachment[T]) TableName() string {
	return "idk_res_attachments"
}

func (a Attachment[T]) GetCore() T {
	return a.Core
}
