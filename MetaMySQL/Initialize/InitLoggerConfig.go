package Initialize

import (
	"MetaMySQL/global"
	"errors"
	"go.uber.org/zap"
)

func InitLoggerConfig() error {
	var logger *zap.Logger
	var err error
	if global.GlobalReleaseEnv == global.ReleaseDevVersion {
		logger, err = zap.NewDevelopment()
	} else if global.GlobalReleaseEnv == global.ReleaseProdVersion {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		return errors.New(global.ErrInitZapLoggerFailed.String())
	}
	zap.ReplaceGlobals(logger)
	return nil
}
