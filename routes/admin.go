package routes

import (
	"check-in/controls"
	"check-in/middlereware"

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
	}

}
