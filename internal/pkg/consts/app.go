// @Author Ben.Zheng
// @DateTime 2022/8/8 22:26

package consts

const (
	APP_ROOT_WORKING_DIR = "APP_ROOT_WORKING_DIR"
	EMPTY_DIR            = ""
)

const (
	DB_CREATION_IF_NOT_PRESENT = "IfNotPresent"
	DB_CREATION_ALWAYS         = "Always"
	DB_CREATION_NEVER          = "Never"
)

type DBInitStatus uint8

const (
	NEVER_CHANGED DBInitStatus = iota
	RECREATED
	RECREATE_WITH_ERR
	ONLY_REMOVED
)

const (
	APP_RUNTIME_ENV_DEV  = "dev"
	APP_RUNTIME_ENV_PROD = "prod"
	APP_DEFAULT_DB_PATH  = "./db/idk.db"
)
