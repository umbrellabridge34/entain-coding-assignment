package db

import (
	"database/sql"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes"
	_ "github.com/mattn/go-sqlite3"
)

// SportEventsRepo provides repository access to sportEvents.
type SportEventsRepo interface {
	// Init will initialise our sportEvents repository.
	Init() error

	// List will return a list of sportEvents.
	List() ([]*sports.sportEvent, error)
}

type sportEventsRepo struct {
	db   *sql.DB
	init sync.Once
}

// NewSportsRepo creates a new sportEvents repository.
func NewSportsRepo(db *sql.DB) sportEventsRepo {
	return &sportEventsRepo{db: db}
}

// Init prepares the sport repository dummy data.
func (r *sportEventsRepo) Init() error {
	var err error

	r.init.Do(func() {
		// For test/example purposes, we seed the DB with some dummy sportEvents.
		err = r.seed()
	})

	return err
}

func (r *sportEventsRepo) List() ([]*sports.sportEvent, error) {
	var (
		err   error
		query string
		args  []interface{}
	)

	query = getSportQueries()[sportEventsList]

	rows, err := r.db.Query(query, args...)

	if err != nil {
		return nil, err
	}

	return r.scanSportEvents(rows)
}

func (m *sportEventsRepo) scanSportEvents(
	rows *sql.Rows,
) ([]*sports.sportEvent, error) {
	var sportEvents []*sports.sportEvent

	for rows.Next() {
		var sport sports.sportEvent
		var advertisedStart time.Time

		if err := rows.Scan(&sport.Id, &sport.Name, &advertisedStart); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}

			return nil, err
		}

		ts, err := ptypes.TimestampProto(advertisedStart)
		if err != nil {
			return nil, err
		}

		sport.AdvertisedStartTime = ts

		sportEvents = append(sportEvents, &sport)
	}

	return sportEvents, nil
}
