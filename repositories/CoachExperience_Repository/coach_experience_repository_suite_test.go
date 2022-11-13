package CoachExperience_Repository

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

type coachExperienceSuite struct {
	suite.Suite
	mock                sqlmock.Sqlmock
	coachExperienceRepo CoachExperienceRepositories
}

func (s *coachExperienceSuite) SetupSuite() {
	db, mock, _ := sqlmock.New()
	s.mock = mock

	DB, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}))

	s.coachExperienceRepo = NewCoachExperienceRepositories(DB)
}

func (s *coachExperienceSuite) TestFindByID() {

	testCase := []struct {
		Name           string
		Id             string
		Mock           func()
		ExpectedReturn models.CoachExperience
		ExpectedErr    error
	}{
		{
			Name: "Success",
			Id:   "1",
			Mock: func() {
				rows := sqlmock.NewRows([]string{"id", "coach_id", "description"}).
					AddRow(uuid.MustParse("00000000-0000-0000-0000-000000000000"), 1, "test")
				s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "coach_experiences" WHERE id = $1 ORDER BY "coach_experiences"."id" LIMIT 1`)).
					WithArgs(sqlmock.AnyArg()).WillReturnRows(rows).WillReturnError(nil)
			},
			ExpectedReturn: models.CoachExperience{
				ID:          uuid.MustParse("00000000-0000-0000-0000-000000000000"),
				CoachID:     "1",
				Description: "test",
			},
			ExpectedErr: nil,
		},
		{
			Name: "Failed",
			Id:   "1",
			Mock: func() {
				s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "coach_experiences" WHERE id = $1 ORDER BY "coach_experiences"."id" LIMIT 1`)).
					WithArgs(sqlmock.AnyArg()).WillReturnError(errors.New("error"))
			},
			ExpectedReturn: models.CoachExperience{},
			ExpectedErr:    errors.New("error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			tc.Mock()
			result, err := s.coachExperienceRepo.FindByID(tc.Id)
			s.Equal(tc.ExpectedReturn, result)
			s.Equal(tc.ExpectedErr, err)
		})
	}
}

func (s *coachExperienceSuite) TestUpdate() {

	testCase := []struct {
		Name            string
		Id              string
		coachExperience models.CoachExperience
		Mock            func()
		ExpectedReturn  models.CoachExperience
		ExpectedErr     error
	}{
		{
			Name: "Success",
			Id:   "1",
			coachExperience: models.CoachExperience{
				ID:          uuid.MustParse("00000000-0000-0000-0000-000000000000"),
				CoachID:     "1",
				Description: "test",
			},
			Mock: func() {
				rows := sqlmock.NewRows([]string{"id", "coach_id", "description"}).
					AddRow(uuid.MustParse("00000000-0000-0000-0000-000000000000"), 1, "test")
				s.mock.ExpectBegin()
				s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "coach_experiences" SET "coach_id"=$1,"description"=$2 WHERE id = $3`)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1)).WillReturnError(nil)
				s.mock.ExpectCommit()
				s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "coach_experiences" WHERE id = $1 ORDER BY "coach_experiences"."id" LIMIT 1`)).
					WithArgs(sqlmock.AnyArg()).WillReturnRows(rows).WillReturnError(nil)
			},
			ExpectedReturn: models.CoachExperience{
				ID:          uuid.MustParse("00000000-0000-0000-0000-000000000000"),
				CoachID:     "1",
				Description: "test",
			},
			ExpectedErr: nil,
		},
		{
			Name: "Failed",
			Id:   "1",
			coachExperience: models.CoachExperience{
				ID:          uuid.MustParse("00000000-0000-0000-0000-000000000000"),
				CoachID:     "1",
				Description: "test",
			},
			Mock: func() {
				s.mock.ExpectBegin()
				s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "coach_experiences" SET "coach_id"=$1,"description"=$2 WHERE id = $3`)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnError(errors.New("error")).WillReturnResult(sqlmock.NewResult(0, 0))

				s.mock.ExpectRollback()
			},
			ExpectedReturn: models.CoachExperience{
				ID:          uuid.MustParse("00000000-0000-0000-0000-000000000000"),
				CoachID:     "1",
				Description: "test",
			},
			ExpectedErr: errors.New("error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			tc.Mock()
			result, err := s.coachExperienceRepo.Update(tc.coachExperience, tc.Id)
			s.Equal(tc.ExpectedReturn, result)
			s.Equal(tc.ExpectedErr, err)
		})
	}
}

func (s *coachExperienceSuite) TestDelete() {

	testCase := []struct {
		Name           string
		Id             string
		Mock           func()
		ExpectedReturn models.CoachExperience
		ExpectedErr    error
	}{
		{
			Name: "Success",
			Id:   "1",
			Mock: func() {
				s.mock.ExpectBegin()
				s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "coach_experiences" WHERE id = $1`)).
					WithArgs(sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1)).WillReturnError(nil)
				s.mock.ExpectCommit()
			},
			ExpectedReturn: models.CoachExperience{
				ID:          uuid.MustParse("00000000-0000-0000-0000-000000000000"),
				CoachID:     "",
				Description: "",
			},
			ExpectedErr: nil,
		},
		{
			Name: "Failed",
			Id:   "1",
			Mock: func() {
				s.mock.ExpectBegin()
				s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "coach_experiences" WHERE id = $1`)).
					WithArgs(sqlmock.AnyArg()).WillReturnError(errors.New("error")).WillReturnResult(sqlmock.NewResult(0, 0))

				s.mock.ExpectRollback()
			},
			ExpectedReturn: models.CoachExperience{
				ID:          uuid.MustParse("00000000-0000-0000-0000-000000000000"),
				CoachID:     "",
				Description: "",
			},
			ExpectedErr: errors.New("error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			tc.Mock()
			result, err := s.coachExperienceRepo.Delete(tc.Id)
			s.Equal(tc.ExpectedReturn, result)
			s.Equal(tc.ExpectedErr, err)
		})
	}
}

func TestCoachExperienceRepositorySuite(t *testing.T) {
	suite.Run(t, new(coachExperienceSuite))
}
