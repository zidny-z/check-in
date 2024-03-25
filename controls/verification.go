package controls

import (
	"crypto/rand"

	"net/http"

	"math/big"
	"strconv"

	"checkin/config"
	"checkin/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func VerifyOTP(email string) string {
	Otp, err := getRandNum()
	if err != nil {
		panic(err)
	}
	return Otp
}
func getRandNum() (string, error) {
	nBig, e := rand.Int(rand.Reader, big.NewInt(8999))
	if e != nil {
		return "", e
	}
	return strconv.FormatInt(nBig.Int64()+1000, 10), nil
}

//-------Otp validtioin------------->
func OtpValidation(c *gin.Context) {
	type User_otp struct {
		Otp   string
		Email string
	}
	var user_otp User_otp
	var userData models.User
	if c.Bind(&user_otp) != nil {
		c.JSON(400, gin.H{
			"Error": "Could not bind the JSON Data",
		})
		return
	}
	db := config.DB
	result := db.First(&userData, "otp LIKE ? AND email LIKE ?", user_otp.Otp, user_otp.Email)

	if result.Error != nil {
		c.JSON(404, gin.H{
			"Error": result.Error.Error(),
		})
		// db.Last(&userData).Delete(&userData)
		c.JSON(422, gin.H{
			"Error":   "Wrong OTP Register Once agian",
			"Message": "Goto /signup/otpvalidate",
		})
		return
	}
	c.JSON(200, gin.H{
		"Message": "New User Successfully Registered",
	})
}

//Generating otp for forgot password
func GenerateOtpForForgotPassword(c *gin.Context) {
	type UserEnterDate struct {
		Email string
	}
	var data UserEnterDate
	if c.Bind(&data) != nil {
		c.JSON(400, gin.H{
			"Error": "Error when the data binding",
		})
		return
	}
	otp := VerifyOTP(data.Email)

	db := config.DB
	var userData models.User

	result := db.Model(userData).Where("email = ?", data.Email).Update("otp", otp)
	if result.Error != nil {
		c.JSON(409, gin.H{
			"Error": "User not exist",
		})
		return
	}
	c.JSON(200, gin.H{
		"Message": "go to /user/changepassword with otp",
		"otp":     otp,
	})
}

//Reseting the password
func ChangePassword(c *gin.Context) {
	type userEnterData struct {
		Email           string
		Otp             string
		Password        string
		ConfirmPassword string
	}
	var data userEnterData
	var userData models.User
	if c.Bind(&data) != nil {
		c.JSON(400, gin.H{
			"Error": "Error when the data binding",
		})
		return
	}
	if data.Password != data.ConfirmPassword {
		c.JSON(400, gin.H{
			"Error": "Password not match",
		})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": "False",
			"Error":  "Hashing password error",
		})
		return
	}
	db := config.DB
	result := db.Find(&userData, "email = ?", data.Email)
	if result.Error != nil {
		c.JSON(409, gin.H{
			"Error": "User not exist",
		})
		return
	}
	if data.Otp != userData.Otp {
		c.JSON(400, gin.H{
			"Error": "Invalide otp",
		})
		return
	}
	result = db.Model(&userData).Where("email = ?", data.Email).Update("password", hash)
	if result.Error != nil {
		c.JSON(409, gin.H{
			"Error": "User not exist",
		})
		return
	}
	c.JSON(200, gin.H{
		"Message": "Password Change successfully",
	})
}
