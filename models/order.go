package models

import "time"

type Order struct {
	OrderID      int    `json:"order_id" gorm:"primary_key"`
	CustomerName string `json:"customer_name" gorm:"type:varchar(100)"`
	OrderedAt    time.Time `json:"ordered_at" gorm:"type:timestamp"`
	Hotel 	 Hotel  `json:"hotel" gorm:"foreignkey:HotelID"`        
}