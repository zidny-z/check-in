package main

import (
	controller "check-in/controllers"
	"check-in/database"
	router "check-in/routers"
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Println("Starting application...")
	fmt.Println()

	db, err := database.Start()
	if err != nil {
		fmt.Println("Error starting database")
		return
	}

	cont := controller.New(db)

	err = router.StartServer(cont)
	if err != nil {
		fmt.Println("Error starting server")
		return
	}

	// contHotel := controller.NewHotelController(db)
	// // make data hotel with hotelRepository
	// hotel := models.Hotel{
	// 	Name:        "Hotel A",
	// 	Location:   "Jakarta",
	// 	RoomCount: 10,
	// 	RoomAvailable: 10,
	// 	Star: 5,
	// 	IsSyariah: false,
	// 	Photo: "photo.jpg",
	// 	Facility: "wifi, tv, ac",
	// }
	// contHotel.Create(hotel)

}