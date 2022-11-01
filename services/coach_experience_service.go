package services

import (
	"Mini-Project_Coaching-Clinic/models"
	"Mini-Project_Coaching-Clinic/repositories"
)

type CoachExperienceService interface {
	FindByID(id string) (models.CoachExperience, error)
	Create(coachExperience models.CoachExperience) (models.CoachExperience, error)
	Update(coachExperience models.CoachExperience, id string) (models.CoachExperience, error)
	Delete(id string) (models.CoachExperience, error)
}

type coachExperienceService struct {
	coachExperienceRepo repositories.CoachExperienceRepositories
}

func NewCoachExperienceService(coachExperienceRepo repositories.CoachExperienceRepositories) *coachExperienceService {
	return &coachExperienceService{coachExperienceRepo}
}

func (s *coachExperienceService) FindByID(id string) (models.CoachExperience, error) {
	coachExperience, err := s.coachExperienceRepo.FindByID(id)
	if err != nil {
		return coachExperience, err
	}
	return coachExperience, nil
}

func (s *coachExperienceService) Create(coachExperience models.CoachExperience) (models.CoachExperience, error) {
	coachExperience, err := s.coachExperienceRepo.Create(coachExperience)
	if err != nil {
		return coachExperience, err
	}
	return coachExperience, nil
}

func (s *coachExperienceService) Update(coachExperience models.CoachExperience, id string) (models.CoachExperience, error) {
	coachExperience, err := s.coachExperienceRepo.Update(coachExperience, id)
	if err != nil {
		return coachExperience, err
	}
	return coachExperience, nil
}

func (s *coachExperienceService) Delete(id string) (models.CoachExperience, error) {
	coachExperience, err := s.coachExperienceRepo.Delete(id)
	if err != nil {
		return coachExperience, err
	}
	return coachExperience, nil
}
