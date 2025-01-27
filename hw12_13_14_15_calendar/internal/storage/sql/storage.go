package sqlstorage

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/RollerM0bster/hw-esheludenko/hw12_13_14_15_calendar/models"
	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) Connect(ctx context.Context, dsn string) error {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	if err = db.PingContext(ctx); err != nil {
		return err
	}
	s.db = db
	return nil
}

func (s *Storage) Close(_ context.Context) error {
	if err := s.db.Close(); err != nil {
		return err
	}
	return s.db.Close()
}

func (s *Storage) CreateEvent(dto models.NewEvent) (int64, error) {
	query := `insert into event_storage (title, "start", "end", description, owner_id, days_before_notify) values ($1, $2, $3, $4, $5, $6) returning id` //nolint:lll
	var id int64
	err := s.db.QueryRow(query, dto.Title, dto.Start, dto.End, dto.Description, dto.OwnerID, dto.DaysAmountTillNotify).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *Storage) ChangeEvent(id int64, dto models.NewEvent) error {
	query := `update event_storage set title = $1, start = $2, "end" = $3, description = $4, owner_id = $5, days_before_notify = $6 where id = $7` //nolint:lll
	res, err := s.db.Exec(query, dto.Title, dto.Start, dto.End, dto.Description, dto.OwnerID, dto.DaysAmountTillNotify, id)                        //nolint:lll
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("event not found")
	}
	return nil
}

func (s *Storage) DeleteEventByID(id int64) error {
	query := `delete from event_storage where id = $1`
	res, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("event not found")
	}
	return nil
}

func (s *Storage) FindEventsByDay(day time.Time) ([]*models.Event, error) {
	start := time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, time.UTC)
	end := start.Add(24 * time.Hour)
	query := `select id, title, start, "end", description, owner_id, days_before_notify from event_storage where start >= $1 and start <= $2` //nolint:lll
	rows, err := s.db.Query(query, start, end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []*models.Event
	for rows.Next() {
		var event models.Event
		if err = rows.Scan(&event.ID, &event.Title, &event.Start, &event.End, &event.Description, &event.OwnerID, &event.DaysAmountTillNotify); err != nil { //nolint:lll
			return nil, err
		}
		events = append(events, &event)
	}
	return events, nil
}

func (s *Storage) FindEventsByWeek(day time.Time) ([]*models.Event, error) {
	year, week := day.ISOWeek()
	query := `select id, title, start, "end", description, owner_id, days_before_notify from event_storage where extract(year from start) = $1 and extract(week from start) = $2` //nolint:lll
	rows, err := s.db.Query(query, year, week)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []*models.Event
	for rows.Next() {
		var event models.Event
		if err = rows.Scan(&event.ID, &event.Title, &event.Start, &event.End, &event.Description, &event.OwnerID, &event.DaysAmountTillNotify); err != nil { //nolint:lll
			return nil, err
		}
		events = append(events, &event)
	}
	return events, nil
}

func (s *Storage) FindEventsByMonth(day time.Time) ([]*models.Event, error) {
	query := `select id, title, start, "end", description, owner_id, days_before_notify from event_storage where extract(year from start) = $1 and extract(month from start) = $2` //nolint:lll
	rows, err := s.db.Query(query, day.Year(), day.Month())
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []*models.Event
	for rows.Next() {
		var event models.Event
		if err = rows.Scan(&event.ID, &event.Title, &event.Start, &event.End, &event.Description, &event.OwnerID, &event.DaysAmountTillNotify); err != nil { //nolint:lll
			return nil, err
		}
		events = append(events, &event)
	}
	return events, nil
}
