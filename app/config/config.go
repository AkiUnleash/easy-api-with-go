package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gopkg.in/ini.v1"
)

// SQLConnect DB接続
func SqlConnect() (database *gorm.DB, err error) {

	cfg, err := ini.Load("config.ini")
	if err != nil {
		panic(err.Error())
	}

	DBMS := cfg.Section("db").Key("host").String()
	USER := cfg.Section("db").Key("user_name").String()
	PASS := cfg.Section("db").Key("password").String()
	PROTOCOL := cfg.Section("db").Key("protocol").String()
	DBNAME := cfg.Section("db").Key("db_name").String()

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	return gorm.Open(DBMS, CONNECT)
}
