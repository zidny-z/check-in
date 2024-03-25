package config

import (
	"fmt"
	"os"

	"checkin/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBconnect() (*gorm.DB, error) {
	db_user := os.Getenv("POSTGRES_USER")
	db_pass := os.Getenv("POSTGRES_PASSWORD")
	db_name := os.Getenv("POSTGRES_DB")
	db_host := os.Getenv("POSTGRES_HOST")
	db_port := os.Getenv("POSTGRES_PORT")
	db_ssl := os.Getenv("POSTGRES_SSL_MODE")
	db_timezone := os.Getenv("POSTGRES_TIMEZONE")
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", db_host, db_user, db_pass, db_name, db_port, db_ssl, db_timezone)
	DB, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Admin{})
	DB.AutoMigrate(&models.Room{})
	DB.AutoMigrate(&models.Hotel{})
	DB.AutoMigrate(&models.Payment{})
	DB.AutoMigrate(&models.Order{})

	return DB, nil

}
