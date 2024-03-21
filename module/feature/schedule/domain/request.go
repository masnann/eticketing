package domain

import "time"

type CreateScheduleRequest struct {
	FilmID    uint64    `json:"film_id," validate:"required"`
	StudioID  uint64    `json:"studio_id" validate:"required"`
	Date      time.Time `json:"date" validate:"required"`
	StartTime string    `json:"start_time" validate:"required"`
	Price     float64   `json:"price" validate:"required"`
}

type UpdateScheduleRequest struct {
	FilmID    uint64    `json:"film_id,"`
	StudioID  uint64    `json:"studio_id"`
	Date      time.Time `json:"date" `
	StartTime string    `json:"start_time"`
	Price     float64   `json:"price"`
}
