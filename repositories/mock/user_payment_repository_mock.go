package mock

import (
	"Mini-Project_Coaching-Clinic/models"

	"github.com/stretchr/testify/mock"
)

type MockUserPaymentRepository struct {
	mock.Mock
}

func (_m *MockUserPaymentRepository) FindByID(id string) (models.UserPayment, error) {
	args := _m.Called(id)
	return args.Get(0).(models.UserPayment), args.Error(1)
}

func (_m *MockUserPaymentRepository) Create(userPayment models.UserPayment) (models.UserPayment, error) {
	args := _m.Called(userPayment)
	return args.Get(0).(models.UserPayment), args.Error(1)
}

func (_m *MockUserPaymentRepository) Update(userPayment models.UserPayment, id string) (models.UserPayment, error) {
	args := _m.Called(userPayment, id)
	return args.Get(0).(models.UserPayment), args.Error(1)
}

func (_m *MockUserPaymentRepository) Delete(id string) (models.UserPayment, error) {
	args := _m.Called(id)
	return args.Get(0).(models.UserPayment), args.Error(1)
}

func (_m *MockUserPaymentRepository) FindAll() ([]models.UserPayment, error) {
	args := _m.Called()
	return args.Get(0).([]models.UserPayment), args.Error(1)
}

func (_m *MockUserPaymentRepository) FindUserBookByUserID(id string) ([]models.UserBook, error) {
	args := _m.Called(id)
	return args.Get(0).([]models.UserBook), args.Error(1)
}

func (_m *MockUserPaymentRepository) FindByInvoiceNumber(invoiceNumber string) (models.UserPayment, error) {
	args := _m.Called(invoiceNumber)
	return args.Get(0).(models.UserPayment), args.Error(1)
}

func (_m *MockUserPaymentRepository) FindByPaidAndProofOfPaymentIsNotNull() ([]models.UserPayment, error) {
	args := _m.Called()
	return args.Get(0).([]models.UserPayment), args.Error(1)
}

func (_m *MockUserPaymentRepository) FindByIDWithAllRelationUserBook(id string) (models.UserPayment, error) {
	args := _m.Called(id)
	return args.Get(0).(models.UserPayment), args.Error(1)
}
