package domain

import (
	"eticketing/module/entities"
	"time"
)

type UserResponse struct {
	ID           uint64    `json:"id"`
	Email        string    `json:"email"`
	Password     string    `json:"-"`
	Phone        string    `json:"phone"`
	Name         string    `json:"name"`
	PhotoProfile string    `json:"photo_profile"`
	Gender       string    `json:"gender"`
	DateOfBirth  time.Time `json:"date_of_birth"`
	Role         string    `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
}

func UserFormatter(user *entities.UserModels) *UserResponse {
	result := &UserResponse{
		ID:           user.ID,
		Email:        user.Email,
		Password:     "",
		Phone:        user.Phone,
		Name:         user.Name,
		PhotoProfile: user.PhotoProfile,
		Gender:       user.Gender,
		DateOfBirth:  user.DateOfBirth,
		Role:         user.Role,
		CreatedAt:    user.CreatedAt,
	}
	return result
}
