package UserBook_Controller

import (
	"Mini-Project_Coaching-Clinic/dto"
	"Mini-Project_Coaching-Clinic/dto/payload"
	"Mini-Project_Coaching-Clinic/services/UserBook_Service"
	"Mini-Project_Coaching-Clinic/services/UserPayment_Service"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type UserBookController interface {
	CreateUserBook(c echo.Context) error
	UpdateUserBook(c echo.Context) error
	DeleteUserBook(c echo.Context) error
	FindAllUserBook(c echo.Context) error
	FindUserBookById(c echo.Context) error
	FindUserBookByUserID(c echo.Context) error
}

type userBookController struct {
	userBookService    UserBook_Service.UserBookService
	userPaymentService UserPayment_Service.UserPaymentService
}

func NewUserBookController(userBookService UserBook_Service.UserBookService, userPaymentService UserPayment_Service.UserPaymentService) *userBookController {
	return &userBookController{userBookService, userPaymentService}
}

func (ub *userBookController) CreateUserBook(c echo.Context) error {
	var userBook payload.UserBookPayloadCreate
	if err := c.Bind(&userBook); err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusBadRequest, baseResponse)
	}

	userBookCreate, err := ub.userBookService.Create(userBook)
	if err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusInternalServerError, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusInternalServerError, baseResponse)
	}

	baseResponse := dto.ConvertToBaseResponse("success create user book", http.StatusOK, userBookCreate)
	return c.JSON(http.StatusOK, baseResponse)
}

func (ub *userBookController) UpdateUserBook(c echo.Context) error {
	var userBook payload.UserBookPayloadUpdate
	idReq := c.Param("id")

	if idReq == "" {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, dto.EmptyObj{}, "id is required")
		return c.JSON(http.StatusBadRequest, baseResponse)
	}

	if err := c.Bind(&userBook); err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusBadRequest, baseResponse)
	}

	file, err := c.FormFile("summary")
	if err != nil {
		if err == http.ErrMissingFile {
			userBook.Summary = nil
		} else {
			baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, dto.EmptyObj{}, err.Error())
			return c.JSON(http.StatusBadRequest, baseResponse)
		}
	} else {
		userBook.Summary = file
	}

	userBookUpdate, err := ub.userBookService.Update(userBook, idReq)
	if err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusInternalServerError, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusInternalServerError, baseResponse)
	}

	baseResponse := dto.ConvertToBaseResponse("success update user book", http.StatusOK, userBookUpdate)
	return c.JSON(http.StatusOK, baseResponse)
}

func (ub *userBookController) DeleteUserBook(c echo.Context) error {
	idReq := c.Param("id")

	if idReq == "" {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, dto.EmptyObj{}, "id is required")
		return c.JSON(http.StatusBadRequest, baseResponse)
	}

	_, err := ub.userBookService.Delete(idReq)
	if err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusInternalServerError, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusInternalServerError, baseResponse)
	}

	baseResponse := dto.ConvertToBaseResponse("success delete user book", http.StatusOK, dto.EmptyObj{})
	return c.JSON(http.StatusOK, baseResponse)
}

func (ub *userBookController) FindAllUserBook(c echo.Context) error {
	userBook, err := ub.userBookService.FindAll()
	if err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusInternalServerError, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusInternalServerError, baseResponse)
	}

	baseResponse := dto.ConvertToBaseResponse("success get all user book", http.StatusOK, userBook)
	return c.JSON(http.StatusOK, baseResponse)
}

func (ub *userBookController) FindUserBookById(c echo.Context) error {
	idReq := c.Param("id")

	if idReq == "" {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, dto.EmptyObj{}, "id is required")
		return c.JSON(http.StatusBadRequest, baseResponse)
	}

	userBook, err := ub.userBookService.FindByID(idReq)
	if err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusInternalServerError, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusInternalServerError, baseResponse)
	}

	userPayment, _ := ub.userPaymentService.FindByID(userBook.UserPaymentID)

	token := c.Get("token").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	role := claims["role"].(string)

	if role == "User" && userID != userPayment.UserID {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusUnauthorized, dto.EmptyObj{}, "unauthorized")
		return c.JSON(http.StatusUnauthorized, baseResponse)
	}

	baseResponse := dto.ConvertToBaseResponse("success get user book by id", http.StatusOK, userBook)
	return c.JSON(http.StatusOK, baseResponse)
}

func (ub *userBookController) FindUserBookByUserID(c echo.Context) error {
	idReq := c.Param("id")

	if idReq == "" {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, dto.EmptyObj{}, "id is required")
		return c.JSON(http.StatusBadRequest, baseResponse)
	}

	token := c.Get("token").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	role := claims["role"].(string)

	if role == "User" && userID != idReq {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusUnauthorized, dto.EmptyObj{}, "unauthorized")
		return c.JSON(http.StatusUnauthorized, baseResponse)
	}

	userBook, err := ub.userBookService.FindByUserID(idReq)
	if err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusInternalServerError, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusInternalServerError, baseResponse)
	}

	baseResponse := dto.ConvertToBaseResponse("success get user book by user id", http.StatusOK, userBook)
	return c.JSON(http.StatusOK, baseResponse)
}
