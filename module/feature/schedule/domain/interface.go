package domain

import (
	"eticketing/module/entities"

	"github.com/gofiber/fiber/v2"
)

type ScheduleRepositoryInterface interface {
	CreateSchedule(schedule *entities.ScheduleModels) (*entities.ScheduleModels, error)
	GetTotalItems() (int64, error)
	GetPaginatedSchedules(page, pageSize int) ([]*entities.ScheduleModels, error)
	GetScheduleByID(scheduleID uint64) (*entities.ScheduleModels, error)
	UpdateSchedule(scheduleID uint64, updatedSchedule *entities.ScheduleModels) error
	DeleteSchedule(scheduleID uint64) error
}

type ScheduleServiceInterface interface {
	CreateSchedule(req *CreateScheduleRequest) (*entities.ScheduleModels, error)
	GetAllSchedules(page, pageSize int) ([]*entities.ScheduleModels, int64, error)
	GetSchedulePage(currentPage, pageSize, totalItems int) (int, int, int, error)
	UpdateSchedule(scheduleID uint64, req *UpdateScheduleRequest) error
	DeleteSchedule(scheduleID uint64) error
}

type ScheduleHandlerInterface interface {
	CreateSchedule(c *fiber.Ctx) error
	GetAllSchedules(c *fiber.Ctx) error
	UpdateSchedule(c *fiber.Ctx) error
	DeleteSchedule(c *fiber.Ctx) error	
}
