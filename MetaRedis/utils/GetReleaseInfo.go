package utils

import "github.com/spf13/viper"

func GetReleaseInfo(param string) string {
	viper.AutomaticEnv()
	paramValue := viper.GetString(param)
	return paramValue
}
