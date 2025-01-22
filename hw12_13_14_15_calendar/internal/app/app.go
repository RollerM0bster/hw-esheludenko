package app

import (
	"context"

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

func (a *App) CreateEvent(_ context.Context, _ string) (int64, error) {
	return 0, nil
}
