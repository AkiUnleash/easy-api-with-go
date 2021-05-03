package main

import (
	"crudapi/controllers"
	"crudapi/server"

	_ "crudapi/docs"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Easy API with Go
// @version 1.0.0
// @description This is program I developed as a lerning process before creating an API in Golang.

// @license.name MIT
// @license.url https://github.com/tcnksm/tool/blob/master/LICENCE

// @host http://localhost:8081
// @BasePath /

func main() {

	db, err := server.SqlConnect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	routing(db)

}

func routing(db *gorm.DB) {
	e := echo.New()

	e.GET("/users", controllers.Handler_user(db))
	e.GET("/user", controllers.Handler_show(db))
	e.POST("/create", controllers.Handler_create(db))
	e.PUT("/update", controllers.Handler_update(db))
	e.DELETE("/delete", controllers.Handler_delete(db))

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8081"))
}
