package CoachAvailability_Controller

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

type suiteCoachAvailability struct {
	suite.Suite
	controller CoachAvailabilityController
	mock       *mock.MockCoachAvailabilityService
}

func (s *suiteCoachAvailability) SetupTest() {
	s.mock = new(mock.MockCoachAvailabilityService)
	s.controller = NewCoachAvailabilityController(s.mock)
}

func (s *suiteCoachAvailability) TestCreateCoachAvailability() {

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
			Name:               "Create Coach Availability Success",
			ExpectedStatusCode: 200,
			Method:             "POST",
			HasReturnBody:      true,
			RequestBody:        `{"coach_id":"1","start_time":"12:00","end_time":"13:00", "day":"Senin"}`,
			ExpectedBody: dto.BaseResponse{
				Message: "success create coach availability",
				Data:    models.CoachAvailability{Day: "Senin", StartTime: "12:00", EndTime: "13:00", CoachID: "1"},
				Status:  http.StatusOK,
			},
			FunctionError: nil,
		},
		{
			Name:               "Create Coach Availability Failed - Bad Request",
			ExpectedStatusCode: 400,
			Method:             "POST",
			HasReturnBody:      true,
			RequestBody:        `{"coach_id":"1","start_time":"12:00","end_time":"13:00", "day":Senin}`,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Data:    dto.EmptyObj{},
				Status:  http.StatusBadRequest,
			},
			FunctionError: errors.New("error"),
		},
		{
			Name:               "Create Coach Availability Failed - Internal Server Error",
			ExpectedStatusCode: 500,
			Method:             "POST",
			HasReturnBody:      true,
			RequestBody:        `{"coach_id":"1","start_time":"12:00","end_time":"13:00", "day":"Senin"}`,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Data:    dto.EmptyObj{},
				Status:  http.StatusInternalServerError,
			},
			FunctionError: errors.New("error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mock.On("Create", models.CoachAvailability{Day: "Senin", StartTime: "12:00", EndTime: "13:00", CoachID: "1"}).Return(dto.CoachAvailabilityResponse{Day: "Senin", StartTime: "12:00", EndTime: "13:00", CoachID: "1"}, tc.FunctionError).Once()

			req := httptest.NewRequest(tc.Method, "/", strings.NewReader(tc.RequestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			res := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, res)

			err := s.controller.CreateCoachAvailability(c)
			s.NoError(err)

			s.Equal(tc.ExpectedStatusCode, res.Code)

			if tc.FunctionError == nil {
				if tc.HasReturnBody {
					var bodyResponse dto.BaseResponse
					s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
					s.Equal(tc.ExpectedBody.Data.(models.CoachAvailability).Day, bodyResponse.Data.(map[string]interface{})["day"])
					s.Equal(tc.ExpectedBody.Data.(models.CoachAvailability).StartTime, bodyResponse.Data.(map[string]interface{})["start_time"])
					s.Equal(tc.ExpectedBody.Data.(models.CoachAvailability).EndTime, bodyResponse.Data.(map[string]interface{})["end_time"])
					s.Equal(tc.ExpectedBody.Data.(models.CoachAvailability).CoachID, bodyResponse.Data.(map[string]interface{})["coach_id"])
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

func (s *suiteCoachAvailability) TestUpdateAvailability() {

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
			Name:               "Update Coach Availability Success",
			ExpectedStatusCode: 200,
			Method:             "PUT",
			HasReturnBody:      true,
			RequestBody:        `{"coach_id":"1","start_time":"12:00","end_time":"13:00", "day":"Senin"}`,
			ExpectedBody: dto.BaseResponse{
				Message: "success update coach availability",
				Data:    models.CoachAvailability{Day: "Senin", StartTime: "12:00", EndTime: "13:00", CoachID: "1"},
				Status:  http.StatusOK,
				Errors:  errors.New("error"),
			},
			FunctionError: nil,
			HasParams:     true,
			ParamID:       "1",
		},
		{
			Name:               "Update Coach Availability Failed - Bad Request Bind Error",
			ExpectedStatusCode: 400,
			Method:             "PUT",
			HasReturnBody:      true,
			RequestBody:        `{"coach_id":"1","start_time":"12:00","end_time":"13:00", "day":Senin}`,
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
		{
			Name:               "Update Coach Availability Failed - Bad Request ID Empty",
			ExpectedStatusCode: 400,
			Method:             "PUT",
			HasReturnBody:      true,
			RequestBody:        `{"coach_id":"1","start_time":"12:00","end_time":"13:00", "day":"Senin"}`,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Data:    dto.EmptyObj{},
				Status:  http.StatusBadRequest,
				Errors:  errors.New("error"),
			},
			FunctionError: errors.New("error"),
			HasParams:     false,
			ParamID:       "",
		},
		{
			Name:               "Update Coach Availability Failed - Internal Server Error",
			ExpectedStatusCode: 500,
			Method:             "PUT",
			HasReturnBody:      true,
			RequestBody:        `{"coach_id":"1","start_time":"12:00","end_time":"13:00", "day":"Senin"}`,
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
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mock.On("Update", models.CoachAvailability{Day: "Senin", StartTime: "12:00", EndTime: "13:00", CoachID: "1"}, "1").Return(dto.CoachAvailabilityResponse{Day: "Senin", StartTime: "12:00", EndTime: "13:00", CoachID: "1"}, tc.FunctionError).Once()

			req := httptest.NewRequest(tc.Method, "/", strings.NewReader(tc.RequestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			res := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, res)

			if tc.HasParams {
				c.SetParamNames("id")
				c.SetParamValues(tc.ParamID)
			}

			err := s.controller.UpdateCoachAvailability(c)
			s.NoError(err)

			s.Equal(tc.ExpectedStatusCode, res.Code)

			if tc.FunctionError == nil {
				if tc.HasReturnBody {
					var bodyResponse dto.BaseResponse
					s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
					s.Equal(tc.ExpectedBody.Data.(models.CoachAvailability).Day, bodyResponse.Data.(map[string]interface{})["day"])
					s.Equal(tc.ExpectedBody.Data.(models.CoachAvailability).StartTime, bodyResponse.Data.(map[string]interface{})["start_time"])
					s.Equal(tc.ExpectedBody.Data.(models.CoachAvailability).EndTime, bodyResponse.Data.(map[string]interface{})["end_time"])
					s.Equal(tc.ExpectedBody.Data.(models.CoachAvailability).CoachID, bodyResponse.Data.(map[string]interface{})["coach_id"])
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

func (s *suiteCoachAvailability) TestDeleteCoachAvailability() {
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
			Name:               "Delete Coach Availability Success",
			ExpectedStatusCode: 200,
			Method:             "DELETE",
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "success delete coach availability",
				Data:    models.CoachAvailability{Day: "Senin", StartTime: "12:00", EndTime: "13:00", CoachID: "1"},
				Status:  http.StatusOK,
				Errors:  nil,
			},
			FunctionError: nil,
			HasParams:     true,
			ParamID:       "1",
		},
		{
			Name:               "Delete Coach Availability Failed - Bad Request ID Empty",
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
			ParamID:       "",
		},
		{
			Name:               "Delete Coach Availability Failed - Internal Server Error",
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
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.mock.On("Delete", "1").Return(models.CoachAvailability{}, tc.FunctionError).Once()

			req := httptest.NewRequest(tc.Method, "/", strings.NewReader(tc.RequestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			res := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, res)

			if tc.HasParams {
				c.SetParamNames("id")
				c.SetParamValues(tc.ParamID)
			}

			err := s.controller.DeleteCoachAvailability(c)
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

func TestCoachAvailabilitySuite(t *testing.T) {
	suite.Run(t, new(suiteCoachAvailability))
}
