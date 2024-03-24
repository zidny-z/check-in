package models

import (
	"time"
)

type Payment struct {
	PaymentId     uint `JSON:"payment_id" gorm:"primarykey"`
	User          User `gorm:"ForeignKey:UserId"`
	UserId        uint
	Totalamount   uint   `jSON:"total_amount" gorm:"not null"`
	Status        string `jSON:"Status" gorm:"not null"`
	Date          time.Time
	NoKTP	    string `jSON:"NoKTP" gorm:"null"`
}

type Oder_item struct {
	OrderId     uint    `JSON:"OrderId" gorm:"primarykey"`
	User        User    `gorm:"ForeignKey:UserIdNo"`
	UserIdNo    uint    `json:"useridno"  gorm:"not null" `
	TotalAmount uint    `json:"TotalAmount"  gorm:"not null" `
	Room		Room    `gorm:"ForeignKey:RoomId"`
	RoomId	   uint    `json:"RoomId"  gorm:"not null" `
	People	   uint    `json:"People"  gorm:"not null" `
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
