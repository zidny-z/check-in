package controls

import (
	"fmt"

	"strconv"

	"checkin/config"

	"checkin/models"

	"github.com/gin-gonic/gin"
)

//>>>>>>>>>>>>>> Add rooms <<<<<<<<<<<<<<<<<<<<<<<<<<
func AddRoom(c *gin.Context) {
	var Room models.Room
	err := c.Bind(&Room)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Data binding error",
		})
		return
	}

	db := config.DB
	var count int64
	result := db.Find(&Room, "room_name = ?", Room.RoomName).Count(&count)
	if result.Error != nil {
		c.JSON(404, gin.H{
			"Error": result.Error.Error(),
		})
		return
	}
	if count == 0 {
		result := db.Create(&Room)
		if result.Error != nil {
			c.JSON(404, gin.H{
				"Error": result.Error.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"Message": "Successfully Added the Room",
		})
	} else {
		c.JSON(400, gin.H{
			"Message": "Room already exist",
		})
	}
}

//>>>>>>>>>>>>>>>>> View rooms <<<<<<<<<<<<<<<<<<<<<
func ViewRooms(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))
	type Room struct {
		Room_id uint
		Room_name string
		Facilities string
		Stock uint
		Avaliable uint
		Price uint
		Hotel_name string
	}
	var rooms []Room

	db := config.DB
	query := "SELECT rooms.room_id, rooms.room_name, rooms.facilities, rooms.stock, rooms.avaliable, rooms.price, hotels.hotel_name FROM rooms LEFT JOIN hotels ON rooms.hotel_id=hotels.id  GROUP BY rooms.room_id, hotels.hotel_name"

	if limit != 0 || offset != 0 {
		if limit == 0 {
			query = fmt.Sprintf("%s OFFSET %d", query, offset)
		} else if offset == 0 {
			query = fmt.Sprintf("%s LIMIT %d", query, limit)
		} else {
			query = fmt.Sprintf("%s LIMIT %d OFFSET %d", query, limit, offset)
		}
	}
	result := db.Raw(query).Scan(&rooms)
	if result.Error != nil {
		c.JSON(404, gin.H{
			"Error": result.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"rooms": rooms,
	})
}

//>>>>>>>>>>>>>>>>> Edit rooms <<<<<<<<<<<<<<<<<<<<<
func EditRoom(c *gin.Context) {
	bid := c.Param("id")
	id, err := strconv.Atoi(bid)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid ID",
		})
		return
	}
	var editrooms models.Room
	err = c.Bind(&editrooms)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Data binding error",
		})
		return
	}
	editrooms.RoomId = uint(id)
	db := config.DB

	result := db.Save(&editrooms).Updates(models.Room{
		RoomId: editrooms.RoomId,
		RoomName: editrooms.RoomName, 
		Facilities: editrooms.Facilities, 
		Stock: editrooms.Stock, 
		Avaliable: editrooms.Avaliable, 
		Price: editrooms.Price,
		HotelId: editrooms.HotelId,
	})

	if result.Error != nil {
		c.JSON(404, gin.H{
			"Error": result.Error.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"Message": "Successfully Updated the Room",
	})
}