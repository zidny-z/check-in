package controller

import (
	"check-in/database"
)

type Controller struct {
	db database.Database
}


func New(db database.Database) Controller {
	return Controller{
		db: db,
	}
}
