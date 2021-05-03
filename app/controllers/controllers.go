package controllers

import (
	"crudapi/models"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

// Handler_user
// @Summary Show a users
// @Description Display of all user data.
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Users
// @Router /users [get]
func Handler_user(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.User_data_all(db))
	}
}

// Handler_show
// @Summary Show a personal user
// @Description Display the specified person.
// @Param name path string true "Name"
// @NAME string
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Users
// @Router /user/{name} [get]
func Handler_show(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.QueryParam("name")
		return c.JSON(http.StatusOK, models.User_data_name(db, name))
	}
}

// Handler_show
// @Summary Add new personal data.
// @Description Display the specified person.
// @Accept  json
// @Produce  json
// @Success 200
// @Router /create [post]
func Handler_create(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := &models.Users{}
		if err := c.Bind(&u); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, models.User_data_create(db, u))
	}
}

// Handler_update
// @Summary Update personal data.
// @Description Display the specified person.
// @Accept  json
// @Produce  json
// @Success 200
// @Router /update [put]
func Handler_update(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := &models.Users{}
		if err := c.Bind(&u); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, models.User_data_update(db, u))
	}
}

// Handler_delete
// @Summary Delete the personal data.
// @Description Display the specified person.
// @Accept  json
// @Produce  json
// @Success 204
// @Router /delete [delete]
func Handler_delete(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := &models.Users{}
		if err := c.Bind(&u); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, models.User_data_delete(db, u))
	}
}
