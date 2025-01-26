package app

import (
	"github.com/RollerM0bster/hw-esheludenko/hw12_13_14_15_calendar/models"

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

func (cs *App) CreateEvent(event models.NewEvent) (int64, error) {
	return cs.storage.CreateEvent(event)
}

func (cs *App) DeleteEvent(id int64) error {
	return cs.storage.DeleteEventByID(id)
}

func (cs *App) ChangeEvent(id int64, event models.NewEvent) error {
	return cs.storage.ChangeEvent(id, event)
}
