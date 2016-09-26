package models

import (
	"chat/utiles"
	"errors"
)

type User struct {
	CommonModel
	Name   string `json:"name" gorm:"size:255"`
	Avatar string `json:"avatar" gorm:"size:500"`
}

func init() {
	db := utiles.DB
	if (db.HasTable(&User{}) == false) {
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
	}
}

func GetUserById(userID string) (User, error) {
	if len(userID) == 0 {
		return User{}, errors.New("-1")
	}
	db := utiles.DB
	user := User{}
	db.First(&user, userID)
	return user, nil
}
