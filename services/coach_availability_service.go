package services

import (
	"Mini-Project_Coaching-Clinic/models"
	"Mini-Project_Coaching-Clinic/repositories"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type CoachAvailabilityService interface {
	FindByID(id string) (models.CoachAvailability, error)
	Create(coachAvailability models.CoachAvailability) (models.CoachAvailability, error)
	Update(coachAvailability models.CoachAvailability, id string) (models.CoachAvailability, error)
	Delete(id string) (models.CoachAvailability, error)
	CheckInterval(coachAvailability models.CoachAvailability) (bool, error)
}

type coachAvailabilityService struct {
	coachAvailabilityRepo repositories.CoachAvailabilityRepositories
}

func NewCoachAvailabilityService(coachAvailabilityRepo repositories.CoachAvailabilityRepositories) *coachAvailabilityService {
	return &coachAvailabilityService{coachAvailabilityRepo}
}

func (s *coachAvailabilityService) FindByID(id string) (models.CoachAvailability, error) {
	coachAvailability, err := s.coachAvailabilityRepo.FindByID(id)
	if err != nil {
		return coachAvailability, err
	}
	return coachAvailability, nil
}

func (s *coachAvailabilityService) Create(coachAvailability models.CoachAvailability) (models.CoachAvailability, error) {
	_, err := s.CheckInterval(coachAvailability)
	if err != nil {
		return coachAvailability, err
	}

	coachAvailability, err = s.coachAvailabilityRepo.Create(coachAvailability)
	if err != nil {
		return coachAvailability, err
	}
	return coachAvailability, nil
}

func (s *coachAvailabilityService) Update(coachAvailability models.CoachAvailability, id string) (models.CoachAvailability, error) {
	if coachAvailability.StartTime != "" && coachAvailability.EndTime != "" {
		_, err := s.CheckInterval(coachAvailability)
		if err != nil {
			return coachAvailability, err
		}
	}

	coachAvailability, err := s.coachAvailabilityRepo.Update(coachAvailability, id)
	if err != nil {
		return coachAvailability, err
	}
	return coachAvailability, nil
}

func (s *coachAvailabilityService) Delete(id string) (models.CoachAvailability, error) {
	coachAvailability, err := s.coachAvailabilityRepo.Delete(id)
	if err != nil {
		return coachAvailability, err
	}
	return coachAvailability, nil
}

func (s *coachAvailabilityService) CheckInterval(coachAvailability models.CoachAvailability) (bool, error) {
	startTime, err := strconv.Atoi(strings.ReplaceAll(coachAvailability.StartTime, ":", ""))
	if err != nil {
		return false, err
	}

	startTime = startTime / 100

	endTime, err := strconv.Atoi(strings.ReplaceAll(coachAvailability.EndTime, ":", ""))
	if err != nil {
		return false, err
	}

	endTime = endTime / 100

	if endTime-startTime != 1 {
		return false, errors.New("Time interval must be 1 hour")
	}

	fmt.Println(startTime, endTime)

	return true, nil
}
