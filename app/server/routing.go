package server

import (
	"crudapi/controllers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Routing(db *gorm.DB) {
	e := echo.New()

	e.GET("/users", controllers.Handler_user(db))
	e.GET("/user", controllers.Handler_show(db))
	e.POST("/create", controllers.Handler_create(db))
	e.PUT("/update", controllers.Handler_update(db))
	e.DELETE("/delete", controllers.Handler_delete(db))

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8081"))
}
