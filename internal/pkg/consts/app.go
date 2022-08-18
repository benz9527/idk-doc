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

const (
	APP_LOG_ENC_JSON  = "JSON"
	APP_LOG_ENC_PLAIN = "PlainText"
	APP_LOG_LVL_DEBUG = "DEBUG"
	APP_LOG_LVL_INFO  = "INFO" // Product env default level if not present in configuration.
	APP_LOG_LVL_WARN  = "WARN"
	APP_LOG_LVL_ERR   = "ERROR"
)

type FileType uint8

const (
	FILE_TYPE_MD FileType = iota
	FILE_TYPE_CODE
	FILE_TYPE_ISSUE
	// TODO(Ben) DOC, EXCEL
)
