package domain

import (
	"eticketing/module/entities"
	"time"
)

type SchedulesResponse struct {
	ID        uint64    `json:"id"`
	FilmID    uint64    `json:"film_id"`
	StudioID  uint64    `json:"studio_id"`
	Date      time.Time `json:"date"`
	StartTime string    `json:"start_time"`
	Price     float64   `json:"price"`
}

func ResponseArraySchedules(data []*entities.ScheduleModels) []*SchedulesResponse {
	res := make([]*SchedulesResponse, 0)

	for _, schedule := range data {
		scheduleRes := &SchedulesResponse{
			ID:        schedule.ID,
			FilmID:    schedule.FilmID,
			StudioID:  schedule.StudioID,
			Date:      schedule.Date,
			StartTime: schedule.StartTime,
			Price:     schedule.Price,
		}
		res = append(res, scheduleRes)
	}

	return res
}
