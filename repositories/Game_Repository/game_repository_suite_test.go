package Game_Repository

import (
	"Mini-Project_Coaching-Clinic/models"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type gameSuite struct {
	suite.Suite
	mock     sqlmock.Sqlmock
	gameRepo GameRepositories
}

func (s *gameSuite) SetupSuite() {
	db, mock, _ := sqlmock.New()
	s.mock = mock

	DB, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}))

	s.gameRepo = NewGameRepositories(DB)
}

func (s *gameSuite) TestFindByID() {

	testCase := []struct {
		Name           string
		Id             uint
		Mock           func()
		ExpectedReturn models.Game
		ExpectedErr    error
	}{
		{
			Name: "Success",
			Id:   1,
			Mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "description"}).
					AddRow(1, "test", "test")
				s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "games" WHERE id = $1 ORDER BY "games"."id" LIMIT 1`)).
					WithArgs(1).WillReturnRows(rows).WillReturnError(nil)
			},
			ExpectedReturn: models.Game{
				ID:          1,
				Name:        "test",
				Description: "test",
			},
			ExpectedErr: nil,
		},
		{
			Name: "Failed",
			Id:   1,
			Mock: func() {
				s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "games" WHERE id = $1 ORDER BY "games"."id" LIMIT 1`)).
					WithArgs(1).WillReturnError(errors.New("error"))
			},
			ExpectedReturn: models.Game{},
			ExpectedErr:    errors.New("error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			tc.Mock()
			result, err := s.gameRepo.FindByID(tc.Id)
			s.Equal(tc.ExpectedReturn, result)
			s.Equal(tc.ExpectedErr, err)
		})
	}

}

func (s *gameSuite) TestCreate() {

	testCase := []struct {
		Name           string
		Game           models.Game
		Mock           func()
		ExpectedReturn models.Game
		ExpectedErr    error
	}{
		{
			Name: "Success",
			Game: models.Game{
				Name:        "test",
				Description: "test",
			},
			Mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "description"}).
					AddRow(1, "test", "test")

				s.mock.ExpectBegin()
				s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "games" ("name","description") VALUES ($1,$2) RETURNING "id"`)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnRows(rows).WillReturnError(nil)
				s.mock.ExpectCommit()
			},
			ExpectedReturn: models.Game{
				ID:          1,
				Name:        "test",
				Description: "test",
			},
			ExpectedErr: nil,
		},
		{
			Name: "Failed",
			Game: models.Game{
				Name:        "",
				Description: "",
			},
			Mock: func() {
				s.mock.ExpectBegin()
				s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "games" ("name","description") VALUES ($1,$2) RETURNING "id"`)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnError(errors.New("error"))
				s.mock.ExpectRollback()
			},
			ExpectedReturn: models.Game{},
			ExpectedErr:    errors.New("error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			tc.Mock()
			result, err := s.gameRepo.Create(tc.Game)
			s.Equal(tc.ExpectedReturn.Name, result.Name)
			s.Equal(tc.ExpectedErr, err)
		})
	}

}

func (s *gameSuite) TestUpdate() {

	testCase := []struct {
		Name           string
		Game           models.Game
		Mock           func()
		ExpectedReturn models.Game
		ExpectedErr    error
	}{
		{
			Name: "Success",
			Game: models.Game{
				ID:          1,
				Name:        "test",
				Description: "test",
			},
			Mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "description"}).
					AddRow(1, "test", "test")

				s.mock.ExpectBegin()
				s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "games" SET "id"=$1,"name"=$2,"description"=$3 WHERE id = $4`)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnError(nil).WillReturnResult(sqlmock.NewResult(1, 1))
				s.mock.ExpectCommit()
				s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "games" WHERE id = $1 ORDER BY "games"."id" LIMIT 1`)).
					WithArgs(1).WillReturnRows(rows).WillReturnError(nil)
			},
			ExpectedReturn: models.Game{
				ID:          1,
				Name:        "test",
				Description: "test",
			},
			ExpectedErr: nil,
		},
		{
			Name: "Failed",
			Game: models.Game{
				ID:          1,
				Name:        "",
				Description: "",
			},
			Mock: func() {
				s.mock.ExpectBegin()
				s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "games" SET "id"=$1 WHERE id = $2`)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnError(errors.New("error"))
				s.mock.ExpectRollback()
			},
			ExpectedReturn: models.Game{},
			ExpectedErr:    errors.New("error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			tc.Mock()
			result, err := s.gameRepo.Update(tc.Game, tc.Game.ID)
			s.Equal(tc.ExpectedReturn.Name, result.Name)
			s.Equal(tc.ExpectedErr, err)
		})
	}

}

func (s *gameSuite) TestDelete() {

	testCase := []struct {
		Name           string
		Id             uint
		Mock           func()
		ExpectedReturn models.Game
		ExpectedErr    error
	}{
		{
			Name: "Success",
			Id:   1,
			Mock: func() {

				s.mock.ExpectBegin()
				s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "games" WHERE id = $1`)).
					WithArgs(sqlmock.AnyArg()).WillReturnError(nil).WillReturnResult(sqlmock.NewResult(1, 1))
				s.mock.ExpectCommit()
			},
			ExpectedReturn: models.Game{},
			ExpectedErr:    nil,
		},
		{
			Name: "Failed",
			Id:   1,
			Mock: func() {
				s.mock.ExpectBegin()
				s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "games" WHERE id = $1`)).
					WithArgs(sqlmock.AnyArg()).WillReturnError(errors.New("error"))
				s.mock.ExpectRollback()
			},
			ExpectedReturn: models.Game{},
			ExpectedErr:    errors.New("error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			tc.Mock()
			result, err := s.gameRepo.Delete(tc.Id)
			s.Equal(tc.ExpectedReturn.Name, result.Name)
			s.Equal(tc.ExpectedErr, err)
		})
	}

}

func (s *gameSuite) TestGetAll() {

	testCase := []struct {
		Name           string
		Mock           func()
		ExpectedReturn []models.Game
		ExpectedErr    error
	}{
		{
			Name: "Success",
			Mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "description"}).
					AddRow(1, "test", "test")

				s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "games"`)).
					WillReturnRows(rows).WillReturnError(nil)
			},
			ExpectedReturn: []models.Game{
				{
					ID:          1,
					Name:        "test",
					Description: "test",
				},
			},
			ExpectedErr: nil,
		},
		{
			Name: "Failed",
			Mock: func() {
				s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "games"`)).
					WillReturnError(errors.New("error"))
			},
			ExpectedReturn: []models.Game(nil),
			ExpectedErr:    errors.New("error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			tc.Mock()
			result, err := s.gameRepo.FindAll()
			s.Equal(tc.ExpectedReturn, result)
			s.Equal(tc.ExpectedErr, err)
		})
	}

}

func (s *gameSuite) TestGetById() {

	testCase := []struct {
		Name           string
		Id             uint
		Mock           func()
		ExpectedReturn models.Game
		ExpectedErr    error
	}{
		{
			Name: "Success",
			Id:   1,
			Mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "description"}).
					AddRow(1, "test", "test")

				s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "games" WHERE id = $1 ORDER BY "games"."id" LIMIT 1`)).
					WithArgs(1).WillReturnRows(rows).WillReturnError(nil)
			},
			ExpectedReturn: models.Game{
				ID:          1,
				Name:        "test",
				Description: "test",
			},
			ExpectedErr: nil,
		},
		{
			Name: "Failed",
			Id:   1,
			Mock: func() {
				s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "games" WHERE id = $1 ORDER BY "games"."id" LIMIT 1`)).
					WithArgs(1).WillReturnError(errors.New("error"))
			},
			ExpectedReturn: models.Game{},
			ExpectedErr:    errors.New("error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			tc.Mock()
			result, err := s.gameRepo.FindByID(tc.Id)
			s.Equal(tc.ExpectedReturn, result)
			s.Equal(tc.ExpectedErr, err)
		})
	}

}

func TestGameRepositorySuite(t *testing.T) {
	suite.Run(t, new(gameSuite))
}
