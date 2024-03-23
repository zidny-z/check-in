package models

type Room struct {
	RoomID        int     `json:"room_id" gorm:"primary_key"`
	Type          string  `json:"room_type" gorm:"type:varchar(100)"`
	Capacity      int     `json:"room_capacity" gorm:"type:int"`
	Price         int     `json:"room_price" gorm:"type:int"`
	Photo         string  `json:"photo" gorm:"type:varchar(255)"`
	Hotel         []Hotel `json:"hotel" gorm:"foreignkey:HotelID"`
	Facility      string  `json:"facility" gorm:"type:varchar(255)"`
	RoomCount     int     `json:"room_count" gorm:"type:int"`
	RoomAvailable int     `json:"room_available" gorm:"type:int"`
}