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

		
		// //User management routes
		// admin.GET("/user/viewuser", middlereware.AdminAuth, controls.ViewAllUser)
		// admin.GET("/user/searchuser", middlereware.AdminAuth, controls.AdminSearchUser)
		// admin.PUT("/user/edituserprofile/:id", middlereware.AdminAuth, controls.EditUserProfileByadmin)
		// admin.PUT("/user/blockusers", middlereware.AdminAuth, controls.AdminBlockUser)
		// admin.GET("/user/getuserprofile", middlereware.AdminAuth, controls.GetUserProfile)

	}

}
