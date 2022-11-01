package repositories

import (
	"Mini-Project_Coaching-Clinic/models"

	"gorm.io/gorm"
)

type CoachExperienceRepositories interface {
	FindByID(id string) (models.CoachExperience, error)
	Create(coachExperience models.CoachExperience) (models.CoachExperience, error)
	Update(coachExperience models.CoachExperience, id string) (models.CoachExperience, error)
	Delete(id string) (models.CoachExperience, error)
}

type coachExperienceRepo struct {
	db *gorm.DB
}

func NewCoachExperienceRepositories(db *gorm.DB) *coachExperienceRepo {
	return &coachExperienceRepo{db}
}

func (r *coachExperienceRepo) FindByID(id string) (models.CoachExperience, error) {
	var coachExperience models.CoachExperience
	err := r.db.Where("id = ?", id).First(&coachExperience).Error
	if err != nil {
		return coachExperience, err
	}
	return coachExperience, nil
}

func (r *coachExperienceRepo) Create(coachExperience models.CoachExperience) (models.CoachExperience, error) {
	err := r.db.Create(&coachExperience).Error
	if err != nil {
		return coachExperience, err
	}
	return coachExperience, nil
}

func (r *coachExperienceRepo) Update(coachExperienceUpdate models.CoachExperience, id string) (models.CoachExperience, error) {
	var coachExperience models.CoachExperience
	err := r.db.Model(&coachExperience).Where("id = ?", id).Updates(&coachExperienceUpdate).Error
	if err != nil {
		return coachExperience, err
	}
	return r.FindByID(id)
}

func (r *coachExperienceRepo) Delete(id string) (models.CoachExperience, error) {
	var coachExperience models.CoachExperience
	err := r.db.Where("id = ?", id).Delete(&coachExperience).Error
	if err != nil {
		return coachExperience, err
	}
	return coachExperience, nil
}
