package main

import (
	"checkin/config"
	"checkin/routes"

	"gorm.io/gorm"

	"checkin/initializer"

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

	routes.AdminRouts(R)
	routes.UserRouts(R)

	R.Run()
}
