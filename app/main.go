package main

import (
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

	db, err := sqlConnect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, user_data_all(db))
	})

	e.GET("/shows", func(c echo.Context) error {
		name := c.QueryParam("name")
		return c.JSON(http.StatusOK, user_data_name(db, name))
	})

	e.POST("/create", func(c echo.Context) error {
		u := &Users{}
		if err := c.Bind(&u); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, user_data_create(db, u))
	})

	e.PUT("/update", func(c echo.Context) error {
		u := &Users{}
		if err := c.Bind(&u); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, user_data_update(db, u))
	})

	e.DELETE("/delete", func(c echo.Context) error {
		u := &Users{}
		if err := c.Bind(&u); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, user_data_delete(db, u))
	})

	e.Logger.Fatal(e.Start(":8081"))

}

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

// データ挿入
func user_data_create(db *gorm.DB, u *Users) string {
	error := db.Create(&u).Error
	if error != nil {
		return "NG"
	} else {
		return "OK"
	}
}

// アップデート
func user_data_update(db *gorm.DB, u *Users) string {
	error := db.Model(Users{}).Where("id = ?", u.ID).Update(&u).Error

	if error != nil {
		return "NG"
	} else {
		return "OK"
	}
}

// アップデート
func user_data_delete(db *gorm.DB, u *Users) string {
	error := db.Model(Users{}).Where("id = ?", u.ID).Delete(&u).Error

	if error != nil {
		return "NG"
	} else {
		return "OK"
	}
}

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
