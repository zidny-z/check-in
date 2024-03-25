package routes

import (
	"checkin/controls"
	"checkin/middlereware"

	"github.com/gin-gonic/gin"
)

func UserRouts(c *gin.Engine) {
	User := c.Group("/user")
	{

		//User rountes >>
		User.POST("/login", controls.UserLogin)
		User.POST("/signup", controls.UserSignUP)
		User.POST("/signup/otpvalidate", controls.OtpValidation)
		User.GET("/logout", middlereware.UserAuth, controls.UserSignout)

		//User profile routes
		User.GET("/viewprofile", middlereware.UserAuth, controls.ShowUserDetails)
		User.POST("/userchangepassword", middlereware.UserAuth, controls.UserChangePassword)
		User.PUT("/userchangepassword/updatepassword", middlereware.UserAuth, controls.Updatepassword)
		User.PUT("/editprofile", middlereware.UserAuth, controls.EditUserProfilebyUser)

		//User hotel explore
		User.GET("/hotels", middlereware.UserAuth, controls.ViewRooms)
		
		// order room
		User.POST("/orderroom", middlereware.UserAuth, controls.OrderRoom)
		User.POST("/orderpayment", middlereware.UserAuth, controls.OrderPayment)
		User.GET("/myorder", middlereware.UserAuth, controls.ViewOrdersUser)
		User.GET("/mypayment", middlereware.UserAuth, controls.ViewPaymentsUser)

		//Forgot Password >>
		User.PUT("/forgotpassword", middlereware.UserAuth, controls.GenerateOtpForForgotPassword)
		User.POST("/forgotpassword/changepassword", middlereware.UserAuth, controls.ChangePassword)

	}
}
