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
func HandlerUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.UserDataAll(db))
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
func HandlerShow(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.QueryParam("name")
		return c.JSON(http.StatusOK, models.UserDataName(db, name))
	}
}

// Handler_show
// @Summary Add new personal data.
// @Description Display the specified person.
// @Accept  json
// @Produce  json
// @Success 200
// @Router /create [post]
func HandlerCreate(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := &models.Users{}
		if err := c.Bind(&u); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, models.UserDataCreate(db, u))
	}
}

// Handler_update
// @Summary Update personal data.
// @Description Display the specified person.
// @Accept  json
// @Produce  json
// @Success 200
// @Router /update [put]
func HandlerUpdate(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := &models.Users{}
		if err := c.Bind(&u); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, models.UserDataUpdate(db, u))
	}
}

// Handler_delete
// @Summary Delete the personal data.
// @Description Display the specified person.
// @Accept  json
// @Produce  json
// @Success 204
// @Router /delete [delete]
func HandlerDelete(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := &models.Users{}
		if err := c.Bind(&u); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, models.UserDataDelete(db, u))
	}
}
