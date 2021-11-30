package Initialize

import (
	"MetaRedis/global"
	"MetaRedis/utils"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"os"
)

func InitBaseInfo() error {
	getReleaseInfo := utils.GetReleaseInfo(global.ReleaseParam)
	zap.S().Info("版本信息：", getReleaseInfo)
	if getReleaseInfo == global.ReleaseDevVersion {
		global.GlobalConfigFile = fmt.Sprintf("MetaRedis/config/config-%s.yaml", getReleaseInfo)
	} else if getReleaseInfo == global.ReleaseProdVersion {
		global.GlobalConfigFile = fmt.Sprintf("MetaRedis/config/config-%s.yaml", getReleaseInfo)
	} else {
		return errors.New(global.ErrInvalidReleaseParam.String())
	}
	_, err := os.Lstat(global.GlobalConfigFile)
	if err != nil {
		return errors.New(global.ErrConfigFileNotExists.String() + ":" + global.GlobalConfigFile)
	}
	global.GlobalReleaseEnv = getReleaseInfo
	return nil
}
