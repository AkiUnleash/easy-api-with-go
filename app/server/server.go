package server

import (
	"crudapi/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// SQLConnect DB接続
func SqlConnect() (database *gorm.DB, err error) {

	DBMS := config.Config.Dbms
	USER := config.Config.User
	PASS := config.Config.Pass
	PROTOCOL := config.Config.Protocol
	DBNAME := config.Config.Dbname

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	return gorm.Open(DBMS, CONNECT)
}
