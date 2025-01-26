package storage

import (
	"errors"
	"sync"
	"time"

	"github.com/RollerM0bster/hw-esheludenko/hw12_13_14_15_calendar/models"
)

var ErrEventNotFound = errors.New("event not found")

type Storage struct {
	mu     sync.RWMutex
	Events map[int64]models.Event
	nextID int64
}

func New() *Storage {
	return &Storage{
		Events: make(map[int64]models.Event),
		nextID: 1,
	}
}

func (s *Storage) CreateEvent(dto models.NewEvent) (int64, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	event := toEventDto(dto)
	event.ID = s.nextID
	s.nextID++
	s.Events[event.ID] = event
	return event.ID, nil
}

func (s *Storage) ChangeEvent(id int64, dto models.NewEvent) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if _, exists := s.Events[id]; !exists {
		return ErrEventNotFound
	}
	changedEventDto := toEventDto(dto)
	s.Events[id] = changedEventDto
	return nil
}

func (s *Storage) DeleteEventByID(id int64) error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if _, exists := s.Events[id]; !exists {
		return ErrEventNotFound
	}
	delete(s.Events, id)
	return nil
}

func (s *Storage) FindEventsByDay(day time.Time) ([]models.Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var filteredEvents []models.Event
	for _, event := range s.Events {
		start, _ := time.Parse("1900-01-01", event.Start.String())
		if start.Year() == day.Year() && start.Month() == day.Month() && start.Day() == day.Day() {
			filteredEvents = append(filteredEvents, event)
		}
	}
	return filteredEvents, nil
}

func (s *Storage) FindEventsByWeek(day time.Time) ([]models.Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var filteredEvents []models.Event
	year, week := day.ISOWeek()
	for _, event := range s.Events {
		start, _ := time.Parse("1900-01-01", event.Start.String())
		eYear, eWeek := start.ISOWeek()
		if eYear == year && eWeek == week {
			filteredEvents = append(filteredEvents, event)
		}
	}
	return filteredEvents, nil
}

func (s *Storage) FindEventsByMonth(day time.Time) ([]models.Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var filteredEvents []models.Event
	for _, event := range s.Events {
		start, _ := time.Parse("1900-01-01", event.Start.String())
		if start.Year() == day.Year() && start.Month() == day.Month() {
			filteredEvents = append(filteredEvents, event)
		}
	}
	return filteredEvents, nil
}

func toEventDto(dto models.NewEvent) models.Event {
	event := models.Event{
		DaysAmountTillNotify: dto.DaysAmountTillNotify,
		Start:                dto.Start,
		End:                  dto.End,
		Description:          dto.Description,
		Title:                dto.Title,
		OwnerID:              dto.OwnerID,
	}
	return event
}
