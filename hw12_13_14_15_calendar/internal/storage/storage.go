package storage

import (
	"time"

	"github.com/RollerM0bster/hw-esheludenko/hw12_13_14_15_calendar/models"
)

type Storage interface {
	CreateEvent(dto models.NewEvent) (int64, error)
	ChangeEvent(id int64, dto models.NewEvent) error
	DeleteEventByID(id int64) error
	FindEventsByDay(day time.Time) ([]models.Event, error)
	FindEventsByWeek(day time.Time) ([]models.Event, error)
	FindEventsByMonth(day time.Time) ([]models.Event, error)
}
