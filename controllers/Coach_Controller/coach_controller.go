package Coach_Controller

import (
	"Mini-Project_Coaching-Clinic/dto"
	"Mini-Project_Coaching-Clinic/models"
	"Mini-Project_Coaching-Clinic/services/Coach_Service"
	"net/http"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CoachController interface {
	FindCoachByID(c echo.Context) error
	FindCoachByGameID(c echo.Context) error
	FindCoachByCode(c echo.Context) error
	CreateCoach(c echo.Context) error
	UpdateCoach(c echo.Context) error
	DeleteCoach(c echo.Context) error
}

type coachController struct {
	coachService Coach_Service.CoachService
}

func NewCoachController(coachService Coach_Service.CoachService) *coachController {
	return &coachController{coachService}
}

func (co *coachController) CreateCoach(c echo.Context) error {
	var coach models.Coach
	if err := c.Bind(&coach); err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusBadRequest, baseResponse)
	}

	if !reflect.ValueOf(coach.User).IsZero() {
		validate := validator.New()

		if err := validate.Struct(coach.User); err != nil {
			baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, dto.EmptyObj{}, err.Error())
			return c.JSON(http.StatusBadRequest, baseResponse)
		}
	}

	coachResponse, err := co.coachService.Create(coach)
	if err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusInternalServerError, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusInternalServerError, baseResponse)
	}

	baseResponse := dto.ConvertToBaseResponse("success create coach", http.StatusOK, coachResponse)
	return c.JSON(http.StatusOK, baseResponse)
}

func (co *coachController) UpdateCoach(c echo.Context) error {
	var coach models.Coach

	idReq := c.Param("id")

	if idReq == "" {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, dto.EmptyObj{}, "id is required")
		return c.JSON(http.StatusBadRequest, baseResponse)
	}

	if err := c.Bind(&coach); err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusBadRequest, baseResponse)
	}

	coachResponse, err := co.coachService.Update(coach, idReq)
	if err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusInternalServerError, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusInternalServerError, baseResponse)
	}

	baseResponse := dto.ConvertToBaseResponse("success update coach", http.StatusOK, coachResponse)
	return c.JSON(http.StatusOK, baseResponse)
}

func (co *coachController) DeleteCoach(c echo.Context) error {
	idReq := c.Param("id")

	if idReq == "" {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, dto.EmptyObj{}, "id is required")
		return c.JSON(http.StatusBadRequest, baseResponse)
	}

	_, err := co.coachService.Delete(idReq)
	if err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusInternalServerError, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusInternalServerError, baseResponse)
	}

	baseResponse := dto.ConvertToBaseResponse("success delete coach", http.StatusOK, dto.EmptyObj{})
	return c.JSON(http.StatusOK, baseResponse)
}

func (co *coachController) FindCoachByID(c echo.Context) error {
	idReq := c.Param("id")

	if idReq == "" {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, dto.EmptyObj{}, "id is required")
		return c.JSON(http.StatusBadRequest, baseResponse)
	}

	coach, err := co.coachService.FindByID(idReq)
	if err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusInternalServerError, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusInternalServerError, baseResponse)
	}

	baseResponse := dto.ConvertToBaseResponse("success get coach by id", http.StatusOK, coach)
	return c.JSON(http.StatusOK, baseResponse)
}

func (co *coachController) FindCoachByGameID(c echo.Context) error {
	idReq := c.Param("id")

	if idReq == "" {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, dto.EmptyObj{}, "id is required")
		return c.JSON(http.StatusBadRequest, baseResponse)
	}

	coach, err := co.coachService.FindByGameID(idReq)
	if err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusInternalServerError, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusInternalServerError, baseResponse)
	}

	baseResponse := dto.ConvertToBaseResponse("success get coach by game id", http.StatusOK, coach)
	return c.JSON(http.StatusOK, baseResponse)
}

func (co *coachController) FindCoachByCode(c echo.Context) error {
	codeReq := c.Param("code")

	if codeReq == "" {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, dto.EmptyObj{}, "code is required")
		return c.JSON(http.StatusBadRequest, baseResponse)
	}

	coach, err := co.coachService.FindByCode(codeReq)
	if err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusInternalServerError, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusInternalServerError, baseResponse)
	}

	baseResponse := dto.ConvertToBaseResponse("success get coach by code", http.StatusOK, coach)
	return c.JSON(http.StatusOK, baseResponse)
}
