package main

import (
	"check-in/config"
	"check-in/routes"

	"gorm.io/gorm"

	"check-in/initializer"

	"github.com/gin-gonic/gin"
)

var DB *gorm.DB

func init() {
	initializer.LoadEnv()
	var err error
	config.DB, err = config.DBconnect()
	if err != nil {
		panic(err)
	}

	R.LoadHTMLGlob("templates/*.html")
}

var R = gin.Default()

func main() {

	// routes.AdminRouts(R)
	routes.UserRouts(R)

	R.Run()
}
