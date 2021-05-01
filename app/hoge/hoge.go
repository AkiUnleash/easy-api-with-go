package hoge

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// SQLConnect DB接続
func sqlConnect() (database *gorm.DB, err error) {
	DBMS := "mysql"
	USER := "root"
	PASS := "golang"
	PROTOCOL := "tcp(mysql-container:3306)"
	DBNAME := "golang_db"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	return gorm.Open(DBMS, CONNECT)
}
