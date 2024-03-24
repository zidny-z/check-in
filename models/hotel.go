package models

type Hotel struct {
	HotelID       uint   `gorm:"column:hotel_id;primaryKey;autoIncrement" json:"hotel_id"`
	Name          string `gorm:"column:name;type:varchar(255)" json:"name"`
	Location      string `gorm:"column:location;type:varchar(255)" json:"location"`
	RoomCount     int    `gorm:"column:room_count;type:int" json:"room_count"`
	RoomAvailable int    `gorm:"column:room_available;type:int" json:"room_available"`
	Star          int    `gorm:"column:star;type:int" json:"star"`
	IsSyariah     bool   `gorm:"column:is_syariah;type:boolean" json:"is_syariah"`
	Photo         string `gorm:"column:photo;type:varchar(255)" json:"photo"`
	Facility      string `gorm:"column:facility;type:text" json:"facility"`
}