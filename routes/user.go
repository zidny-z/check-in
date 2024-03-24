package routes

import (
	"check-in/controls"
	"check-in/middlereware"

	"github.com/gin-gonic/gin"
)

func UserRouts(c *gin.Engine) {
	User := c.Group("/user")
	{

		//User rountes >>
		User.POST("/login", controls.UserLogin)
		User.POST("/signup", controls.UserSignUP)
		User.POST("/signup/otpvalidate", controls.OtpValidation)

		//User profile routes
		User.GET("/viewprofile", middlereware.UserAuth, controls.ShowUserDetails)
		User.POST("/userchangepassword", middlereware.UserAuth, controls.UserChangePassword)
		User.PUT("/userchangepassword/updatepassword", middlereware.UserAuth, controls.Updatepassword)
		User.PUT("/editprofile", middlereware.UserAuth, controls.EditUserProfilebyUser)

		//Forgot Password >>
		User.PUT("/forgotpassword", middlereware.UserAuth, controls.GenerateOtpForForgotPassword)
		User.POST("/forgotpassword/changepassword", middlereware.UserAuth, controls.ChangePassword)

	}
}
