package service

import (
	"git.neds.sh/matty/entain/sports/db"
	"git.neds.sh/matty/entain/sports/proto/sports"
	"golang.org/x/net/context"
)

type Sports interface {
	ListSportEvents(ctx context.Context, in *sports.ListSportEventsRequest) (*sports.ListSportEventsResponse, error)
}

type sportsService struct {
	sportEventsRepo db.SportEventsRepo
}

func NewSportsService(sportEventsRepo db.SportEventsRepo) Sports {
	return &sportsService{sportEventsRepo}
}

func (s *sportsService) ListSportEvents(ctx context.Context, in *sports.ListSportEventsRequest) (*sports.ListSportEventsResponse, error) {
	sportEvents, err := s.sportEventsRepo.List()
	if err != nil {
		return nil, err
	}

	return &sports.ListSportEventsResponse{SportEvents: sportEvents}, nil
}
