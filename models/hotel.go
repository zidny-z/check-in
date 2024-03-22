package models

type Hotel struct {
	HotelID   int    `json:"hotel_id" gorm:"primary_key"`
	Name      string `json:"name" gorm:"type:varchar(100)"`
	Location  string `json:"location" gorm:"type:varchar(100)"`
	RoomCount int    `json:"room_count" gorm:"type:int"`
}