package entities

import "time"

type StudioModels struct {
	ID        uint64     `gorm:"column:id;primaryKey" json:"id"`
	Name      string     `gorm:"column:name;type:VARCHAR(255)" json:"name"`
	Capacity  int        `gorm:"column:capacity" json:"capacity"`
	CreatedAt time.Time  `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"deleted_at"`
}

func (StudioModels) TableName() string {
	return "studio"
}
