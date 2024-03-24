package models

type Room struct {
	RoomID        uint   `gorm:"column:room_id;primaryKey;autoIncrement" json:"room_id"`
	Type          string `gorm:"column:type;type:varchar(255)" json:"type"`
	Capacity      int    `gorm:"column:capacity;type:int" json:"capacity"`
	Price         int    `gorm:"column:price;type:int" json:"price"`
	Photo         string `gorm:"column:photo;type:varchar(255)" json:"photo"`
	Hotel         Hotel  `gorm:"foreignKey:HotelID"`
	HotelID       uint   `json:"hotel_id"`
	Facility      string `gorm:"column:facility;type:text" json:"facility"`
	RoomCount     int    `gorm:"column:room_count;type:int" json:"room_count"`
	RoomAvailable int    `gorm:"column:room_available;type:int" json:"room_available"`
}