package global

import (
	"MetaMySQL/config"
	"gorm.io/gorm"
)

var (
	ReleaseParam       = "RELEASED_VERSION"
	ReleaseDevVersion  = "dev"
	ReleaseProdVersion = "prod"
)

var (
	GlobalReleaseEnv string
	GlobalConfigFile string
)

var (
	GlobalServerConfigInfo config.ServerInfo
	GlobalClient *gorm.DB
)