package storage

import (
	"time"

	"github.com/RollerM0bster/hw-esheludenko/hw12_13_14_15_calendar/internal/dto"
)

type Storage interface {
	// CreateEvent Добавляет событие в хранилище и возвращает идентификатор события
	CreateEvent(dto dto.Event) (int64, error)
	// ChangeEvent Изменяет событие
	ChangeEvent(id int64, dto dto.Event) error
	// DeleteEventById Удаляет событие по идентификатору
	DeleteEventById(id int64) error
	FindEventsByDay(day time.Time) ([]dto.Event, error)
	FindEventsByWeek(day time.Time) ([]dto.Event, error)
	FindEventsByMonth(day time.Time) ([]dto.Event, error)
}
