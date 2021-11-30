package global

import (
	"MetaRedis/config"
	"github.com/go-redis/redis"
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
	GlobalClient *redis.Client
)