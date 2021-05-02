package main

import (
	"crudapi/config"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// Users ユーザー情報のテーブル情報
type Users struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func main() {

	db, err := config.SqlConnect()
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

	e.GET("/user", handler_user(db))
	e.GET("/shows", handler_show(db))
	e.POST("/create", handler_create(db))
	e.PUT("/update", handler_update(db))
	e.DELETE("/delete", handler_delete(db))

	e.Logger.Fatal(e.Start(":8081"))
}

// ------- Controllers
func handler_user(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, user_data_all(db))
	}
}

func handler_show(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.QueryParam("name")
		return c.JSON(http.StatusOK, user_data_name(db, name))
	}
}

func handler_create(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := &Users{}
		if err := c.Bind(&u); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, user_data_create(db, u))
	}
}

func handler_update(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := &Users{}
		if err := c.Bind(&u); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, user_data_update(db, u))
	}
}

func handler_delete(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := &Users{}
		if err := c.Bind(&u); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, user_data_delete(db, u))
	}
}

// ------- Models
func user_data_name(db *gorm.DB, name string) []Users {
	result := []Users{}
	db.Where("name = ?", name).Find(&result)
	return result
}

func user_data_all(db *gorm.DB) []Users {
	result := []Users{}
	db.Find(&result)
	return result
}

func user_data_create(db *gorm.DB, u *Users) string {
	error := db.Create(&u).Error
	if error != nil {
		return "NG"
	} else {
		return "OK"
	}
}

func user_data_update(db *gorm.DB, u *Users) string {
	error := db.Model(Users{}).Where("id = ?", u.ID).Update(&u).Error

	if error != nil {
		return "NG"
	} else {
		return "OK"
	}
}

func user_data_delete(db *gorm.DB, u *Users) string {
	error := db.Model(Users{}).Where("id = ?", u.ID).Delete(&u).Error

	if error != nil {
		return "NG"
	} else {
		return "OK"
	}
}
