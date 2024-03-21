package entities

import "time"

	type FilmModels struct {
		ID        uint64     `gorm:"column:id;primaryKey" json:"id"`
		Title     string     `gorm:"column:title;type:VARCHAR(255)" json:"title"`
		Genre     string     `gorm:"column:genre;type:VARCHAR(255)" json:"genre"`
		Year      int        `gorm:"column:year" json:"year"`
		CreatedAt time.Time  `gorm:"column:created_at;type:timestamp" json:"created_at"`
		UpdatedAt time.Time  `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
		DeletedAt *time.Time `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"deleted_at"`
	}

func (FilmModels) TableName() string {
	return "film"
}
