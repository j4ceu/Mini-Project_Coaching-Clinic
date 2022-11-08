package controllers

import (
	"Mini-Project_Coaching-Clinic/dto"
	"Mini-Project_Coaching-Clinic/models"
	"Mini-Project_Coaching-Clinic/services"

	"net/http"

	"github.com/labstack/echo/v4"
)

type CoachAvailabilityController interface {
	CreateCoachAvailability(c echo.Context) error
	UpdateCoachAvailability(c echo.Context) error
	DeleteCoachAvailability(c echo.Context) error
}

type coachAvailabilityController struct {
	coachAvailabilityService services.CoachAvailabilityService
}

func NewCoachAvailabilityController(coachAvailabilityService services.CoachAvailabilityService) *coachAvailabilityController {
	return &coachAvailabilityController{coachAvailabilityService}
}

func (ca *coachAvailabilityController) CreateCoachAvailability(c echo.Context) error {
	var coachAvailability models.CoachAvailability
	if err := c.Bind(&coachAvailability); err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusBadRequest, baseResponse)
	}

	coachAvailability, err := ca.coachAvailabilityService.Create(coachAvailability)
	if err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusInternalServerError, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusInternalServerError, baseResponse)
	}

	baseResponse := dto.ConvertToBaseResponse("success create coach availability", http.StatusOK, coachAvailability)
	return c.JSON(http.StatusOK, baseResponse)
}

func (ca *coachAvailabilityController) UpdateCoachAvailability(c echo.Context) error {
	var coachAvailability models.CoachAvailability
	if err := c.Bind(&coachAvailability); err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusBadRequest, baseResponse)
	}

	id := c.Param("id")
	if id == "" {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, dto.EmptyObj{}, "id is required")
		return c.JSON(http.StatusBadRequest, baseResponse)
	}

	coachAvailability, err := ca.coachAvailabilityService.Update(coachAvailability, id)
	if err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusInternalServerError, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusInternalServerError, baseResponse)
	}

	baseResponse := dto.ConvertToBaseResponse("success update coach availability", http.StatusOK, coachAvailability)
	return c.JSON(http.StatusOK, baseResponse)
}

func (ca *coachAvailabilityController) DeleteCoachAvailability(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, dto.EmptyObj{}, "id is required")
		return c.JSON(http.StatusBadRequest, baseResponse)
	}

	_, err := ca.coachAvailabilityService.Delete(id)
	if err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusInternalServerError, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusInternalServerError, baseResponse)
	}

	baseResponse := dto.ConvertToBaseResponse("success delete coach availability", http.StatusOK, dto.EmptyObj{})
	return c.JSON(http.StatusOK, baseResponse)
}
