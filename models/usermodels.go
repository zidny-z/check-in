package models

import "time"

type User struct {
	ID          uint   `json:"id" gorm:"primaryKey;unique"  `
	FirstName   string `json:"first_name"  gorm:"not null" validate:"required,min=2,max=50"`
	LastName    string `json:"last_name"    gorm:"not null"    validate:"required,min=1,max=50"`
	Email       string `json:"email"   gorm:"not null;unique"  validate:"email,required"`
	Password    string `json:"password" gorm:"not null"  validate:"required"`
	PhoneNumber int    `json:"phone"   gorm:"not null;unique" validate:"required"`
	IsAdmin     bool   `JSON:"isadmin" gorm:"default:false"`
	Otp         string `JSON:"otp"`
	Isblocked   bool   `JSON:"isblocked" gorm:"default:false"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

