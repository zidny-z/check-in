package controls

import (
	"fmt"

	"strconv"

	"checkin/config"
	"checkin/models"

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

// View Hotels and the rooms
// func ViewHotelRooms(c *gin.Context) {
//   limit, _ := strconv.Atoi(c.Query("limit"))
//   offset, _ := strconv.Atoi(c.Query("offset"))

//   type Room struct {
//     RoomID     uint   `json:"Room_id"`
//     RoomName   string `json:"room_name"`
//     Facilities string `json:"facilities"`
//     Stock      uint   `json:"stock"`
//     Available  uint   `json:"available"`
//     Price      uint   `json:"price"`
//     HotelID    uint   `json:"Hotel_id"` // Renamed for consistency with model
//   }

//   type Hotel struct {
//     ID        uint   `json:"id"`
//     HotelName string `json:"hotel_name"`
//     Location  string `json:"location"`
//     Phone     string `json:"phone"`
//     Rooms     []Room `json:"rooms"`
//   }

//   var hotels []Hotel

//   db := config.DB
//   query := `SELECT rooms.room_id, rooms.room_name, rooms.facilities, rooms.stock, rooms.avaliable, rooms.price, hotels.id, hotels.hotel_name, hotels.location, hotels.phone
//             FROM rooms
//             LEFT JOIN hotels ON rooms.hotel_id=hotels.id
//             GROUP BY hotels.id, rooms.room_id` // Group by hotel.id first

//   if limit != 0 || offset != 0 {
//     if limit == 0 {
//       query = fmt.Sprintf("%s OFFSET %d", query, offset)
//     } else if offset == 0 {
//       query = fmt.Sprintf("%s LIMIT %d", query, limit)
//     } else {
//       query = fmt.Sprintf("%s LIMIT %d OFFSET %d", query, limit, offset)
//     }
//   }

//   rows, err := db.Raw(query).Rows()
//   if err != nil {
//     c.JSON(404, gin.H{
//       "Error": err.Error(),
//     })
//     return
//   }
//   defer rows.Close()

//   var currentHotelID uint
//   var currentHotelData Hotel

//   for rows.Next() {
//     var room Room
//     var hotel Hotel
//     err := rows.Scan(&room.RoomID, &room.RoomName, &room.Facilities, &room.Stock, &room.Available, &room.Price, &hotel.ID, &hotel.HotelName, &hotel.Location, &hotel.Phone)
//     if err != nil {
//       c.JSON(404, gin.H{
//         "Error": err.Error(),
//       })
//       return
//     }

//     if currentHotelID != currentHotelData.ID {
//       // New hotel encountered, start a new record
//       currentHotelData = Hotel{
//         ID:        hotel.ID,
//         HotelName: hotel.HotelName,
//         Location:  hotel.Location,
//         Phone:     hotel.Phone,
//         Rooms:     []Room{},
//       }
//       hotels = append(hotels, currentHotelData)
//     }

//     currentHotelData.Rooms = append(currentHotelData.Rooms, room)
//     currentHotelID = hotel.ID
//   }

//   c.JSON(200, gin.H{
//     "hotels": hotels,
//   })
// }
