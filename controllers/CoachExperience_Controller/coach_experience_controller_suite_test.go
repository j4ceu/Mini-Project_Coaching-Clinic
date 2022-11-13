package CoachExperience_Controller

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

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type suiteCoachExperience struct {
	suite.Suite
	controller CoachExperienceController
	mock       *mock.MockCoachExperienceService
}

func (s *suiteCoachExperience) SetupSuite() {
	s.mock = new(mock.MockCoachExperienceService)
	s.controller = NewCoachExperienceController(s.mock)
}

func (s *suiteCoachExperience) TestCreateCoachExperience() {
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
			Name:               "success create coach experience",
			ExpectedStatusCode: http.StatusOK,
			Method:             http.MethodPost,
			HasReturnBody:      true,
			RequestBody:        `{"coach_id": "1", "title": "test", "description": "test"}`,
			ExpectedBody: dto.BaseResponse{
				Message: "success create coach experience",
				Status:  http.StatusOK,
				Data:    models.CoachExperience{Title: "test", Description: "test"},
			},
			FunctionError: nil,
		},
		{
			Name:               "failed create coach experience bad request",
			ExpectedStatusCode: http.StatusBadRequest,
			Method:             http.MethodPost,
			HasReturnBody:      true,
			RequestBody:        `{"coach_id": "1", "title": test, "description": "test"}`,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Status:  http.StatusBadRequest,
				Data:    dto.EmptyObj{},
				Errors:  errors.New("error"),
			},
			FunctionError: errors.New("error"),
		},
		{
			Name:               "failed create coach experience internal server error",
			ExpectedStatusCode: http.StatusInternalServerError,
			Method:             http.MethodPost,
			HasReturnBody:      true,
			RequestBody:        `{"coach_id": "1", "title": "test", "description": "test"}`,
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
			s.mock.On("Create", models.CoachExperience{CoachID: "1", Title: "test", Description: "test"}).Return(models.CoachExperience{CoachID: "1", Title: "test", Description: "test"}, tc.FunctionError).Once()
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

			err := s.controller.CreateCoachExperience(c)
			s.NoError(err)

			s.Equal(tc.ExpectedStatusCode, res.Code)

			if tc.FunctionError == nil {
				if tc.HasReturnBody {
					var bodyResponse dto.BaseResponse
					s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
					s.Equal(tc.ExpectedBody.Data.(models.CoachExperience).Title, bodyResponse.Data.(map[string]interface{})["title"])
					s.Equal(tc.ExpectedBody.Data.(models.CoachExperience).Description, bodyResponse.Data.(map[string]interface{})["description"])
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

func (s *suiteCoachExperience) TestUpdateExperience() {

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
			Name:               "success update coach experience",
			ExpectedStatusCode: http.StatusOK,
			Method:             http.MethodPut,
			HasReturnBody:      true,
			RequestBody:        `{"title": "test", "description": "test"}`,
			ExpectedBody: dto.BaseResponse{
				Message: "success update coach experience",
				Status:  http.StatusOK,
				Data:    models.CoachExperience{Title: "test", Description: "test"},
			},
			FunctionError: nil,
			HasParams:     true,
			ParamID:       "1",
		},
		{
			Name:               "failed update coach experience bad request",
			ExpectedStatusCode: http.StatusBadRequest,
			Method:             http.MethodPut,
			HasReturnBody:      true,
			RequestBody:        `{"title": test, "description": "test"}`,
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
			Name:               "failed update coach bad request id empty",
			ExpectedStatusCode: http.StatusBadRequest,
			Method:             http.MethodPut,
			HasReturnBody:      true,
			RequestBody:        `{"title": "test", "description": "test"}`,
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
			Name:               "failed update coach experience internal server error",
			ExpectedStatusCode: http.StatusInternalServerError,
			Method:             http.MethodPut,
			HasReturnBody:      true,
			RequestBody:        `{"title": "test", "description": "test"}`,
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
			s.mock.On("Update", models.CoachExperience{Title: "test", Description: "test"}, "1").Return(models.CoachExperience{Title: "test", Description: "test"}, tc.FunctionError).Once()
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

			err := s.controller.UpdateCoachExperience(c)
			s.NoError(err)

			s.Equal(tc.ExpectedStatusCode, res.Code)

			if tc.FunctionError == nil {
				if tc.HasReturnBody {
					var bodyResponse dto.BaseResponse
					s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
					s.Equal(tc.ExpectedBody.Data.(models.CoachExperience).Title, bodyResponse.Data.(map[string]interface{})["title"])
					s.Equal(tc.ExpectedBody.Data.(models.CoachExperience).Description, bodyResponse.Data.(map[string]interface{})["description"])
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

func (s *suiteCoachExperience) TestDeleteExperience() {

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
			Name:               "success delete coach experience",
			ExpectedStatusCode: http.StatusOK,
			Method:             http.MethodDelete,
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "success delete coach experience",
				Status:  http.StatusOK,
				Data:    dto.EmptyObj{},
			},
			FunctionError: nil,
			HasParams:     true,
			ParamID:       "1",
		},
		{
			Name:               "failed delete coach experience bad request",
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
			Name:               "failed delete coach experience internal server error",
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
			s.mock.On("Delete", "1").Return(models.CoachExperience{}, tc.FunctionError).Once()
			req := httptest.NewRequest(tc.Method, "/", strings.NewReader(tc.RequestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			res := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, res)

			if tc.HasParams {
				c.SetParamNames("id")
				c.SetParamValues(tc.ParamID)
			}

			err := s.controller.DeleteCoachExperience(c)
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

func TestCoachExperienceSuite(t *testing.T) {
	suite.Run(t, new(suiteCoachExperience))
}
