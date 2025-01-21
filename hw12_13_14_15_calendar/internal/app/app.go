package app

import (
	"context"
	"github.com/RollerM0bster/hw-esheludenko/hw12_13_14_15_calendar/internal/dto"

	"github.com/RollerM0bster/hw-esheludenko/hw12_13_14_15_calendar/internal/logger"
	"github.com/RollerM0bster/hw-esheludenko/hw12_13_14_15_calendar/internal/storage"
)

type App struct {
	logger  *logger.Logger
	storage storage.Storage
}

func New(logger *logger.Logger, storage storage.Storage) *App {
	return &App{
		logger:  logger,
		storage: storage,
	}
}

func (a *App) CreateEvent(ctx context.Context, title string) (int64, error) {
	a.logger.Info("Creating event " + title)
	event := dto.Event{Title: title}
	return a.storage.CreateEvent(event)
}
