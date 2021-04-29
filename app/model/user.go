package user

import (
	"database/sql"
	"net/http"

	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

// User db users
type User struct {
	ID       int
	Name     string
	Password string
}

var (
	db, err = sql.Open("mysql", "root:golang@tcp(localhost:3306)/golang_db")
	dbmap   = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
)

const (
	DriverName     = "mysql"
	DataSourceName = "root:golang@tcp(mysql-container:3306)/golang_db"
)

func selectUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var users = []User{}
		// select * from members where name = name;を実行し、membersに格納する
		dbmap.Select(&users, "SELECT * FROM user"+";")
		return c.JSON(http.StatusOK, users)
	}
}
