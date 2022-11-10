package mock

import (
	"Mini-Project_Coaching-Clinic/models"

	"github.com/stretchr/testify/mock"
)

type MockUserBookRepository struct {
	mock.Mock
}

func (_m *MockUserBookRepository) FindByID(id string) (models.UserBook, error) {
	args := _m.Called(id)
	return args.Get(0).(models.UserBook), args.Error(1)
}

func (_m *MockUserBookRepository) Create(userBook models.UserBook) (models.UserBook, error) {
	args := _m.Called(userBook)
	return args.Get(0).(models.UserBook), args.Error(1)
}

func (_m *MockUserBookRepository) Update(userBook models.UserBook, id string) (models.UserBook, error) {
	args := _m.Called(userBook, id)
	return args.Get(0).(models.UserBook), args.Error(1)
}

func (_m *MockUserBookRepository) Delete(id string) (models.UserBook, error) {
	args := _m.Called(id)
	return args.Get(0).(models.UserBook), args.Error(1)
}

func (_m *MockUserBookRepository) FindAll() ([]models.UserBook, error) {
	args := _m.Called()
	return args.Get(0).([]models.UserBook), args.Error(1)
}
