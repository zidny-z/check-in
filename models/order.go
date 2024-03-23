package models

import "time"

type Order struct {
	OrderID      int    `json:"order_id" gorm:"primary_key"`
	User 	  []User   `json:"user" gorm:"foreignkey:UserID"`
	Room 	  []Room   `json:"room" gorm:"foreignkey:RoomID"`
	OrderedAt    time.Time `json:"ordered_at" gorm:"type:timestamp"`
	CheckinDate  time.Time `json:"checkin_date" gorm:"type:date"`
	CheckoutDate time.Time `json:"checkout_date" gorm:"type:date"`
	PeopleCount  int    `json:"people_count" gorm:"type:int"`

}