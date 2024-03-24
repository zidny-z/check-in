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

	}

}
