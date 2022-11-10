package services

import (
	"Mini-Project_Coaching-Clinic/dto/payload"
	"Mini-Project_Coaching-Clinic/models"
	"Mini-Project_Coaching-Clinic/repositories/mock"
	"errors"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	mockTest "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type suiteUserPayment struct {
	suite.Suite
	mockUserPaymentRepository *mock.MockUserPaymentRepository
	userPaymentService        UserPaymentService
	mockUserRepo              *mock.MockUserRepository
}

func (s *suiteUserPayment) SetupTest() {
	s.mockUserPaymentRepository = new(mock.MockUserPaymentRepository)
	s.mockUserRepo = new(mock.MockUserRepository)
	s.userPaymentService = NewUserPaymentServices(s.mockUserPaymentRepository, s.mockUserRepo)
}

func (s *suiteUserPayment) TestFindAll() {

	testCase := []struct {
		Name           string
		ExpectedReturn []models.UserPayment
		ExpectedError  error
		MockReturn     []models.UserPayment
		MockError      error
	}{
		{
			Name: "Success",
			ExpectedReturn: []models.UserPayment{
				{
					ID: uuid.MustParse("c9e7f7c0-5c5a-4b4f-8f9c-0c9a9f9f9f9f"),
				},
			},
			ExpectedError: nil,
			MockReturn: []models.UserPayment{
				{
					ID: uuid.MustParse("c9e7f7c0-5c5a-4b4f-8f9c-0c9a9f9f9f9f"),
				},
			},
		},
		{
			Name:           "Error",
			ExpectedReturn: nil,
			ExpectedError:  errors.New("Error"),
			MockReturn:     nil,
			MockError:      errors.New("Error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mockUserPaymentRepository.On("FindAll").Return(tc.MockReturn, tc.MockError).Once()
			result, err := s.userPaymentService.FindAll()
			if tc.MockError == nil {
				s.Equal(tc.ExpectedReturn[0].ID.String(), result[0].ID)
			}
			s.Equal(tc.ExpectedError, err)
		})
	}
}

func (s *suiteUserPayment) TestFindByID() {

	testCase := []struct {
		Name           string
		ExpectedReturn models.UserPayment
		ExpectedError  error
		MockReturn     models.UserPayment
		MockError      error
	}{
		{
			Name: "Success",
			ExpectedReturn: models.UserPayment{
				ID: uuid.MustParse("c9e7f7c0-5c5a-4b4f-8f9c-0c9a9f9f9f9f"),
			},
			ExpectedError: nil,
			MockReturn: models.UserPayment{
				ID: uuid.MustParse("c9e7f7c0-5c5a-4b4f-8f9c-0c9a9f9f9f9f"),
			},
		},
		{
			Name:           "Error",
			ExpectedReturn: models.UserPayment{},
			ExpectedError:  errors.New("Error"),
			MockReturn:     models.UserPayment{},
			MockError:      errors.New("Error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mockUserPaymentRepository.On("FindByID", "c9e7f7c0-5c5a-4b4f-8f9c-0c9a9f9f9f9f").Return(tc.MockReturn, tc.MockError).Once()
			result, err := s.userPaymentService.FindByID("c9e7f7c0-5c5a-4b4f-8f9c-0c9a9f9f9f9f")
			if tc.MockError == nil {
				s.Equal(tc.ExpectedReturn.ID.String(), result.ID)
			}
			s.Equal(tc.ExpectedError, err)
		})
	}
}

func (s *suiteUserPayment) TestFindByInvoiceNumber() {

	testCase := []struct {
		Name           string
		ExpectedReturn models.UserPayment
		ExpectedError  error
		MockReturn     models.UserPayment
		MockError      error
	}{
		{
			Name: "Success",
			ExpectedReturn: models.UserPayment{
				ID: uuid.MustParse("c9e7f7c0-5c5a-4b4f-8f9c-0c9a9f9f9f9f"),
			},
			ExpectedError: nil,
			MockReturn: models.UserPayment{
				ID: uuid.MustParse("c9e7f7c0-5c5a-4b4f-8f9c-0c9a9f9f9f9f"),
			},
		},
		{
			Name:           "Error",
			ExpectedReturn: models.UserPayment{},
			ExpectedError:  errors.New("Error"),
			MockReturn:     models.UserPayment{},
			MockError:      errors.New("Error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mockUserPaymentRepository.On("FindByInvoiceNumber", "J4CEU").Return(tc.MockReturn, tc.MockError).Once()
			result, err := s.userPaymentService.FindByInvoiceNumber("J4CEU")
			if tc.MockError == nil {
				s.Equal(tc.ExpectedReturn.ID.String(), result.ID)
			}
			s.Equal(tc.ExpectedError, err)
		})
	}
}

func (s *suiteUserPayment) TestFindByPaidAndProofOfPaymentIsNotNull() {

	testCase := []struct {
		Name           string
		ExpectedReturn []models.UserPayment
		ExpectedError  error
		MockReturn     []models.UserPayment
		MockError      error
	}{
		{
			Name: "Success",
			ExpectedReturn: []models.UserPayment{
				{
					ID: uuid.MustParse("c9e7f7c0-5c5a-4b4f-8f9c-0c9a9f9f9f9f"),
				},
			},
			ExpectedError: nil,
			MockReturn: []models.UserPayment{
				{
					ID: uuid.MustParse("c9e7f7c0-5c5a-4b4f-8f9c-0c9a9f9f9f9f"),
				},
			},
		},
		{
			Name:           "Error",
			ExpectedReturn: nil,
			ExpectedError:  errors.New("Error"),
			MockReturn:     nil,
			MockError:      errors.New("Error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mockUserPaymentRepository.On("FindByPaidAndProofOfPaymentIsNotNull").Return(tc.MockReturn, tc.MockError).Once()
			result, err := s.userPaymentService.FindByPaidAndProofOfPaymentIsNotNull()
			if tc.MockError == nil {
				s.Equal(tc.ExpectedReturn[0].ID.String(), result[0].ID)
			}
			s.Equal(tc.ExpectedError, err)
		})
	}
}

func (s *suiteUserPayment) TestCreateUserPayment() {

	now := time.Now()
	testCase := []struct {
		Name           string
		ExpectedReturn models.UserPayment
		ExpectedError  error
		MockReturn     models.UserPayment
		MockError      error
	}{
		{
			Name: "Success",
			ExpectedReturn: models.UserPayment{

				UserID:    "1",
				ExpiredAt: now.Local().Add(time.Minute * 15).Format("2006-01-02 15:04:05"),
			},
			ExpectedError: nil,
			MockReturn: models.UserPayment{

				UserID:    "1",
				ExpiredAt: now.Local().Add(time.Minute * 15).Format("2006-01-02 15:04:05"),
			},
			MockError: nil,
		},
		{
			Name: "Error",
			ExpectedReturn: models.UserPayment{

				UserID:    "1",
				ExpiredAt: now.Local().Add(time.Minute * 15).Format("2006-01-02 15:04:05"),
			},
			ExpectedError: errors.New("Error"),
			MockReturn: models.UserPayment{

				UserID:    "1",
				ExpiredAt: now.Local().Add(time.Minute * 15).Format("2006-01-02 15:04:05"),
			},
			MockError: errors.New("Error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mockUserPaymentRepository.On("Create", tc.MockReturn).Return(tc.MockReturn, tc.MockError).Once()
			s.mockUserRepo.On("FindByID", "1").Return(models.User{}, nil).Once()
			result, err := s.userPaymentService.Create(payload.UserPaymentPayload{
				UserID: "1",
			})
			if tc.MockError == nil {
				s.Equal(tc.ExpectedReturn.ID.String(), result.ID)
			}
			s.Equal(tc.ExpectedError, err)
		})
	}

}

func (s *suiteUserPayment) TestDeleteUserPayment() {

	testCase := []struct {
		Name           string
		ExpectedReturn models.UserPayment
		ExpectedError  error
		MockReturn     models.UserPayment
		MockError      error
	}{
		{
			Name: "Success",
			ExpectedReturn: models.UserPayment{
				ID: uuid.MustParse("c9e7f7c0-5c5a-4b4f-8f9c-0c9a9f9f9f9f"),
			},
			ExpectedError: nil,
			MockReturn: models.UserPayment{
				ID: uuid.MustParse("c9e7f7c0-5c5a-4b4f-8f9c-0c9a9f9f9f9f"),
			},
			MockError: nil,
		},
		{
			Name:           "Error",
			ExpectedReturn: models.UserPayment{},
			ExpectedError:  errors.New("Error"),
			MockReturn:     models.UserPayment{},
			MockError:      errors.New("Error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mockUserPaymentRepository.On("Delete", tc.MockReturn.ID.String()).Return(tc.MockReturn, tc.MockError).Once()
			result, err := s.userPaymentService.Delete(tc.ExpectedReturn.ID.String())
			if tc.MockError == nil {
				s.Equal(tc.ExpectedReturn.ID, result.ID)
			}
			s.Equal(tc.ExpectedError, err)
		})
	}
}

func (s *suiteUserPayment) TestUpdateUserPayment() {

	testCase := []struct {
		Name           string
		ExpectedReturn models.UserPayment
		ExpectedError  error
		MockReturn     models.UserPayment
		MockError      error
	}{
		{
			Name: "Success",
			ExpectedReturn: models.UserPayment{
				ID:     uuid.MustParse("c9e7f7c0-5c5a-4b4f-8f9c-0c9a9f9f9f9f"),
				UserID: "1",
			},
			ExpectedError: nil,
			MockReturn: models.UserPayment{
				ID:     uuid.MustParse("c9e7f7c0-5c5a-4b4f-8f9c-0c9a9f9f9f9f"),
				UserID: "1",
			},
			MockError: nil,
		},
		{
			Name:           "Error Unauthorized",
			ExpectedReturn: models.UserPayment{},
			ExpectedError:  errors.New("unauthorized"),
			MockReturn:     models.UserPayment{},
			MockError:      errors.New("unauthorized"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mockUserPaymentRepository.On("Update", mockTest.Anything, mockTest.Anything).Return(tc.MockReturn, tc.MockError).Once()
			s.mockUserPaymentRepository.On("FindByID", mockTest.Anything).Return(tc.ExpectedReturn, nil).Once()
			s.mockUserRepo.On("FindByID", mockTest.Anything).Return(models.User{}, nil).Once()
			claims := jwt.MapClaims{
				"role":    "User",
				"user_id": "1",
			}
			result, err := s.userPaymentService.Update(payload.UserPaymentPayload{}, "1", claims)
			if tc.MockError == nil {
				s.Equal(tc.ExpectedReturn.ID.String(), result.ID)
			}
			s.Equal(tc.ExpectedError, err)
		})
	}
}

func TestUserPaymentService(t *testing.T) {
	suite.Run(t, new(suiteUserPayment))
}
