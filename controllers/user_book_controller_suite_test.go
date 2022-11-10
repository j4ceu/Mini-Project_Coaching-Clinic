package controllers

import (
	"Mini-Project_Coaching-Clinic/dto"
	"Mini-Project_Coaching-Clinic/dto/payload"
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

type suiteUserBook struct {
	suite.Suite
	controller      UserBookController
	mock            *mock.UserBookServiceMock
	mockUserPayment *mock.UserPaymentServiceMock
}

func (s *suiteUserBook) SetupSuite() {
	s.mock = new(mock.UserBookServiceMock)
	s.mockUserPayment = new(mock.UserPaymentServiceMock)
	s.controller = NewUserBookController(s.mock, s.mockUserPayment)
}

func (s *suiteUserBook) TestCreateUserBook() {

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
			Name:               "success create user book",
			ExpectedStatusCode: http.StatusOK,
			Method:             http.MethodPost,
			HasReturnBody:      true,
			RequestBody:        `{"title":"Coaching Clinic", "coach_availability_id":"1", "user_payment_id":"1"}`,
			ExpectedBody: dto.BaseResponse{
				Message: "success create user book",
				Status:  http.StatusOK,
				Data:    dto.UserBookResponse{Title: "Coaching Clinic", UserPaymentID: "1", CoachAvailabilityID: "1"},
			},
			FunctionError: nil,
		},
		{
			Name:               "failed create user book - bad request",
			ExpectedStatusCode: http.StatusBadRequest,
			Method:             http.MethodPost,
			HasReturnBody:      true,
			RequestBody:        `{"title":"Coaching Clinic", "coach_availability_id":1, "user_payment_id":1}`,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Status:  http.StatusBadRequest,
				Data:    dto.EmptyObj{},
				Errors:  errors.New("error"),
			},
			FunctionError: errors.New("error"),
		},
		{
			Name:               "failed create user book - internal server error",
			ExpectedStatusCode: http.StatusInternalServerError,
			Method:             http.MethodPost,
			HasReturnBody:      true,
			RequestBody:        `{"title":"Coaching Clinic", "coach_availability_id":"1", "user_payment_id":"1"}`,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Status:  http.StatusInternalServerError,
				Data:    dto.EmptyObj{},
				Errors:  errors.New("error"),
			},
			FunctionError: errors.New("error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mock.On("Create", payload.UserBookPayloadCreate{Title: "Coaching Clinic", UserPaymentID: "1", CoachAvailabilityID: "1"}).Return(dto.UserBookResponse{Title: "Coaching Clinic", UserPaymentID: "1", CoachAvailabilityID: "1"}, tc.FunctionError).Once()
			req := httptest.NewRequest(tc.Method, "/", strings.NewReader(tc.RequestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			res := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, res)

			err := s.controller.CreateUserBook(c)
			s.NoError(err)

			s.Equal(tc.ExpectedStatusCode, res.Code)

			if tc.FunctionError == nil {
				if tc.HasReturnBody {
					var bodyResponse dto.BaseResponse
					s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
					s.Equal(tc.ExpectedBody.Data.(dto.UserBookResponse).Title, bodyResponse.Data.(map[string]interface{})["title"])
					s.Equal(tc.ExpectedBody.Data.(dto.UserBookResponse).UserPaymentID, bodyResponse.Data.(map[string]interface{})["user_payment_id"])
					s.Equal(tc.ExpectedBody.Data.(dto.UserBookResponse).CoachAvailabilityID, bodyResponse.Data.(map[string]interface{})["coach_availability_id"])
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

func (s *suiteUserBook) TestUpdateUserBook() {

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
			Name:               "failed update user book - bad request bind error",
			ExpectedStatusCode: http.StatusBadRequest,
			Method:             http.MethodPut,
			HasReturnBody:      true,
			RequestBody:        `{"title":"Coaching Clinic", "coach_availability_id":1, "user_payment_id":1}`,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Status:  http.StatusBadRequest,
				Data:    dto.EmptyObj{},
				Errors:  errors.New("error"),
			},
			FunctionError: errors.New("error"),
			HasParams:     true,
			ParamID:       "1",
		},
		{
			Name:               "failed update user book - bad request id empty",
			ExpectedStatusCode: http.StatusBadRequest,
			Method:             http.MethodPut,
			HasReturnBody:      true,
			RequestBody:        `{"title":"Coaching Clinic", "coach_availability_id":"1", "user_payment_id":"1"}`,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Status:  http.StatusBadRequest,
				Data:    dto.EmptyObj{},
				Errors:  errors.New("error"),
			},
			FunctionError: errors.New("error"),
			HasParams:     false,
			ParamID:       "",
		},
		{
			Name:               "failed update user book - bad request payload summary error",
			ExpectedStatusCode: http.StatusBadRequest,
			Method:             http.MethodPut,
			HasReturnBody:      true,
			RequestBody:        `{"summary":1}`,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Status:  http.StatusBadRequest,
				Data:    dto.EmptyObj{},
				Errors:  errors.New("error"),
			},
			FunctionError: errors.New("error"),
			HasParams:     true,
			ParamID:       "1",
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mock.On("Update", payload.UserBookPayloadUpdate{Title: "Coaching Clinic", UserPaymentID: "1", CoachAvailabilityID: "1"}).Return(dto.UserBookResponse{Title: "Coaching Clinic", UserPaymentID: "1", CoachAvailabilityID: "1"}, tc.FunctionError)

			req := httptest.NewRequest(tc.Method, "/", strings.NewReader(tc.RequestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEMultipartForm)
			res := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, res)

			if tc.HasParams {
				c.SetParamNames("id")
				c.SetParamValues(tc.ParamID)
			}

			err := s.controller.UpdateUserBook(c)
			s.NoError(err)

			s.Equal(tc.ExpectedStatusCode, res.Code)

			if tc.FunctionError == nil {
				if tc.HasReturnBody {
					var bodyResponse dto.BaseResponse
					s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
					fmt.Println(bodyResponse.Errors)
					s.Equal(tc.ExpectedBody.Data.(dto.UserBookResponse).Title, bodyResponse.Data.(map[string]interface{})["title"])
					s.Equal(tc.ExpectedBody.Data.(dto.UserBookResponse).UserPaymentID, bodyResponse.Data.(map[string]interface{})["user_payment_id"])
					s.Equal(tc.ExpectedBody.Data.(dto.UserBookResponse).CoachAvailabilityID, bodyResponse.Data.(map[string]interface{})["coach_availability_id"])
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

func (s *suiteUserBook) TestDeleteUserBook() {

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
			Name:               "success delete user book",
			ExpectedStatusCode: http.StatusOK,
			Method:             http.MethodDelete,
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "success delete user book",
				Status:  http.StatusOK,
				Data:    dto.EmptyObj{},
				Errors:  nil,
			},
			FunctionError: nil,
			HasParams:     true,
			ParamID:       "1",
		},
		{
			Name:               "failed delete user book - bad request id empty",
			ExpectedStatusCode: http.StatusBadRequest,
			Method:             http.MethodDelete,
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Status:  http.StatusBadRequest,
				Data:    dto.EmptyObj{},
				Errors:  errors.New("error"),
			},
			FunctionError: errors.New("error"),
			HasParams:     false,
			ParamID:       "",
		},
		{
			Name:               "failed delete user book - internal server error",
			ExpectedStatusCode: http.StatusInternalServerError,
			Method:             http.MethodDelete,
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Status:  http.StatusInternalServerError,
				Data:    dto.EmptyObj{},
				Errors:  errors.New("error"),
			},
			FunctionError: errors.New("error"),
			HasParams:     true,
			ParamID:       "1",
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mock.On("Delete", "1").Return(models.UserBook{}, tc.FunctionError).Once()

			req := httptest.NewRequest(tc.Method, "/", strings.NewReader(tc.RequestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEMultipartForm)
			res := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, res)

			if tc.HasParams {
				c.SetParamNames("id")
				c.SetParamValues(tc.ParamID)
			}

			err := s.controller.DeleteUserBook(c)
			s.NoError(err)

			s.Equal(tc.ExpectedStatusCode, res.Code)

			if tc.HasReturnBody {
				var bodyResponse dto.BaseResponse
				s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
				s.Equal(tc.ExpectedBody.Message, bodyResponse.Message)
			}

		})
	}
}

func (s *suiteUserBook) TestGetAllUserBook() {

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
			Name:               "success get all user book",
			ExpectedStatusCode: http.StatusOK,
			Method:             http.MethodGet,
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "success get all user book",
				Status:  http.StatusOK,
				Data: []dto.UserBookResponse{
					{Title: "Coaching Clinic", UserPaymentID: "1", CoachAvailabilityID: "1"},
				},
				Errors: nil,
			},
			FunctionError: nil,
		},
		{
			Name:               "failed get all user book - internal server error",
			ExpectedStatusCode: http.StatusInternalServerError,
			Method:             http.MethodGet,
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Status:  http.StatusInternalServerError,
				Data:    dto.EmptyObj{},
				Errors:  errors.New("error"),
			},
			FunctionError: errors.New("error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mock.On("FindAll").Return([]dto.UserBookResponse{{Title: "Coaching Clinic", UserPaymentID: "1", CoachAvailabilityID: "1"}}, tc.FunctionError).Once()

			req := httptest.NewRequest(tc.Method, "/", strings.NewReader(tc.RequestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEMultipartForm)
			res := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, res)

			err := s.controller.FindAllUserBook(c)
			s.NoError(err)

			s.Equal(tc.ExpectedStatusCode, res.Code)

			if tc.FunctionError == nil {
				if tc.HasReturnBody {
					var bodyResponse dto.BaseResponse
					s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
					s.Equal(tc.ExpectedBody.Message, bodyResponse.Message)
					s.Equal(tc.ExpectedBody.Data.([]dto.UserBookResponse)[0].Title, bodyResponse.Data.([]interface{})[0].(map[string]interface{})["title"])
					s.Equal(tc.ExpectedBody.Data.([]dto.UserBookResponse)[0].UserPaymentID, bodyResponse.Data.([]interface{})[0].(map[string]interface{})["user_payment_id"])
					s.Equal(tc.ExpectedBody.Data.([]dto.UserBookResponse)[0].CoachAvailabilityID, bodyResponse.Data.([]interface{})[0].(map[string]interface{})["coach_availability_id"])
				}
			} else {
				var bodyResponse dto.BaseResponse
				s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
				s.Equal(tc.ExpectedBody.Message, bodyResponse.Message)

			}

		})
	}
}

func (s *suiteUserBook) TestGetUserBookByID() {

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
			Name:               "success get user book by id",
			ExpectedStatusCode: http.StatusOK,
			Method:             http.MethodGet,
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "success get user book by id",
				Status:  http.StatusOK,
				Data:    dto.UserBookResponse{Title: "Coaching Clinic", UserPaymentID: "1", CoachAvailabilityID: "1"},
				Errors:  nil,
			},
			FunctionError: nil,
			HasParams:     true,
			ParamID:       "1",
		},
		{
			Name:               "failed get user book by id - internal server error",
			ExpectedStatusCode: http.StatusInternalServerError,
			Method:             http.MethodGet,
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Status:  http.StatusInternalServerError,
				Data:    dto.EmptyObj{},
				Errors:  errors.New("error"),
			},
			FunctionError: errors.New("error"),
			HasParams:     true,
			ParamID:       "1",
		},
		{
			Name:               "failed get user book by id - bad request id empty",
			ExpectedStatusCode: http.StatusBadRequest,
			Method:             http.MethodGet,
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Status:  http.StatusBadRequest,
				Data:    dto.EmptyObj{},
				Errors:  errors.New("id is empty"),
			},
			FunctionError: errors.New("id is empty"),
			HasParams:     true,
			ParamID:       "",
		},
		{
			Name:               "failed get user book by id - unauthorized",
			ExpectedStatusCode: http.StatusUnauthorized,
			Method:             http.MethodGet,
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Status:  http.StatusUnauthorized,
				Data:    dto.EmptyObj{},
				Errors:  errors.New("unauthorized"),
			},
			FunctionError: nil,
			HasParams:     true,
			ParamID:       "2",
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mock.On("FindByID", tc.ParamID).Return(dto.UserBookResponse{Title: "Coaching Clinic", UserPaymentID: "1", CoachAvailabilityID: "1"}, tc.FunctionError).Once()
			s.mockUserPayment.On("FindByID", "1").Return(dto.UserPaymentResponse{UserID: "1"}, nil).Once()

			req := httptest.NewRequest(tc.Method, "/", strings.NewReader(tc.RequestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEMultipartForm)
			res := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, res)

			token := new(jwt.Token)
			token.Claims = jwt.MapClaims{
				"user_id": tc.ParamID,
				"role":    "User",
			}
			c.Set("token", token)

			if tc.HasParams {
				c.SetParamNames("id")
				c.SetParamValues(tc.ParamID)
			}

			err := s.controller.FindUserBookById(c)
			s.NoError(err)

			s.Equal(tc.ExpectedStatusCode, res.Code)

			if tc.FunctionError == nil && tc.ExpectedStatusCode != http.StatusUnauthorized {
				if tc.HasReturnBody {
					var bodyResponse dto.BaseResponse
					s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
					s.Equal(tc.ExpectedBody.Message, bodyResponse.Message)
					s.Equal(tc.ExpectedBody.Data.(dto.UserBookResponse).Title, bodyResponse.Data.(map[string]interface{})["title"])
					s.Equal(tc.ExpectedBody.Data.(dto.UserBookResponse).UserPaymentID, bodyResponse.Data.(map[string]interface{})["user_payment_id"])
					s.Equal(tc.ExpectedBody.Data.(dto.UserBookResponse).CoachAvailabilityID, bodyResponse.Data.(map[string]interface{})["coach_availability_id"])
				}
			} else {
				var bodyResponse dto.BaseResponse
				s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
				s.Equal(tc.ExpectedBody.Message, bodyResponse.Message)

			}

		})
	}

}

func (s *suiteUserBook) TestGetUserBookByUserID() {

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
			Name:               "success get user book by user id",
			ExpectedStatusCode: http.StatusOK,
			Method:             http.MethodGet,
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "success get user book by user id",
				Status:  http.StatusOK,
				Data:    []dto.UserBookResponse{{Title: "Coaching Clinic", UserPaymentID: "1", CoachAvailabilityID: "1"}},
				Errors:  nil,
			},
			FunctionError: nil,
			HasParams:     true,
			ParamID:       "1",
		},
		{
			Name:               "failed get user book by user id - internal server error",
			ExpectedStatusCode: http.StatusInternalServerError,
			Method:             http.MethodGet,
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Status:  http.StatusInternalServerError,
				Data:    dto.EmptyObj{},
				Errors:  errors.New("error"),
			},
			FunctionError: errors.New("error"),
			HasParams:     true,
			ParamID:       "1",
		},
		{
			Name:               "failed get user book by user id - bad request id empty",
			ExpectedStatusCode: http.StatusBadRequest,
			Method:             http.MethodGet,
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Status:  http.StatusBadRequest,
				Data:    dto.EmptyObj{},
				Errors:  errors.New("id is empty"),
			},
			FunctionError: errors.New("id is empty"),
			HasParams:     true,
			ParamID:       "",
		},
		{
			Name:               "failed get user book by user id - unauthorized",
			ExpectedStatusCode: http.StatusUnauthorized,
			Method:             http.MethodGet,
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Status:  http.StatusUnauthorized,
				Data:    dto.EmptyObj{},
				Errors:  errors.New("unauthorized"),
			},
			FunctionError: errors.New("unauthorized"),
			HasParams:     true,
			ParamID:       "2",
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mock.On("FindByUserID", tc.ParamID).Return([]dto.UserBookResponse{{Title: "Coaching Clinic", UserPaymentID: "1", CoachAvailabilityID: "1"}}, tc.FunctionError).Once()

			req := httptest.NewRequest(tc.Method, "/", strings.NewReader(tc.RequestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEMultipartForm)
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

			err := s.controller.FindUserBookByUserID(c)
			s.NoError(err)

			s.Equal(tc.ExpectedStatusCode, res.Code)

			if tc.FunctionError == nil && tc.ExpectedStatusCode != http.StatusUnauthorized {
				if tc.HasReturnBody {
					var bodyResponse dto.BaseResponse
					s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
					s.Equal(tc.ExpectedBody.Message, bodyResponse.Message)
					s.Equal(tc.ExpectedBody.Data.([]dto.UserBookResponse)[0].Title, bodyResponse.Data.([]interface{})[0].(map[string]interface{})["title"])
					s.Equal(tc.ExpectedBody.Data.([]dto.UserBookResponse)[0].UserPaymentID, bodyResponse.Data.([]interface{})[0].(map[string]interface{})["user_payment_id"])
					s.Equal(tc.ExpectedBody.Data.([]dto.UserBookResponse)[0].CoachAvailabilityID, bodyResponse.Data.([]interface{})[0].(map[string]interface{})["coach_availability_id"])
				}
			} else {
				var bodyResponse dto.BaseResponse
				s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
				s.Equal(tc.ExpectedBody.Message, bodyResponse.Message)

			}

		})
	}

}

func TestUserBookSuite(t *testing.T) {
	suite.Run(t, new(suiteUserBook))
}
