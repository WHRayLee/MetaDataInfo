package Initialize

import (
	"MetaRedis/global"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func InitRedisServerConfig() error {
	zap.S().Info("初始化Redis服务器配置")
	//viper.SetConfigFile(global.GlobalConfigFile)
	viper.SetConfigFile("./config/config-dev.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		zap.S().Fatal("读取配置信息失败:", err.Error())
		return err
	}

	err = viper.Unmarshal(&global.GlobalServerConfigInfo)
	if err != nil {
		zap.S().Fatal("解析配置信息失败")
		return err
	}
	return nil
}
