package controllers

import (
	"Mini-Project_Coaching-Clinic/dto"
	"Mini-Project_Coaching-Clinic/models"
	"Mini-Project_Coaching-Clinic/services/mock"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type suiteUser struct {
	suite.Suite
	controller UserController
	mock       *mock.MockUserService
}

func (s *suiteUser) SetupSuite() {
	s.mock = new(mock.MockUserService)
	s.controller = NewUserController(s.mock)
}

func (s *suiteUser) TestCreateUser() {

	testCase := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		HasReturnBody      bool
		RequestBody        string
		ExpectedBody       dto.BaseResponse
		FunctionError      error
	}{
		{
			Name:               "Create User Success",
			ExpectedStatusCode: 200,
			Method:             "POST",
			HasReturnBody:      true,
			RequestBody:        `{"email":"test@gmail.com", "password":"123456", "firstname":"test", "lastname":"test"}`,
			ExpectedBody: dto.BaseResponse{
				Message: "success",
				Data:    map[string]interface{}{"email": "test@gmail.com", "password": "123456", "firstname": "test", "lastname": "test"},
				Status:  http.StatusOK,
				Errors:  nil,
			},
			FunctionError: nil,
		},
		{
			Name:               "Create User Error Bad Request",
			ExpectedStatusCode: 400,
			Method:             "POST",
			HasReturnBody:      true,
			RequestBody:        `{"email":"test@gmail.com", "password":123456, "firstname":"test", "lastname":"test"}`,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Data:    dto.EmptyObj{},
				Status:  http.StatusBadRequest,
				Errors: errors.New("error"),
			},
			FunctionError: errors.New("Status Bad Request"),
		},
		{
			Name:               "Create User Error Internal Server Error",
			ExpectedStatusCode: 500,
			Method:             "POST",
			HasReturnBody:      true,
			RequestBody:        `{"email":"test@gmail.com", "password":"123456", "firstname":"test", "lastname":"test"}`,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Data:    dto.EmptyObj{},
				Status:  http.StatusInternalServerError,
				Errors: errors.New("error"),
			},
			FunctionError: errors.New("Internal Server Error"),
		},
	}

	for _, tc := range testCase {
		s.T().Run(tc.Name, func(t *testing.T) {
			s.mock.On("Create", models.User{Email: "test@gmail.com", Password: "123456", FirstName: "test", LastName: "test"}).Return(models.User{Email: "test@gmail.com", Password: "123456", FirstName: "test", LastName: "test"}, tc.FunctionError).Once()

			req := httptest.NewRequest(tc.Method, "/", strings.NewReader(tc.RequestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			res := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, res)

			fmt.Println(tc.FunctionError)
			err := s.controller.CreateUser(c)
			s.NoError(err)

			s.Equal(tc.ExpectedStatusCode, res.Code)

			if tc.FunctionError == nil {
				if tc.HasReturnBody {
					var bodyResponse dto.BaseResponse
					s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
					s.Equal(tc.ExpectedBody.Data.(map[string]interface{})["email"], bodyResponse.Data.(map[string]interface{})["email"])
					s.Equal(tc.ExpectedBody.Data.(map[string]interface{})["firstname"], bodyResponse.Data.(map[string]interface{})["firstname"])
					s.Equal(tc.ExpectedBody.Data.(map[string]interface{})["lastname"], bodyResponse.Data.(map[string]interface{})["lastname"])
				}
			} else {
				if tc.HasReturnBody {
					var bodyResponse dto.BaseResponse
					s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
					s.Equal(tc.ExpectedBody.Message, bodyResponse.Message)

				}
			}
		})

	}

}

func (s *suiteUser) TestUpdateUser() {

	testCase := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		HasReturnBody      bool
		RequestBody        string
		ExpectedBody       dto.BaseResponse
		FunctionError      error
		HasParams          bool
		ParamID            string
	}{
		{
			Name:               "Update User Success",
			ExpectedStatusCode: 200,
			Method:             "PUT",
			HasReturnBody:      true,
			RequestBody:        `{"password":"123456", "firstname":"test", "lastname":"test"}`,
			ExpectedBody: dto.BaseResponse{
				Message: "success",
				Data:    map[string]interface{}{"email": "test@gmail.com", "password": "123456", "firstname": "test", "lastname": "test"},
				Status:  http.StatusOK,
				Errors:  nil,
			},
			FunctionError: nil,
			HasParams:     true,
			ParamID:       "1",
		},
		{
			Name:               "Update User Error Bad Request cause id not found",
			ExpectedStatusCode: 400,
			Method:             "PUT",
			HasReturnBody:      true,
			RequestBody:        `{"password":"123456", "firstname":"test", "lastname":"test"}`,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Data:    dto.EmptyObj{},
				Status:  http.StatusBadRequest,
			},
			FunctionError: errors.New("Status Bad Request"),
			HasParams:     false,
			ParamID:       "1",
		},
		{
			Name:               "Update User Error Internal Server Error",
			ExpectedStatusCode: 500,
			Method:             "PUT",
			HasReturnBody:      true,
			RequestBody:        `{"password":"123456", "firstname":"test", "lastname":"test"}`,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Data:    dto.EmptyObj{},
				Status:  http.StatusInternalServerError,
				Errors:  "Internal Server Error",
			},
			FunctionError: errors.New("Internal Server Error"),
			HasParams:     true,
			ParamID:       "1",
		},
		{
			Name:               "Update User Error Bad Request cause binding error",
			ExpectedStatusCode: 400,
			Method:             "PUT",
			HasReturnBody:      true,
			RequestBody:        `{"password":123456, "firstname":"test", "lastname":"test"}`,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Data:    dto.EmptyObj{},
				Status:  http.StatusBadRequest,
				Errors: errors.New("error"),
			},
			FunctionError: errors.New("Status Bad Request"),
			HasParams:     false,
			ParamID:       "1",
		},
		{
			Name:               "Update User Error Unauthorized",
			ExpectedStatusCode: 401,
			Method:             "PUT",
			HasReturnBody:      true,
			RequestBody:        `{"password":"123456", "firstname":"test", "lastname":"test"}`,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Data:    dto.EmptyObj{},
				Status:  http.StatusUnauthorized,
				Errors: errors.New("error"),
			},
			FunctionError: errors.New("Status Unauthorized"),
			HasParams:     true,
			ParamID:       "2",
		},
	}

	for _, tc := range testCase {
		s.T().Run(tc.Name, func(t *testing.T) {
			s.mock.On("Update", models.User{Password: "123456", FirstName: "test", LastName: "test"}, tc.ParamID).Return(models.User{Password: "123456", FirstName: "test", LastName: "test"}, tc.FunctionError).Once()
			req := httptest.NewRequest(tc.Method, "/", strings.NewReader(tc.RequestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			res := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, res)

			token := new(jwt.Token)
			token.Claims = jwt.MapClaims{
				"user_id": "1",
				"role":    "User",
			}
			c.Set("token", token)

			if tc.HasParams {
				c.SetParamNames("id")
				c.SetParamValues(tc.ParamID)
			}

			err := s.controller.UpdateUser(c)
			s.NoError(err)

			s.Equal(tc.ExpectedStatusCode, res.Code)

			if tc.FunctionError == nil {
				if tc.HasReturnBody {
					var bodyResponse dto.BaseResponse
					s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
					s.Equal(tc.ExpectedBody.Data.(map[string]interface{})["firstname"], bodyResponse.Data.(map[string]interface{})["firstname"])
					s.Equal(tc.ExpectedBody.Data.(map[string]interface{})["lastname"], bodyResponse.Data.(map[string]interface{})["lastname"])
				}
			} else {
				if tc.HasReturnBody {
					var bodyResponse dto.BaseResponse
					s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
					s.Equal(tc.ExpectedBody.Message, bodyResponse.Message)

				}
			}
		})
	}

}

func (s *suiteUser) TestDeleteUser() {

	testCase := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		HasReturnBody      bool
		RequestBody        string
		ExpectedBody       dto.BaseResponse
		FunctionError      error
		HasParams          bool
		ParamID            string
	}{
		{
			Name:               "Delete User Success",
			ExpectedStatusCode: 200,
			Method:             "DELETE",
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "success",
				Data:    dto.EmptyObj{},
				Status:  http.StatusOK,
			},
			FunctionError: nil,
			HasParams:     true,
			ParamID:       "1",
		},
		{
			Name:               "Delete User Bad Request",
			ExpectedStatusCode: 400,
			Method:             "DELETE",
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Data:    dto.EmptyObj{},
				Status:  http.StatusBadRequest,
				Errors: errors.New("error"),
			},
			FunctionError: errors.New("Status Bad Request"),
			HasParams:     false,
			ParamID:       "1",
		},
		{
			Name:               "Delete User Error Unauthorized",
			ExpectedStatusCode: 401,
			Method:             "DELETE",
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Data:    dto.EmptyObj{},
				Status:  http.StatusUnauthorized,
				Errors: errors.New("error"),
			},
			FunctionError: errors.New("Status Unauthorized"),
			HasParams:     true,
			ParamID:       "2",
		},
		{
			Name:               "Delete User Internal Server Error",
			ExpectedStatusCode: 500,
			Method:             "DELETE",
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Data:    dto.EmptyObj{},
				Status:  http.StatusInternalServerError,
				Errors: errors.New("error"),
			},
			FunctionError: errors.New("Status Internal Server Error"),
			HasParams:     true,
			ParamID:       "1",
		},
	}

	for _, tc := range testCase {
		s.T().Run(tc.Name, func(t *testing.T) {
			s.mock.On("Delete", tc.ParamID).Return(models.User{Password: "123456", FirstName: "test", LastName: "test"}, tc.FunctionError).Once()
			req := httptest.NewRequest(tc.Method, "/", strings.NewReader(tc.RequestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			res := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, res)

			token := new(jwt.Token)
			token.Claims = jwt.MapClaims{
				"user_id": "1",
				"role":    "User",
			}
			c.Set("token", token)

			if tc.HasParams {
				c.SetParamNames("id")
				c.SetParamValues(tc.ParamID)
			}

			err := s.controller.DeleteUser(c)
			s.NoError(err)

			s.Equal(tc.ExpectedStatusCode, res.Code)

			if tc.FunctionError == nil {
				if tc.HasReturnBody {
					var bodyResponse dto.BaseResponse
					s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
					s.Equal(tc.ExpectedBody.Message, bodyResponse.Message)
				}
			} else {
				if tc.HasReturnBody {
					var bodyResponse dto.BaseResponse
					s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
					s.Equal(tc.ExpectedBody.Message, bodyResponse.Message)
				}
			}
		})
	}
}

func (s *suiteUser) TestGetAllUser() {

	testCase := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		HasReturnBody      bool
		RequestBody        string
		ExpectedBody       dto.BaseResponse
		FunctionError      error
	}{
		{
			Name:               "Get All User Success",
			ExpectedStatusCode: 200,
			Method:             "GET",
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "success",
				Data:    map[string]interface{}{"email": "test@gmail.com", "password": "123456", "firstname": "test", "lastname": "test"},
				Status:  http.StatusOK,
			},
			FunctionError: nil,
		},
		{
			Name:               "Get All User Internal Server Error",
			ExpectedStatusCode: 500,
			Method:             "GET",
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Data:    dto.EmptyObj{},
				Status:  http.StatusInternalServerError,
				Errors: errors.New("error"),
			},
			FunctionError: errors.New("Status Internal Server Error"),
		},
	}

	for _, tc := range testCase {
		s.T().Run(tc.Name, func(t *testing.T) {
			s.mock.On("FindAll").Return([]models.User{{Email: "test@gmail.com", Password: "123456", FirstName: "test", LastName: "test"}}, tc.FunctionError).Once()
			req := httptest.NewRequest(tc.Method, "/", strings.NewReader(tc.RequestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			res := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, res)

			err := s.controller.GetAllUser(c)
			s.NoError(err)

			s.Equal(tc.ExpectedStatusCode, res.Code)

			if tc.FunctionError == nil {
				if tc.HasReturnBody {
					var bodyResponse dto.BaseResponse
					s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
					s.Equal(tc.ExpectedBody.Message, bodyResponse.Message)
					s.Equal(tc.ExpectedBody.Data.(map[string]interface{})["firstname"], bodyResponse.Data.([]interface{})[0].(map[string]interface{})["firstname"])
					s.Equal(tc.ExpectedBody.Data.(map[string]interface{})["lastname"], bodyResponse.Data.([]interface{})[0].(map[string]interface{})["lastname"])

				}
			} else {
				if tc.HasReturnBody {
					var bodyResponse dto.BaseResponse
					s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
					s.Equal(tc.ExpectedBody.Message, bodyResponse.Message)
				}
			}
		})
	}
}

func (s *suiteUser) TestGetUserByID() {

	testCase := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		HasReturnBody      bool
		RequestBody        string
		ExpectedBody       dto.BaseResponse
		FunctionError      error
		HasParams          bool
		ParamID            string
	}{
		{
			Name:               "Get User By ID Success",
			ExpectedStatusCode: 200,
			Method:             "GET",
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "success",
				Data:    map[string]interface{}{"email": "test@gmail.com", "password": "123456", "firstname": "test", "lastname": "test"},
				Status:  http.StatusOK,
			},
			FunctionError: nil,
			HasParams:     true,
			ParamID:       "1",
		},
		{
			Name:               "Get User Bad Request",
			ExpectedStatusCode: 400,
			Method:             "GET",
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Data:    dto.EmptyObj{},
				Status:  http.StatusBadRequest,
				Errors: errors.New("error"),
			},
			FunctionError: errors.New("Status Bad Request"),
			HasParams:     true,
			ParamID:       "",
		},
		{
			Name:               "Get User Internal Server Error",
			ExpectedStatusCode: 500,
			Method:             "GET",
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Data:    dto.EmptyObj{},
				Status:  http.StatusInternalServerError,
				Errors: errors.New("error"),
			},
			FunctionError: errors.New("Status Internal Server Error"),
			HasParams:     true,
			ParamID:       "1",
		},
	}

	for _, tc := range testCase {
		s.T().Run(tc.Name, func(t *testing.T) {
			s.mock.On("FindByID", tc.ParamID).Return(models.User{Email: "test@gmail.com", Password: "123456", FirstName: "test", LastName: "test"}, tc.FunctionError).Once()
			req := httptest.NewRequest(tc.Method, "/", strings.NewReader(tc.RequestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			res := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, res)
			c.SetParamNames("id")
			c.SetParamValues(tc.ParamID)

			err := s.controller.GetDetailUser(c)
			s.NoError(err)

			s.Equal(tc.ExpectedStatusCode, res.Code)

			if tc.FunctionError == nil {
				if tc.HasReturnBody {
					var bodyResponse dto.BaseResponse
					s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
					s.Equal(tc.ExpectedBody.Message, bodyResponse.Message)
					s.Equal(tc.ExpectedBody.Data.(map[string]interface{})["firstname"], bodyResponse.Data.(map[string]interface{})["firstname"])
					s.Equal(tc.ExpectedBody.Data.(map[string]interface{})["lastname"], bodyResponse.Data.(map[string]interface{})["lastname"])
				}
			} else {
				if tc.HasReturnBody {
					var bodyResponse dto.BaseResponse
					s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
					s.Equal(tc.ExpectedBody.Message, bodyResponse.Message)
				}
			}
		})
	}
}

func (s *suiteUser) TestLoginUser() {
	testCase := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		HasReturnBody      bool
		RequestBody        string
		ExpectedBody       dto.BaseResponse
		FunctionError      error
	}{
		{
			Name:               "Login User Success",
			ExpectedStatusCode: 200,
			Method:             "POST",
			HasReturnBody:      true,
			RequestBody:        `{"email":"test@gmail.com", "password":"123456"}`,
			ExpectedBody: dto.BaseResponse{
				Message: "success",
				Data:    map[string]interface{}{"email": "test@gmail.com", "token": "123"},
				Status:  http.StatusOK,
			},
			FunctionError: nil,
		},
		{
			Name:               "Login User Bad Request",
			ExpectedStatusCode: 400,
			Method:             "POST",
			HasReturnBody:      true,
			RequestBody:        `{"email":asd, "password":"123456"}`,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Data:    dto.EmptyObj{},
				Status:  http.StatusBadRequest,
				Errors: errors.New("error"),
			},
			FunctionError: errors.New("Status Bad Request"),
		},
		{
			Name:               "Login User Internal Server Error",
			ExpectedStatusCode: 500,
			Method:             "POST",
			HasReturnBody:      true,
			RequestBody:        `{"email":"test@gmail.com", "password":"123456"}`,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Data:    dto.EmptyObj{},
				Status:  http.StatusInternalServerError,
				Errors: errors.New("error"),
			},
			FunctionError: errors.New("Status Internal Server Error"),
		},
	}

	for _, tc := range testCase {
		s.T().Run(tc.Name, func(t *testing.T) {
			s.mock.On("LoginUser", "test@gmail.com", "123456").Return(dto.UserResponse{Email: "test@gmail.com", Token: "123"}, tc.FunctionError).Once()
			req := httptest.NewRequest(tc.Method, "/", strings.NewReader(tc.RequestBody))

			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			res := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, res)

			err := s.controller.LoginUser(c)
			s.NoError(err)

			s.Equal(tc.ExpectedStatusCode, res.Code)

			if tc.FunctionError == nil {
				if tc.HasReturnBody {
					var bodyResponse dto.BaseResponse
					s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
					s.Equal(tc.ExpectedBody.Message, bodyResponse.Message)
					s.Equal(tc.ExpectedBody.Data.(map[string]interface{})["email"], bodyResponse.Data.(map[string]interface{})["email"])
					s.Equal(tc.ExpectedBody.Data.(map[string]interface{})["token"], bodyResponse.Data.(map[string]interface{})["token"])

				}
			} else {
				if tc.HasReturnBody {
					var bodyResponse dto.BaseResponse
					s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
					s.Equal(tc.ExpectedBody.Message, bodyResponse.Message)
				}
			}

		})
	}
}

func TestUserSuite(t *testing.T) {
	suite.Run(t, new(suiteUser))
}
