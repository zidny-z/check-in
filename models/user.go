package models

type User struct {
	UserID   uint   `json:"user_id" gorm:"primary_key"`
	Name     string `json:"name" gorm:"type:varchar(100)"`
	Email    string `json:"email" gorm:"type:varchar(100); unique"`
	Password string `json:"password" gorm:"type:varchar(255)"`
	Role     string `json:"role" gorm:"type:varchar(10)"`
}