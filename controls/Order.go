package controls

import (
	"checkin/config"
	"checkin/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// import (
// 	"fmt"
// 	"strconv"
// 	"time"

// 	"checkin/config"
// 	"checkin/models"

// 	"github.com/gin-gonic/gin"
// )

// //>>>>>>>>>> MakeORder <<<<<<<<<<<<<<<<
func OrderRoom(c *gin.Context) {
	userId, err := strconv.Atoi(c.GetString("userid"))
	var requestBody struct{
		RoomId uint `json:"RoomId"`
		People uint `json:"people"`
	}

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{
			"Error": "Error in binding json",
		})
		return
	}

	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Error in string conversion userid",
		})
	}

	// check room capacity
	if requestBody.People > 3 {
		c.JSON(400, gin.H{
			"Error": "Room can only accomodate 3 people",
		})
		return
	}

	var room models.Room
	var user models.User
	var oderItem models.Order

	db := config.DB
	result := db.Find(&room, "room_id = ?", requestBody.RoomId)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"Error": result.Error.Error(),
		})
		return
	}

	if room.Avaliable <= 0{
		c.JSON(400,gin.H{
			"Error": "Kamar tida tersedia",
		})
		return
	}

	result = db.Find(&user, "id = ?", userId)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"Error": result.Error.Error(),
		})
		return
	}

	oderItem = models.Order{
		UserIdNo:    uint(userId),
		TotalAmount: uint(room.Price),
		RoomId:      uint(requestBody.RoomId),
		People:      uint(requestBody.People),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	result = db.Create(&oderItem)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"Error": result.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"Message": "Room odered successfully",
		"orderID":  oderItem.OrderId,    
	})
}

// //>>>>>>>>>> make payment for order <<<<<<<<<<<<<<<<
func OrderPayment(c *gin.Context){
	userId, err := strconv.Atoi(c.GetString("userid"))
	var requestBody struct{
		OrderId uint `json:"OrderId"`
	}

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{
			"Error": "Error in binding json",
		})
		return
	}

	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Error in string conversion userid",
		})
	}

	var order models.Order
	var user models.User
	var payment models.Payment

	db := config.DB
	result := db.Find(&order, "order_id = ?", requestBody.OrderId)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"Error": result.Error.Error(),
		})
		return
	}

	result = db.Find(&user, "id = ?", userId)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"Error": result.Error.Error(),
		})
		return
	}

	payment = models.Payment{
		UserId:    uint(userId),
		Totalamount: uint(order.TotalAmount),
		Status: "Belum Lunas",
		Date:   time.Now(),
		OrderId: uint(requestBody.OrderId),
	}

	result = db.Create(&payment)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"Error": result.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"Message": "payment created successfully",
		"paymentID":  payment.PaymentId,    
	})
}

// //>>>>>>>>>> Validate payment <<<<<<<<<<<<<<<<
func ValidatePayment(c *gin.Context){
	bid := c.Param("id")
	id, err := strconv.Atoi(bid)
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Error in string conversion",
		})		
	}

	var editpayment models.Payment
	if c.Bind(&editpayment) != nil {
		c.JSON(400, gin.H{
			"Error": "Error in binding the JSON data",
		})
		return
	}

	editpayment.PaymentId = uint(id)
	DB := config.DB

	result := DB.Model(&editpayment).Updates(models.Payment{
		Status: "Lunas",
		Date:   time.Now(),
		NoKTP: editpayment.NoKTP,
	})

	if result.Error != nil {
		c.JSON(404, gin.H{
			"Error": result.Error.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"Message": "Validate the payment",
	})
}

