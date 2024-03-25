package controls

import (
	"strconv"

	"checkin/config"
	"checkin/models"

	"github.com/gin-gonic/gin"
)

type ProfileData struct {
	Firstname   string
	Lastname    string
	Email       string
	PhoneNumber int
}

//>>>>>>>>>>> Admin profile <<<<<<<<<<<<<<<<<<<<<<<
func AdminProfile(c *gin.Context) {
	adminid, err := strconv.Atoi(c.GetString("adminid"))
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Error in string conversion ",
		})
		return
	}
	var user_data models.Admin
	db := config.DB
	result := db.First(&user_data, adminid)
	if result.Error != nil {
		c.JSON(404, gin.H{
			"Error":   result.Error.Error(),
			"Message": "Admin does't exist",
			"Status":  "false",
		})
		return
	}
	c.JSON(200, gin.H{

		"First name ":  user_data.Firstname,
		"Last Name":    user_data.Lastname,
		"Email":        user_data.Email,
		"Phone number": user_data.PhoneNumber,
	})
}
