package mock

import (
	"Mini-Project_Coaching-Clinic/models"

	"github.com/stretchr/testify/mock"
)

type MockGameRepository struct {
	mock.Mock
}

func (_m *MockGameRepository) FindByID(id uint) (models.Game, error) {
	args := _m.Called(id)
	return args.Get(0).(models.Game), args.Error(1)
}

func (_m *MockGameRepository) Create(game models.Game) (models.Game, error) {
	args := _m.Called(game)
	return args.Get(0).(models.Game), args.Error(1)
}

func (_m *MockGameRepository) FindAll() ([]models.Game, error) {
	args := _m.Called()
	return args.Get(0).([]models.Game), args.Error(1)
}

func (_m *MockGameRepository) Update(game models.Game, id uint) (models.Game, error) {
	args := _m.Called(game, id)
	return args.Get(0).(models.Game), args.Error(1)
}

func (_m *MockGameRepository) Delete(id uint) (models.Game, error) {
	args := _m.Called(id)
	return args.Get(0).(models.Game), args.Error(1)
}
