package services

import (
	"Mini-Project_Coaching-Clinic/models"
	"Mini-Project_Coaching-Clinic/repositories"
)

type GameService interface {
	FindByID(id uint) (models.Game, error)
	FindAll() ([]models.Game, error)
	Create(game models.Game) (models.Game, error)
	Update(game models.Game, id uint) (models.Game, error)
	Delete(id uint) (models.Game, error)
}

type gameService struct {
	gameRepo repositories.GameRepositories
}

func NewGameService(gameRepo repositories.GameRepositories) *gameService {
	return &gameService{gameRepo}
}

func (s *gameService) FindByID(id uint) (models.Game, error) {
	game, err := s.gameRepo.FindByID(id)
	if err != nil {
		return game, err
	}
	return game, nil
}

func (s *gameService) FindAll() ([]models.Game, error) {
	games, err := s.gameRepo.FindAll()
	if err != nil {
		return games, err
	}
	return games, nil
}

func (s *gameService) Create(game models.Game) (models.Game, error) {
	game, err := s.gameRepo.Create(game)
	if err != nil {
		return game, err
	}
	return game, nil
}

func (s *gameService) Update(game models.Game, id uint) (models.Game, error) {
	game, err := s.gameRepo.Update(game, id)
	if err != nil {
		return game, err
	}
	return game, nil
}

func (s *gameService) Delete(id uint) (models.Game, error) {
	game, err := s.gameRepo.Delete(id)
	if err != nil {
		return game, err
	}
	return game, nil
}

// Path: services\game_service.go
