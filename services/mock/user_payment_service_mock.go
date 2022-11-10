package mock

import (
	"Mini-Project_Coaching-Clinic/dto"
	"Mini-Project_Coaching-Clinic/dto/payload"
	"Mini-Project_Coaching-Clinic/models"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/mock"
)

type UserPaymentServiceMock struct {
	mock.Mock
}

func (m *UserPaymentServiceMock) FindAll() ([]dto.UserPaymentResponse, error) {
	args := m.Called()
	return args.Get(0).([]dto.UserPaymentResponse), args.Error(1)
}

func (m *UserPaymentServiceMock) FindByPaidAndProofOfPaymentIsNotNull() ([]dto.UserPaymentResponse, error) {
	args := m.Called()
	return args.Get(0).([]dto.UserPaymentResponse), args.Error(1)
}

func (m *UserPaymentServiceMock) FindByID(id string) (dto.UserPaymentResponse, error) {
	args := m.Called(id)
	return args.Get(0).(dto.UserPaymentResponse), args.Error(1)
}

func (m *UserPaymentServiceMock) FindByInvoiceNumber(invoiceNumber string) (dto.UserPaymentResponse, error) {
	args := m.Called(invoiceNumber)
	return args.Get(0).(dto.UserPaymentResponse), args.Error(1)
}

func (m *UserPaymentServiceMock) Create(userPayment payload.UserPaymentPayload) (dto.UserPaymentResponse, error) {
	args := m.Called(userPayment)
	return args.Get(0).(dto.UserPaymentResponse), args.Error(1)
}

func (m *UserPaymentServiceMock) Update(userPayment payload.UserPaymentPayload, id string, claims jwt.MapClaims) (dto.UserPaymentResponse, error) {
	args := m.Called(userPayment, id, claims)
	return args.Get(0).(dto.UserPaymentResponse), args.Error(1)
}

func (m *UserPaymentServiceMock) Delete(id string) (models.UserPayment, error) {
	args := m.Called(id)
	return args.Get(0).(models.UserPayment), args.Error(1)
}
