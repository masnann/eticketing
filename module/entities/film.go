package entities

import "time"

type FilmModels struct {
	ID           uint64       `gorm:"column:id;primaryKey" json:"id"`
	Title        string       `gorm:"column:title;type:VARCHAR(255)" json:"title"`
	Genre        string       `gorm:"column:genre;type:VARCHAR(255)" json:"genre"`
	Year         int          `gorm:"column:year" json:"year"`
	Duration     int          `gorm:"column:duration" json:"duration"`
	Director     string       `gorm:"column:director;type:VARCHAR(255)" json:"director"`
	Description  string       `gorm:"column:description;type:TEXT" json:"description"`
	Country      string       `gorm:"column:country;type:VARCHAR(255)" json:"country"`
	Language     string       `gorm:"column:language;type:VARCHAR(255)" json:"language"`
	PosterURL    string       `gorm:"column:poster_url;type:VARCHAR(255)" json:"poster_url"`
	TrailerURL   string       `gorm:"column:trailer_url;type:VARCHAR(255)" json:"trailer_url"`
	CreatedAt    time.Time    `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt    time.Time    `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
	DeletedAt    *time.Time   `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"deleted_at"`
}

func (FilmModels) TableName() string {
	return "film"
}
