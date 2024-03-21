package entities

import "time"

type UserModels struct {
	ID           uint64     `gorm:"column:id;primaryKey" json:"id"`
	Email        string     `gorm:"column:email;type:VARCHAR(255)" json:"email"`
	Password     string     `gorm:"column:password;type:VARCHAR(255)" json:"password"`
	Phone        string     `gorm:"column:phone;type:VARCHAR(255)" json:"phone"`
	Role         string     `gorm:"column:role;type:VARCHAR(255)" json:"role"`
	Name         string     `gorm:"column:name;type:VARCHAR(255)" json:"name"`
	PhotoProfile string     `gorm:"column:photo_profile;type:VARCHAR(255)" json:"photo_profile"`
	Gender       string     `gorm:"column:gender;type:VARCHAR(255)" json:"gender"`
	DateOfBirth  time.Time  `gorm:"column:date_of_birth;type:DATE" json:"date_of_birth"`
	CreatedAt    time.Time  `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
	DeletedAt    *time.Time `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"deleted_at"`
}

func (UserModels) TableName() string {
	return "users"
}