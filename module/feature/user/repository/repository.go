package repository

import (
	"eticketing/module/entities"
	"eticketing/module/feature/user/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepositoryInterface {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetUserByID(addressID uint64) (*entities.UserModels, error) {
	var users *entities.UserModels

	if err := r.db.Where("id = ? AND deleted_at IS NULL", addressID).First(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
