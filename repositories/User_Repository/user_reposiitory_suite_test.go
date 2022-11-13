package User_Repository

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

type suiteUser struct {
	suite.Suite
	mock     sqlmock.Sqlmock
	userRepo UserRepositories
}

func (s *suiteUser) SetupSuite() {
	db, mock, _ := sqlmock.New()
	s.mock = mock

	DB, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}))

	s.userRepo = NewUserRepositories(DB)
}

func (s *suiteUser) TestFindByID() {

	testCase := []struct {
		Name           string
		Id             string
		Mock           func()
		ExpectedReturn models.User
		ExpectedErr    error
	}{
		{
			Name: "success",
			Id:   "1",
			Mock: func() {
				rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "password", "role"}).
					AddRow("00000000-0000-0000-0000-000000000000", "Jace", "Herondale", "test@gmail.com", "test", "admin")
				s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE id = $1 ORDER BY "users"."id" LIMIT 1`)).
					WithArgs(sqlmock.AnyArg()).WillReturnRows(rows).WillReturnError(nil)
			},
			ExpectedReturn: models.User{
				ID:        uuid.MustParse("00000000-0000-0000-0000-000000000000"),
				FirstName: "Jace",
				LastName:  "Herondale",
				Email:     "test@gmail.com",
				Password:  "test",
				Role:      "admin",
			},
			ExpectedErr: nil,
		},
		{
			Name: "failed",
			Id:   "1",
			Mock: func() {
				rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "password", "role"}).
					AddRow("00000000-0000-0000-0000-000000000000", "Jace", "Herondale", "test@gmail.com", "test", "admin")
				s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE id = $1 ORDER BY "users"."id" LIMIT 1`)).
					WithArgs(sqlmock.AnyArg()).WillReturnRows(rows).WillReturnError(errors.New("error"))
			},
			ExpectedReturn: models.User{},
			ExpectedErr:    errors.New("error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			tc.Mock()
			result, err := s.userRepo.FindByID(tc.Id)
			s.Equal(tc.ExpectedErr, err)
			s.Equal(tc.ExpectedReturn, result)

			if err := s.mock.ExpectationsWereMet(); err != nil {
				s.T().Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func (s *suiteUser) TestFindByEmail() {

	testCase := []struct {
		Name           string
		Email          string
		Mock           func()
		ExpectedReturn models.User
		ExpectedErr    error
	}{
		{
			Name:  "success",
			Email: "test@gmail.com",
			Mock: func() {
				rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "password", "role"}).
					AddRow("00000000-0000-0000-0000-000000000000", "Jace", "Herondale", "test@gmail.com", "test", "admin")
				s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE email = $1 ORDER BY "users"."id" LIMIT 1`)).
					WithArgs(sqlmock.AnyArg()).WillReturnRows(rows).WillReturnError(nil)
			},
			ExpectedReturn: models.User{
				ID:        uuid.MustParse("00000000-0000-0000-0000-000000000000"),
				FirstName: "Jace",
				LastName:  "Herondale",
				Email:     "test@gmail.com",
				Password:  "test",
				Role:      "admin",
			},
			ExpectedErr: nil,
		},
		{
			Name:  "failed",
			Email: "test@gmail.com",
			Mock: func() {
				rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "password", "role"}).
					AddRow("00000000-0000-0000-0000-000000000000", "Jace", "Herondale", "test@gmail.com", "test", "admin")
				s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE email = $1 ORDER BY "users"."id" LIMIT 1`)).
					WithArgs(sqlmock.AnyArg()).WillReturnRows(rows).WillReturnError(errors.New("error"))
			},
			ExpectedReturn: models.User{},
			ExpectedErr:    errors.New("error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			tc.Mock()
			result, err := s.userRepo.FindByEmail(tc.Email)
			s.Equal(tc.ExpectedErr, err)
			s.Equal(tc.ExpectedReturn, result)

			if err := s.mock.ExpectationsWereMet(); err != nil {
				s.T().Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func (s *suiteUser) TestFindAll() {

	testCase := []struct {
		Name           string
		Mock           func()
		ExpectedReturn []models.User
		ExpectedErr    error
	}{
		{
			Name: "success",
			Mock: func() {
				rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "password", "role"}).
					AddRow("00000000-0000-0000-0000-000000000000", "Jace", "Herondale", "test@gmail.com", "test", "admin")
				s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users"`)).
					WillReturnRows(rows).WillReturnError(nil)
			},
			ExpectedReturn: []models.User{
				{
					ID:        uuid.MustParse("00000000-0000-0000-0000-000000000000"),
					FirstName: "Jace",
					LastName:  "Herondale",
					Email:     "test@gmail.com",
					Password:  "test",
					Role:      "admin",
				},
			},
			ExpectedErr: nil,
		},
		{
			Name: "failed",
			Mock: func() {
				rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "password", "role"}).
					AddRow("00000000-0000-0000-0000-000000000000", "Jace", "Herondale", "test@gmail.com", "test", "admin")
				s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users"`)).
					WillReturnRows(rows).WillReturnError(errors.New("error"))
			},
			ExpectedReturn: []models.User(nil),
			ExpectedErr:    errors.New("error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			tc.Mock()
			result, err := s.userRepo.FindAll()
			s.Equal(tc.ExpectedErr, err)
			s.Equal(tc.ExpectedReturn, result)

			if err := s.mock.ExpectationsWereMet(); err != nil {
				s.T().Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func (s *suiteUser) TestUpdate() {

	testCase := []struct {
		Name           string
		User           models.User
		Mock           func()
		ExpectedReturn models.User
		ExpectedErr    error
	}{
		{
			Name: "success",
			User: models.User{
				ID:        uuid.MustParse("00000000-0000-0000-0000-000000000000"),
				FirstName: "Jace",
				LastName:  "Herondale",
				Email:     "test@gmail.com",
				Password:  "test",
				Role:      "admin",
			},
			Mock: func() {
				s.mock.ExpectBegin()
				s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "users" SET "first_name"=$1,"last_name"=$2,"email"=$3,"password"=$4,"role"=$5 WHERE id = $6`)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1)).WillReturnError(nil)
				s.mock.ExpectCommit()
				rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "password", "role"}).
					AddRow("00000000-0000-0000-0000-000000000000", "Jace", "Herondale", "test@gmail.com", "testUpdate", "admin")
				s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE id = $1 ORDER BY "users"."id" LIMIT 1`)).
					WithArgs(sqlmock.AnyArg()).WillReturnRows(rows).WillReturnError(nil)

			},
			ExpectedReturn: models.User{
				ID:        uuid.MustParse("00000000-0000-0000-0000-000000000000"),
				FirstName: "Jace",
				LastName:  "Herondale",
				Email:     "test@gmail.com",
				Password:  "testUpdate",
				Role:      "admin",
			},
			ExpectedErr: nil,
		},
		{
			Name: "failed",
			User: models.User{
				ID:        uuid.MustParse("00000000-0000-0000-0000-000000000000"),
				FirstName: "Jace",
				LastName:  "Herondale",
				Email:     "test@gmail.com",
				Password:  "test",
				Role:      "admin",
			},
			Mock: func() {
				s.mock.ExpectBegin()
				s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "users" SET "first_name"=$1,"last_name"=$2,"email"=$3,"password"=$4,"role"=$5 WHERE id = $6`)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(0, 0)).WillReturnError(errors.New("error"))
				s.mock.ExpectRollback()
			},
			ExpectedReturn: models.User{
				ID:        uuid.MustParse("00000000-0000-0000-0000-000000000000"),
				FirstName: "Jace",
				LastName:  "Herondale",
				Email:     "test@gmail.com",
				Password:  "test",
				Role:      "admin",
			},
			ExpectedErr: errors.New("error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			tc.Mock()
			result, err := s.userRepo.Update(tc.User, tc.User.ID.String())
			s.Equal(tc.ExpectedErr, err)
			s.Equal(tc.ExpectedReturn, result)

			if err := s.mock.ExpectationsWereMet(); err != nil {
				s.T().Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func (s *suiteUser) TestDelete() {

	testCase := []struct {
		Name           string
		User           models.User
		Mock           func()
		ExpectedReturn models.User
		ExpectedErr    error
	}{
		{
			Name: "success",
			User: models.User{
				ID:        uuid.MustParse("00000000-0000-0000-0000-000000000000"),
				FirstName: "Jace",
				LastName:  "Herondale",
				Email:     "test@gmail.com",
				Password:  "test",
				Role:      "admin",
			},
			Mock: func() {
				s.mock.ExpectBegin()
				s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "users" WHERE id = $1`)).
					WithArgs(sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1)).WillReturnError(nil)
				s.mock.ExpectCommit()
			},
			ExpectedReturn: models.User{
				ID:        uuid.MustParse("00000000-0000-0000-0000-000000000000"),
				FirstName: "",
				LastName:  "",
				Email:     "",
				Password:  "",
				Role:      "",
			},
			ExpectedErr: nil,
		},
		{
			Name: "failed",
			User: models.User{
				ID:        uuid.MustParse("00000000-0000-0000-0000-000000000000"),
				FirstName: "Jace",
				LastName:  "Herondale",
				Email:     "test@gmail.com",
				Password:  "test",
				Role:      "admin",
			},
			Mock: func() {
				s.mock.ExpectBegin()
				s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "users" WHERE id = $1`)).
					WithArgs(sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(0, 0)).WillReturnError(errors.New("error"))
				s.mock.ExpectRollback()
			},
			ExpectedReturn: models.User{
				ID:        uuid.MustParse("00000000-0000-0000-0000-000000000000"),
				FirstName: "",
				LastName:  "",
				Email:     "",
				Password:  "",
				Role:      "",
			},
			ExpectedErr: errors.New("error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			tc.Mock()
			result, err := s.userRepo.Delete(tc.User.ID.String())
			s.Equal(tc.ExpectedErr, err)
			s.Equal(tc.ExpectedReturn, result)

			if err := s.mock.ExpectationsWereMet(); err != nil {
				s.T().Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func (s *suiteUser) TestCreate() {

	testCase := []struct {
		Name           string
		User           models.User
		Mock           func()
		ExpectedReturn models.User
		ExpectedErr    error
	}{
		{
			Name: "success",
			User: models.User{
				FirstName: "Jace",
				LastName:  "Herondale",
				Email:     "test@gmail.com",
				Password:  "test",
				Role:      "admin",
			},
			Mock: func() {
				rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "password", "role"}).
					AddRow("00000000-0000-0000-0000-000000000000", "Jace", "Herondale", "test@gmail.com", "test", "admin")

				s.mock.ExpectBegin()
				s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users" ("first_name","last_name","email","password","role","id") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnError(nil).WillReturnRows(rows)
				s.mock.ExpectCommit()
			},
			ExpectedReturn: models.User{
				FirstName: "Jace",
				LastName:  "Herondale",
				Email:     "test@gmail.com",
				Password:  "test",
				Role:      "admin",
			},
			ExpectedErr: nil,
		},
		{
			Name: "failed",
			User: models.User{
				FirstName: "Jace",
				LastName:  "Herondale",
				Email:     "test@gmail.com",
				Password:  "test",
				Role:      "admin",
			},
			Mock: func() {
				rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "password", "role"}).
					AddRow("00000000-0000-0000-0000-000000000000", "Jace", "Herondale", "test@gmail.com", "test", "admin")

				s.mock.ExpectBegin()
				s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users" ("first_name","last_name","email","password","role","id") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnError(errors.New("error")).WillReturnRows(rows)
				s.mock.ExpectRollback()
			},
			ExpectedErr: errors.New("error"),
			ExpectedReturn: models.User{
				FirstName: "Jace",
				LastName:  "Herondale",
				Email:     "test@gmail.com",
				Password:  "test",
				Role:      "admin",
			},
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			tc.Mock()
			result, err := s.userRepo.Create(tc.User)

			s.Equal(tc.ExpectedErr, err)
			s.Equal(tc.ExpectedReturn.Email, result.Email)

			if err := s.mock.ExpectationsWereMet(); err != nil {
				s.T().Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(suiteUser))
}
