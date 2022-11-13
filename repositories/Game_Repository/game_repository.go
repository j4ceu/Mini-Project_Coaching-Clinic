package Game_Repository

import (
	"Mini-Project_Coaching-Clinic/models"

	"gorm.io/gorm"
)

type GameRepositories interface {
	FindByID(id uint) (models.Game, error)
	Create(game models.Game) (models.Game, error)
	Update(game models.Game, id uint) (models.Game, error)
	Delete(id uint) (models.Game, error)
	FindAll() ([]models.Game, error)
}

type gameRepo struct {
	db *gorm.DB
}

func NewGameRepositories(db *gorm.DB) *gameRepo {
	return &gameRepo{db}
}

func (r *gameRepo) FindByID(id uint) (models.Game, error) {
	var game models.Game
	err := r.db.Where("id = ?", id).First(&game).Error
	if err != nil {
		return game, err
	}
	return game, nil
}

func (r *gameRepo) FindAll() ([]models.Game, error) {
	var games []models.Game
	err := r.db.Find(&games).Error
	if err != nil {
		return games, err
	}
	return games, nil
}

func (r *gameRepo) Create(game models.Game) (models.Game, error) {
	err := r.db.Create(&game).Error
	if err != nil {
		return game, err
	}
	return game, nil
}

func (r *gameRepo) Update(gameUpdate models.Game, id uint) (models.Game, error) {
	var game models.Game
	err := r.db.Model(&game).Where("id = ?", id).Updates(&gameUpdate).Error
	if err != nil {
		return game, err
	}
	return r.FindByID(id)
}

func (r *gameRepo) Delete(id uint) (models.Game, error) {
	var game models.Game
	err := r.db.Where("id = ?", id).Delete(&game).Error
	if err != nil {
		return game, err
	}
	return game, nil
}
