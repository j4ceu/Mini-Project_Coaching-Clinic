package services

import (
	"Mini-Project_Coaching-Clinic/models"
	"Mini-Project_Coaching-Clinic/repositories"
)

type CoachAvailibilityService interface {
	FindByID(id string) (models.CoachAvailibility, error)
	Create(coachAvailibility models.CoachAvailibility) (models.CoachAvailibility, error)
	Update(coachAvailibility models.CoachAvailibility, id string) (models.CoachAvailibility, error)
	Delete(id string) (models.CoachAvailibility, error)
}

type coachAvailibilityService struct {
	coachAvailibilityRepo repositories.CoachAvailibilityRepositories
}

func NewCoachAvailibilityService(coachAvailibilityRepo repositories.CoachAvailibilityRepositories) *coachAvailibilityService {
	return &coachAvailibilityService{coachAvailibilityRepo}
}

func (s *coachAvailibilityService) FindByID(id string) (models.CoachAvailibility, error) {
	coachAvailibility, err := s.coachAvailibilityRepo.FindByID(id)
	if err != nil {
		return coachAvailibility, err
	}
	return coachAvailibility, nil
}

func (s *coachAvailibilityService) Create(coachAvailibility models.CoachAvailibility) (models.CoachAvailibility, error) {
	coachAvailibility, err := s.coachAvailibilityRepo.Create(coachAvailibility)
	if err != nil {
		return coachAvailibility, err
	}
	return coachAvailibility, nil
}

func (s *coachAvailibilityService) Update(coachAvailibility models.CoachAvailibility, id string) (models.CoachAvailibility, error) {
	coachAvailibility, err := s.coachAvailibilityRepo.Update(coachAvailibility, id)
	if err != nil {
		return coachAvailibility, err
	}
	return coachAvailibility, nil
}

func (s *coachAvailibilityService) Delete(id string) (models.CoachAvailibility, error) {
	coachAvailibility, err := s.coachAvailibilityRepo.Delete(id)
	if err != nil {
		return coachAvailibility, err
	}
	return coachAvailibility, nil
}
