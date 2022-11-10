package services

import (
	"Mini-Project_Coaching-Clinic/models"
	"Mini-Project_Coaching-Clinic/repositories/mock"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type suiteGame struct {
	suite.Suite
	mockGameRepository *mock.MockGameRepository
	gameService        GameService
}

func (s *suiteGame) SetupTest() {
	s.mockGameRepository = new(mock.MockGameRepository)
	s.gameService = NewGameService(s.mockGameRepository)
}

func (s *suiteGame) TestFindAll() {

	testCase := []struct {
		Name           string
		ExpectedReturn []models.Game
		ExpectedError  error
		MockReturn     []models.Game
		MockError      error
	}{
		{
			Name: "Success",
			ExpectedReturn: []models.Game{
				{
					ID:   1,
					Name: "Game 1",
				},
			},
			ExpectedError: nil,
			MockReturn: []models.Game{
				{
					ID:   1,
					Name: "Game 1",
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
			s.mockGameRepository.On("FindAll").Return(tc.MockReturn, tc.MockError).Once()
			result, err := s.gameService.FindAll()
			s.Equal(tc.ExpectedReturn, result)
			s.Equal(tc.ExpectedError, err)
		})
	}
}

func (s *suiteGame) TestFindById() {

	testCase := []struct {
		Name           string
		ExpectedReturn models.Game
		ExpectedError  error
		MockReturn     models.Game
		MockError      error
	}{
		{
			Name: "Success",
			ExpectedReturn: models.Game{
				ID:   1,
				Name: "Game 1",
			},
			ExpectedError: nil,
			MockReturn: models.Game{
				ID:   1,
				Name: "Game 1",
			},
			MockError: nil,
		},
		{
			Name:           "Error",
			ExpectedReturn: models.Game{},
			ExpectedError:  errors.New("Error"),
			MockReturn:     models.Game{},
			MockError:      errors.New("Error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mockGameRepository.On("FindByID", uint(1)).Return(tc.MockReturn, tc.MockError).Once()
			result, err := s.gameService.FindByID(uint(1))
			s.Equal(tc.ExpectedReturn, result)
			s.Equal(tc.ExpectedError, err)
		})
	}
}

func (s *suiteGame) TestCreateGame() {

	testCase := []struct {
		Name           string
		ExpectedReturn models.Game
		ExpectedError  error
		MockReturn     models.Game
		MockError      error
	}{
		{
			Name: "Success",
			ExpectedReturn: models.Game{
				ID:   1,
				Name: "Game 1",
			},
			ExpectedError: nil,
			MockReturn: models.Game{
				ID:   1,
				Name: "Game 1",
			},
			MockError: nil,
		},
		{
			Name:           "Error",
			ExpectedReturn: models.Game{},
			ExpectedError:  errors.New("Error"),
			MockReturn:     models.Game{},
			MockError:      errors.New("Error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mockGameRepository.On("Create", tc.MockReturn).Return(tc.MockReturn, tc.MockError).Once()
			result, err := s.gameService.Create(tc.MockReturn)
			s.Equal(tc.ExpectedReturn, result)
			s.Equal(tc.ExpectedError, err)
		})
	}
}

func (s *suiteGame) TestUpdateGame() {

	testCase := []struct {
		Name           string
		ExpectedReturn models.Game
		ExpectedError  error
		MockReturn     models.Game
		MockError      error
	}{
		{
			Name: "Success",
			ExpectedReturn: models.Game{
				ID:   1,
				Name: "Game 1",
			},
			ExpectedError: nil,
			MockReturn: models.Game{
				ID:   1,
				Name: "Game 1",
			},
			MockError: nil,
		},
		{
			Name:           "Error",
			ExpectedReturn: models.Game{},
			ExpectedError:  errors.New("Error"),
			MockReturn:     models.Game{},
			MockError:      errors.New("Error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mockGameRepository.On("Update", tc.MockReturn, uint(1)).Return(tc.MockReturn, tc.MockError).Once()
			result, err := s.gameService.Update(tc.MockReturn, uint(1))
			s.Equal(tc.ExpectedReturn, result)
			s.Equal(tc.ExpectedError, err)
		})
	}
}

func (s *suiteGame) TestDeleteGame() {

	testCase := []struct {
		Name           string
		ExpectedReturn models.Game
		ExpectedError  error
		MockReturn     models.Game
		MockError      error
	}{
		{
			Name: "Success",
			ExpectedReturn: models.Game{
				ID:   1,
				Name: "Game 1",
			},
			ExpectedError: nil,
			MockReturn: models.Game{
				ID:   1,
				Name: "Game 1",
			},
			MockError: nil,
		},
		{
			Name:           "Error",
			ExpectedReturn: models.Game{},
			ExpectedError:  errors.New("Error"),
			MockReturn:     models.Game{},
			MockError:      errors.New("Error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mockGameRepository.On("Delete", uint(1)).Return(tc.MockReturn, tc.MockError).Once()
			result, err := s.gameService.Delete(uint(1))
			fmt.Println(result)
			s.Equal(tc.ExpectedReturn, result)
			s.Equal(tc.ExpectedError, err)
		})
	}
}

func TestGameSuite(t *testing.T) {
	suite.Run(t, new(suiteGame))
}
