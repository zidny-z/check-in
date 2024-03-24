package models

type Payment struct {
	PaymentID uint   `json:"payment_id" gorm:"primary_key"`
	Order     Order  `gorm:"foreignkey:OrderID"`
	OrderID   uint   `json:"order_id"`
	Amount    int    `json:"amount" gorm:"type:int"`
	PaymentAt string `json:"payment_at" gorm:"type:timestamp"`
	File      string `json:"file" gorm:"type:varchar(255)"`
	Status    string `json:"status" gorm:"type:varchar(10)"`
	KTP       string `json:"ktp" gorm:"type:varchar(255)"`
}