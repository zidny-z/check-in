package router

import (
	"fmt"
	"check-in/config"
	"check-in/controller"

	"github.com/gin-gonic/gin"
)

func StartServer(c controller.Controller) error {
	port := config.GetConfigPort()
	host := config.GetConfigDB().DB_HOST
	serverInfo := fmt.Sprintf("%s:%s", host, port)

	r := gin.Default()

	r.POST("/orders", c.CreateOrder)
	r.GET("/orders", c.GetOrders)
	r.PUT("/orders/:orderId", c.UpdateOrder)
	r.DELETE("/orders/:orderId", c.DeleteOrder)

	return r.Run(serverInfo)
}