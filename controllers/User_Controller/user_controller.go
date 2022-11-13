package User_Controller

import (
	"Mini-Project_Coaching-Clinic/dto"
	"Mini-Project_Coaching-Clinic/models"
	"Mini-Project_Coaching-Clinic/services/User_Service"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type UserController interface {
	CreateUser(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
	GetAllUser(c echo.Context) error
	LoginUser(c echo.Context) error
	GetDetailUser(c echo.Context) error
}

type userController struct {
	userService User_Service.UserService
}

func NewUserController(userService User_Service.UserService) *userController {
	return &userController{userService}
}

func (u *userController) CreateUser(c echo.Context) error {
	var user models.User

	if err := c.Bind(&user); err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusBadRequest, baseResponse)
	}

	user, err := u.userService.Create(user)
	if err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusInternalServerError, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusInternalServerError, baseResponse)
	}

	baseResponse := dto.ConvertToBaseResponse("success", http.StatusOK, user)
	return c.JSON(http.StatusOK, baseResponse)
}

func (u *userController) UpdateUser(c echo.Context) error {
	var user models.User

	idReq := c.Param("id")

	if idReq == "" {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, dto.EmptyObj{}, "id is required")
		return c.JSON(http.StatusBadRequest, baseResponse)
	}

	token := c.Get("token").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	role := claims["role"].(string)

	if userID != idReq && role != "Admin" {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusUnauthorized, dto.EmptyObj{}, "unauthorized")
		return c.JSON(http.StatusUnauthorized, baseResponse)
	}

	if err := c.Bind(&user); err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusBadRequest, baseResponse)
	}

	user, err := u.userService.Update(user, idReq)
	if err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusInternalServerError, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusInternalServerError, baseResponse)
	}

	baseResponse := dto.ConvertToBaseResponse("success", http.StatusOK, user)
	return c.JSON(http.StatusOK, baseResponse)
}

func (u *userController) DeleteUser(c echo.Context) error {
	idReq := c.Param("id")

	if idReq == "" {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, dto.EmptyObj{}, "id is required")
		return c.JSON(http.StatusBadRequest, baseResponse)
	}

	token := c.Get("token").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	role := claims["role"].(string)

	if userID != idReq && role != "Admin" {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusUnauthorized, dto.EmptyObj{}, "unauthorized")
		return c.JSON(http.StatusUnauthorized, baseResponse)
	}

	_, err := u.userService.Delete(idReq)
	if err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusInternalServerError, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusInternalServerError, baseResponse)
	}

	baseResponse := dto.ConvertToBaseResponse("success", http.StatusOK, dto.EmptyObj{})
	return c.JSON(http.StatusOK, baseResponse)
}

func (u *userController) GetAllUser(c echo.Context) error {
	users, err := u.userService.FindAll()
	if err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusInternalServerError, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusInternalServerError, baseResponse)
	}

	baseResponse := dto.ConvertToBaseResponse("success", http.StatusOK, users)
	return c.JSON(http.StatusOK, baseResponse)
}

func (u *userController) LoginUser(c echo.Context) error {
	var user models.User

	if err := c.Bind(&user); err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusBadRequest, baseResponse)
	}

	userResponse, err := u.userService.LoginUser(user.Email, user.Password)
	if err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusInternalServerError, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusInternalServerError, baseResponse)
	}

	baseResponse := dto.ConvertToBaseResponse("success", http.StatusOK, userResponse)
	return c.JSON(http.StatusOK, baseResponse)
}

func (u *userController) GetDetailUser(c echo.Context) error {
	idReq := c.Param("id")

	if idReq == "" {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, dto.EmptyObj{}, "id is required")
		return c.JSON(http.StatusBadRequest, baseResponse)
	}

	user, err := u.userService.FindByID(idReq)
	if err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusInternalServerError, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusInternalServerError, baseResponse)
	}

	baseResponse := dto.ConvertToBaseResponse("success", http.StatusOK, user)
	return c.JSON(http.StatusOK, baseResponse)
}
