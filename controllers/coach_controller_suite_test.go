package controllers

import (
	"Mini-Project_Coaching-Clinic/dto"
	"Mini-Project_Coaching-Clinic/models"
	"Mini-Project_Coaching-Clinic/services/mock"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type suiteCoach struct {
	suite.Suite
	controller CoachController
	mock       *mock.MockCoachService
}

func (s *suiteCoach) SetupSuite() {
	s.mock = new(mock.MockCoachService)
	s.controller = NewCoachController(s.mock)
}

func (s *suiteCoach) TestCreateCoach() {

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
			Name:               "Create coach success",
			ExpectedStatusCode: 200,
			Method:             "POST",
			HasReturnBody:      true,
			RequestBody:        `{"position":"support","price":50000,"game_id":1,"user_id":"1"}`,
			ExpectedBody: dto.BaseResponse{
				Message: "success create coach",
				Data:    map[string]interface{}{"position": "support", "price": float64(50000)},
				Status:  http.StatusOK,
				Errors:  nil,
			},
			FunctionError: nil,
		},
		{
			Name:               "Create coach bad request",
			ExpectedStatusCode: 400,
			Method:             "POST",
			HasReturnBody:      true,
			RequestBody:        `{"position":support,"price":50000}`,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Data:    dto.EmptyObj{},
				Status:  http.StatusBadRequest,
				Errors:  "invalid character 's' looking for beginning of object key string",
			},
			FunctionError: errors.New("error"),
		},
		{
			Name:               "Create coach internal server error",
			ExpectedStatusCode: 500,
			Method:             "POST",
			HasReturnBody:      true,
			RequestBody:        `{"position":"support","price":50000,"game_id":1,"user_id":"1"}`,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Data:    dto.EmptyObj{},
				Status:  http.StatusInternalServerError,
				Errors:  "error",
			},
			FunctionError: errors.New("error"),
		},

		{
			Name:               "Create coach bad request",
			ExpectedStatusCode: 400,
			Method:             "POST",
			HasReturnBody:      true,
			RequestBody:        `{"position":"support","price":50000,"game_id":"1","user":{"first_name":"te","last_name":"t"}}`,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Data:    dto.EmptyObj{},
				Status:  http.StatusBadRequest,
				Errors:  errors.New("error"),
			},
			FunctionError: errors.New("error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mock.On("Create", models.Coach{Position: "support", Price: 50000, GameID: 1, UserID: "1"}).Return(dto.CoachResponse{Position: "support", Price: 50000}, tc.FunctionError).Once()
			req := httptest.NewRequest(tc.Method, "/", strings.NewReader(tc.RequestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			res := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, res)

			err := s.controller.CreateCoach(c)
			s.NoError(err)

			s.Equal(tc.ExpectedStatusCode, res.Code)

			if tc.FunctionError == nil {
				if tc.HasReturnBody {
					var bodyResponse dto.BaseResponse
					s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
					s.Equal(tc.ExpectedBody.Data.(map[string]interface{})["position"], bodyResponse.Data.(map[string]interface{})["position"])
					s.Equal(tc.ExpectedBody.Data.(map[string]interface{})["price"], bodyResponse.Data.(map[string]interface{})["price"])
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

func (s *suiteCoach) TestUpdateCoach() {

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
			Name:               "Update coach success",
			ExpectedStatusCode: 200,
			Method:             "PUT",
			HasReturnBody:      true,
			RequestBody:        `{"position":"support","price":50000,"game_id":1,"user_id":"1"}`,
			ExpectedBody: dto.BaseResponse{
				Message: "success update coach",
				Data:    map[string]interface{}{"position": "support", "price": float64(50000)},
				Status:  http.StatusOK,
				Errors:  nil,
			},
			FunctionError: nil,
			HasParams:     true,
			ParamID:       "1",
		},
		{
			Name:               "Update coach bad request id empty",
			ExpectedStatusCode: 400,
			Method:             "PUT",
			HasReturnBody:      true,
			RequestBody:        `{"position":"support","price":50000}`,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Data:    dto.EmptyObj{},
				Status:  http.StatusBadRequest,
				Errors:  "invalid character 's' looking for beginning of object key string",
			},
			FunctionError: errors.New("error"),
			HasParams:     false,
		},
		{
			Name:               "Update coach internal server error",
			ExpectedStatusCode: 500,
			Method:             "PUT",
			HasReturnBody:      true,
			RequestBody:        `{"position":"support","price":50000,"game_id":1,"user_id":"1"}`,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Data:    dto.EmptyObj{},
				Status:  http.StatusInternalServerError,
				Errors:  "error",
			},
			FunctionError: errors.New("error"),
			HasParams:     true,
			ParamID:       "1",
		},
		{
			Name:               "Update coach bad request error cause bind error",
			ExpectedStatusCode: 400,
			Method:             "PUT",
			HasReturnBody:      true,
			RequestBody:        `{"position":support,"price":50000,"game_id":1,"user_id":"1"}`,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Data:    dto.EmptyObj{},
				Status:  http.StatusBadRequest,
				Errors:  errors.New("error"),
			},
			FunctionError: errors.New("error"),
			HasParams:     true,
			ParamID:       "a",
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mock.On("Update", models.Coach{Position: "support", Price: 50000, GameID: 1, UserID: "1"}, "1").Return(dto.CoachResponse{Position: "support", Price: 50000}, tc.FunctionError).Once()
			req := httptest.NewRequest(tc.Method, "/", strings.NewReader(tc.RequestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			res := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, res)

			if tc.HasParams {
				c.SetParamNames("id")
				c.SetParamValues(tc.ParamID)
			}

			err := s.controller.UpdateCoach(c)
			s.NoError(err)

			s.Equal(tc.ExpectedStatusCode, res.Code)

			if tc.FunctionError == nil {
				if tc.HasReturnBody {
					var bodyResponse dto.BaseResponse
					s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
					s.Equal(tc.ExpectedBody.Data.(map[string]interface{})["position"], bodyResponse.Data.(map[string]interface{})["position"])
					s.Equal(tc.ExpectedBody.Data.(map[string]interface{})["price"], bodyResponse.Data.(map[string]interface{})["price"])
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

func (s *suiteCoach) TestDeleteCoach() {

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
			Name:               "Delete coach success",
			ExpectedStatusCode: 200,
			Method:             "DELETE",
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "success delete coach",
				Data:    dto.EmptyObj{},
				Status:  http.StatusOK,
				Errors:  nil,
			},
			FunctionError: nil,
			HasParams:     true,
			ParamID:       "1",
		},
		{
			Name:               "Delete coach internal server error",
			ExpectedStatusCode: 500,
			Method:             "DELETE",
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Data:    dto.EmptyObj{},
				Status:  http.StatusInternalServerError,
				Errors:  "error",
			},
			FunctionError: errors.New("error"),
			HasParams:     true,
			ParamID:       "1",
		},
		{
			Name:               "Delete coach bad request error cause bind error",
			ExpectedStatusCode: 400,
			Method:             "DELETE",
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Data:    dto.EmptyObj{},
				Status:  http.StatusBadRequest,
				Errors:  errors.New("error"),
			},
			FunctionError: errors.New("error"),
			HasParams:     false,
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mock.On("Delete", "1").Return(models.Coach{}, tc.FunctionError).Once()
			req := httptest.NewRequest(tc.Method, "/", strings.NewReader(tc.RequestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			res := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, res)

			if tc.HasParams {
				c.SetParamNames("id")
				c.SetParamValues(tc.ParamID)
			}

			err := s.controller.DeleteCoach(c)
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

func (s *suiteCoach) TestGetCoachByID() {
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
			Name:               "Get coach by id success",
			ExpectedStatusCode: 200,
			Method:             "GET",
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "success get coach by id",
				Data:    dto.CoachResponse{Position: "support", Price: 50000},
				Status:  http.StatusOK,
				Errors:  nil,
			},
			FunctionError: nil,
			HasParams:     true,
			ParamID:       "1",
		},
		{
			Name:               "Get coach by id internal server error",
			ExpectedStatusCode: 500,
			Method:             "GET",
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Data:    dto.EmptyObj{},
				Status:  http.StatusInternalServerError,
				Errors:  "error",
			},
			FunctionError: errors.New("error"),
			HasParams:     true,
			ParamID:       "1",
		},
		{
			Name:               "Get coach by id bad request error cause id empty",
			ExpectedStatusCode: 400,
			Method:             "GET",
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Data:    dto.EmptyObj{},
				Status:  http.StatusBadRequest,
				Errors:  errors.New("error"),
			},
			FunctionError: errors.New("error"),
			HasParams:     false,
		},


	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mock.On("FindByID", "1").Return(dto.CoachResponse{Position: "support", Price: 50000}, tc.FunctionError).Once()
			req := httptest.NewRequest(tc.Method, "/", strings.NewReader(tc.RequestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			res := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, res)

			if tc.HasParams {
				c.SetParamNames("id")
				c.SetParamValues(tc.ParamID)
			}

			err := s.controller.FindCoachByID(c)
			s.NoError(err)

			s.Equal(tc.ExpectedStatusCode, res.Code)

			if tc.FunctionError == nil {
				if tc.HasReturnBody {
					var bodyResponse dto.BaseResponse
					s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
					s.Equal(tc.ExpectedBody.Message, bodyResponse.Message)
					s.Equal(tc.ExpectedBody.Data.(dto.CoachResponse).Position, bodyResponse.Data.(map[string]interface{})["position"])
					s.Equal(float64(tc.ExpectedBody.Data.(dto.CoachResponse).Price), bodyResponse.Data.(map[string]interface{})["price"])
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

func (s *suiteCoach) TestGetCoachByGameID() {

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
			Name:               "Get coach by game id success",
			ExpectedStatusCode: 200,
			Method:             "GET",
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "success get coach by game id",
				Data:    []dto.CoachResponse{{Position: "support", Price: 50000}},
				Status:  http.StatusOK,
				Errors:  nil,
			},
			FunctionError: nil,
			HasParams:     true,
			ParamID:       "1",
		},
		{
			Name:               "Get coach by game id internal server error",
			ExpectedStatusCode: 500,
			Method:             "GET",
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Data:    dto.EmptyObj{},
				Status:  http.StatusInternalServerError,
				Errors:  "error",
			},
			FunctionError: errors.New("error"),
			HasParams:     true,
			ParamID:       "1",
		},
		{
			Name:               "Get coach by game id bad request error cause id empty",
			ExpectedStatusCode: 400,
			Method:             "GET",
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Data:    dto.EmptyObj{},
				Status:  http.StatusBadRequest,
				Errors:  errors.New("error"),
			},
			FunctionError: errors.New("error"),
			HasParams:     false,
		},

	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mock.On("FindByGameID", "1").Return([]dto.CoachResponse{{Position: "support", Price: 50000}}, tc.FunctionError).Once()
			req := httptest.NewRequest(tc.Method, "/", strings.NewReader(tc.RequestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			res := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, res)

			if tc.HasParams {
				c.SetParamNames("id")
				c.SetParamValues(tc.ParamID)
			}

			err := s.controller.FindCoachByGameID(c)
			s.NoError(err)

			s.Equal(tc.ExpectedStatusCode, res.Code)

			if tc.FunctionError == nil {
				if tc.HasReturnBody {
					var bodyResponse dto.BaseResponse
					s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
					s.Equal(tc.ExpectedBody.Message, bodyResponse.Message)
					s.Equal(tc.ExpectedBody.Data.([]dto.CoachResponse)[0].Position, bodyResponse.Data.([]interface{})[0].(map[string]interface{})["position"])
					s.Equal(float64(tc.ExpectedBody.Data.([]dto.CoachResponse)[0].Price), bodyResponse.Data.([]interface{})[0].(map[string]interface{})["price"])
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

func (s *suiteCoach) TestGetCoachByCode() {
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
			Name:               "Get coach by code success",
			ExpectedStatusCode: 200,
			Method:             "GET",
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "success get coach by code",
				Data:    dto.CoachResponse{Position: "support", Price: 50000},
				Status:  http.StatusOK,
				Errors:  nil,
			},
			FunctionError: nil,
			HasParams:     true,
			ParamID:       "J4CEU",
		},
		{
			Name:               "Get coach by code internal server error",
			ExpectedStatusCode: 500,
			Method:             "GET",
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Data:    dto.EmptyObj{},
				Status:  http.StatusInternalServerError,
				Errors:  "error",
			},
			FunctionError: errors.New("error"),
			HasParams:     true,
			ParamID:       "J4CEU",
		},
		{
			Name:               "Get coach by code bad request error cause id empty",
			ExpectedStatusCode: 400,
			Method:             "GET",
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Data:    dto.EmptyObj{},
				Status:  http.StatusBadRequest,
				Errors:  errors.New("error"),
			},
			FunctionError: errors.New("error"),
			HasParams:     false,
		},


	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mock.On("FindByCode", "J4CEU").Return(dto.CoachResponse{Position: "support", Price: 50000}, tc.FunctionError).Once()
			req := httptest.NewRequest(tc.Method, "/", strings.NewReader(tc.RequestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			res := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, res)

			if tc.HasParams {
				c.SetParamNames("code")
				c.SetParamValues(tc.ParamID)
			}

			err := s.controller.FindCoachByCode(c)
			s.NoError(err)

			s.Equal(tc.ExpectedStatusCode, res.Code)

			if tc.FunctionError == nil {
				if tc.HasReturnBody {
					var bodyResponse dto.BaseResponse
					s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
					s.Equal(tc.ExpectedBody.Message, bodyResponse.Message)
					s.Equal(tc.ExpectedBody.Data.(dto.CoachResponse).Position, bodyResponse.Data.(map[string]interface{})["position"])
					s.Equal(float64(tc.ExpectedBody.Data.(dto.CoachResponse).Price), bodyResponse.Data.(map[string]interface{})["price"])
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
	

func TestCoachSuite(t *testing.T) {
	suite.Run(t, new(suiteCoach))
}
