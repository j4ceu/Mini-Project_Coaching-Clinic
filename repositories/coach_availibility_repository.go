package repositories

import (
	"Mini-Project_Coaching-Clinic/models"

	"gorm.io/gorm"
)

type CoachAvailibilityRepositories interface {
	FindByID(id string) (models.CoachAvailibility, error)
	Create(CoachAvailibility models.CoachAvailibility) (models.CoachAvailibility, error)
	Update(CoachAvailibility models.CoachAvailibility, id string) (models.CoachAvailibility, error)
	Delete(id string) (models.CoachAvailibility, error)
}

type CoachAvailibilityRepo struct {
	db *gorm.DB
}

func NewCoachAvailibilityRepositories(db *gorm.DB) *CoachAvailibilityRepo {
	return &CoachAvailibilityRepo{db}
}

func (r *CoachAvailibilityRepo) FindByID(id string) (models.CoachAvailibility, error) {
	var CoachAvailibility models.CoachAvailibility
	err := r.db.Where("id = ?", id).First(&CoachAvailibility).Error
	if err != nil {
		return CoachAvailibility, err
	}
	return CoachAvailibility, nil
}

func (r *CoachAvailibilityRepo) Create(CoachAvailibility models.CoachAvailibility) (models.CoachAvailibility, error) {
	err := r.db.Create(&CoachAvailibility).Error
	if err != nil {
		return CoachAvailibility, err
	}
	return CoachAvailibility, nil
}

func (r *CoachAvailibilityRepo) Update(CoachAvailibilityUpdate models.CoachAvailibility, id string) (models.CoachAvailibility, error) {
	var CoachAvailibility models.CoachAvailibility
	err := r.db.Model(&CoachAvailibility).Where("id = ?", id).Updates(&CoachAvailibilityUpdate).Error
	if err != nil {
		return CoachAvailibility, err
	}
	return r.FindByID(id)
}

func (r *CoachAvailibilityRepo) Delete(id string) (models.CoachAvailibility, error) {
	var CoachAvailibility models.CoachAvailibility
	err := r.db.Where("id = ?", id).Delete(&CoachAvailibility).Error
	if err != nil {
		return CoachAvailibility, err
	}
	return CoachAvailibility, nil
}
