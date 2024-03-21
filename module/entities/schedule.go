package entities

import "time"

type ScheduleModels struct {
	ID        uint64       `gorm:"column:id;primaryKey" json:"id"`
	FilmID    uint64       `gorm:"column:film_id" json:"film_id"`
	StudioID  uint64       `gorm:"column:studio_id" json:"studio_id"`
	Date      time.Time    `gorm:"column:date" json:"date"`
	StartTime string       `gorm:"column:start_time" json:"start_time"`
	Price     float64      `gorm:"column:price" json:"price"`
	CreatedAt time.Time    `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt time.Time    `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
	DeletedAt *time.Time   `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"deleted_at"`
	Film      FilmModels   `gorm:"foreignKey:FilmID" json:"film"`
	Studio    StudioModels `gorm:"foreignKey:StudioID" json:"studio"`
}

func (ScheduleModels) TableName() string {
	return "schedule"
}
