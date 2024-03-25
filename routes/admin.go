package routes

import (
	"checkin/controls"
	"checkin/middlereware"

	"github.com/gin-gonic/gin"
)

func AdminRouts(c *gin.Engine) {
	admin := c.Group("/admin")
	{
		//Admin rounts
		admin.POST("/login", controls.AdminLogin)
		admin.POST("/signup", controls.AdminSignup)
		admin.GET("/logout", middlereware.AdminAuth, controls.AdminSignout)
		admin.GET("/profile", middlereware.AdminAuth, controls.AdminProfile)
		admin.GET("/adminvalidate", middlereware.AdminAuth, controls.ValidateAdmin)

		//specification hotel management routes
		admin.PUT("/hotel/edit/:id", middlereware.AdminAuth, controls.EditHotel)
		admin.GET("/hotel", middlereware.AdminAuth, controls.ViewHotel)
		admin.POST("/hotel", middlereware.AdminAuth, controls.AddHotel)

		// room management routes
		admin.POST("/rooms", middlereware.AdminAuth, controls.AddRoom)
		admin.GET("/rooms", middlereware.AdminAuth, controls.ViewRooms)
		admin.PUT("/rooms/edit/:id", middlereware.AdminAuth, controls.EditRoom)

		// order n payment management
		admin.PUT("/validatepayment/:id", middlereware.AdminAuth, controls.ValidatePayment)
		admin.GET("/order", middlereware.AdminAuth, controls.ViewAllOrders)
		admin.GET("/payment", middlereware.AdminAuth, controls.ViewAllPayments)

		
		// //User management routes
		admin.GET("/user", middlereware.AdminAuth, controls.ViewAllUser)
		// admin.PUT("/user/blockuser/:id", middlereware.AdminAuth, controls.AdminBlockUser)

	}

}
