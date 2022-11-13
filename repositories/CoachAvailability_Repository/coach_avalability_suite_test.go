package CoachAvailability_Repository

import (
	"Mini-Project_Coaching-Clinic/models"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type coachAvailabilitySuite struct {
	suite.Suite
	mock                  sqlmock.Sqlmock
	coachAvailabilityRepo CoachAvailabilityRepositories
}

func (s *coachAvailabilitySuite) SetupSuite() {
	db, mock, _ := sqlmock.New()
	s.mock = mock

	DB, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}))

	s.coachAvailabilityRepo = NewCoachAvailabilityRepositories(DB)
}

func (s *coachAvailabilitySuite) TestFindByID() {

	testCase := []struct {
		Name           string
		Id             string
		Mock           func()
		ExpectedReturn models.CoachAvailability
		ExpectedErr    error
	}{
		{
			Name: "Success",
			Id:   "1",
			Mock: func() {
				rows := sqlmock.NewRows([]string{"id", "coach_id", "day", "start_time", "end_time"}).
					AddRow(uuid.MustParse("00000000-0000-0000-0000-000000000000"), 1, "test", "test", "test")
				s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "coach_availabilities" WHERE id = $1 ORDER BY "coach_availabilities"."id" LIMIT 1`)).
					WithArgs(sqlmock.AnyArg()).WillReturnRows(rows).WillReturnError(nil)
			},
			ExpectedReturn: models.CoachAvailability{
				ID:        uuid.MustParse("00000000-0000-0000-0000-000000000000"),
				CoachID:   "1",
				Day:       "test",
				StartTime: "test",
				EndTime:   "test",
			},
			ExpectedErr: nil,
		},
		{
			Name: "Failed",
			Id:   "1",
			Mock: func() {
				s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "coach_availabilities" WHERE id = $1 ORDER BY "coach_availabilities"."id" LIMIT 1`)).
					WithArgs(sqlmock.AnyArg()).WillReturnError(errors.New("error"))
			},
			ExpectedReturn: models.CoachAvailability{},
			ExpectedErr:    errors.New("error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			tc.Mock()
			result, err := s.coachAvailabilityRepo.FindByID(tc.Id)
			s.Equal(tc.ExpectedReturn.Day, result.Day)
			s.Equal(tc.ExpectedErr, err)
		})
	}
}

func (s *coachAvailabilitySuite) TestUpdate() {

	testCase := []struct {
		Name              string
		Id                string
		coachAvailability models.CoachAvailability
		Mock              func()
		ExpectedReturn    models.CoachAvailability
		ExpectedErr       error
	}{
		{
			Name: "Success",
			Id:   "1",
			coachAvailability: models.CoachAvailability{
				ID:        uuid.MustParse("00000000-0000-0000-0000-000000000000"),
				CoachID:   "1",
				Day:       "test",
				StartTime: "test",
				EndTime:   "test",
			},
			Mock: func() {

				rows := sqlmock.NewRows([]string{"id", "coach_id", "day", "start_time", "end_time"}).
					AddRow(uuid.MustParse("00000000-0000-0000-0000-000000000000"), 1, "test", "test", "test")
				s.mock.ExpectBegin()
				s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "coach_availabilities" SET "coach_id"=$1,"day"=$2,"start_time"=$3,"end_time"=$4 WHERE id = $5`)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))
				s.mock.ExpectCommit()
				s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "coach_availabilities" WHERE id = $1 ORDER BY "coach_availabilities"."id" LIMIT 1`)).
					WithArgs(sqlmock.AnyArg()).WillReturnRows(rows).WillReturnError(nil)
			},
		},
		{
			Name: "Failed",
			Id:   "1",
			coachAvailability: models.CoachAvailability{
				ID:        uuid.MustParse("00000000-0000-0000-0000-000000000000"),
				CoachID:   "1",
				Day:       "test",
				StartTime: "test",
				EndTime:   "test",
			},
			Mock: func() {
				s.mock.ExpectBegin()
				s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "coach_availabilities" SET "coach_id"=$1,"day"=$2,"start_time"=$3,"end_time"=$4 WHERE id = $5`)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnError(errors.New("error"))
				s.mock.ExpectRollback()
			},
			ExpectedErr: errors.New("error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			tc.Mock()
			result, err := s.coachAvailabilityRepo.Update(tc.coachAvailability, tc.coachAvailability.ID.String())
			s.Equal(tc.ExpectedErr, err)
			s.Equal(tc.coachAvailability.Day, result.Day)
		})
	}
}

func (s *coachAvailabilitySuite) TestDelete() {

	testCase := []struct {
		Name           string
		Id             string
		Mock           func()
		ExpectedReturn models.CoachAvailability
		ExpectedErr    error
	}{
		{
			Name: "Success",
			Id:   "1",
			Mock: func() {
				s.mock.ExpectBegin()
				s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "coach_availabilities" WHERE id = $1`)).
					WithArgs(sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))
				s.mock.ExpectCommit()

			},
		},
		{
			Name: "Failed",
			Id:   "1",
			Mock: func() {
				s.mock.ExpectBegin()
				s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "coach_availabilities" WHERE id = $1`)).
					WithArgs(sqlmock.AnyArg()).WillReturnError(errors.New("error"))
				s.mock.ExpectRollback()
			},
			ExpectedErr: errors.New("error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			tc.Mock()
			result, err := s.coachAvailabilityRepo.Delete(tc.Id)
			s.Equal(tc.ExpectedErr, err)
			s.Equal(tc.ExpectedReturn, result)
		})
	}
}

func TestCoachAvailabilityRepositorySuite(t *testing.T) {
	suite.Run(t, new(coachAvailabilitySuite))
}
