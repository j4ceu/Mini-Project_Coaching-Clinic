package Game_Controller

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

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type suiteGame struct {
	suite.Suite
	controller GameController
	mock       *mock.MockGameService
}

func (s *suiteGame) SetupSuite() {
	s.mock = new(mock.MockGameService)
	s.controller = NewGameController(s.mock)
}

func (s *suiteGame) TestCreateGame() {

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
			Name:               "Create Game Success",
			ExpectedStatusCode: 200,
			Method:             "POST",
			HasReturnBody:      true,
			RequestBody:        `{"name":"Game 1","description":"Game 1 Description"}`,
			ExpectedBody: dto.BaseResponse{
				Message: "success create game",
				Data:    map[string]interface{}{"description": "Game 1 Description", "name": "Game 1"},
				Status:  http.StatusOK,
			},
			FunctionError: nil,
		},
		{
			Name:               "Create Game Internal Server Error",
			ExpectedStatusCode: 500,
			Method:             "POST",
			HasReturnBody:      true,
			RequestBody:        `{"name":"Game 1","description":"Game 1 Description"}`,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Data:    dto.EmptyObj{},
				Status:  http.StatusInternalServerError,
				Errors:  errors.New("error"),
			},
			FunctionError: errors.New("error"),
		},
		{
			Name:               "Create Game Bad Request",
			ExpectedStatusCode: 400,
			Method:             "POST",
			HasReturnBody:      true,
			RequestBody:        `{"name":Game,"description":"Game 1 Description"}`,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Data:    dto.EmptyObj{},
				Status:  http.StatusBadRequest,
				Errors:  errors.New("Bad Request"),
			},
			FunctionError: errors.New("Bad Request"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mock.On("Create", models.Game{Name: "Game 1", Description: "Game 1 Description"}).Return(models.Game{Name: "Game 1", Description: "Game 1 Description"}, tc.FunctionError).Once()

			req := httptest.NewRequest(tc.Method, "/", strings.NewReader(tc.RequestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			res := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, res)

			fmt.Println(tc.FunctionError)
			err := s.controller.CreateGame(c)
			s.NoError(err)

			s.Equal(tc.ExpectedStatusCode, res.Code)

			if tc.FunctionError == nil {
				if tc.HasReturnBody {
					var bodyResponse dto.BaseResponse
					s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
					s.Equal(tc.ExpectedBody.Data.(map[string]interface{})["name"], bodyResponse.Data.(map[string]interface{})["name"])
					s.Equal(tc.ExpectedBody.Data.(map[string]interface{})["description"], bodyResponse.Data.(map[string]interface{})["description"])
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

func (s *suiteGame) TestUpdateGame() {

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
			Name:               "Update Game Success",
			ExpectedStatusCode: 200,
			Method:             "PUT",
			HasReturnBody:      true,
			RequestBody:        `{"name":"Game 1","description":"Game 1 Description"}`,
			ExpectedBody: dto.BaseResponse{
				Message: "success update game",
				Data:    map[string]interface{}{"description": "Game 1 Description", "name": "Game 1"},
				Status:  http.StatusOK,
			},
			FunctionError: nil,
			HasParams:     true,
			ParamID:       "1",
		},
		{
			Name:               "Update Game Internal Server Error",
			ExpectedStatusCode: 500,
			Method:             "PUT",
			HasReturnBody:      true,
			RequestBody:        `{"name":"Game 1","description":"Game 1 Description"}`,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Data:    dto.EmptyObj{},
				Status:  http.StatusInternalServerError,
				Errors:  errors.New("error"),
			},
			FunctionError: errors.New("error"),
			HasParams:     true,
			ParamID:       "1",
		},
		{
			Name:               "Update Game Bad Request Cause ID Empty",
			ExpectedStatusCode: 400,
			Method:             "PUT",
			HasReturnBody:      true,
			RequestBody:        `{"name":"Game 1","description":"Game 1 Description"}`,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Data:    dto.EmptyObj{},
				Status:  http.StatusBadRequest,
				Errors:  errors.New("Bad Request"),
			},
			FunctionError: errors.New("Bad Request"),
			HasParams:     false,
			ParamID:       "",
		},
		{
			Name:               "Update Game Bad Request Cause ID Not Integer",
			ExpectedStatusCode: 400,
			Method:             "PUT",
			HasReturnBody:      true,
			RequestBody:        `{"name":"Game 1","description":"Game 1 Description"}`,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Data:    dto.EmptyObj{},
				Status:  http.StatusBadRequest,
				Errors:  errors.New("Bad Request"),
			},
			FunctionError: errors.New("Bad Request"),
			HasParams:     true,
			ParamID:       "a",
		},
		{
			Name:               "Update Game Bad Request Case Bind Error",
			ExpectedStatusCode: 400,
			Method:             "PUT",
			HasReturnBody:      true,
			RequestBody:        `{"name":Game 1,"description":"Game 1 Description"}`,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Data:    dto.EmptyObj{},
				Status:  http.StatusBadRequest,
				Errors:  errors.New("error"),
			},
			FunctionError: errors.New("error"),
			HasParams:     true,
			ParamID:       "1",
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mock.On("Update", models.Game{Name: "Game 1", Description: "Game 1 Description"}, uint(1)).Return(models.Game{Name: "Game 1", Description: "Game 1 Description"}, tc.FunctionError).Once()

			req := httptest.NewRequest(tc.Method, "/", strings.NewReader(tc.RequestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			res := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, res)

			if tc.HasParams {
				c.SetParamNames("id")
				c.SetParamValues(tc.ParamID)
			}

			err := s.controller.UpdateGame(c)
			s.NoError(err)

			s.Equal(tc.ExpectedStatusCode, res.Code)

			if tc.FunctionError == nil {
				if tc.HasReturnBody {
					var bodyResponse dto.BaseResponse
					s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
					s.Equal(tc.ExpectedBody.Data.(map[string]interface{})["name"], bodyResponse.Data.(map[string]interface{})["name"])
					s.Equal(tc.ExpectedBody.Data.(map[string]interface{})["description"], bodyResponse.Data.(map[string]interface{})["description"])
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

func (s *suiteGame) TestDeleteGame() {

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
			Name:               "Delete Game Success",
			ExpectedStatusCode: 200,
			Method:             "DELETE",
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "success delete game",
				Data:    dto.EmptyObj{},
				Status:  http.StatusOK,
				Errors:  nil,
			},
			FunctionError: nil,
			HasParams:     true,
			ParamID:       "1",
		},
		{
			Name:               "Delete Game Internal Server Error",
			ExpectedStatusCode: 500,
			Method:             "DELETE",
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Data:    dto.EmptyObj{},
				Status:  http.StatusInternalServerError,
				Errors:  errors.New("error"),
			},
			FunctionError: errors.New("error"),
			HasParams:     true,
			ParamID:       "1",
		},
		{
			Name:               "Delete Game Bad Request Cause ID Empty",
			ExpectedStatusCode: 400,
			Method:             "DELETE",
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Data:    dto.EmptyObj{},
				Status:  http.StatusBadRequest,
				Errors:  errors.New("Bad Request"),
			},
			FunctionError: errors.New("Bad Request"),
			HasParams:     false,
			ParamID:       "",
		},
		{
			Name:               "Delete Game Bad Request Cause ID Not Integer",
			ExpectedStatusCode: 400,
			Method:             "DELETE",
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Data:    dto.EmptyObj{},
				Status:  http.StatusBadRequest,
				Errors:  errors.New("Bad Request"),
			},
			FunctionError: errors.New("Bad Request"),
			HasParams:     true,
			ParamID:       "a",
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mock.On("Delete", uint(1)).Return(models.Game{Name: "Game 1", Description: "Game 1 Description"}, tc.FunctionError).Once()

			req := httptest.NewRequest(tc.Method, "/", strings.NewReader(tc.RequestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			res := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, res)

			if tc.HasParams {
				c.SetParamNames("id")
				c.SetParamValues(tc.ParamID)
			}

			err := s.controller.DeleteGame(c)
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

func (s *suiteGame) TestGetAllGame() {

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
			Name:               "Get All Game Success",
			ExpectedStatusCode: 200,
			Method:             "GET",
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "success get all game",
				Data:    map[string]interface{}{"name": "Game 1", "description": "Game 1 Description"},
				Status:  http.StatusOK,
				Errors:  nil,
			},
			FunctionError: nil,
		},
		{
			Name:               "Get All Game Internal Server Error",
			ExpectedStatusCode: 500,
			Method:             "GET",
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Data:    dto.EmptyObj{},
				Status:  http.StatusInternalServerError,
				Errors:  errors.New("error"),
			},
			FunctionError: errors.New("error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mock.On("FindAll").Return([]models.Game{{Name: "Game 1", Description: "Game 1 Description"}}, tc.FunctionError).Once()

			req := httptest.NewRequest(tc.Method, "/", strings.NewReader(tc.RequestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			res := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, res)

			err := s.controller.FindAllGame(c)
			s.NoError(err)

			s.Equal(tc.ExpectedStatusCode, res.Code)

			if tc.FunctionError == nil {
				if tc.HasReturnBody {
					var bodyResponse dto.BaseResponse
					s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
					s.Equal(tc.ExpectedBody.Message, bodyResponse.Message)
					s.Equal(tc.ExpectedBody.Data.(map[string]interface{})["name"], bodyResponse.Data.([]interface{})[0].(map[string]interface{})["name"])
					s.Equal(tc.ExpectedBody.Data.(map[string]interface{})["description"], bodyResponse.Data.([]interface{})[0].(map[string]interface{})["description"])
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

func (s *suiteGame) TestGetGameByID() {

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
			Name:               "Get Game By ID Success",
			ExpectedStatusCode: 200,
			Method:             "GET",
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "success get game",
				Data:    map[string]interface{}{"name": "Game 1", "description": "Game 1 Description"},
				Status:  http.StatusOK,
				Errors:  nil,
			},
			FunctionError: nil,
			HasParams:     true,
			ParamID:       "1",
		},
		{
			Name:               "Get Game By ID Internal Server Error",
			ExpectedStatusCode: 500,
			Method:             "GET",
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Data:    dto.EmptyObj{},
				Status:  http.StatusInternalServerError,
				Errors:  errors.New("error"),
			},
			FunctionError: errors.New("error"),
			HasParams:     true,
			ParamID:       "1",
		},
		{
			Name:               "Get Game By ID Bad Request ID Empty",
			ExpectedStatusCode: 400,
			Method:             "GET",
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Data:    dto.EmptyObj{},
				Status:  http.StatusBadRequest,
				Errors:  errors.New("id is empty"),
			},
			FunctionError: errors.New("error"),
			HasParams:     false,
			ParamID:       "",
		},
		{
			Name:               "Get Game By ID Bad Request ID Not Integer",
			ExpectedStatusCode: 400,
			Method:             "GET",
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Data:    dto.EmptyObj{},
				Status:  http.StatusBadRequest,
				Errors:  errors.New("id is not integer"),
			},
			FunctionError: errors.New("error"),
			HasParams:     true,
			ParamID:       "a",
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mock.On("FindByID", uint(1)).Return(models.Game{Name: "Game 1", Description: "Game 1 Description"}, tc.FunctionError).Once()

			req := httptest.NewRequest(tc.Method, "/", strings.NewReader(tc.RequestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			res := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, res)

			if tc.HasParams {
				c.SetParamNames("id")
				c.SetParamValues(tc.ParamID)
			}

			err := s.controller.FindGameByID(c)
			s.NoError(err)

			s.Equal(tc.ExpectedStatusCode, res.Code)

			if tc.FunctionError == nil {
				if tc.HasReturnBody {
					var bodyResponse dto.BaseResponse
					s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
					s.Equal(tc.ExpectedBody.Message, bodyResponse.Message)
					s.Equal(tc.ExpectedBody.Data.(map[string]interface{})["name"], bodyResponse.Data.(map[string]interface{})["name"])
					s.Equal(tc.ExpectedBody.Data.(map[string]interface{})["description"], bodyResponse.Data.(map[string]interface{})["description"])
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

func TestGameSuite(t *testing.T) {
	suite.Run(t, new(suiteGame))
}
