package storage

import (
	"errors"
	"sync"
	"time"

	"github.com/RollerM0bster/hw-esheludenko/hw12_13_14_15_calendar/internal/dto"
)

var ErrEventNotFound = errors.New("event not found")

type Storage struct {
	mu     sync.RWMutex //nolint:unused
	Events map[int64]dto.Event
	nextId int64
}

func New() *Storage {
	return &Storage{
		Events: make(map[int64]dto.Event),
		nextId: 1,
	}
}

func (s *Storage) CreateEvent(dto dto.Event) (int64, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	dto.ID = s.nextId
	s.nextId++
	s.Events[dto.ID] = dto
	return dto.ID, nil
}

func (s *Storage) ChangeEvent(id int64, dto dto.Event) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if _, exists := s.Events[id]; !exists {
		return ErrEventNotFound
	}
	s.Events[id] = dto
	return nil
}

func (s *Storage) DeleteEventById(id int64) error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if _, exists := s.Events[id]; !exists {
		return ErrEventNotFound
	}
	delete(s.Events, id)
	return nil
}

func (s *Storage) FindEventsByDay(day time.Time) ([]dto.Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var filteredEvents []dto.Event
	for _, event := range s.Events {
		if event.Start.Year() == day.Year() && event.Start.Month() == day.Month() && event.Start.Day() == day.Day() {
			filteredEvents = append(filteredEvents, event)
		}
	}
	return filteredEvents, nil
}

func (s *Storage) FindEventsByWeek(day time.Time) ([]dto.Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var filteredEvents []dto.Event
	year, week := day.ISOWeek()
	for _, event := range s.Events {
		eYear, eWeek := event.Start.ISOWeek()
		if eYear == year && eWeek == week {
			filteredEvents = append(filteredEvents, event)
		}
	}
	return filteredEvents, nil
}

func (s *Storage) FindEventsByMonth(day time.Time) ([]dto.Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var filteredEvents []dto.Event
	for _, event := range s.Events {
		if event.Start.Year() == day.Year() && event.Start.Month() == day.Month() {
			filteredEvents = append(filteredEvents, event)
		}
	}
	return filteredEvents, nil
}
