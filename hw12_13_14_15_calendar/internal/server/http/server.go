package internalhttp

import (
	"context"
	"errors"
	"net/http"
	"sync"
	"time"

	"github.com/RollerM0bster/hw-esheludenko/hw12_13_14_15_calendar/models"
	"github.com/go-openapi/runtime/middleware"

	"github.com/RollerM0bster/hw-esheludenko/hw12_13_14_15_calendar/restapi"
	operations "github.com/RollerM0bster/hw-esheludenko/hw12_13_14_15_calendar/restapi/operations"
	"github.com/go-openapi/loads"

	"github.com/RollerM0bster/hw-esheludenko/hw12_13_14_15_calendar/internal/config"

	"github.com/RollerM0bster/hw-esheludenko/hw12_13_14_15_calendar/internal/app"

	"github.com/RollerM0bster/hw-esheludenko/hw12_13_14_15_calendar/internal/logger"
)

type Server struct {
	logger  *logger.Logger
	server  *http.Server
	app     *app.App
	wg      sync.WaitGroup
	stopped bool
	mu      sync.Mutex
}

func NewServer(logger *logger.Logger, app *app.App) *Server {
	return &Server{
		logger: logger,
		app:    app,
	}
}

func (s *Server) Start(ctx context.Context, cfg config.Config) error {
	s.logger.Info("Starting server")
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return err
	}
	api := operations.NewCalendarAPIAPI(swaggerSpec)

	//хендлеры
	api.PostEventsHandler = operations.PostEventsHandlerFunc(s.CreateEventHandler)
	api.PutEventsIDHandler = operations.PutEventsIDHandlerFunc(s.UpdateEventHandler)
	api.DeleteEventsIDHandler = operations.DeleteEventsIDHandlerFunc(s.DeleteEventByIDHandler)
	api.GetEventsByMonthHandler = operations.GetEventsByMonthHandlerFunc(s.GetEventsByMonthHandler)
	api.GetEventsByDayHandler = operations.GetEventsByDayHandlerFunc(s.GetEventsByDayHandler)
	api.GetEventsByWeekHandler = operations.GetEventsByWeekHandlerFunc(s.GetEventsByWeekHandler)

	handler := api.Serve(nil)

	s.server = &http.Server{
		Addr:    cfg.ServerConfig.Host + ":" + cfg.ServerConfig.Port,
		Handler: handler,
	}
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		if err := s.server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			s.logger.Error("Server error: " + err.Error())
		}
	}()

	<-ctx.Done()
	return s.Stop(ctx)
}

func (s *Server) Stop(ctx context.Context) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.stopped {
		return nil
	}
	s.stopped = true
	s.logger.Info("Stopping server")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.server.Shutdown(shutdownCtx); err != nil {
		s.logger.Error("Server error: " + err.Error())
		return err
	}
	s.wg.Wait()
	s.logger.Info("Server stopped")
	return nil
}

func (s *Server) CreateEventHandler(params operations.PostEventsParams) middleware.Responder {
	event := models.NewEvent{
		Description:          params.Body.Description,
		Title:                params.Body.Title,
		Start:                params.Body.Start,
		End:                  params.Body.End,
		OwnerID:              params.Body.OwnerID,
		DaysAmountTillNotify: params.Body.DaysAmountTillNotify,
	}
	id, err := s.app.CreateEvent(event)
	if err != nil {
		return &operations.PostEventsInternalServerError{Payload: &models.Error{Message: err.Error()}}
	}
	return &operations.PostEventsCreated{Payload: &models.EventCreated{ID: id}}
}

func (s *Server) GetEventsByDayHandler(params operations.GetEventsByDayParams) middleware.Responder {
	date, _ := time.Parse("2006-01-02", params.Date.String())
	events, err := s.app.FindEventsByDay(date)
	if err != nil {
		return &operations.GetEventsByDayInternalServerError{Payload: &models.Error{Message: err.Error()}}
	}
	return &operations.GetEventsByDayOK{Payload: events}
}

func (s *Server) GetEventsByMonthHandler(params operations.GetEventsByMonthParams) middleware.Responder {
	date, _ := time.Parse("2006-01-02", params.Date.String())
	events, err := s.app.FindEventsByMonth(date)
	if err != nil {
		return &operations.GetEventsByMonthInternalServerError{Payload: &models.Error{Message: err.Error()}}
	}
	return &operations.GetEventsByMonthOK{Payload: events}
}

func (s *Server) GetEventsByWeekHandler(params operations.GetEventsByWeekParams) middleware.Responder {
	date, _ := time.Parse("2006-01-02", params.WeekStart.String())
	events, err := s.app.FindEventsByWeek(date)
	if err != nil {
		return &operations.GetEventsByWeekInternalServerError{Payload: &models.Error{Message: err.Error()}}
	}
	return &operations.GetEventsByWeekOK{Payload: events}
}

func (s *Server) DeleteEventByIDHandler(params operations.DeleteEventsIDParams) middleware.Responder {
	err := s.app.DeleteEvent(params.ID)
	if err != nil {
		return &operations.DeleteEventsIDInternalServerError{Payload: &models.Error{Message: err.Error()}}
	}
	return operations.NewDeleteEventsIDNoContent()
}

func (s *Server) UpdateEventHandler(params operations.PutEventsIDParams) middleware.Responder {
	upd := models.NewEvent{
		Description:          params.Body.Description,
		Title:                params.Body.Title,
		Start:                params.Body.Start,
		End:                  params.Body.End,
		OwnerID:              params.Body.OwnerID,
		DaysAmountTillNotify: params.Body.DaysAmountTillNotify,
	}
	err := s.app.ChangeEvent(params.ID, upd)
	if err != nil {
		return &operations.PutEventsIDInternalServerError{Payload: &models.Error{Message: err.Error()}}
	}
	return operations.NewPutEventsIDNoContent()
}
