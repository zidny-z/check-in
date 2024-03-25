package models

import (
	"time"
)

type Payment struct {
	PaymentId     uint `json:"payment_id" gorm:"primarykey"`
	User          User `gorm:"ForeignKey:UserId"`
	UserId        uint
	Totalamount   uint   `json:"total_amount" gorm:"not null"`
	Status        string `json:"Status" gorm:"not null"`
	Date          time.Time
	NoKTP	    string `json:"NoKTP" gorm:"null"`
	Order	Order  `gorm:"ForeignKey:OrderId`
	OrderId		uint `json:"OrderID`
}

type Order struct {
	OrderId     uint    `json:"OrderId" gorm:"primarykey"`
	User        User    `gorm:"ForeignKey:UserIdNo"`
	UserIdNo    uint    `json:"useridno"  gorm:"not null" `
	TotalAmount uint    `json:"TotalAmount"  gorm:"not null" `
	Room		Room    `gorm:"ForeignKey:RoomId"`
	RoomId	   uint    `json:"RoomId"  gorm:"not null" `
	People	   uint    `json:"People"  gorm:"not null" `
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
