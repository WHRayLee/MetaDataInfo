package Initialize

import (
	"MetaRedis/global"
	"fmt"
	"github.com/go-redis/redis"
)

func InitRedisServerConn() error {
	dsn := fmt.Sprintf("%s:%d", global.GlobalServerConfigInfo.RedisConfigInfo.Host,
		global.GlobalServerConfigInfo.RedisConfigInfo.Port)
	options := redis.Options{
		Network:      "tcp",
		Addr:         dsn,
		DB:           0,
		//DialTimeout:  time.Duration(global.GlobalServerConfigInfo.RedisConfigInfo.DialTimeout * 1000),
		//ReadTimeout:  time.Duration(global.GlobalServerConfigInfo.RedisConfigInfo.ReadTimeout),
		//WriteTimeout: time.Duration(global.GlobalServerConfigInfo.RedisConfigInfo.WriteTimeout),
	}
	global.GlobalClient = redis.NewClient(&options)
	return nil
}
