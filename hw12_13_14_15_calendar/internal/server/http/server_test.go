package internalhttp

import (
	"github.com/RollerM0bster/hw-esheludenko/hw12_13_14_15_calendar/models"
	"github.com/RollerM0bster/hw-esheludenko/hw12_13_14_15_calendar/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type MockApp struct {
	mock.Mock
}

func (m *MockApp) FindEventsByWeek(weekStart time.Time) ([]*models.Event, error) {
	args := m.Called(weekStart)
	return args.Get(0).([]*models.Event), args.Error(1)
}

func (m *MockApp) CreateEvent(event models.NewEvent) (int64, error) {
	args := m.Called(event)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockApp) DeleteEvent(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockApp) ChangeEvent(id int64, event models.NewEvent) error {
	args := m.Called(id, event)
	return args.Error(0)
}

func (m *MockApp) FindEventsByDay(day time.Time) ([]*models.Event, error) {
	args := m.Called(day)
	return args.Get(0).([]*models.Event), args.Error(1)
}

func (m *MockApp) FindEventsByMonth(day time.Time) ([]*models.Event, error) {
	args := m.Called(day)
	return args.Get(0).([]*models.Event), args.Error(1)
}

func TestGetEventsByWeekHandler_Unit(t *testing.T) {
	mockApp := new(MockApp)

	mockApp.On("FindEventsByWeek", time.Date(2024, 1, 27, 0, 0, 0, 0, time.UTC)).Return(
		[]*models.Event{
			{ID: 1, Title: "Event 1"},
			{ID: 2, Title: "Event 2"},
		}, nil,
	)

	handler := operations.GetEventsByWeekHandlerFunc(
		func(params operations.GetEventsByWeekParams) middleware.Responder {
			weekStart, err := time.Parse("2006-01-02", params.WeekStart.String())
			if err != nil {
				return &operations.GetEventsByWeekInternalServerError{Payload: &models.Error{Message: "Invalid date"}}
			}

			events, err := mockApp.FindEventsByWeek(weekStart)
			if err != nil {
				return &operations.GetEventsByWeekInternalServerError{Payload: &models.Error{Message: err.Error()}}
			}

			return &operations.GetEventsByWeekOK{Payload: events}
		},
	)

	params := operations.GetEventsByWeekParams{
		WeekStart: strfmt.Date(time.Date(2024, 1, 27, 0, 0, 0, 0, time.UTC)),
	}

	res := handler.Handle(params)

	okRes, ok := res.(*operations.GetEventsByWeekOK)
	assert.True(t, ok)
	assert.Equal(t, 2, len(okRes.Payload))
	assert.Equal(t, "Event 1", okRes.Payload[0].Title)

	// Verify the mock call
	mockApp.AssertCalled(t, "FindEventsByWeek", time.Date(2024, 1, 27, 0, 0, 0, 0, time.UTC))
}
