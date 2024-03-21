package service

import (
	"errors"
	"eticketing/module/entities"
	"eticketing/module/feature/schedule/domain"
	"eticketing/module/feature/schedule/mocks"
	"math"

	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func setupScheduleServiceTest(t *testing.T) (*mocks.ScheduleRepositoryInterface, domain.ScheduleServiceInterface) {
	repo := mocks.NewScheduleRepositoryInterface(t)
	service := NewScheduleService(repo)
	return repo, service
}

func TestCreateSchedule(t *testing.T) {
	filmID := uint64(1)
	studioID := uint64(1)
	date := time.Now()
	startTime := "19:00:00"
	price := 15.50

	t.Run("Success Case - Schedule Created", func(t *testing.T) {
		repo, service := setupScheduleServiceTest(t)
		req := &domain.CreateScheduleRequest{
			FilmID:    filmID,
			StudioID:  studioID,
			Date:      date,
			StartTime: startTime,
			Price:     price,
		}
		expectedSchedule := &entities.ScheduleModels{
			FilmID:    filmID,
			StudioID:  studioID,
			Date:      date,
			StartTime: startTime,
			Price:     price,
			CreatedAt: time.Now(),
		}

		repo.On("CreateSchedule", expectedSchedule).Return(expectedSchedule, nil)

		schedule, err := service.CreateSchedule(req)

		assert.Nil(t, err)
		assert.NotNil(t, schedule)
		assert.Equal(t, expectedSchedule, schedule)

		repo.AssertExpectations(t)
	})

	t.Run("Error Case - Schedule Creation Failure", func(t *testing.T) {
		repo, service := setupScheduleServiceTest(t)
		req := &domain.CreateScheduleRequest{
			FilmID:    filmID,
			StudioID:  studioID,
			Date:      date,
			StartTime: startTime,
			Price:     price,
		}
		expectedSchedule := &entities.ScheduleModels{
			FilmID:    filmID,
			StudioID:  studioID,
			Date:      date,
			StartTime: startTime,
			Price:     price,
			CreatedAt: time.Now(),
		}
		expectedErr := errors.New("schedule creation failed")

		repo.On("CreateSchedule", expectedSchedule).Return(nil, expectedErr)

		schedule, err := service.CreateSchedule(req)

		assert.Error(t, err)
		assert.Nil(t, schedule)
		assert.EqualError(t, err, expectedErr.Error())

		repo.AssertExpectations(t)
	})
}

func TestGetAllSchedules(t *testing.T) {
	t.Run("Success Case - Schedules Found", func(t *testing.T) {
		repo, service := setupScheduleServiceTest(t)
		page := 1
		pageSize := 10
		totalItems := int64(20)
		expectedSchedules := []*entities.ScheduleModels{
			{ID: 1, FilmID: 1, StudioID: 1, Date: time.Now(), StartTime: "19:00:00", Price: 15.50},
			{ID: 2, FilmID: 2, StudioID: 2, Date: time.Now(), StartTime: "20:00:00", Price: 20.00},
		}

		repo.On("GetPaginatedSchedules", page, pageSize).Return(expectedSchedules, nil)
		repo.On("GetTotalItems").Return(totalItems, nil)

		schedules, total, err := service.GetAllSchedules(page, pageSize)

		assert.NoError(t, err)
		assert.NotNil(t, schedules)
		assert.Equal(t, expectedSchedules, schedules)
		assert.Equal(t, totalItems, total)

		repo.AssertExpectations(t)
	})

	t.Run("Error Case - Failed to Get Schedules", func(t *testing.T) {
		repo, service := setupScheduleServiceTest(t)
		page := 1
		pageSize := 10
		expectedErr := errors.New("failed to get schedules")

		repo.On("GetPaginatedSchedules", page, pageSize).Return(nil, expectedErr)

		schedules, total, err := service.GetAllSchedules(page, pageSize)

		assert.Error(t, err)
		assert.Nil(t, schedules)
		assert.Zero(t, total)
		assert.EqualError(t, err, expectedErr.Error())

		repo.AssertExpectations(t)
	})

	t.Run("Error Case - Failed to Get Total Items", func(t *testing.T) {
		repo, service := setupScheduleServiceTest(t)
		page := 1
		pageSize := 10
		expectedErr := errors.New("failed to get total items")

		repo.On("GetPaginatedSchedules", page, pageSize).Return(nil, nil)
		repo.On("GetTotalItems").Return(int64(0), expectedErr)

		schedules, total, err := service.GetAllSchedules(page, pageSize)

		assert.Error(t, err)
		assert.Nil(t, schedules)
		assert.Zero(t, total)
		assert.EqualError(t, err, expectedErr.Error())

		repo.AssertExpectations(t)
	})
}

func TestGetSchedulePage(t *testing.T) {
	service := &ScheduleService{} // Inisialisasi service

	t.Run("Success Case - Current Page within Bounds", func(t *testing.T) {
		currentPage := 1
		pageSize := 10
		totalItems := 20

		expectedTotalPages := int(math.Ceil(float64(totalItems) / float64(pageSize)))
		expectedNextPage := currentPage + 1
		expectedPrevPage := currentPage - 1

		totalPages, nextPage, prevPage, err := service.GetSchedulePage(currentPage, pageSize, totalItems)

		assert.NoError(t, err)
		assert.Equal(t, expectedTotalPages, totalPages)
		assert.Equal(t, expectedNextPage, nextPage)
		assert.Equal(t, expectedPrevPage, prevPage)
	})

	t.Run("Success Case - Current Page Zero", func(t *testing.T) {
		currentPage := 0
		pageSize := 10
		totalItems := 20

		expectedTotalPages := int(math.Ceil(float64(totalItems) / float64(pageSize)))
		expectedNextPage := currentPage + 1
		expectedPrevPage := 0

		totalPages, nextPage, prevPage, err := service.GetSchedulePage(currentPage, pageSize, totalItems)

		assert.NoError(t, err)
		assert.Equal(t, expectedTotalPages, totalPages)
		assert.Equal(t, expectedNextPage, nextPage)
		assert.Equal(t, expectedPrevPage, prevPage)
	})

	t.Run("Success Case - Next Page Beyond Total Pages", func(t *testing.T) {
		currentPage := 2
		pageSize := 10
		totalItems := 20

		expectedTotalPages := int(math.Ceil(float64(totalItems) / float64(pageSize)))
		expectedNextPage := 0
		expectedPrevPage := currentPage - 1

		totalPages, nextPage, prevPage, err := service.GetSchedulePage(currentPage, pageSize, totalItems)

		assert.NoError(t, err)
		assert.Equal(t, expectedTotalPages, totalPages)
		assert.Equal(t, expectedNextPage, nextPage)
		assert.Equal(t, expectedPrevPage, prevPage)
	})

	t.Run("Success Case - Previous Page Before First Page", func(t *testing.T) {
		currentPage := 1
		pageSize := 10
		totalItems := 20

		expectedTotalPages := int(math.Ceil(float64(totalItems) / float64(pageSize)))
		expectedNextPage := currentPage + 1
		expectedPrevPage := 0

		totalPages, nextPage, prevPage, err := service.GetSchedulePage(currentPage, pageSize, totalItems)

		assert.NoError(t, err)
		assert.Equal(t, expectedTotalPages, totalPages)
		assert.Equal(t, expectedNextPage, nextPage)
		assert.Equal(t, expectedPrevPage, prevPage)
	})
}

func TestUpdateSchedule(t *testing.T) {
	scheduleID := uint64(1)
	schedule := &entities.ScheduleModels{
		ID:        scheduleID,
		FilmID:    1,
		StudioID:  1,
		Date:      time.Now(),
		StartTime: "09:00",
		Price:     9.50,
	}
	req := &domain.UpdateScheduleRequest{
		FilmID:    2,
		StudioID:  2,
		Date:      time.Now(),
		StartTime: "10:00",
		Price:     10.50,
	}

	repo, service := setupScheduleServiceTest(t)
	t.Run("Failed Case - Schedule Not Found", func(t *testing.T) {
		expectedErr := errors.New("schedule not found")
		repo.On("GetScheduleByID", scheduleID).Return(nil, expectedErr).Once()

		err := service.UpdateSchedule(scheduleID, req)

		assert.Equal(t, expectedErr, err)
		repo.AssertExpectations(t)
	})

	t.Run("Success Case - Schedule Updated", func(t *testing.T) {
		repo.On("GetScheduleByID", scheduleID).Return(schedule, nil).Once()
		repo.On("UpdateSchedule", scheduleID, &entities.ScheduleModels{
			FilmID:    req.FilmID,
			StudioID:  req.StudioID,
			Date:      req.Date,
			StartTime: req.StartTime,
			Price:     req.Price,
			UpdatedAt: time.Now(),
		}).Return(nil).Once()

		err := service.UpdateSchedule(scheduleID, req)

		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Case - Error Updating Schedule", func(t *testing.T) {
		expectedErr := errors.New("failed to update schedule")
		repo.On("GetScheduleByID", scheduleID).Return(schedule, nil).Once()
		repo.On("UpdateSchedule", scheduleID, &entities.ScheduleModels{
			FilmID:    req.FilmID,
			StudioID:  req.StudioID,
			Date:      req.Date,
			StartTime: req.StartTime,
			Price:     req.Price,
			UpdatedAt: time.Now(),
		}).Return(expectedErr).Once()

		err := service.UpdateSchedule(scheduleID, req)

		assert.Equal(t, expectedErr, err)
		repo.AssertExpectations(t)
	})
}


func TestScheduleService_DeleteSchedule(t *testing.T) {
	scheduleID := uint64(1)

	repo, service := setupScheduleServiceTest(t)

	t.Run("Failed Case - Schedule Not Found", func(t *testing.T) {
		expectedErr := errors.New("schedule not found")
		repo.On("GetScheduleByID", scheduleID).Return(nil, expectedErr).Once()

		err := service.DeleteSchedule(scheduleID)

		assert.Equal(t, expectedErr, err)
		repo.AssertExpectations(t)
	})

	t.Run("Success Case - Schedule Deleted", func(t *testing.T) {
		schedule := &entities.ScheduleModels{
			ID: scheduleID,
		}
		repo.On("GetScheduleByID", scheduleID).Return(schedule, nil).Once()
		repo.On("DeleteSchedule", scheduleID).Return(nil).Once()

		err := service.DeleteSchedule(scheduleID)

		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Case - Error Deleting Schedule", func(t *testing.T) {
		expectedErr := errors.New("failed to delete schedule")
		schedule := &entities.ScheduleModels{
			ID: scheduleID,
		}
		repo.On("GetScheduleByID", scheduleID).Return(schedule, nil).Once()
		repo.On("DeleteSchedule", scheduleID).Return(expectedErr).Once()

		err := service.DeleteSchedule(scheduleID)

		assert.Equal(t, expectedErr, err)
		repo.AssertExpectations(t)
	})
}
