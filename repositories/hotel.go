package repositories

import (
	"check-in/models"

	"gorm.io/gorm"
)

type HotelRepository struct {
	db *gorm.DB
}

func NewHotelRepository(db *gorm.DB) HotelRepository {
	return HotelRepository{
		db: db,
	}
}

func (r HotelRepository) FindAll() ([]models.Hotel, error) {
	var hotels []models.Hotel
	if error := r.db.Find(&hotels).Error; error != nil {
		return []models.Hotel{}, error
	}

	return hotels, nil
}

func (r HotelRepository) FindByID(id int) (models.Hotel, error) {
	var hotel models.Hotel
	if error := r.db.First(&hotel, id).Error; error != nil {
		return models.Hotel{}, error
	}

	return hotel, nil
}

func (r HotelRepository) Create(hotel models.Hotel) (models.Hotel, error) {
	if error := r.db.Create(&hotel).Error; error != nil {
		return models.Hotel{}, error
	}

	return hotel, nil
}

func (r HotelRepository) Update(hotel models.Hotel) (models.Hotel, error) {
	if error := r.db.Save(&hotel).Error; error != nil {
		return models.Hotel{}, error
	}

	return hotel, nil
}

func (r HotelRepository) Delete(id int) error {
	if error := r.db.Delete(&models.Hotel{}, id).Error; error != nil {
		return error
	}

	return nil
}
