package controllers

import (
	"Mini-Project_Coaching-Clinic/dto"
	"Mini-Project_Coaching-Clinic/models"
	"Mini-Project_Coaching-Clinic/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type GameController interface {
	CreateGame(c echo.Context) error
	UpdateGame(c echo.Context) error
	DeleteGame(c echo.Context) error
	FindGameByID(c echo.Context) error
	FindAllGame(c echo.Context) error
}

type gameController struct {
	gameService services.GameService
}

func NewGameController(gameService services.GameService) *gameController {
	return &gameController{gameService}
}

func (g *gameController) CreateGame(c echo.Context) error {
	var game models.Game

	if err := c.Bind(&game); err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusBadRequest, baseResponse)
	}

	game, err := g.gameService.Create(game)
	if err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusInternalServerError, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusInternalServerError, baseResponse)
	}

	baseResponse := dto.ConvertToBaseResponse("success create game", http.StatusOK, game)
	return c.JSON(http.StatusOK, baseResponse)
}

func (g *gameController) UpdateGame(c echo.Context) error {
	var game models.Game

	idReq := c.Param("id")

	if idReq == "" {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, dto.EmptyObj{}, "id is required")
		return c.JSON(http.StatusBadRequest, baseResponse)
	}

	id, err := strconv.Atoi(idReq)
	if err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, dto.EmptyObj{}, "id must be number")
		return c.JSON(http.StatusBadRequest, baseResponse)
	}

	if err := c.Bind(&game); err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusBadRequest, baseResponse)
	}

	game, err = g.gameService.Update(game, uint(id))
	if err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusInternalServerError, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusInternalServerError, baseResponse)
	}

	baseResponse := dto.ConvertToBaseResponse("success update game", http.StatusOK, game)
	return c.JSON(http.StatusOK, baseResponse)
}

func (g *gameController) DeleteGame(c echo.Context) error {
	idReq := c.Param("id")

	if idReq == "" {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, dto.EmptyObj{}, "id is required")
		return c.JSON(http.StatusBadRequest, baseResponse)
	}

	id, err := strconv.Atoi(idReq)
	if err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, dto.EmptyObj{}, "id must be number")
		return c.JSON(http.StatusBadRequest, baseResponse)
	}

	_, err = g.gameService.Delete(uint(id))
	if err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusInternalServerError, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusInternalServerError, baseResponse)
	}

	baseResponse := dto.ConvertToBaseResponse("success delete game", http.StatusOK, dto.EmptyObj{})
	return c.JSON(http.StatusOK, baseResponse)
}

func (g *gameController) FindGameByID(c echo.Context) error {
	idReq := c.Param("id")

	if idReq == "" {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, dto.EmptyObj{}, "id is required")
		return c.JSON(http.StatusBadRequest, baseResponse)
	}

	id, err := strconv.Atoi(idReq)
	if err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, dto.EmptyObj{}, "id must be number")
		return c.JSON(http.StatusBadRequest, baseResponse)
	}

	game, err := g.gameService.FindByID(uint(id))
	if err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusInternalServerError, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusInternalServerError, baseResponse)
	}

	baseResponse := dto.ConvertToBaseResponse("success get game", http.StatusOK, game)
	return c.JSON(http.StatusOK, baseResponse)
}

func (g *gameController) FindAllGame(c echo.Context) error {
	games, err := g.gameService.FindAll()
	if err != nil {
		baseResponse := dto.ConvertErrorToBaseResponse("failed", http.StatusInternalServerError, dto.EmptyObj{}, err.Error())
		return c.JSON(http.StatusInternalServerError, baseResponse)
	}

	baseResponse := dto.ConvertToBaseResponse("success get all game", http.StatusOK, games)
	return c.JSON(http.StatusOK, baseResponse)
}
