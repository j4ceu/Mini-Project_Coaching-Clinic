package controllers

import (
	"Mini-Project_Coaching-Clinic/dto"
	"Mini-Project_Coaching-Clinic/dto/payload"
	"Mini-Project_Coaching-Clinic/services"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type UserPaymentController interface {
	CreateUserPayment(c echo.Context) error
	UpdateUserPayment(c echo.Context) error
	DeleteUserPayment(c echo.Context) error
	FindUserPaymentById(c echo.Context) error
	FindAllUserPayment(c echo.Context) error
	FindAllUserPaymentByPaid(c echo.Context) error
	FindUserPaymentByInvoice(c echo.Context) error
}

type userPaymentController struct {
	userPaymentService services.UserPaymentService
}

func NewUserPaymentController(userPaymentService services.UserPaymentService) *userPaymentController {
	return &userPaymentController{userPaymentService}
}

func (up *userPaymentController) CreateUserPayment(c echo.Context) error {
	var userPayment payload.UserPaymentPayload

	token := c.Get("token").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	role := claims["role"].(string)
	userPayment.UserID = claims["user_id"].(string)

	if role != "User" {
		if err := c.Bind(&userPayment); err != nil {
			baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, dto.EmptyObj{}, err.Error())
			return c.JSON(http.StatusBadRequest, baseResponse)
		}
	}

	userPaymentResponse, err := up.userPaymentService.Create(userPayment)
	if err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusInternalServerError, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusInternalServerError, baseResponse)
	}

	baseResponse := dto.ConvertToBaseResponse("success create user payment", http.StatusOK, userPaymentResponse)
	return c.JSON(http.StatusOK, baseResponse)

}

func (up *userPaymentController) UpdateUserPayment(c echo.Context) error {
	var userPayment payload.UserPaymentPayload
	idReq := c.Param("id")

	if idReq == "" {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, dto.EmptyObj{}, "id is required")
		return c.JSON(http.StatusBadRequest, baseResponse)
	}

	token := c.Get("token").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	contentType := c.Request().Header.Get("Content-Type")
	if claims["role"].(string) != "Admin" {
		if contentType == "application/json" {
			baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusUnauthorized, dto.EmptyObj{}, "unauthorized")
			return c.JSON(http.StatusUnauthorized, baseResponse)
		}
		file, err := c.FormFile("proof_of_payment")
		if err != nil {
			baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, dto.EmptyObj{}, "proof of payment is required")
			return c.JSON(http.StatusBadRequest, baseResponse)
		}
		userPayment.ProofOfPayment = file
	}

	if err := c.Bind(&userPayment); err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusBadRequest, baseResponse)
	}

	userPaymentResponse, err := up.userPaymentService.Update(userPayment, idReq, claims)
	if err != nil {
		if err.Error() == "unauthorized" {
			baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusUnauthorized, dto.EmptyObj{}, err.Error())
			return c.JSON(http.StatusUnauthorized, baseResponse)
		}
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusInternalServerError, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusInternalServerError, baseResponse)
	}

	baseResponse := dto.ConvertToBaseResponse("success update user payment", http.StatusOK, userPaymentResponse)
	return c.JSON(http.StatusOK, baseResponse)
}

func (up *userPaymentController) DeleteUserPayment(c echo.Context) error {
	idReq := c.Param("id")

	if idReq == "" {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, dto.EmptyObj{}, "id is required")
		return c.JSON(http.StatusBadRequest, baseResponse)
	}

	_, err := up.userPaymentService.Delete(idReq)
	if err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusInternalServerError, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusInternalServerError, baseResponse)
	}

	baseResponse := dto.ConvertToBaseResponse("success delete user payment", http.StatusOK, dto.EmptyObj{})
	return c.JSON(http.StatusOK, baseResponse)
}

func (up *userPaymentController) FindUserPaymentById(c echo.Context) error {
	idReq := c.Param("id")

	if idReq == "" {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, dto.EmptyObj{}, "id is required")
		return c.JSON(http.StatusBadRequest, baseResponse)
	}

	userPayment, err := up.userPaymentService.FindByID(idReq)
	if err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusInternalServerError, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusInternalServerError, baseResponse)
	}

	token := c.Get("token").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	role := claims["role"].(string)

	if userPayment.UserID != userID && role != "Admin" {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusUnauthorized, dto.EmptyObj{}, "unauthorized")
		return c.JSON(http.StatusUnauthorized, baseResponse)
	}

	baseResponse := dto.ConvertToBaseResponse("success get user payment by id", http.StatusOK, userPayment)
	return c.JSON(http.StatusOK, baseResponse)
}

func (up *userPaymentController) FindAllUserPayment(c echo.Context) error {
	paid := c.QueryParam("paid")
	if paid == "false" {
		return up.FindAllUserPaymentByPaid(c)
	}

	userPayment, err := up.userPaymentService.FindAll()
	if err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusInternalServerError, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusInternalServerError, baseResponse)
	}

	baseResponse := dto.ConvertToBaseResponse("success get all user payment", http.StatusOK, userPayment)
	return c.JSON(http.StatusOK, baseResponse)
}

func (up *userPaymentController) FindAllUserPaymentByPaid(c echo.Context) error {
	userPayment, err := up.userPaymentService.FindByPaidAndProofOfPaymentIsNotNull()
	if err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusInternalServerError, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusInternalServerError, baseResponse)
	}

	baseResponse := dto.ConvertToBaseResponse("success get all user payment by paid", http.StatusOK, userPayment)
	return c.JSON(http.StatusOK, baseResponse)
}

func (up *userPaymentController) FindUserPaymentByInvoice(c echo.Context) error {
	invoice := c.Param("invoiceNumber")

	if invoice == "" {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, dto.EmptyObj{}, "invoice is required")
		return c.JSON(http.StatusBadRequest, baseResponse)
	}

	userPayment, err := up.userPaymentService.FindByInvoiceNumber(invoice)
	if err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusInternalServerError, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusInternalServerError, baseResponse)
	}

	baseResponse := dto.ConvertToBaseResponse("success get user payment by invoice", http.StatusOK, userPayment)
	return c.JSON(http.StatusOK, baseResponse)
}
