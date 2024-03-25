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

		

		// //User management routes
		// admin.GET("/user/viewuser", middlereware.AdminAuth, controls.ViewAllUser)
		// admin.GET("/user/searchuser", middlereware.AdminAuth, controls.AdminSearchUser)
		// admin.PUT("/user/edituserprofile/:id", middlereware.AdminAuth, controls.EditUserProfileByadmin)
		// admin.PUT("/user/blockusers", middlereware.AdminAuth, controls.AdminBlockUser)
		// admin.GET("/user/getuserprofile", middlereware.AdminAuth, controls.GetUserProfile)

		// //product management
		// admin.POST("/addcatogeries", middlereware.UserAuth, controls.AddCatogeries)
		// admin.POST("/product/addimage", middlereware.UserAuth, controls.AddImages)

		// //coupon routes
		// admin.POST("/coupon/add", middlereware.AdminAuth, controls.AddCoupon)
		// admin.POST("/coupon/checkcoupon", middlereware.AdminAuth, controls.CheckCoupon)

		// //Salse Report
		// admin.GET("/order/salesreport", middlereware.AdminAuth, controls.SalesReport)
		// admin.GET("/order/salesreport/download/excel", middlereware.AdminAuth, controls.DownloadExel)
		// admin.GET("/order/salesreport/download/pdf", middlereware.AdminAuth, controls.Downloadpdf)
	}

}
