package repository

import (
	"eticketing/module/entities"
	"eticketing/module/feature/schedule/domain"
	"time"

	"gorm.io/gorm"
)

type ScheduleRepository struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) domain.ScheduleRepositoryInterface {
	return &ScheduleRepository{
		db: db,
	}
}

func (r *ScheduleRepository) CreateSchedule(schedule *entities.ScheduleModels) (*entities.ScheduleModels, error) {
	err := r.db.Create(&schedule).Error
	if err != nil {
		return nil, err
	}
	return schedule, nil
}

func (r *ScheduleRepository) UpdateSchedule(scheduleID uint64, updatedSchedule *entities.ScheduleModels) error {
	var schedule *entities.ScheduleModels
	if err := r.db.Where("id = ? AND deleted_at IS NULL", scheduleID).First(&schedule).Error; err != nil {
		return err
	}

	if err := r.db.Model(schedule).Updates(updatedSchedule).Error; err != nil {
		return err
	}

	return nil
}

func (r *ScheduleRepository) DeleteSchedule(scheduleID uint64) error {
	schedule := &entities.ScheduleModels{}
	if err := r.db.First(schedule, scheduleID).Error; err != nil {
		return err
	}

	if err := r.db.Model(schedule).Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}

	return nil
}

func (r *ScheduleRepository) GetTotalItems() (int64, error) {
	var totalItems int64

	if err := r.db.Where("deleted_at IS NULL").
		Model(&entities.ScheduleModels{}).Count(&totalItems).Error; err != nil {
		return 0, err
	}

	return totalItems, nil
}

func (r *ScheduleRepository) GetPaginatedSchedules(page, pageSize int) ([]*entities.ScheduleModels, error) {
	var schedules []*entities.ScheduleModels

	offset := (page - 1) * pageSize

	if err := r.db.Where("deleted_at IS NULL").
		Offset(offset).Limit(pageSize).Find(&schedules).Error; err != nil {
		return nil, err
	}

	return schedules, nil
}

func (r *ScheduleRepository) GetScheduleByID(scheduleID uint64) (*entities.ScheduleModels, error) {
	var schedule *entities.ScheduleModels

	if err := r.db.Where("id = ? AND deleted_at IS NULL", scheduleID).First(&schedule).Error; err != nil {
		return nil, err
	}
	return schedule, nil
}
