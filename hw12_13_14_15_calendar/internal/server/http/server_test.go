package internalhttp

import (
	"github.com/RollerM0bster/hw-esheludenko/hw12_13_14_15_calendar/models"
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

func TestGetEventsByWeekHandler(t *testing.T) {
	//mockApp := new(MockApp)
	//mockApp.On("FindEventsByWeek", mock.AnythingOfType("time.Time")).Return([]*models.Event{
	//	{ID: 1, Title: "Event 1"},
	//	{ID: 2, Title: "Event 2"},
	//}, nil)
	//server := &Server{app: mockApp}
	//req, err := http.NewRequest("GET", "/events-by-week?weekStart=2024-10-25", nil)
	//rec := httptest.NewRecorder()
	//
	//handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	params := operations.GetEventsByWeekParams{
	//		WeekStart: strfmt.Date(time.Date(2024-10-25, 0, 0, 0, 0, 0, 0, time.UTC)),
	//	}
	//	res := server.GetEventsByWeekHandler(params)
	//	if responder, ok := res.(middleware.ResponderFunc); ok {
	//		responder(w, r)
	//	}
	//})
	//handler.ServeHTTP(rec, req)
}
