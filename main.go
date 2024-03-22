package main

import (
	"fmt"
	"check-in/controller"
	"check-in/database"
	router "check-in/routers"
)

func main() {
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
}