package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Users ユーザー情報のテーブル情報
type Users struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// ------- Models
func UserDataName(db *gorm.DB, name string) []Users {
	result := []Users{}
	db.Where("name = ?", name).Find(&result)
	return result
}

func UserDataAll(db *gorm.DB) []Users {
	result := []Users{}
	db.Find(&result)
	return result
}

func UserDataCreate(db *gorm.DB, u *Users) string {
	error := db.Create(&u).Error
	if error != nil {
		return "NG"
	} else {
		return "OK"
	}
}

func UserDataUpdate(db *gorm.DB, u *Users) string {
	error := db.Model(Users{}).Where("id = ?", u.ID).Update(&u).Error

	if error != nil {
		return "NG"
	} else {
		return "OK"
	}
}

func UserDataDelete(db *gorm.DB, u *Users) string {
	error := db.Model(Users{}).Where("id = ?", u.ID).Delete(&u).Error

	if error != nil {
		return "NG"
	} else {
		return "OK"
	}
}
