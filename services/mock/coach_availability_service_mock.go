package mock

import (
	"Mini-Project_Coaching-Clinic/dto"
	"Mini-Project_Coaching-Clinic/models"

	"github.com/stretchr/testify/mock"
)

type MockCoachAvailabilityService struct {
	mock.Mock
}

func (m *MockCoachAvailabilityService) FindByID(id string) (dto.CoachAvailabilityResponse, error) {
	args := m.Called(id)
	return args.Get(0).(dto.CoachAvailabilityResponse), args.Error(1)
}

func (m *MockCoachAvailabilityService) Create(coachAvailability models.CoachAvailability) (dto.CoachAvailabilityResponse, error) {
	args := m.Called(coachAvailability)
	return args.Get(0).(dto.CoachAvailabilityResponse), args.Error(1)
}

func (m *MockCoachAvailabilityService) Update(coachAvailability models.CoachAvailability, id string) (dto.CoachAvailabilityResponse, error) {
	args := m.Called(coachAvailability, id)
	return args.Get(0).(dto.CoachAvailabilityResponse), args.Error(1)
}

func (m *MockCoachAvailabilityService) Delete(id string) (models.CoachAvailability, error) {
	args := m.Called(id)
	return args.Get(0).(models.CoachAvailability), args.Error(1)
}

func (m *MockCoachAvailabilityService) CheckInterval(coachAvailability models.CoachAvailability) (bool, error) {
	args := m.Called(coachAvailability)
	return args.Bool(0), args.Error(1)
}
