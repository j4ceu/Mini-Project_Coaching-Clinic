package CoachExperience_Service

import (
	"Mini-Project_Coaching-Clinic/models"
	"Mini-Project_Coaching-Clinic/repositories/mock"
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
)

type suiteCoachExperience struct {
	suite.Suite
	mockCoachExperienceRepository *mock.MockCoachExperienceRepository
	coachExperienceService        CoachExperienceService
}

func (s *suiteCoachExperience) SetupTest() {
	s.mockCoachExperienceRepository = new(mock.MockCoachExperienceRepository)
	s.coachExperienceService = NewCoachExperienceService(s.mockCoachExperienceRepository)
}

func (s *suiteCoachExperience) TestFindByID() {

	testCase := []struct {
		Name           string
		ExpectedReturn models.CoachExperience
		ExpectedError  error
		MockReturn     models.CoachExperience
		MockError      error
	}{
		{
			Name: "Success",
			ExpectedReturn: models.CoachExperience{
				Title:       "Test",
				Description: "Test",
			},
			ExpectedError: nil,
			MockReturn: models.CoachExperience{
				Title:       "Test",
				Description: "Test",
			},
			MockError: nil,
		},
		{
			Name:           "Error",
			ExpectedReturn: models.CoachExperience{},
			ExpectedError:  errors.New("Error"),
			MockReturn:     models.CoachExperience{},
			MockError:      errors.New("Error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mockCoachExperienceRepository.On("FindByID", "1").Return(tc.MockReturn, tc.MockError).Once()
			result, err := s.coachExperienceService.FindByID("1")
			s.Equal(tc.ExpectedReturn, result)
			s.Equal(tc.ExpectedError, err)
		})
	}

}

func (s *suiteCoachExperience) TestCreateCoachExperience() {

	testCase := []struct {
		Name           string
		ExpectedReturn models.CoachExperience
		ExpectedError  error
		MockReturn     models.CoachExperience
		MockError      error
	}{
		{
			Name: "Success",
			ExpectedReturn: models.CoachExperience{
				Title:       "Test",
				Description: "Test",
			},
			ExpectedError: nil,
			MockReturn: models.CoachExperience{
				Title:       "Test",
				Description: "Test",
			},
			MockError: nil,
		},
		{
			Name:           "Error",
			ExpectedReturn: models.CoachExperience{},
			ExpectedError:  errors.New("Error"),
			MockReturn:     models.CoachExperience{},
			MockError:      errors.New("Error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mockCoachExperienceRepository.On("Create", tc.ExpectedReturn).Return(tc.MockReturn, tc.MockError).Once()
			result, err := s.coachExperienceService.Create(tc.ExpectedReturn)
			s.Equal(tc.ExpectedReturn, result)
			s.Equal(tc.ExpectedError, err)
		})
	}
}

func (s *suiteCoachExperience) TestUpdateCoachExperience() {

	testCase := []struct {
		Name           string
		ExpectedReturn models.CoachExperience
		ExpectedError  error
		MockReturn     models.CoachExperience
		MockError      error
	}{
		{
			Name: "Success",
			ExpectedReturn: models.CoachExperience{
				Title:       "Test",
				Description: "Test",
			},
			ExpectedError: nil,
			MockReturn: models.CoachExperience{
				Title:       "Test",
				Description: "Test",
			},
			MockError: nil,
		},
		{
			Name:           "Error",
			ExpectedReturn: models.CoachExperience{},
			ExpectedError:  errors.New("Error"),
			MockReturn:     models.CoachExperience{},
			MockError:      errors.New("Error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mockCoachExperienceRepository.On("Update", tc.ExpectedReturn, "1").Return(tc.MockReturn, tc.MockError).Once()
			result, err := s.coachExperienceService.Update(tc.ExpectedReturn, "1")
			s.Equal(tc.ExpectedReturn, result)
			s.Equal(tc.ExpectedError, err)
		})
	}
}

func (s *suiteCoachExperience) TestDeleteCoachExperience() {

	testCase := []struct {
		Name           string
		ExpectedReturn models.CoachExperience
		ExpectedError  error
		MockReturn     models.CoachExperience
		MockError      error
	}{
		{
			Name: "Success",
			ExpectedReturn: models.CoachExperience{
				Title:       "Test",
				Description: "Test",
			},
			ExpectedError: nil,
			MockReturn: models.CoachExperience{
				Title:       "Test",
				Description: "Test",
			},
			MockError: nil,
		},
		{
			Name:           "Error",
			ExpectedReturn: models.CoachExperience{},
			ExpectedError:  errors.New("Error"),
			MockReturn:     models.CoachExperience{},
			MockError:      errors.New("Error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mockCoachExperienceRepository.On("Delete", "1").Return(tc.MockReturn, tc.MockError).Once()
			result, err := s.coachExperienceService.Delete("1")
			s.Equal(tc.ExpectedReturn, result)
			s.Equal(tc.ExpectedError, err)
		})
	}
}

func TestCoachExperienceSuite(t *testing.T) {
	suite.Run(t, new(suiteCoachExperience))
}
