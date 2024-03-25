package models

type Room struct {
	RoomId     uint   `JSON:"RoomId" gorm:"primarykey;unique"`
	RoomName   string `JSON:"roomname" gorm:"not null"`
	Facilities string `JSON:"facilities" gorm:"not null"`
	Stock      uint   `JSON:"stock" gorm:"not null"`
	Avaliable  uint   `JSON:"avaliable" gorm:"not null"`
	Price      uint   `JSON:"price" gorm:"not null"`
	Hotel      Hotel  `gorm:"ForeignKey:HotelId"`
	HotelId    uint   `JSON:"HotelId"`
}

type Hotel struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	HotelName string `json:"hotel_name" gorm:"not null"`
	Location  string `json:"location" gorm:"not null"`
	Phone     string `json:"phone" gorm:"not null"`
}
