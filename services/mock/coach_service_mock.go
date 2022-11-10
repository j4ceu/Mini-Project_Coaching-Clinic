package mock

import (
	"Mini-Project_Coaching-Clinic/dto"
	"Mini-Project_Coaching-Clinic/models"

	"github.com/stretchr/testify/mock"
)

type MockCoachService struct {
	mock.Mock
}

func (_m *MockCoachService) Create(coach models.Coach) (dto.CoachResponse, error) {
	args := _m.Called(coach)
	return args.Get(0).(dto.CoachResponse), args.Error(1)
}

func (_m *MockCoachService) Delete(id string) (models.Coach, error) {
	args := _m.Called(id)
	return args.Get(0).(models.Coach), args.Error(1)
}

func (_m *MockCoachService) FindByCode(code string) (dto.CoachResponse, error) {
	args := _m.Called(code)
	return args.Get(0).(dto.CoachResponse), args.Error(1)
}

func (_m *MockCoachService) FindByID(id string) (dto.CoachResponse, error) {
	args := _m.Called(id)
	return args.Get(0).(dto.CoachResponse), args.Error(1)
}

func (_m *MockCoachService) FindByGameID(gameID string) ([]dto.CoachResponse, error) {
	args := _m.Called(gameID)
	return args.Get(0).([]dto.CoachResponse), args.Error(1)
}

func (_m *MockCoachService) Update(coach models.Coach, id string) (dto.CoachResponse, error) {
	args := _m.Called(coach, id)
	return args.Get(0).(dto.CoachResponse), args.Error(1)
}
