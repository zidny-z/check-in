package controls

import (
	"fmt"

	"strconv"

	"check-in/config"
	"check-in/models"

	"github.com/gin-gonic/gin"
)

//Admin adding the hotel
func AddHotel(c *gin.Context) {
	var addhotel models.Hotel
	if c.Bind(&addhotel) != nil {
		c.JSON(400, gin.H{
			"Error": "Could not bind JSON data",
		})
		return
	}
	fmt.Println("================", addhotel.HotelName)
	DB := config.DB
	result := DB.Create(&addhotel)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"Error": result.Error.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"Message":       "New Hotel added Successfully",
		"Hotel details": addhotel,
	})
}

//view hotel
func ViewHotel(c *gin.Context) {
	var hotelData []models.Hotel
	db := config.DB
	result := db.First(&hotelData)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"Message": "Hotel is empty",
		})
		return
	}
	c.JSON(200, gin.H{
		"Hotels data": hotelData,
	})
}

//Edit hotel by admin
func EditHotel(c *gin.Context) {
	bid := c.Param("id")
	id, err := strconv.Atoi(bid)
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Error in string conversion",
		})
	}
	var edithotels models.Hotel
	if c.Bind(&edithotels) != nil {
		c.JSON(400, gin.H{
			"Error": "Error in binding the JSON data",
		})
		return
	}
	edithotels.ID = uint(id)
	DB := config.DB

	result := DB.Model(&edithotels).Updates(models.Hotel{
		HotelName: edithotels.HotelName,
		Location:  edithotels.Location,
		Phone:    edithotels.Phone,
	})

	if result.Error != nil {
		c.JSON(404, gin.H{
			"Error": result.Error.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"Message": "Successfully updated the Hotel",
	})
}