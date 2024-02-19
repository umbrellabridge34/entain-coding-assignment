package service

import (
	"context"
	"testing"

	"git.neds.sh/matty/entain/sports/proto/sports"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockSportEventsRepo struct {
	mock.Mock
}

func (m *MockSportEventsRepo) Init() error {
	return nil
}

func (m *MockSportEventsRepo) List() ([]*sports.SportEvent, error) {
	args := m.Called()
	return args.Get(0).([]*sports.SportEvent), args.Error(1)
}

func TestListSportEvents(t *testing.T) {
	mockRepo := new(MockSportEventsRepo)
	testService := NewSportsService(mockRepo)

	mockEvents := []*sports.SportEvent{
		{Id: 1, Name: "Football"},
		{Id: 2, Name: "Basketball"},
	}
	mockRepo.On("List").Return(mockEvents, nil)

	response, err := testService.ListSportEvents(context.Background(), &sports.ListSportEventsRequest{})

	mockRepo.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, mockEvents, response.SportEvents)
}
