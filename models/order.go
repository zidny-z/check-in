package models

import "time"

type Order struct {
	OrderID      uint    `json:"order_id" gorm:"primary_key;unique;autoIncrement"`
	User 	  User   `gorm:"ForeignKey:UserID"`
	UserID        uint    `json:"user_id"`
	Room 	  Room   `gorm:"ForeignKey:RoomID"`
	RoomID        int    `json:"room_id"`
	OrderedAt    time.Time `json:"ordered_at" gorm:"type:timestamp"`
	CheckinDate  time.Time `json:"checkin_date" gorm:"type:date"`
	CheckoutDate time.Time `json:"checkout_date" gorm:"type:date"`
	PeopleCount  int    `json:"people_count" gorm:"type:int"`

}