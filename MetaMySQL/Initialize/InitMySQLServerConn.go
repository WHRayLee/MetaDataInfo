package Initialize

import (
	"MetaMySQL/global"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMySQLServerConn() error {
	var err error
	//
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		global.GlobalServerConfigInfo.MySQLConfigInfo.User,
		global.GlobalServerConfigInfo.MySQLConfigInfo.Password,
		global.GlobalServerConfigInfo.MySQLConfigInfo.Host,
		global.GlobalServerConfigInfo.MySQLConfigInfo.Port,
		global.GlobalServerConfigInfo.MySQLConfigInfo.Db,
		global.GlobalServerConfigInfo.MySQLConfigInfo.Charset)
	global.GlobalClient, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		errors.New(global.ErrDBConnectFailed.String())
	}
	return nil
}
