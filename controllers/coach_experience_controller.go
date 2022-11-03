package controllers

import (
	"Mini-Project_Coaching-Clinic/dto"
	"Mini-Project_Coaching-Clinic/models"
	"Mini-Project_Coaching-Clinic/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CoachExperienceController interface {
	CreateCoachExperience(c echo.Context) error
	UpdateCoachExperience(c echo.Context) error
	DeleteCoachExperience(c echo.Context) error
}

type coachExperienceController struct {
	coachExperienceService services.CoachExperienceService
}

func NewCoachExperienceController(coachExperienceService services.CoachExperienceService) *coachExperienceController {
	return &coachExperienceController{coachExperienceService}
}

func (ce *coachExperienceController) CreateCoachExperience(c echo.Context) error {
	var coachExperience models.CoachExperience
	if err := c.Bind(&coachExperience); err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusBadRequest, baseResponse)
	}

	coachExperience, err := ce.coachExperienceService.Create(coachExperience)
	if err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusInternalServerError, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusInternalServerError, baseResponse)
	}

	baseResponse := dto.ConvertToBaseResponse("success create coach experience", http.StatusOK, coachExperience)
	return c.JSON(http.StatusOK, baseResponse)
}

func (ce *coachExperienceController) UpdateCoachExperience(c echo.Context) error {
	var coachExperience models.CoachExperience
	idReq := c.Param("id")

	if idReq == "" {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, dto.EmptyObj{}, "id is required")
		return c.JSON(http.StatusBadRequest, baseResponse)
	}

	if err := c.Bind(&coachExperience); err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusBadRequest, baseResponse)
	}

	coachExperience, err := ce.coachExperienceService.Update(coachExperience, idReq)
	if err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusInternalServerError, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusInternalServerError, baseResponse)
	}

	baseResponse := dto.ConvertToBaseResponse("success update coach experience", http.StatusOK, coachExperience)
	return c.JSON(http.StatusOK, baseResponse)
}

func (ce *coachExperienceController) DeleteCoachExperience(c echo.Context) error {
	idReq := c.Param("id")

	if idReq == "" {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, dto.EmptyObj{}, "id is required")
		return c.JSON(http.StatusBadRequest, baseResponse)
	}

	_, err := ce.coachExperienceService.Delete(idReq)
	if err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusInternalServerError, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusInternalServerError, baseResponse)
	}

	baseResponse := dto.ConvertToBaseResponse("success delete coach experience", http.StatusOK, dto.EmptyObj{})
	return c.JSON(http.StatusOK, baseResponse)
}
