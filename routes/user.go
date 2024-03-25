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

		// User.GET("/viewbrand", middlereware.UserAuth, controls.ViewBrand)
		// User.GET("/search", middlereware.UserAuth, controls.SearchProduct)
		// User.GET("/viewproducts", middlereware.UserAuth, controls.ViewProducts)

		// //User carts routes
		// User.GET("/profile/viewcart", middlereware.UserAuth, controls.ViewCart)
		// User.POST("/profile/addtocart", middlereware.UserAuth, controls.AddToCart)
		// User.GET("/fileterbycatogery", middlereware.UserAuth, controls.FilteringByCatogery)
		// User.GET("/cart/checkout", middlereware.UserAuth, controls.CheckOut)
		// User.DELETE("/deletecart/:id", middlereware.UserAuth, controls.DeleteCart)

		// //Oder managements by user
		// User.GET("/showorder", middlereware.UserAuth, controls.ShowOder)
		// User.GET("/order/return", middlereware.UserAuth, controls.ReturnOrderByUser)
		// User.GET("/order/cancelorder", middlereware.UserAuth, controls.CancelOrder)

		// //Coupon management
		// User.POST("/applycoupon", middlereware.UserAuth, controls.Applycoupon)
		// User.POST("/checkcoupon", middlereware.UserAuth, controls.CheckCoupon)

		//Forgot Password >>
		User.PUT("/forgotpassword", middlereware.UserAuth, controls.GenerateOtpForForgotPassword)
		User.POST("/forgotpassword/changepassword", middlereware.UserAuth, controls.ChangePassword)

		//payments route
		// User.GET("/payment/cashOnDelivery", middlereware.UserAuth, controls.CashOnDelivery)
		// User.GET("/payment/walletpayment", middlereware.UserAuth, controls.WalletPay)
		// User.GET("/payment/razorpay", middlereware.UserAuth, controls.Razorpay)
		// User.GET("/payment/success", middlereware.UserAuth, controls.RazorpaySuccess)
		// User.GET("/success", middlereware.UserAuth, controls.Success)
		// User.GET("payment/showwallet", middlereware.UserAuth, controls.ShowWallet)
		// User.GET("payment/wallethistory", middlereware.UserAuth, controls.WalletHistory)

		//invoice
		// User.GET("/invoice", middlereware.UserAuth, controls.InvoiceF)
		// User.GET("/invoice/download", middlereware.UserAuth, controls.Download)

	}
}
