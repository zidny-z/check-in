package router

import (
	"check-in/config"
	"check-in/controller"
	"fmt"

	"github.com/gin-gonic/gin"
)

func StartServer(c controller.Controller) error {
	port := config.GetConfigPort()
	host := config.GetConfigDB().DB_HOST
	serverInfo := fmt.Sprintf("%s%s", host, port)

	r := gin.Default()

	return r.Run(serverInfo)
}