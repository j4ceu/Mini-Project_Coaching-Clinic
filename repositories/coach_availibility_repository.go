package repositories

import (
	"Mini-Project_Coaching-Clinic/models"

	"gorm.io/gorm"
)

type CoachAvailabilityRepositories interface {
	FindByID(id string) (models.CoachAvailability, error)
	Create(CoachAvailability models.CoachAvailability) (models.CoachAvailability, error)
	Update(CoachAvailability models.CoachAvailability, id string) (models.CoachAvailability, error)
	Delete(id string) (models.CoachAvailability, error)
}

type CoachAvailabilityRepo struct {
	db *gorm.DB
}

func NewCoachAvailabilityRepositories(db *gorm.DB) *CoachAvailabilityRepo {
	return &CoachAvailabilityRepo{db}
}

func (r *CoachAvailabilityRepo) FindByID(id string) (models.CoachAvailability, error) {
	var CoachAvailability models.CoachAvailability
	err := r.db.Where("id = ?", id).First(&CoachAvailability).Error
	if err != nil {
		return CoachAvailability, err
	}
	return CoachAvailability, nil
}

func (r *CoachAvailabilityRepo) Create(CoachAvailability models.CoachAvailability) (models.CoachAvailability, error) {
	err := r.db.Create(&CoachAvailability).Error
	if err != nil {
		return CoachAvailability, err
	}
	return CoachAvailability, nil
}

func (r *CoachAvailabilityRepo) Update(CoachAvailabilityUpdate models.CoachAvailability, id string) (models.CoachAvailability, error) {
	var CoachAvailability models.CoachAvailability
	err := r.db.Model(&CoachAvailability).Where("id = ?", id).Updates(&CoachAvailabilityUpdate).Error
	if err != nil {
		return CoachAvailability, err
	}
	return r.FindByID(id)
}

func (r *CoachAvailabilityRepo) Delete(id string) (models.CoachAvailability, error) {
	var CoachAvailability models.CoachAvailability
	err := r.db.Where("id = ?", id).Delete(&CoachAvailability).Error
	if err != nil {
		return CoachAvailability, err
	}
	return CoachAvailability, nil
}
