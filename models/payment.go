package models

type Payment struct {
	PaymentID int     `json:"payment_id" gorm:"primary_key"`
	Order     []Order `json:"order" gorm:"foreignkey:OrderID"`
	Amount    int     `json:"amount" gorm:"type:int"`
	PaymentAt string  `json:"payment_at" gorm:"type:timestamp"`
	File      string  `json:"file" gorm:"type:varchar(255)"`
	Status    string  `json:"status" gorm:"type:varchar(10)"`
	KTP       string  `json:"ktp" gorm:"type:varchar(255)"`
}