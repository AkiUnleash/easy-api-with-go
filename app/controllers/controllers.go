package controllers

import (
	"crudapi/models"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// ------- Controllers
func Handler_user(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.User_data_all(db))
	}
}

func Handler_show(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.QueryParam("name")
		return c.JSON(http.StatusOK, models.User_data_name(db, name))
	}
}

func Handler_create(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := &models.Users{}
		if err := c.Bind(&u); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, models.User_data_create(db, u))
	}
}

func Handler_update(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := &models.Users{}
		if err := c.Bind(&u); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, models.User_data_update(db, u))
	}
}

func Handler_delete(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := &models.Users{}
		if err := c.Bind(&u); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, models.User_data_delete(db, u))
	}
}
