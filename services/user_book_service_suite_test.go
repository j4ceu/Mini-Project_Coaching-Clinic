package services

import (
	"Mini-Project_Coaching-Clinic/dto"
	"Mini-Project_Coaching-Clinic/dto/payload"
	"Mini-Project_Coaching-Clinic/models"
	"Mini-Project_Coaching-Clinic/repositories/mock"
	"errors"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type suiteUserBook struct {
	suite.Suite
	mockUserBookRepository          *mock.MockUserBookRepository
	mockUserPaymentRepository       *mock.MockUserPaymentRepository
	mockCoachAvailabilityRepository *mock.MockCoachAvailabilityRepository
	mockCoachRepository             *mock.MockCoachRepository
	userBookService                 UserBookService
}

func (s *suiteUserBook) SetupTest() {
	s.mockUserBookRepository = new(mock.MockUserBookRepository)
	s.mockUserPaymentRepository = new(mock.MockUserPaymentRepository)
	s.mockCoachAvailabilityRepository = new(mock.MockCoachAvailabilityRepository)
	s.mockCoachRepository = new(mock.MockCoachRepository)
	s.userBookService = NewUserBookServices(s.mockUserBookRepository, s.mockUserPaymentRepository, s.mockCoachAvailabilityRepository, s.mockCoachRepository)
}

func (s *suiteUserBook) TestFindAllUserBook() {

	testCase := []struct {
		Name           string
		ExpectedReturn []dto.UserBookResponse
		ExpectedError  error
		MockReturn     []dto.UserBookResponse
		MockError      error
	}{
		{
			Name: "Success",
			ExpectedReturn: []dto.UserBookResponse{
				{
					Title: "Test",
				},
			},
			ExpectedError: nil,
			MockReturn: []dto.UserBookResponse{
				{
					Title: "Test",
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
			s.mockUserBookRepository.On("FindAll").Return([]models.UserBook{{Title: "Test"}}, tc.MockError).Once()
			result, err := s.userBookService.FindAll()
			if tc.MockError == nil {
				s.Equal(tc.ExpectedReturn[0].Title, result[0].Title)
			}
			s.Equal(tc.ExpectedError, err)
		})
	}
}

func (s *suiteUserBook) TestFindUserBookByID() {

	testCase := []struct {
		Name           string
		ExpectedReturn dto.UserBookResponse
		ExpectedError  error
		MockReturn     dto.UserBookResponse
		MockError      error
	}{
		{
			Name: "Success",
			ExpectedReturn: dto.UserBookResponse{
				Title: "Test",
			},
			ExpectedError: nil,
			MockReturn: dto.UserBookResponse{
				Title: "Test",
			},
			MockError: nil,
		},
		{
			Name:           "Error",
			ExpectedReturn: dto.UserBookResponse{},
			ExpectedError:  errors.New("Error"),
			MockReturn:     dto.UserBookResponse{},
			MockError:      errors.New("Error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mockUserBookRepository.On("FindByID", "1").Return(models.UserBook{Title: "Test"}, tc.MockError).Once()
			result, err := s.userBookService.FindByID("1")
			if tc.MockError == nil {
				s.Equal(tc.ExpectedReturn.Title, result.Title)
			}
			s.Equal(tc.ExpectedError, err)
		})
	}
}

func (s *suiteUserBook) TestFindByUserID() {

	testCase := []struct {
		Name           string
		ExpectedReturn []dto.UserBookResponse
		ExpectedError  error
		MockReturn     []dto.UserBookResponse
		MockError      error
	}{
		{
			Name: "Success",
			ExpectedReturn: []dto.UserBookResponse{
				{
					Title: "Test",
				},
			},
			ExpectedError: nil,
			MockReturn: []dto.UserBookResponse{
				{
					Title: "Test",
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
			s.mockUserPaymentRepository.On("FindUserBookByUserID", "1").Return([]models.UserBook{{Title: "Test"}}, tc.MockError).Once()
			result, err := s.userBookService.FindByUserID("1")
			if tc.MockError == nil {
				s.Equal(tc.ExpectedReturn[0].Title, result[0].Title)
			}
			s.Equal(tc.ExpectedError, err)
		})
	}
}

func (s *suiteUserBook) TestCreateUserBookErrorCreate() {

	testCase := []struct {
		Name                             string
		ExpectedReturn                   dto.UserBookResponse
		ExpectedError                    error
		MockReturn                       dto.UserBookResponse
		MockError                        error
		MockUserBookReturn               models.UserBook
		MockUserBookError                error
		MockCoachAvailabilityReturn      models.CoachAvailability
		MockCoachAvailabilityError       error
		MockCoachReturn                  models.Coach
		MockCoachError                   error
		MockUserPaymentReturn            models.UserPayment
		MockUserPaymentError             error
		MockCoachAvailabilityUpdateError error
	}{
		{
			Name:           "Error - Create User Book",
			ExpectedReturn: dto.UserBookResponse{},
			ExpectedError:  errors.New("Error"),
			MockReturn:     dto.UserBookResponse{},
			MockError:      errors.New("Error"),
			MockUserBookReturn: models.UserBook{
				CoachAvailabilityID: "1",
			},
			MockUserBookError: errors.New("Error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mockUserBookRepository.On("Create", tc.MockUserBookReturn).Return(tc.MockUserBookReturn, tc.MockUserBookError)

			result, err := s.userBookService.Create(payload.UserBookPayloadCreate{
				CoachAvailabilityID: "1",
			})
			fmt.Println(err)
			if tc.MockError == nil {
				s.Equal(tc.ExpectedReturn.CoachAvailabilityID, result.CoachAvailabilityID)
			}
			s.Equal(tc.ExpectedError, err)
		})
	}
}

func (s *suiteUserBook) TestCreateUserBook_ErrorWhenFindCoachByID() {

	testCase := []struct {
		Name                             string
		ExpectedReturn                   dto.UserBookResponse
		ExpectedError                    error
		MockReturn                       dto.UserBookResponse
		MockError                        error
		MockUserBookReturn               models.UserBook
		MockUserBookError                error
		MockCoachAvailabilityReturn      models.CoachAvailability
		MockCoachAvailabilityError       error
		MockCoachReturn                  models.Coach
		MockCoachError                   error
		MockUserPaymentReturn            models.UserPayment
		MockUserPaymentError             error
		MockCoachAvailabilityUpdateError error
	}{
		{
			Name:           "Error - Find Coach By ID",
			ExpectedReturn: dto.UserBookResponse{},
			ExpectedError:  errors.New("Error"),
			MockReturn:     dto.UserBookResponse{},
			MockError:      errors.New("Error"),
			MockUserBookReturn: models.UserBook{
				CoachAvailabilityID: "1",
			},
			MockUserBookError: nil,
			MockCoachAvailabilityReturn: models.CoachAvailability{
				ID:      uuid.MustParse("69359037-9599-48e7-b8f2-48393c019135"),
				Book:    true,
				CoachID: "69359037-9599-48e7-b8f2-48393c019135",
			},
			MockCoachAvailabilityError:       nil,
			MockCoachAvailabilityUpdateError: nil,
			MockCoachReturn: models.Coach{
				ID: uuid.MustParse("69359037-9599-48e7-b8f2-48393c019135"),
			},
			MockCoachError: errors.New("Error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mockUserBookRepository.On("Create", tc.MockUserBookReturn).Return(tc.MockUserBookReturn, tc.MockUserBookError)
			s.mockCoachAvailabilityRepository.On("FindByID", "1").Return(tc.MockCoachAvailabilityReturn, tc.MockCoachAvailabilityError)
			s.mockCoachAvailabilityRepository.On("Update", tc.MockCoachAvailabilityReturn, tc.MockCoachAvailabilityReturn.ID.String()).Return(tc.MockCoachAvailabilityReturn, tc.MockCoachAvailabilityUpdateError)
			s.mockCoachRepository.On("FindByID", tc.MockCoachReturn.ID.String()).Return(tc.MockCoachReturn, tc.MockCoachError)

			result, err := s.userBookService.Create(payload.UserBookPayloadCreate{
				CoachAvailabilityID: "1",
			})
			fmt.Println(err)
			if tc.MockError == nil {
				s.Equal(tc.ExpectedReturn.CoachAvailabilityID, result.CoachAvailabilityID)
			}
			s.Equal(tc.ExpectedError, err)
		})
	}
}

func (s *suiteUserBook) TestCreateUserBookErrorWhenUpdateCoachAvail() {

	testCase := []struct {
		Name                             string
		ExpectedReturn                   dto.UserBookResponse
		ExpectedError                    error
		MockReturn                       dto.UserBookResponse
		MockError                        error
		MockUserBookReturn               models.UserBook
		MockUserBookError                error
		MockCoachAvailabilityReturn      models.CoachAvailability
		MockCoachAvailabilityError       error
		MockCoachReturn                  models.Coach
		MockCoachError                   error
		MockUserPaymentReturn            models.UserPayment
		MockUserPaymentError             error
		MockCoachAvailabilityUpdateError error
	}{
		{
			Name:           "Error - Update Coach Availability",
			ExpectedReturn: dto.UserBookResponse{},
			ExpectedError:  errors.New("Error"),
			MockReturn:     dto.UserBookResponse{},
			MockError:      errors.New("Error"),
			MockUserBookReturn: models.UserBook{
				CoachAvailabilityID: "1",
			},
			MockUserBookError: nil,
			MockCoachAvailabilityReturn: models.CoachAvailability{
				ID:   uuid.MustParse("69359037-9599-48e7-b8f2-48393c019135"),
				Book: true,
			},
			MockCoachAvailabilityError:       nil,
			MockCoachAvailabilityUpdateError: errors.New("Error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mockUserBookRepository.On("Create", tc.MockUserBookReturn).Return(tc.MockUserBookReturn, tc.MockUserBookError)
			s.mockCoachAvailabilityRepository.On("FindByID", "1").Return(tc.MockCoachAvailabilityReturn, tc.MockCoachAvailabilityError).Once()
			s.mockCoachAvailabilityRepository.On("Update", tc.MockCoachAvailabilityReturn, tc.MockCoachAvailabilityReturn.ID.String()).Return(tc.MockCoachAvailabilityReturn, tc.MockCoachAvailabilityUpdateError).Once()

			result, err := s.userBookService.Create(payload.UserBookPayloadCreate{
				CoachAvailabilityID: "1",
			})
			fmt.Println(err)
			if tc.MockError == nil {
				s.Equal(tc.ExpectedReturn.CoachAvailabilityID, result.CoachAvailabilityID)
			}
			s.Equal(tc.ExpectedError, err)
		})
	}
}

func (s *suiteUserBook) TestCreateUserBookErrorWhenFindCoach() {

	testCase := []struct {
		Name                             string
		ExpectedReturn                   dto.UserBookResponse
		ExpectedError                    error
		MockReturn                       dto.UserBookResponse
		MockError                        error
		MockUserBookReturn               models.UserBook
		MockUserBookError                error
		MockCoachAvailabilityReturn      models.CoachAvailability
		MockCoachAvailabilityError       error
		MockCoachReturn                  models.Coach
		MockCoachError                   error
		MockUserPaymentReturn            models.UserPayment
		MockUserPaymentError             error
		MockCoachAvailabilityUpdateError error
	}{
		{
			Name:           "Error - Find Coach Availability",
			ExpectedReturn: dto.UserBookResponse{},
			ExpectedError:  errors.New("Error"),
			MockReturn:     dto.UserBookResponse{},
			MockError:      errors.New("Error"),
			MockUserBookReturn: models.UserBook{
				CoachAvailabilityID: "1",
			},
			MockUserBookError:           nil,
			MockCoachAvailabilityReturn: models.CoachAvailability{},
			MockCoachAvailabilityError:  errors.New("Error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mockUserBookRepository.On("Create", tc.MockUserBookReturn).Return(tc.MockUserBookReturn, tc.MockUserBookError)
			s.mockCoachAvailabilityRepository.On("FindByID", "1").Return(tc.MockCoachAvailabilityReturn, tc.MockCoachAvailabilityError).Once()

			result, err := s.userBookService.Create(payload.UserBookPayloadCreate{
				CoachAvailabilityID: "1",
			})
			fmt.Println(err)
			if tc.MockError == nil {
				s.Equal(tc.ExpectedReturn.CoachAvailabilityID, result.CoachAvailabilityID)
			}
			s.Equal(tc.ExpectedError, err)
		})
	}
}

func (s *suiteUserBook) TestUpdateUserBook() {

	testCase := []struct {
		Name           string
		ExpectedReturn dto.UserBookResponse
		ExpectedError  error
		MockReturn     dto.UserBookResponse
		MockError      error
	}{
		{
			Name: "Success",
			ExpectedReturn: dto.UserBookResponse{
				Title: "Title",
			},
			ExpectedError: nil,
			MockReturn: dto.UserBookResponse{
				Title: "Title",
			},
			MockError: nil,
		},
		{
			Name:           "Error ",
			ExpectedReturn: dto.UserBookResponse{},
			ExpectedError:  errors.New("Error"),
			MockReturn:     dto.UserBookResponse{},
			MockError:      errors.New("Error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mockUserBookRepository.On("Update", models.UserBook{Title: "Title"}, "1").Return(models.UserBook{Title: "Title"}, tc.MockError).Once()
			result, err := s.userBookService.Update(payload.UserBookPayloadUpdate{Title: "Title"}, "1")
			if tc.MockError == nil {
				s.Equal(tc.ExpectedReturn.Title, result.Title)
			}
			s.Equal(tc.ExpectedError, err)
		})
	}

}

func (s *suiteUserBook) TestDeleteUserBook() {

	testCase := []struct {
		Name           string
		ExpectedReturn models.UserBook
		ExpectedError  error
		MockReturn     models.UserBook
		MockError      error
	}{
		{
			Name: "Success",
			ExpectedReturn: models.UserBook{
				Title: "Title",
			},
			ExpectedError: nil,
			MockReturn: models.UserBook{
				Title: "Title",
			},
			MockError: nil,
		},
		{
			Name:           "Error ",
			ExpectedReturn: models.UserBook{},
			ExpectedError:  errors.New("Error"),
			MockReturn:     models.UserBook{},
			MockError:      errors.New("Error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mockUserBookRepository.On("Delete", "1").Return(tc.MockReturn, tc.MockError).Once()
			result, err := s.userBookService.Delete("1")
			if tc.MockError == nil {
				s.Equal(tc.ExpectedReturn.Title, result.Title)
			}
			s.Equal(tc.ExpectedError, err)
		})
	}
}

func TestUserBook(t *testing.T) {
	suite.Run(t, new(suiteUserBook))
}
