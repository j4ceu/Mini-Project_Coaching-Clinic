package CoachAvailability_Service

import (
	"Mini-Project_Coaching-Clinic/dto"
	"Mini-Project_Coaching-Clinic/models"
	"Mini-Project_Coaching-Clinic/repositories/mock"
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/suite"
)

type suiteCoachAvailability struct {
	suite.Suite
	mockCoachAvailabilityRepository *mock.MockCoachAvailabilityRepository
	coachAvailabilityService        CoachAvailabilityService
}

func (s *suiteCoachAvailability) SetupTest() {
	s.mockCoachAvailabilityRepository = new(mock.MockCoachAvailabilityRepository)
	s.coachAvailabilityService = NewCoachAvailabilityService(s.mockCoachAvailabilityRepository)
}

func (s *suiteCoachAvailability) TestFindByID() {

	testCase := []struct {
		Name           string
		ExpectedReturn dto.CoachAvailabilityResponse
		ExpectedError  error
		MockReturn     models.CoachAvailability
		MockError      error
	}{
		{
			Name: "Success",
			ExpectedReturn: dto.CoachAvailabilityResponse{
				ID:        "00000000-0000-0000-0000-000000000000",
				StartTime: "12:00",
				EndTime:   "13:00",
			},
			ExpectedError: nil,
			MockReturn: models.CoachAvailability{
				StartTime: "12:00",
				EndTime:   "13:00",
			},
			MockError: nil,
		},
		{
			Name:           "Error",
			ExpectedReturn: dto.CoachAvailabilityResponse{},
			ExpectedError:  errors.New("Error"),
			MockReturn:     models.CoachAvailability{},
			MockError:      errors.New("Error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mockCoachAvailabilityRepository.On("FindByID", "1").Return(tc.MockReturn, tc.MockError).Once()
			result, err := s.coachAvailabilityService.FindByID("1")
			s.Equal(tc.ExpectedReturn, result)
			s.Equal(tc.ExpectedError, err)
		})
	}
}

func (s *suiteCoachAvailability) TestCreateCoachAvailability() {

	testCase := []struct {
		Name           string
		ExpectedReturn dto.CoachAvailabilityResponse
		ExpectedError  error
		MockReturn     models.CoachAvailability
		MockError      error
	}{
		{
			Name: "Success",
			ExpectedReturn: dto.CoachAvailabilityResponse{
				ID:        "00000000-0000-0000-0000-000000000000",
				StartTime: "12:00",
				EndTime:   "13:00",
			},
			ExpectedError: nil,
			MockReturn: models.CoachAvailability{
				StartTime: "12:00",
				EndTime:   "13:00",
			},
			MockError: nil,
		},
		{
			Name:           "Error",
			ExpectedReturn: dto.CoachAvailabilityResponse{},
			ExpectedError:  errors.New("Error"),
			MockReturn: models.CoachAvailability{
				StartTime: "12:00",
				EndTime:   "13:00",
			},
			MockError: errors.New("Error"),
		},
		{
			Name:           "Error Interval",
			ExpectedReturn: dto.CoachAvailabilityResponse{},
			ExpectedError:  &strconv.NumError{Func: "ParseFloat", Num: "", Err: errors.New("invalid syntax")},
			MockReturn:     models.CoachAvailability{},
			MockError:      &strconv.NumError{Func: "ParseFloat", Num: "", Err: errors.New("invalid syntax")},
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mockCoachAvailabilityRepository.On("Create", tc.MockReturn).Return(tc.MockReturn, tc.MockError).Once()
			result, err := s.coachAvailabilityService.Create(tc.MockReturn)
			s.Equal(tc.ExpectedReturn, result)
			s.Equal(tc.ExpectedError, err)
		})
	}

}

func (s *suiteCoachAvailability) TestUpdateCoachAvailability() {

	testCase := []struct {
		Name           string
		ExpectedReturn dto.CoachAvailabilityResponse
		ExpectedError  error
		MockReturn     models.CoachAvailability
		MockError      error
	}{
		{
			Name: "Success",
			ExpectedReturn: dto.CoachAvailabilityResponse{
				ID:        "00000000-0000-0000-0000-000000000000",
				StartTime: "12:00",
				EndTime:   "13:00",
			},
			ExpectedError: nil,
			MockReturn: models.CoachAvailability{
				StartTime: "12:00",
				EndTime:   "13:00",
			},
			MockError: nil,
		},
		{
			Name:           "Error",
			ExpectedReturn: dto.CoachAvailabilityResponse{},
			ExpectedError:  errors.New("Error"),
			MockReturn: models.CoachAvailability{
				StartTime: "12:00",
				EndTime:   "13:00",
			},
			MockError: errors.New("Error"),
		},
		{
			Name:           "Error Interval",
			ExpectedReturn: dto.CoachAvailabilityResponse{},
			ExpectedError:  &strconv.NumError{Func: "ParseFloat", Num: "", Err: errors.New("invalid syntax")},
			MockReturn:     models.CoachAvailability{},
			MockError:      &strconv.NumError{Func: "ParseFloat", Num: "", Err: errors.New("invalid syntax")},
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mockCoachAvailabilityRepository.On("Update", tc.MockReturn, "1").Return(tc.MockReturn, tc.MockError).Once()
			result, err := s.coachAvailabilityService.Update(tc.MockReturn, "1")
			s.Equal(tc.ExpectedReturn, result)
			s.Equal(tc.ExpectedError, err)
		})
	}

}

func (s *suiteCoachAvailability) TestDeleteCoachAvailability() {

	testCase := []struct {
		Name           string
		ExpectedReturn models.CoachAvailability
		ExpectedError  error
		MockReturn     models.CoachAvailability
		MockError      error
	}{
		{
			Name: "Success",
			ExpectedReturn: models.CoachAvailability{
				StartTime: "12:00",
				EndTime:   "13:00",
			},
			ExpectedError: nil,
			MockReturn: models.CoachAvailability{
				StartTime: "12:00",
				EndTime:   "13:00",
			},
			MockError: nil,
		},
		{
			Name:           "Error",
			ExpectedReturn: models.CoachAvailability{},
			ExpectedError:  errors.New("Error"),
			MockReturn: models.CoachAvailability{
				StartTime: "12:00",
				EndTime:   "13:00",
			},
			MockError: errors.New("Error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mockCoachAvailabilityRepository.On("Delete", "1").Return(tc.ExpectedReturn, tc.MockError).Once()
			result, err := s.coachAvailabilityService.Delete("1")
			s.Equal(tc.ExpectedReturn, result)
			s.Equal(tc.ExpectedError, err)
		})
	}
}

func (s *suiteCoachAvailability) TestCheckInterval() {

	testCase := []struct {
		Name           string
		ExpectedReturn bool
		ExpectedError  error
		MockReturn     bool
		MockError      error
		Payload        models.CoachAvailability
	}{
		{
			Name:           "Success",
			ExpectedReturn: true,
			ExpectedError:  nil,
			MockReturn:     true,
			MockError:      nil,
			Payload: models.CoachAvailability{
				StartTime: "12:00",
				EndTime:   "13:00",
			},
		},
		{
			Name:           "Error",
			ExpectedReturn: false,
			ExpectedError:  &strconv.NumError{Func: "ParseFloat", Num: "", Err: errors.New("invalid syntax")},
			MockReturn:     false,
			MockError:      &strconv.NumError{Func: "ParseFloat", Num: "", Err: errors.New("invalid syntax")},
			Payload: models.CoachAvailability{
				StartTime: "",
				EndTime:   "13:00",
			},
		},
		{
			Name:           "Error",
			ExpectedReturn: false,
			ExpectedError:  &strconv.NumError{Func: "ParseFloat", Num: "", Err: errors.New("invalid syntax")},
			MockReturn:     false,
			MockError:      &strconv.NumError{Func: "ParseFloat", Num: "", Err: errors.New("invalid syntax")},
			Payload: models.CoachAvailability{
				StartTime: "12:00",
				EndTime:   "",
			},
		},
		{
			Name:           "Error",
			ExpectedReturn: false,
			ExpectedError:  errors.New("Time interval must be 1 hour"),
			MockReturn:     false,
			MockError:      errors.New("Time interval must be 1 hour"),
			Payload: models.CoachAvailability{
				StartTime: "12:00",
				EndTime:   "14:00",
			},
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			result, err := s.coachAvailabilityService.CheckInterval(tc.Payload)
			s.Equal(tc.ExpectedReturn, result)
			s.Equal(tc.ExpectedError, err)
		})
	}
}

func TestCoachAvailability(t *testing.T) {
	suite.Run(t, new(suiteCoachAvailability))
}
