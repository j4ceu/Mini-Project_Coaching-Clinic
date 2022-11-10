package controllers

import (
	"Mini-Project_Coaching-Clinic/dto"
	"Mini-Project_Coaching-Clinic/dto/payload"
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

type suiteUserPayment struct {
	suite.Suite
	controller UserPaymentController
	mock       *mock.UserPaymentServiceMock
}

func (s *suiteUserPayment) SetupSuite() {
	s.mock = new(mock.UserPaymentServiceMock)
	s.controller = NewUserPaymentController(s.mock)
}

func (s *suiteUserPayment) TestCreateUserPayment() {

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
			Name:               "Create User Payment Success",
			ExpectedStatusCode: 200,
			Method:             "POST",
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "success create user payment",
				Status:  http.StatusOK,
				Data: dto.UserPaymentResponse{
					Paid:          false,
					InvoiceNumber: "INV-1",
				},
				Errors: nil,
			},
			FunctionError: nil,
		},
		{
			Name:               "Create User Payment Failed - Bind Error",
			ExpectedStatusCode: 400,
			Method:             "POST",
			HasReturnBody:      true,
			RequestBody:        `{"invoice_number" : 1}`,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Status:  http.StatusBadRequest,
				Data:    dto.EmptyObj{},
				Errors:  errors.New("error"),
			},
			FunctionError: errors.New("error"),
		},
		{
			Name:               "Create User Payment Failed - Internal Service Error",
			ExpectedStatusCode: 500,
			Method:             "POST",
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
			s.mock.On("Create", payload.UserPaymentPayload{UserID: "1"}).Return(dto.UserPaymentResponse{UserID: "1", InvoiceNumber: "INV-1"}, tc.FunctionError).Once()

			req := httptest.NewRequest(tc.Method, "/", strings.NewReader(tc.RequestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			res := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, res)

			token := new(jwt.Token)
			token.Claims = jwt.MapClaims{
				"user_id": "1",
				"role":    "Admin",
			}
			c.Set("token", token)

			err := s.controller.CreateUserPayment(c)
			s.NoError(err)

			s.Equal(tc.ExpectedStatusCode, res.Code)

			if tc.FunctionError == nil {
				if tc.HasReturnBody {
					var bodyResponse dto.BaseResponse
					s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
					s.Equal(tc.ExpectedBody.Data.(dto.UserPaymentResponse).Paid, bodyResponse.Data.(map[string]interface{})["paid"])
					s.Equal(tc.ExpectedBody.Data.(dto.UserPaymentResponse).InvoiceNumber, bodyResponse.Data.(map[string]interface{})["invoice_number"])
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

func (s *suiteUserPayment) TestUpdateUserPayment() {

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
		Role               string
		PayloadContentType string
	}{
		{
			Name:               "Update User Payment Success",
			ExpectedStatusCode: 200,
			Method:             "PUT",
			HasReturnBody:      true,
			RequestBody:        `{"paid" : true, "invoice_number" : "INV-1"}`,
			ExpectedBody: dto.BaseResponse{
				Message: "success update user payment",
				Status:  http.StatusOK,
				Data: dto.UserPaymentResponse{
					Paid:          true,
					InvoiceNumber: "INV-1",
				},
				Errors: nil,
			},
			FunctionError:      nil,
			HasParams:          true,
			ParamID:            "1",
			Role:               "Admin",
			PayloadContentType: echo.MIMEApplicationJSON,
		},
		{
			Name:               "Update User Payment Failed - Bind Error",
			ExpectedStatusCode: 400,
			Method:             "PUT",
			HasReturnBody:      true,
			RequestBody:        `{"paid" : true, "invoice_number" : 1}`,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Status:  http.StatusBadRequest,
				Data:    dto.EmptyObj{},
				Errors:  errors.New("error"),
			},
			FunctionError:      errors.New("error"),
			HasParams:          true,
			ParamID:            "1",
			Role:               "Admin",
			PayloadContentType: echo.MIMEApplicationJSON,
		},
		{
			Name:               "Update User Payment Failed - Internal Service Error",
			ExpectedStatusCode: 500,
			Method:             "PUT",
			HasReturnBody:      true,
			RequestBody:        `{"paid" : true, "invoice_number" : "INV-1"}`,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Status:  http.StatusInternalServerError,
				Data:    dto.EmptyObj{},
				Errors:  errors.New("error"),
			},
			FunctionError:      errors.New("error"),
			HasParams:          true,
			ParamID:            "1",
			Role:               "Admin",
			PayloadContentType: echo.MIMEApplicationJSON,
		},
		{
			Name:               "Update User Payment Failed - Unauthorized User Content Type Cant JSON",
			ExpectedStatusCode: 401,
			Method:             "PUT",
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Status:  http.StatusUnauthorized,
				Data:    dto.EmptyObj{},
				Errors:  errors.New("error"),
			},
			FunctionError:      errors.New("error"),
			HasParams:          true,
			ParamID:            "1",
			Role:               "User",
			PayloadContentType: echo.MIMEApplicationJSON,
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			pointerBool := new(bool)
			*pointerBool = true

			pointerString := new(string)
			*pointerString = "INV-1"

			req := httptest.NewRequest(tc.Method, "/", strings.NewReader(tc.RequestBody))
			req.Header.Set(echo.HeaderContentType, tc.PayloadContentType)
			res := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, res)

			if tc.HasParams {
				c.SetParamNames("id")
				c.SetParamValues(tc.ParamID)
			}

			token := new(jwt.Token)
			token.Claims = jwt.MapClaims{
				"user_id": "1",
				"role":    tc.Role,
			}
			c.Set("token", token)

			s.mock.On("Update", payload.UserPaymentPayload{InvoiceNumber: pointerString, Paid: pointerBool}, tc.ParamID, token.Claims).Return(dto.UserPaymentResponse{UserID: tc.ParamID, Paid: true, InvoiceNumber: "INV-1"}, tc.FunctionError).Once()

			err := s.controller.UpdateUserPayment(c)
			s.NoError(err)

			s.Equal(tc.ExpectedStatusCode, res.Code)

			if tc.FunctionError == nil {
				if tc.HasReturnBody {
					var bodyResponse dto.BaseResponse
					s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
					s.Equal(tc.ExpectedBody.Data.(dto.UserPaymentResponse).Paid, bodyResponse.Data.(map[string]interface{})["paid"])
					s.Equal(tc.ExpectedBody.Data.(dto.UserPaymentResponse).InvoiceNumber, bodyResponse.Data.(map[string]interface{})["invoice_number"])
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

func (s *suiteUserPayment) TestDeleteUserPayment() {

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
			Name:               "Delete User Payment Success",
			ExpectedStatusCode: 200,
			Method:             "DELETE",
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "success delete user payment",
				Status:  http.StatusOK,
				Data:    dto.EmptyObj{},
				Errors:  nil,
			},
			FunctionError: nil,
			HasParams:     true,
			ParamID:       "1",
		},
		{
			Name:               "Delete User Payment Failed - Bad Request",
			ExpectedStatusCode: 400,
			Method:             "DELETE",
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
		},
		{
			Name:               "Delete User Payment Failed - Internal Server Error",
			ExpectedStatusCode: 500,
			Method:             "DELETE",
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
			req := httptest.NewRequest(tc.Method, "/", strings.NewReader(tc.RequestBody))
			res := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, res)

			if tc.HasParams {
				c.SetParamNames("id")
				c.SetParamValues(tc.ParamID)
			}

			s.mock.On("Delete", tc.ParamID).Return(models.UserPayment{}, tc.FunctionError).Once()

			err := s.controller.DeleteUserPayment(c)
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

func (s *suiteUserPayment) TestGetUserPaymentByID() {

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
			Name:               "Get User Payment By ID Success",
			ExpectedStatusCode: 200,
			Method:             "GET",
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "success get user payment by id",
				Status:  http.StatusOK,
				Data: dto.UserPaymentResponse{
					UserID:        "1",
					InvoiceNumber: "INV-001",
					Paid:          true,
				},
				Errors: nil,
			},
			FunctionError: nil,
			HasParams:     true,
			ParamID:       "1",
		},
		{
			Name:               "Get User Payment By ID Failed - Bad Request",
			ExpectedStatusCode: 400,
			Method:             "GET",
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
		},
		{
			Name:               "Get User Payment By ID Failed - Internal Server Error",
			ExpectedStatusCode: 500,
			Method:             "GET",
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
			Name:               "Get User Payment By ID Failed - Unauthorized",
			ExpectedStatusCode: 401,
			Method:             "GET",
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "failed",
				Status:  http.StatusUnauthorized,
				Data:    dto.EmptyObj{},
				Errors:  errors.New("error"),
			},
			FunctionError: nil,
			HasParams:     true,
			ParamID:       "2",
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			req := httptest.NewRequest(tc.Method, "/", strings.NewReader(tc.RequestBody))
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

			s.mock.On("FindByID", tc.ParamID).Return(dto.UserPaymentResponse{UserID: tc.ParamID, InvoiceNumber: "INV-001", Paid: true}, tc.FunctionError).Once()

			err := s.controller.FindUserPaymentById(c)
			s.NoError(err)

			s.Equal(tc.ExpectedStatusCode, res.Code)

			if tc.FunctionError == nil && tc.ExpectedStatusCode != 401 {
				var bodyResponse dto.BaseResponse
				s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
				s.Equal(tc.ExpectedBody.Message, bodyResponse.Message)
				s.Equal(tc.ExpectedBody.Data.(dto.UserPaymentResponse).UserID, bodyResponse.Data.(map[string]interface{})["user_id"])
				s.Equal(tc.ExpectedBody.Data.(dto.UserPaymentResponse).InvoiceNumber, bodyResponse.Data.(map[string]interface{})["invoice_number"])
				s.Equal(tc.ExpectedBody.Data.(dto.UserPaymentResponse).Paid, bodyResponse.Data.(map[string]interface{})["paid"])
			} else {
				var bodyResponse dto.BaseResponse
				s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
				s.Equal(tc.ExpectedBody.Message, bodyResponse.Message)
			}

		})

	}

}

func (s *suiteUserPayment) TestFindAllUserPayment() {

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
			Name:               "Get All User Payment Success",
			ExpectedStatusCode: 200,
			Method:             "GET",
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "success get all user payment",
				Status:  http.StatusOK,
				Data: []dto.UserPaymentResponse{
					{
						UserID:        "1",
						InvoiceNumber: "INV-001",
						Paid:          true,
					},
				},
				Errors: nil,
			},
			FunctionError: nil,
		},
		{
			Name:               "Get All User Payment Failed - Internal Server Error",
			ExpectedStatusCode: 500,
			Method:             "GET",
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
			req := httptest.NewRequest(tc.Method, "/", strings.NewReader(tc.RequestBody))
			res := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, res)

			s.mock.On("FindAll").Return([]dto.UserPaymentResponse{
				{
					UserID:        "1",
					InvoiceNumber: "INV-001",
					Paid:          true,
				},
			}, tc.FunctionError).Once()

			err := s.controller.FindAllUserPayment(c)
			s.NoError(err)

			s.Equal(tc.ExpectedStatusCode, res.Code)

			if tc.FunctionError == nil {
				var bodyResponse dto.BaseResponse
				s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
				s.Equal(tc.ExpectedBody.Message, bodyResponse.Message)
				s.Equal(tc.ExpectedBody.Data.([]dto.UserPaymentResponse)[0].UserID, bodyResponse.Data.([]interface{})[0].(map[string]interface{})["user_id"])
				s.Equal(tc.ExpectedBody.Data.([]dto.UserPaymentResponse)[0].InvoiceNumber, bodyResponse.Data.([]interface{})[0].(map[string]interface{})["invoice_number"])
				s.Equal(tc.ExpectedBody.Data.([]dto.UserPaymentResponse)[0].Paid, bodyResponse.Data.([]interface{})[0].(map[string]interface{})["paid"])
			} else {
				var bodyResponse dto.BaseResponse
				s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
				s.Equal(tc.ExpectedBody.Message, bodyResponse.Message)
			}

		})

	}

}

func (s *suiteUserPayment) TestFindAllUserPaymentPaid() {

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
			Name:               "Get All User Payment Paid Success",
			ExpectedStatusCode: 200,
			Method:             "GET",
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "success get all user payment by paid",
				Status:  http.StatusOK,
				Data: []dto.UserPaymentResponse{
					{
						UserID:        "1",
						InvoiceNumber: "INV-001",
						Paid:          true,
					},
				},
				Errors: nil,
			},
			FunctionError: nil,
		},
		{
			Name:               "Get All User Payment Paid Failed - Internal Server Error",
			ExpectedStatusCode: 500,
			Method:             "GET",
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
			req := httptest.NewRequest(tc.Method, "/", strings.NewReader(tc.RequestBody))
			res := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, res)

			s.mock.On("FindByPaidAndProofOfPaymentIsNotNull").Return([]dto.UserPaymentResponse{
				{
					UserID:        "1",
					InvoiceNumber: "INV-001",
					Paid:          true,
				},
			}, tc.FunctionError).Once()

			err := s.controller.FindAllUserPaymentByPaid(c)
			s.NoError(err)

			s.Equal(tc.ExpectedStatusCode, res.Code)

			if tc.FunctionError == nil {
				var bodyResponse dto.BaseResponse
				s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
				s.Equal(tc.ExpectedBody.Message, bodyResponse.Message)
				s.Equal(tc.ExpectedBody.Data.([]dto.UserPaymentResponse)[0].UserID, bodyResponse.Data.([]interface{})[0].(map[string]interface{})["user_id"])
				s.Equal(tc.ExpectedBody.Data.([]dto.UserPaymentResponse)[0].InvoiceNumber, bodyResponse.Data.([]interface{})[0].(map[string]interface{})["invoice_number"])
				s.Equal(tc.ExpectedBody.Data.([]dto.UserPaymentResponse)[0].Paid, bodyResponse.Data.([]interface{})[0].(map[string]interface{})["paid"])
			} else {
				var bodyResponse dto.BaseResponse
				s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
				s.Equal(tc.ExpectedBody.Message, bodyResponse.Message)
			}

		})

	}

}

func (s *suiteUserPayment) TestFindAllUserPaymentByInvoice() {

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
			Name:               "Get All User Payment By Invoice Success",
			ExpectedStatusCode: 200,
			Method:             "GET",
			HasReturnBody:      true,
			RequestBody:        ``,
			ExpectedBody: dto.BaseResponse{
				Message: "success get user payment by invoice",
				Status:  http.StatusOK,
				Data: dto.UserPaymentResponse{
					UserID:        "1",
					InvoiceNumber: "INV-001",
					Paid:          true,
				},
				Errors: nil,
			},
			FunctionError: nil,
			HasParams:     true,
			ParamID:       "INV-001",
		},
		{
			Name:               "Get All User Payment By Invoice Failed - Internal Server Error",
			ExpectedStatusCode: 500,
			Method:             "GET",
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
			ParamID:       "INV-001",
		},
		{
			Name:               "Get All User Payment By Invoice Failed - Bad Request",
			ExpectedStatusCode: 400,
			Method:             "GET",
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
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			req := httptest.NewRequest(tc.Method, "/", strings.NewReader(tc.RequestBody))
			res := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, res)

			if tc.HasParams {
				c.SetParamNames("invoiceNumber")
				c.SetParamValues(tc.ParamID)
			}

			s.mock.On("FindByInvoiceNumber", tc.ParamID).Return(dto.UserPaymentResponse{
				UserID:        "1",
				InvoiceNumber: "INV-001",
				Paid:          true,
			}, tc.FunctionError).Once()

			err := s.controller.FindUserPaymentByInvoice(c)
			s.NoError(err)

			s.Equal(tc.ExpectedStatusCode, res.Code)

			if tc.FunctionError == nil {
				var bodyResponse dto.BaseResponse
				s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
				s.Equal(tc.ExpectedBody.Message, bodyResponse.Message)
				s.Equal(tc.ExpectedBody.Data.(dto.UserPaymentResponse).UserID, bodyResponse.Data.(map[string]interface{})["user_id"])
				s.Equal(tc.ExpectedBody.Data.(dto.UserPaymentResponse).InvoiceNumber, bodyResponse.Data.(map[string]interface{})["invoice_number"])
				s.Equal(tc.ExpectedBody.Data.(dto.UserPaymentResponse).Paid, bodyResponse.Data.(map[string]interface{})["paid"])
			} else {
				var bodyResponse dto.BaseResponse
				s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
				s.Equal(tc.ExpectedBody.Message, bodyResponse.Message)
			}

		})

	}

}

func TestUserPaymentSuite(t *testing.T) {
	suite.Run(t, new(suiteUserPayment))
}
