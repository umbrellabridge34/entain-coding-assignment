package service

import (
	"golang.org/x/net/context"
)

type Sports interface {
	// ListSportEvents will return a collection of sportEvents.
	ListSportEvents(ctx context.Context, in *sports.ListSportEventsRequest) (*sports.ListSportEventsResponse, error)
}

// sportsService implements the Sports interface.
type sportsService struct {
	sportEventsRepo db.SportEventsRepo
}

// NewsportsService instantiates and returns a new sportsService.
func NewsportsService(sportEventsRepo db.SportEventsRepo) Sports {
	return &sportsService{sportEventsRepo}
}

func (s *sportsService) ListSportEvents(ctx context.Context, in *sports.ListSportEventsRequest) (*sports.ListSportEventsResponse, error) {
	sportEvents, err := s.sportEventsRepo.List(in.Filter)
	if err != nil {
		return nil, err
	}

	return &sports.ListSportEventsResponse{sportEvents: sportEvents}, nil
}
