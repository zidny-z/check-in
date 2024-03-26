package config

import (
	"fmt"
	"os"

	"checkin/models"

	// mysql
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBconnect() (*gorm.DB, error) {
	// db_user := os.Getenv("POSTGRES_USER")
	// db_pass := os.Getenv("POSTGRES_PASSWORD")
	// db_name := os.Getenv("POSTGRES_DB")
	// db_host := os.Getenv("POSTGRES_HOST")
	// db_port := os.Getenv("POSTGRES_PORT")
	// db_ssl := os.Getenv("POSTGRES_SSL_MODE")
	// db_timezone := os.Getenv("POSTGRES_TIMEZONE")
	// dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", db_host, db_user, db_pass, db_name, db_port, db_ssl, db_timezone)
	// DB, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	// mysql
	db_user := os.Getenv("MYSQL_USER")
	db_pass := os.Getenv("MYSQL_PASSWORD")
	db_name := os.Getenv("MYSQL_DATABASE")
	db_host := os.Getenv("MYSQL_HOST")
	db_port := os.Getenv("MYSQL_PORT")
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_user, db_pass, db_host, db_port, db_name)
	DB, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

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
