package mock

import (
	"Mini-Project_Coaching-Clinic/models"

	"github.com/stretchr/testify/mock"
)

type MockGameService struct {
	mock.Mock
}

func (_m *MockGameService) Create(game models.Game) (models.Game, error) {
	args := _m.Called(game)
	return args.Get(0).(models.Game), args.Error(1)
}

func (_m *MockGameService) Delete(id uint) (models.Game, error) {
	args := _m.Called(id)
	return args.Get(0).(models.Game), args.Error(1)
}

func (_m *MockGameService) FindAll() ([]models.Game, error) {
	args := _m.Called()
	return args.Get(0).([]models.Game), args.Error(1)
}

func (_m *MockGameService) FindByID(id uint) (models.Game, error) {
	args := _m.Called(id)
	return args.Get(0).(models.Game), args.Error(1)
}

func (_m *MockGameService) Update(game models.Game, id uint) (models.Game, error) {
	args := _m.Called(game, id)
	return args.Get(0).(models.Game), args.Error(1)
}

