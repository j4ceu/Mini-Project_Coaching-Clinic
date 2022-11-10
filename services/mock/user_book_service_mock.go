package mock

import (
	"Mini-Project_Coaching-Clinic/dto"
	"Mini-Project_Coaching-Clinic/dto/payload"
	"Mini-Project_Coaching-Clinic/models"

	"github.com/stretchr/testify/mock"
)

type UserBookServiceMock struct {
	mock.Mock
}

func (m *UserBookServiceMock) FindAll() ([]dto.UserBookResponse, error) {
	args := m.Called()
	return args.Get(0).([]dto.UserBookResponse), args.Error(1)
}

func (m *UserBookServiceMock) FindByUserID(id string) ([]dto.UserBookResponse, error) {
	args := m.Called(id)
	return args.Get(0).([]dto.UserBookResponse), args.Error(1)
}

func (m *UserBookServiceMock) FindByID(id string) (dto.UserBookResponse, error) {
	args := m.Called(id)
	return args.Get(0).(dto.UserBookResponse), args.Error(1)
}

func (m *UserBookServiceMock) Create(userBook payload.UserBookPayloadCreate) (dto.UserBookResponse, error) {
	args := m.Called(userBook)
	return args.Get(0).(dto.UserBookResponse), args.Error(1)
}

func (m *UserBookServiceMock) Update(userBook payload.UserBookPayloadUpdate, id string) (dto.UserBookResponse, error) {
	args := m.Called(userBook, id)
	return args.Get(0).(dto.UserBookResponse), args.Error(1)
}

func (m *UserBookServiceMock) Delete(id string) (models.UserBook, error) {
	args := m.Called(id)
	return args.Get(0).(models.UserBook), args.Error(1)
}
