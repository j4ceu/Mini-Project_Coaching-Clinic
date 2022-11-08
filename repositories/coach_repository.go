package repositories

import (
	"Mini-Project_Coaching-Clinic/models"

	"gorm.io/gorm"
)

type CoachRepositories interface {
	FindByID(id string) (models.Coach, error)
	FindByGameID(gameID string) ([]models.Coach, error)
	Create(coach models.Coach) (models.Coach, error)
	Update(coach models.Coach, id string) (models.Coach, error)
	Delete(id string) (models.Coach, error)
	FindByCode(code string) (models.Coach, error)
}

type coachRepo struct {
	db *gorm.DB
}

func NewCoachRepositories(db *gorm.DB) *coachRepo {
	return &coachRepo{db}
}

func (r *coachRepo) FindByID(id string) (models.Coach, error) {
	var coach models.Coach
	err := r.db.Preload("User").Preload("Game").Preload("CoachExperience").Preload("CoachAvailability").Where("id = ?", id).First(&coach).Error
	if err != nil {
		return coach, err
	}
	return coach, nil
}

func (r *coachRepo) FindByGameID(gameID string) ([]models.Coach, error) {
	var coaches []models.Coach
	err := r.db.Where("game_id = ?", gameID).Preload("User").Preload("Game").Preload("CoachExperience").Preload("CoachAvailability").Find(&coaches).Error
	if err != nil {
		return coaches, err
	}
	return coaches, nil
}

func (r *coachRepo) Create(coach models.Coach) (models.Coach, error) {
	err := r.db.Create(&coach).Error
	if err != nil {
		return coach, err
	}
	return coach, nil
}

func (r *coachRepo) Update(coachUpdate models.Coach, id string) (models.Coach, error) {
	var coach models.Coach
	err := r.db.Model(&coach).Where("id = ?", id).Updates(&coachUpdate).Error
	if err != nil {
		return coach, err
	}
	return r.FindByID(id)
}

func (r *coachRepo) Delete(id string) (models.Coach, error) {
	var coach models.Coach
	err := r.db.Where("id = ?", id).Delete(&coach).Error
	if err != nil {
		return coach, err
	}
	return coach, nil

}

func (r *coachRepo) FindByCode(code string) (models.Coach, error) {
	var coach models.Coach
	err := r.db.Where("code = ?", code).Preload("User").Preload("Game").Preload("CoachExperience").Preload("CoachAvailability").First(&coach).Error
	if err != nil {
		return coach, err
	}
	return coach, nil
}
