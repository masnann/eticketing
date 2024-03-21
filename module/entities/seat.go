package entities

import "time"

type SeatModels struct {
	ID        uint64       `gorm:"column:id;primaryKey" json:"id"`
	StudioID  uint64       `gorm:"column:studio_id" json:"studio_id"`
	Number    string       `gorm:"column:number" json:"number"`
	Status    string       `gorm:"column:status" json:"status"`
	CreatedAt time.Time    `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt time.Time    `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
	DeletedAt *time.Time   `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"deleted_at"`
	Studio    StudioModels `gorm:"foreignKey:StudioID" json:"studio"`
}

func (SeatModels) TableName() string {
	return "seat"
}