package service

import (
	"errors"
	"eticketing/module/entities"
	"eticketing/module/feature/schedule/domain"
	"math"
	"time"
)

type ScheduleService struct {
	repo domain.ScheduleRepositoryInterface
}

func NewScheduleService(repo domain.ScheduleRepositoryInterface) domain.ScheduleServiceInterface {
	return &ScheduleService{
		repo: repo,
	}
}

func (s *ScheduleService) CreateSchedule(req *domain.CreateScheduleRequest) (*entities.ScheduleModels, error) {
	newData := &entities.ScheduleModels{
		FilmID:    req.FilmID,
		StudioID:  req.StudioID,
		Date:      req.Date,
		StartTime: req.StartTime,
		Price:     req.Price,
		CreatedAt: time.Now(),
	}

	createdSchedule, err := s.repo.CreateSchedule(newData)
	if err != nil {
		return nil, err
	}
	return createdSchedule, nil
}

func (s *ScheduleService) GetAllSchedules(page, pageSize int) ([]*entities.ScheduleModels, int64, error) {
	result, err := s.repo.GetPaginatedSchedules(page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	totalItems, err := s.repo.GetTotalItems()
	if err != nil {
		return nil, 0, err
	}

	return result, totalItems, nil
}

func (s *ScheduleService) GetSchedulePage(currentPage, pageSize, totalItems int) (int, int, int, error) {
	totalPages := int(math.Ceil(float64(totalItems) / float64(pageSize)))
	nextPage := currentPage + 1
	prevPage := currentPage - 1

	if nextPage > totalPages {
		nextPage = 0
	}

	if prevPage < 1 {
		prevPage = 0
	}

	return totalPages, nextPage, prevPage, nil
}

func (s *ScheduleService) UpdateSchedule(scheduleID uint64, req *domain.UpdateScheduleRequest) error {
	schedule, err := s.repo.GetScheduleByID(scheduleID)
	if err != nil {
		return errors.New("schedule not found")
	}

	newData := &entities.ScheduleModels{
		FilmID:    req.FilmID,
		StudioID:  req.StudioID,
		Date:      req.Date,
		StartTime: req.StartTime,
		Price:     req.Price,
		UpdatedAt:   time.Now(),
	}

	err = s.repo.UpdateSchedule(schedule.ID, newData)
	if err != nil {
		return err
	}

	return nil
}

func (s *ScheduleService) DeleteSchedule(scheduleID uint64) error {
	schedule, err := s.repo.GetScheduleByID(scheduleID)
	if err != nil {
		return errors.New("schedule not found")
	}

	err = s.repo.DeleteSchedule(schedule.ID)
	if err != nil {
		return err
	}

	return nil
}


