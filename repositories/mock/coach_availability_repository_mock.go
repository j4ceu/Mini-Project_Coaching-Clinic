package mock

import (
	"Mini-Project_Coaching-Clinic/models"

	"github.com/stretchr/testify/mock"
)

type MockCoachAvailabilityRepository struct {
	mock.Mock
}

func (_m *MockCoachAvailabilityRepository) FindByID(id string) (models.CoachAvailability, error) {
	args := _m.Called(id)
	return args.Get(0).(models.CoachAvailability), args.Error(1)
}

func (_m *MockCoachAvailabilityRepository) Create(coachAvailability models.CoachAvailability) (models.CoachAvailability, error) {
	args := _m.Called(coachAvailability)
	return args.Get(0).(models.CoachAvailability), args.Error(1)
}

func (_m *MockCoachAvailabilityRepository) Update(coachAvailability models.CoachAvailability, id string) (models.CoachAvailability, error) {
	args := _m.Called(coachAvailability, id)
	return args.Get(0).(models.CoachAvailability), args.Error(1)
}

func (_m *MockCoachAvailabilityRepository) Delete(id string) (models.CoachAvailability, error) {
	args := _m.Called(id)
	return args.Get(0).(models.CoachAvailability), args.Error(1)
}
