package main

import (
	"crudapi/controllers"
	"crudapi/server"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func main() {

	db, err := server.SqlConnect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	routing(db)

}

// Server
// ------- Routing
func routing(db *gorm.DB) {
	e := echo.New()

	e.GET("/users", controllers.Handler_user(db))
	e.GET("/user", controllers.Handler_show(db))
	e.POST("/create", controllers.Handler_create(db))
	e.PUT("/update", controllers.Handler_update(db))
	e.DELETE("/delete", controllers.Handler_delete(db))

	e.Logger.Fatal(e.Start(":8081"))
}
