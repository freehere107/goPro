package models

import "chat/utiles"

type Room struct {
	CommonModel
	RoomName string `json:"room_name"`
}

func init() {
	db := utiles.DB
	if (db.HasTable(&Room{}) == false) {
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Room{})
	}
}

func (room *Room) CreateRoom() {

}
