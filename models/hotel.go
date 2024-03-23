package models

type Hotel struct {
	HotelID       int    `json:"hotel_id" gorm:"primary_key"`
	Name          string `json:"name" gorm:"type:varchar(100)"`
	Location      string `json:"location" gorm:"type:varchar(255)"`
	RoomCount     int    `json:"room_count" gorm:"type:int"`
	RoomAvailable int    `json:"room_available" gorm:"type:int"`
	Star          int    `json:"star" gorm:"type:int"`
	IsSyariah     bool   `json:"is_syariah" gorm:"type:boolean"`
	Photo         string `json:"photo" gorm:"type:varchar(255)"`
	Facility      string `json:"facility" gorm:"type:varchar(255)"`
}