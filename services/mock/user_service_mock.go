package mock

import (
	"Mini-Project_Coaching-Clinic/dto"
	"Mini-Project_Coaching-Clinic/models"

	"github.com/stretchr/testify/mock"
)

// MockUserService is a mock of UserService interface
type MockUserService struct {
	mock.Mock
}

func (_m *MockUserService) Create(user models.User) (models.User, error) {
	args := _m.Called(user)
	return args.Get(0).(models.User), args.Error(1)
}

func (_m *MockUserService) Delete(id string) (models.User, error) {
	args := _m.Called(id)
	return args.Get(0).(models.User), args.Error(1)
}

func (_m *MockUserService) FindAll() ([]models.User, error) {
	args := _m.Called()
	return args.Get(0).([]models.User), args.Error(1)
}

func (_m *MockUserService) FindByID(id string) (models.User, error) {
	args := _m.Called(id)
	return args.Get(0).(models.User), args.Error(1)
}

func (_m *MockUserService) LoginUser(email string, password string) (dto.UserResponse, error) {
	args := _m.Called(email, password)
	return args.Get(0).(dto.UserResponse), args.Error(1)
}

func (_m *MockUserService) Update(user models.User, id string) (models.User, error) {
	args := _m.Called(user, id)
	return args.Get(0).(models.User), args.Error(1)
}
