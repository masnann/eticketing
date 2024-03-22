package database

import (
	"eticketing/module/entities"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		entities.UserModels{},
		entities.FilmModels{},
		entities.ScheduleModels{},
		entities.StudioModels{},
		entities.SeatModels{},
		entities.OrderModels{},
		entities.OrderDetailModels{},
	)

	if err != nil {
		return
	}
	return
}
