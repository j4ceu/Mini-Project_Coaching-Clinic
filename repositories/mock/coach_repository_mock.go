package mock

import (
	"Mini-Project_Coaching-Clinic/models"

	"github.com/stretchr/testify/mock"
)

type MockCoachRepository struct {
	mock.Mock
}

func (_m *MockCoachRepository) FindByID(id string) (models.Coach, error) {
	args := _m.Called(id)
	return args.Get(0).(models.Coach), args.Error(1)
}

func (_m *MockCoachRepository) Create(coach models.Coach) (models.Coach, error) {
	args := _m.Called(coach)
	return args.Get(0).(models.Coach), args.Error(1)
}

func (_m *MockCoachRepository) FindByCode(code string) (models.Coach, error) {
	args := _m.Called(code)
	return args.Get(0).(models.Coach), args.Error(1)
}

func (_m *MockCoachRepository) Update(coach models.Coach, id string) (models.Coach, error) {
	args := _m.Called(coach, id)
	return args.Get(0).(models.Coach), args.Error(1)
}

func (_m *MockCoachRepository) Delete(id string) (models.Coach, error) {
	args := _m.Called(id)
	return args.Get(0).(models.Coach), args.Error(1)
}

func (_m *MockCoachRepository) FindByGameID(gameID string) ([]models.Coach, error) {
	args := _m.Called(gameID)
	return args.Get(0).([]models.Coach), args.Error(1)
}
