package Initialize

import "go.uber.org/zap"

func Init() {
	err := InitBaseInfo()
	zap.S().Info("InitBaseInfo: ", err.Error())
	err = InitLoggerConfig()
	//zap.S().Info("InitLoggerConfig: ", err.Error())
	err = InitRedisServerConfig()
	//zap.S().Info("InitRedisServerConfig: ", err.Error())
	err = InitRedisServerConn()
	//zap.S().Info("InitRedisServerConfig: ", err.Error())
}
