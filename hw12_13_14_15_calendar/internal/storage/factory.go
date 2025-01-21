package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/RollerM0bster/hw-esheludenko/hw12_13_14_15_calendar/internal/config"
	memorystorage "github.com/RollerM0bster/hw-esheludenko/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/RollerM0bster/hw-esheludenko/hw12_13_14_15_calendar/internal/storage/sql"
)

func NewStorage(ctx context.Context, cfg config.Config) (Storage, error) {
	switch cfg.StorageType {
	case "memory":
		return memorystorage.New(), nil
	case "sql":
		sqlStorage := sqlstorage.New()
		if cfg.Db.Login == "" || cfg.Db.Pass == "" || cfg.Db.Host == "" || cfg.Db.Port == "" || cfg.Db.Database == "" {
			return nil, errors.New("missing required database configuration parameters")
		}
		if err := sqlStorage.Connect(ctx, fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			cfg.Db.Login, cfg.Db.Pass, cfg.Db.Host, cfg.Db.Port, cfg.Db.Database)); err != nil {
			return nil, fmt.Errorf("failed to connect to the database: %w", err)
		}
		return sqlStorage, nil
	default:
		return nil, errors.New("invalid storage type")
	}
}
