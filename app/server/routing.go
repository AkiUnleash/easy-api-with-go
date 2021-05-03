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

	e.GET("/users", controllers.HandlerUser(db))
	e.GET("/user", controllers.HandlerShow(db))
	e.POST("/create", controllers.HandlerCreate(db))
	e.PUT("/update", controllers.HandlerUpdate(db))
	e.DELETE("/delete", controllers.HandlerDelete(db))

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8081"))
}
