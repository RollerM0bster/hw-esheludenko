package storage_test

import (
	"errors"
	"reflect"
	"testing"
	"time"

	memorystorage "github.com/RollerM0bster/hw-esheludenko/hw12_13_14_15_calendar/internal/storage/memory"
	"github.com/RollerM0bster/hw-esheludenko/hw12_13_14_15_calendar/models"
	"github.com/go-openapi/strfmt"
)

func TestStorage_CreateEvent(t *testing.T) {
	s := memorystorage.New()
	event := models.NewEvent{
		Start: strfmt.Date(time.Now()),
		End:   strfmt.Date(time.Now().Add(1 * time.Hour)),
		Title: "Test Event",
	}

	id, err := s.CreateEvent(event)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if id != 1 {
		t.Errorf("expected id 1, got %v", id)
	}

	createdEvent, exists := s.Events[id]
	if !exists {
		t.Fatalf("expected event to exist, but it does not")
	}

	if !reflect.DeepEqual(event.Title, createdEvent.Title) {
		t.Errorf("expected event title %v, got %v", event.Title, createdEvent.Title)
	}
}

func TestStorage_ChangeEvent(t *testing.T) {
	s := memorystorage.New()
	event := models.NewEvent{
		Start: strfmt.Date(time.Now()),
		End:   strfmt.Date(time.Now().Add(1 * time.Hour)),
		Title: "Original Event",
	}
	id, _ := s.CreateEvent(event)

	updatedEvent := models.NewEvent{
		Start: strfmt.Date(time.Now().Add(2 * time.Hour)),
		End:   strfmt.Date(time.Now().Add(3 * time.Hour)),
		Title: "Updated Event",
	}
	err := s.ChangeEvent(id, updatedEvent)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	changedEvent, exists := s.Events[id]
	if !exists {
		t.Fatalf("expected event to exist, but it does not")
	}

	if !reflect.DeepEqual(updatedEvent.Title, changedEvent.Title) {
		t.Errorf("expected event title %v, got %v", updatedEvent.Title, changedEvent.Title)
	}
}

func TestStorage_DeleteEventById(t *testing.T) {
	s := memorystorage.New()
	event := models.NewEvent{
		Start: strfmt.Date(time.Now()),
		End:   strfmt.Date(time.Now().Add(1 * time.Hour)),
		Title: "Event to Delete",
	}
	id, _ := s.CreateEvent(event)

	err := s.DeleteEventByID(id)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if _, exists := s.Events[id]; exists {
		t.Fatalf("expected event to be deleted, but it still exists")
	}

	err = s.DeleteEventByID(id)
	if !errors.Is(err, memorystorage.ErrEventNotFound) {
		t.Errorf("expected 'event not found' error, got %v", err)
	}
}

func TestStorage_FindEventsByDay(t *testing.T) {
	s := memorystorage.New()
	day := time.Now()
	event1 := models.NewEvent{
		Start: strfmt.Date(day),
		End:   strfmt.Date(day.Add(1 * time.Hour)),
		Title: "Event 1",
	}
	event2 := models.NewEvent{
		Start: strfmt.Date(day),
		End:   strfmt.Date(day.Add(2 * time.Hour)),
		Title: "Event 2",
	}
	event3 := models.NewEvent{
		Start: strfmt.Date(day.AddDate(0, 0, 1)),
		End:   strfmt.Date(day.AddDate(0, 0, 1).Add(1 * time.Hour)),
		Title: "Event 3",
	}

	s.CreateEvent(event1)
	s.CreateEvent(event2)
	s.CreateEvent(event3)

	events, err := s.FindEventsByDay(day)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(events) != 2 {
		t.Errorf("expected 2 events, got %d", len(events))
	}
}

func TestStorage_FindEventsByWeek(t *testing.T) {
	s := memorystorage.New()
	week := time.Now()
	event1 := models.NewEvent{
		Start: strfmt.Date(week),
		End:   strfmt.Date(week.Add(1 * time.Hour)),
		Title: "Event 1",
	}
	event2 := models.NewEvent{
		Start: strfmt.Date(week.AddDate(0, 0, 3)),
		End:   strfmt.Date(week.AddDate(0, 0, 3).Add(2 * time.Hour)),
		Title: "Event 2",
	}
	event3 := models.NewEvent{
		Start: strfmt.Date(week.AddDate(0, 0, 10)),
		End:   strfmt.Date(week.AddDate(0, 0, 10).Add(1 * time.Hour)),
		Title: "Event 3",
	}

	s.CreateEvent(event1)
	s.CreateEvent(event2)
	s.CreateEvent(event3)

	events, err := s.FindEventsByWeek(week)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(events) != 2 {
		t.Errorf("expected 2 events, got %d", len(events))
	}
}

func TestStorage_FindEventsByMonth(t *testing.T) {
	s := memorystorage.New()
	month := time.Now()
	event1 := models.NewEvent{
		Start: strfmt.Date(month),
		End:   strfmt.Date(month.Add(1 * time.Hour)),
		Title: "Event 1",
	}
	event2 := models.NewEvent{
		Start: strfmt.Date(month.AddDate(0, 0, 5)),
		End:   strfmt.Date(month.AddDate(0, 0, 5).Add(2 * time.Hour)),
		Title: "Event 2",
	}
	event3 := models.NewEvent{
		Start: strfmt.Date(month.AddDate(0, 1, 0)),
		End:   strfmt.Date(month.AddDate(0, 1, 0).Add(1 * time.Hour)),
		Title: "Event 3",
	}

	s.CreateEvent(event1)
	s.CreateEvent(event2)
	s.CreateEvent(event3)

	events, err := s.FindEventsByMonth(month)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(events) != 2 {
		t.Errorf("expected 2 events, got %d", len(events))
	}
}
