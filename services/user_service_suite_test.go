package services

import (
	"Mini-Project_Coaching-Clinic/dto"
	"Mini-Project_Coaching-Clinic/models"
	"Mini-Project_Coaching-Clinic/repositories/mock"
	"errors"
	"testing"

	mockv2 "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type suiteUser struct {
	suite.Suite
	mockUserRepository *mock.MockUserRepository
	userService        UserService
}

func (s *suiteUser) SetupTest() {
	s.mockUserRepository = new(mock.MockUserRepository)
	s.userService = NewUserService(s.mockUserRepository)
}

func (s *suiteUser) TestFindAll() {

	testCase := []struct {
		Name           string
		ExpectedReturn []models.User
		ExpectedError  error
		MockReturn     []models.User
		MockError      error
	}{
		{
			Name: "Success",
			ExpectedReturn: []models.User{
				{

					FirstName: "User 1",
					Email:     "test@gmail.com",
					Password:  "123456",
				},
			},
			ExpectedError: nil,
			MockReturn: []models.User{
				{

					FirstName: "User 1",
					Email:     "test@gmail.com",
					Password:  "123456",
				},
			},
			MockError: nil,
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
			s.mockUserRepository.On("FindAll").Return(tc.MockReturn, tc.MockError).Once()
			result, err := s.userService.FindAll()
			s.Equal(tc.ExpectedReturn, result)
			s.Equal(tc.ExpectedError, err)
		})
	}

}

func (s *suiteUser) TestFindByID() {

	testCase := []struct {
		Name           string
		ExpectedReturn models.User
		ExpectedError  error
		MockReturn     models.User
		MockError      error
	}{
		{
			Name: "Success",
			ExpectedReturn: models.User{
				FirstName: "User 1",
				Email:     "test@gmail.com",
				Password:  "123456",
			},
			ExpectedError: nil,
			MockReturn: models.User{
				FirstName: "User 1",
				Email:     "test@gmail.com",
				Password:  "123456",
			},
			MockError: nil,
		},
		{
			Name:           "Error",
			ExpectedReturn: models.User{},
			ExpectedError:  errors.New("Error"),
			MockReturn:     models.User{},
			MockError:      errors.New("Error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mockUserRepository.On("FindByID", "1").Return(tc.MockReturn, tc.MockError).Once()
			result, err := s.userService.FindByID("1")
			s.Equal(tc.ExpectedReturn, result)
			s.Equal(tc.ExpectedError, err)
		})
	}
}

func (s *suiteUser) TestLoginUser() {

	testCase := []struct {
		Name           string
		ExpectedReturn dto.UserResponse
		ExpectedError  error
		MockReturn     dto.UserResponse
		MockError      error
		UserModel      models.User
	}{
		{
			Name: "Success",
			ExpectedReturn: dto.UserResponse{
				Email: "test@gmail.com",
			},
			ExpectedError: nil,
			MockReturn: dto.UserResponse{
				Email: "test@gmail.com",
			},
			MockError: nil,
			UserModel: models.User{Email: "test@gmail.com", Password: "123456"},
		},
		{
			Name:           "Error",
			ExpectedReturn: dto.UserResponse{},
			ExpectedError:  errors.New("Error"),
			MockReturn:     dto.UserResponse{},
			MockError:      errors.New("Error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			if tc.UserModel.Email != "" {
				_ = tc.UserModel.HashPassword(tc.UserModel.Password)
			}

			s.mockUserRepository.On("FindByEmail", "test@gmail.com").Return(tc.UserModel, tc.MockError).Once()
			result, err := s.userService.LoginUser("test@gmail.com", "123456")
			s.Equal(tc.ExpectedReturn.Email, result.Email)
			s.Equal(tc.ExpectedError, err)
		})
	}
}

func (s *suiteUser) TestDeleteUser() {

	testCase := []struct {
		Name           string
		ExpectedReturn models.User
		ExpectedError  error
		MockReturn     models.User
		MockError      error
	}{
		{
			Name:           "Success",
			ExpectedReturn: models.User{},
			ExpectedError:  nil,
			MockReturn:     models.User{},
			MockError:      nil,
		},
		{
			Name:           "Error",
			ExpectedReturn: models.User{},
			ExpectedError:  errors.New("Error"),
			MockReturn:     models.User{},
			MockError:      errors.New("Error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mockUserRepository.On("Delete", "1").Return(tc.MockReturn, tc.MockError).Once()
			result, err := s.userService.Delete("1")
			s.Equal(tc.ExpectedReturn, result)
			s.Equal(tc.ExpectedError, err)
		})
	}

}

func (s *suiteUser) TestCreateUser() {

	testCase := []struct {
		Name           string
		ExpectedReturn models.User
		ExpectedError  error
		MockReturn     models.User
		MockError      error
		UserModel      models.User
	}{
		{
			Name: "Success",
			ExpectedReturn: models.User{
				FirstName: "User 1",
				Email:     "test@gmail.com",
				Password:  "123456",
			},
			ExpectedError: nil,
			MockReturn: models.User{
				FirstName: "User 1",
				Email:     "test@gmail.com",
				Password:  "123456",
			},
			MockError: nil,
			UserModel: models.User{
				FirstName: "User 1",
				Email:     "test@gmail.com",
				Password:  "123456",
			},
		},
		{
			Name:           "Error",
			ExpectedReturn: models.User{},
			ExpectedError:  errors.New("Error"),
			MockReturn:     models.User{},
			MockError:      errors.New("Error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {

			s.mockUserRepository.On("Create", mockv2.Anything).Return(tc.MockReturn, tc.MockError).Once()
			result, err := s.userService.Create(tc.UserModel)
			s.Equal(tc.ExpectedReturn, result)
			s.Equal(tc.ExpectedError, err)
		})
	}
}

func (s *suiteUser) TestUpdateUser() {

	testCase := []struct {
		Name           string
		ExpectedReturn models.User
		ExpectedError  error
		MockReturn     models.User
		MockError      error
		UserModel      models.User
	}{
		{
			Name: "Success",
			ExpectedReturn: models.User{
				FirstName: "User 1",
				Email:     "test@gmail.com",
			},
			ExpectedError: nil,
			MockReturn: models.User{
				FirstName: "User 1",
				Email:     "test@gmail.com",
			},
			MockError: nil,
			UserModel: models.User{
				FirstName: "User 1",
				Email:     "test@gmail.com",
			},
		},
		{
			Name:           "Error",
			ExpectedReturn: models.User{},
			ExpectedError:  errors.New("Error"),
			MockReturn:     models.User{},
			MockError:      errors.New("Error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {

			s.mockUserRepository.On("Update", mockv2.Anything, mockv2.Anything).Return(tc.MockReturn, tc.MockError).Once()
			result, err := s.userService.Update(tc.UserModel, "1")
			s.Equal(tc.ExpectedReturn, result)
			s.Equal(tc.ExpectedError, err)
		})
	}
}

func TestUserService(t *testing.T) {
	suite.Run(t, new(suiteUser))
}
