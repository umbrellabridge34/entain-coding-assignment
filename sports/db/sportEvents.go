package db

import (
	"database/sql"
	"sync"
	"time"

	"git.neds.sh/matty/entain/sports/proto/sports"
	"github.com/golang/protobuf/ptypes"
	_ "github.com/mattn/go-sqlite3"
)

type SportEventsRepo interface {
	Init() error

	List() ([]*sports.SportEvent, error)
}

type sportEventsRepo struct {
	db   *sql.DB
	init sync.Once
}

func NewSportsRepo(db *sql.DB) SportEventsRepo {
	return &sportEventsRepo{db: db}
}

func (r *sportEventsRepo) Init() error {
	var err error

	r.init.Do(func() {
		err = r.seed()
	})

	return err
}

func (r *sportEventsRepo) List() ([]*sports.SportEvent, error) {
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
) ([]*sports.SportEvent, error) {
	var sportEvents []*sports.SportEvent

	for rows.Next() {
		var sport sports.SportEvent
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
