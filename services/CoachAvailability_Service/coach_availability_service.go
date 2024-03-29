package CoachAvailability_Service

import (
	"Mini-Project_Coaching-Clinic/dto"
	"Mini-Project_Coaching-Clinic/models"
	"Mini-Project_Coaching-Clinic/repositories/CoachAvailability_Repository"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type CoachAvailabilityService interface {
	FindByID(id string) (dto.CoachAvailabilityResponse, error)
	Create(coachAvailability models.CoachAvailability) (dto.CoachAvailabilityResponse, error)
	Update(coachAvailability models.CoachAvailability, id string) (dto.CoachAvailabilityResponse, error)
	Delete(id string) (models.CoachAvailability, error)
	CheckInterval(coachAvailability models.CoachAvailability) (bool, error)
}

type coachAvailabilityService struct {
	coachAvailabilityRepo CoachAvailability_Repository.CoachAvailabilityRepositories
}

func NewCoachAvailabilityService(coachAvailabilityRepo CoachAvailability_Repository.CoachAvailabilityRepositories) *coachAvailabilityService {
	return &coachAvailabilityService{coachAvailabilityRepo}
}

func (s *coachAvailabilityService) FindByID(id string) (dto.CoachAvailabilityResponse, error) {
	coachAvailability, err := s.coachAvailabilityRepo.FindByID(id)
	if err != nil {
		return dto.CoachAvailabilityResponse{}, err
	}
	coachAvailabilityResponse := dto.CoachAvailabilityResponse{
		ID:        coachAvailability.ID.String(),
		StartTime: coachAvailability.StartTime,
		EndTime:   coachAvailability.EndTime,
		Day:       coachAvailability.Day,
		CoachID:   coachAvailability.CoachID,
	}
	return coachAvailabilityResponse, nil
}

func (s *coachAvailabilityService) Create(coachAvailability models.CoachAvailability) (dto.CoachAvailabilityResponse, error) {
	_, err := s.CheckInterval(coachAvailability)
	if err != nil {
		return dto.CoachAvailabilityResponse{}, err
	}

	coachAvailability, err = s.coachAvailabilityRepo.Create(coachAvailability)
	if err != nil {
		return dto.CoachAvailabilityResponse{}, err
	}
	coachAvailabilityResponse := dto.CoachAvailabilityResponse{
		ID:        coachAvailability.ID.String(),
		StartTime: coachAvailability.StartTime,
		EndTime:   coachAvailability.EndTime,
		Day:       coachAvailability.Day,
		CoachID:   coachAvailability.CoachID,
	}

	return coachAvailabilityResponse, nil
}

func (s *coachAvailabilityService) Update(coachAvailability models.CoachAvailability, id string) (dto.CoachAvailabilityResponse, error) {
	if coachAvailability.StartTime != "" && coachAvailability.EndTime != "" {
		_, err := s.CheckInterval(coachAvailability)
		if err != nil {
			return dto.CoachAvailabilityResponse{}, err
		}
	}

	coachAvailability, err := s.coachAvailabilityRepo.Update(coachAvailability, id)
	if err != nil {
		return dto.CoachAvailabilityResponse{}, err
	}

	coachAvailabilityResponse := dto.CoachAvailabilityResponse{
		ID:        coachAvailability.ID.String(),
		StartTime: coachAvailability.StartTime,
		EndTime:   coachAvailability.EndTime,
		Day:       coachAvailability.Day,
		CoachID:   coachAvailability.CoachID,
	}

	return coachAvailabilityResponse, nil
}

func (s *coachAvailabilityService) Delete(id string) (models.CoachAvailability, error) {
	coachAvailability, err := s.coachAvailabilityRepo.Delete(id)
	if err != nil {
		return coachAvailability, err
	}
	return coachAvailability, nil
}

func (s *coachAvailabilityService) CheckInterval(coachAvailability models.CoachAvailability) (bool, error) {
	startTime, err := strconv.ParseFloat(strings.ReplaceAll(coachAvailability.StartTime, ":", ""), 64)
	if err != nil {
		return false, err
	}

	startTime = startTime / 100

	endTime, err := strconv.ParseFloat(strings.ReplaceAll(coachAvailability.EndTime, ":", ""), 64)
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
