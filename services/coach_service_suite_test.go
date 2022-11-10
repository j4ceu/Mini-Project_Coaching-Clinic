package services

import (
	"Mini-Project_Coaching-Clinic/dto"
	"Mini-Project_Coaching-Clinic/models"
	"Mini-Project_Coaching-Clinic/repositories/mock"
	"errors"
	"fmt"
	"testing"

	mockTest "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type suiteCoach struct {
	suite.Suite
	mockCoachRepository *mock.MockCoachRepository
	mockGameRepository  *mock.MockGameRepository
	mockUserRepository  *mock.MockUserRepository
	coachService        CoachService
}

func (s *suiteCoach) SetupTest() {
	s.mockCoachRepository = new(mock.MockCoachRepository)
	s.mockGameRepository = new(mock.MockGameRepository)
	s.mockUserRepository = new(mock.MockUserRepository)
	s.coachService = NewCoachService(s.mockCoachRepository, s.mockGameRepository, s.mockUserRepository)
}

func (s *suiteCoach) TestFindByID() {

	testCase := []struct {
		Name           string
		ExpectedReturn dto.CoachResponse
		ExpectedError  error
		MockReturn     dto.CoachResponse
		MockError      error
	}{
		{
			Name: "Success",
			ExpectedReturn: dto.CoachResponse{
				Code:     "J4CEU",
				Position: "Coach",
			},
			ExpectedError: nil,
			MockReturn: dto.CoachResponse{
				Code:     "J4CEU",
				Position: "Coach",
			},
			MockError: nil,
		},
		{
			Name:           "Error",
			ExpectedReturn: dto.CoachResponse{},
			ExpectedError:  errors.New("Error"),
			MockReturn:     dto.CoachResponse{},
			MockError:      errors.New("Error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mockCoachRepository.On("FindByID", mockTest.Anything).Return(models.Coach{Code: "J4CEU", Position: "Coach"}, tc.MockError).Once()
			result, err := s.coachService.FindByID("1")
			s.Equal(tc.ExpectedReturn.Code, result.Code)
			s.Equal(tc.ExpectedReturn.Position, result.Position)
			s.Equal(tc.ExpectedError, err)
		})
	}

}

func (s *suiteCoach) TestFindByCode() {

	testCase := []struct {
		Name           string
		ExpectedReturn dto.CoachResponse
		ExpectedError  error
		MockReturn     dto.CoachResponse
		MockError      error
	}{
		{
			Name: "Success",
			ExpectedReturn: dto.CoachResponse{
				Code:     "J4CEU",
				Position: "Coach",
			},
			ExpectedError: nil,
			MockReturn: dto.CoachResponse{
				Code:     "J4CEU",
				Position: "Coach",
			},
			MockError: nil,
		},
		{
			Name:           "Error",
			ExpectedReturn: dto.CoachResponse{},
			ExpectedError:  errors.New("Error"),
			MockReturn:     dto.CoachResponse{},
			MockError:      errors.New("Error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mockCoachRepository.On("FindByCode", mockTest.Anything).Return(models.Coach{Code: "J4CEU", Position: "Coach"}, tc.MockError).Once()
			result, err := s.coachService.FindByCode("J4CEU")
			s.Equal(tc.ExpectedReturn.Code, result.Code)
			s.Equal(tc.ExpectedReturn.Position, result.Position)
			s.Equal(tc.ExpectedError, err)
		})
	}

}

func (s *suiteCoach) TestFindByGameID() {

	testCase := []struct {
		Name           string
		ExpectedReturn []dto.CoachResponse
		ExpectedError  error
		MockReturn     []dto.CoachResponse
		MockError      error
	}{
		{
			Name: "Success",
			ExpectedReturn: []dto.CoachResponse{
				{
					Code:     "J4CEU",
					Position: "Coach",
				},
			},
			ExpectedError: nil,
			MockReturn: []dto.CoachResponse{
				{
					Code:     "J4CEU",
					Position: "Coach",
				},
			},
			MockError: nil,
		},
		{
			Name:           "Error",
			ExpectedReturn: []dto.CoachResponse{},
			ExpectedError:  errors.New("Error"),
			MockReturn:     []dto.CoachResponse{},
			MockError:      errors.New("Error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mockCoachRepository.On("FindByGameID", mockTest.Anything).Return([]models.Coach{{Code: "J4CEU", Position: "Coach"}}, tc.MockError).Once()
			result, err := s.coachService.FindByGameID("1")
			if tc.ExpectedError == nil {
				s.Equal(tc.ExpectedReturn[0].Code, result[0].Code)
				s.Equal(tc.ExpectedReturn[0].Position, result[0].Position)
			}
			s.Equal(tc.ExpectedError, err)
		})
	}

}

func (s *suiteCoach) TestUpdateCoach() {

	testCase := []struct {
		Name           string
		ExpectedReturn dto.CoachResponse
		ExpectedError  error
		MockReturn     dto.CoachResponse
		MockError      error
	}{
		{
			Name: "Success",
			ExpectedReturn: dto.CoachResponse{
				Code:     "J4CEU",
				Position: "Coach",
			},
			ExpectedError: nil,
			MockReturn: dto.CoachResponse{
				Code:     "J4CEU",
				Position: "Coach",
			},
			MockError: nil,
		},
		{
			Name:           "Error",
			ExpectedReturn: dto.CoachResponse{},
			ExpectedError:  errors.New("Error"),
			MockReturn:     dto.CoachResponse{},
			MockError:      errors.New("Error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mockCoachRepository.On("Update", mockTest.Anything, mockTest.Anything).Return(models.Coach{Code: "J4CEU", Position: "Coach"}, tc.MockError).Once()
			result, err := s.coachService.Update(models.Coach{
				Code:     "J4CEU",
				Position: "Coach",
			}, "1")
			s.Equal(tc.ExpectedReturn.Code, result.Code)
			s.Equal(tc.ExpectedReturn.Position, result.Position)
			s.Equal(tc.ExpectedError, err)
		})
	}

}

func (s *suiteCoach) TestDeleteCoach() {

	testCase := []struct {
		Name           string
		ExpectedReturn models.Coach
		ExpectedError  error
		MockReturn     models.Coach
		MockError      error
	}{
		{
			Name: "Success",
			ExpectedReturn: models.Coach{
				Code:     "J4CEU",
				Position: "Coach",
			},
			ExpectedError: nil,
			MockReturn: models.Coach{
				Code:     "J4CEU",
				Position: "Coach",
			},
			MockError: nil,
		},
		{
			Name:           "Error",
			ExpectedReturn: models.Coach{},
			ExpectedError:  errors.New("Error"),
			MockReturn:     models.Coach{},
			MockError:      errors.New("Error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mockCoachRepository.On("Delete", mockTest.Anything).Return(tc.MockReturn, tc.MockError).Once()
			result, err := s.coachService.Delete("1")
			fmt.Println(err)
			s.Equal(tc.ExpectedReturn, result)
			s.Equal(tc.ExpectedError, err)
		})
	}
}

func (s *suiteCoach) TestCreateCoach() {

	testCase := []struct {
		Name            string
		ExpectedReturn  dto.CoachResponse
		ExpectedError   error
		MockReturn      dto.CoachResponse
		MockError       error
		MockCoachModels models.Coach
	}{
		{
			Name: "Success With Create User",
			ExpectedReturn: dto.CoachResponse{
				Code:     "J4CEU",
				Position: "Coach",
			},
			ExpectedError: nil,
			MockReturn: dto.CoachResponse{
				Code:     "J4CEU",
				Position: "Coach",
			},
			MockError: nil,
			MockCoachModels: models.Coach{
				Code:     "J4CEU",
				Position: "Coach",
				User: models.User{
					FirstName: "jace",
					LastName:  "Herondale",
				},
			},
		},
		{
			Name: "Success With Existing User",
			ExpectedReturn: dto.CoachResponse{
				Code:     "J4CEU",
				Position: "Coach",
			},
			ExpectedError: nil,
			MockReturn: dto.CoachResponse{
				Code:     "J4CEU",
				Position: "Coach",
			},
			MockError: nil,
			MockCoachModels: models.Coach{
				Code:     "J4CEU",
				Position: "Coach",
				UserID:   "1",
			},
		},
		{
			Name:           "Error",
			ExpectedReturn: dto.CoachResponse{},
			ExpectedError:  errors.New("Error"),
			MockReturn:     dto.CoachResponse{},
			MockError:      errors.New("Error"),
			MockCoachModels: models.Coach{
				Code:     "J4CEU",
				Position: "Coach",
				User: models.User{
					FirstName: "jace",
					LastName:  "Herondale",
				},
			},
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mockGameRepository.On("FindByID", mockTest.Anything).Return(models.Game{}, nil).Once()
			s.mockCoachRepository.On("Create", mockTest.Anything).Return(tc.MockCoachModels, tc.MockError).Once()
			s.mockUserRepository.On("FindByID", mockTest.Anything).Return(models.User{FirstName: "Jace", LastName: "Herondale"}, nil).Once()
			s.mockUserRepository.On("Update", mockTest.Anything, mockTest.Anything).Return(models.User{FirstName: "Jace", LastName: "Herondale"}, nil).Once()
			result, err := s.coachService.Create(tc.MockCoachModels)
			s.Equal(tc.ExpectedReturn.Code, result.Code)
			s.Equal(tc.ExpectedReturn.Position, result.Position)
			s.Equal(tc.ExpectedError, err)
		})
	}
}

func TestCoachSuite(t *testing.T) {
	suite.Run(t, new(suiteCoach))
}
