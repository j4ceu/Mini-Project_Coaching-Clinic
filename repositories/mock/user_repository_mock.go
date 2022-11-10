package mock

import (
	"Mini-Project_Coaching-Clinic/models"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (_m *MockUserRepository) FindByID(id string) (models.User, error) {
	args := _m.Called(id)
	return args.Get(0).(models.User), args.Error(1)
}

func (_m *MockUserRepository) Create(user models.User) (models.User, error) {
	args := _m.Called(user)
	return args.Get(0).(models.User), args.Error(1)
}

func (_m *MockUserRepository) FindAll() ([]models.User, error) {
	args := _m.Called()
	return args.Get(0).([]models.User), args.Error(1)
}

func (_m *MockUserRepository) Update(user models.User, id string) (models.User, error) {
	args := _m.Called(user, id)
	return args.Get(0).(models.User), args.Error(1)
}

func (_m *MockUserRepository) Delete(id string) (models.User, error) {
	args := _m.Called(id)
	return args.Get(0).(models.User), args.Error(1)
}

func (_m *MockUserRepository) FindByEmail(email string) (models.User, error) {
	args := _m.Called(email)
	return args.Get(0).(models.User), args.Error(1)
}
