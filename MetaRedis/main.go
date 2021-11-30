package main

import (
	"MetaRedis/Initialize"
	"MetaRedis/global"
)

func main()  {
	Initialize.Init()
	statusCmd := global.GlobalClient.Ping()
	if statusCmd.Err() != nil {
		panic(statusCmd.Err())
	}
	defer global.GlobalClient.Close()

}
