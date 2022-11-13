package Coach_Service

import (
	"Mini-Project_Coaching-Clinic/dto"
	"Mini-Project_Coaching-Clinic/models"
	"Mini-Project_Coaching-Clinic/repositories/Coach_Repository"
	"Mini-Project_Coaching-Clinic/repositories/Game_Repository"
	"Mini-Project_Coaching-Clinic/repositories/User_Repository"
	"strings"
)

type CoachService interface {
	FindByID(id string) (dto.CoachResponse, error)
	FindByCode(code string) (dto.CoachResponse, error)
	FindByGameID(gameID string) ([]dto.CoachResponse, error)
	Create(coach models.Coach) (dto.CoachResponse, error)
	Update(coach models.Coach, id string) (dto.CoachResponse, error)
	Delete(id string) (models.Coach, error)
}

type coachService struct {
	coachRepo Coach_Repository.CoachRepositories
	gameRepo  Game_Repository.GameRepositories
	userRepo  User_Repository.UserRepositories
}

func NewCoachService(coachRepo Coach_Repository.CoachRepositories, gameRepo Game_Repository.GameRepositories, userRepo User_Repository.UserRepositories) *coachService {
	return &coachService{coachRepo, gameRepo, userRepo}
}

func (s *coachService) FindByID(id string) (dto.CoachResponse, error) {
	coach, err := s.coachRepo.FindByID(id)
	if err != nil {
		return dto.CoachResponse{}, err
	}

	var coachAvailabilities []dto.CoachAvailabilityResponse

	for _, ca := range coach.CoachAvailability {
		coachAvailability := dto.CoachAvailabilityResponse{
			ID:        ca.ID.String(),
			Day:       ca.Day,
			StartTime: ca.StartTime,
			EndTime:   ca.EndTime,
		}
		coachAvailabilities = append(coachAvailabilities, coachAvailability)
	}

	coachResponse := dto.CoachResponse{
		ID:                coach.ID.String(),
		FirstName:         coach.User.FirstName,
		LastName:          coach.User.LastName,
		Position:          coach.Position,
		Code:              coach.Code,
		Price:             coach.Price,
		UserID:            coach.UserID,
		GameID:            coach.GameID,
		CoachExperience:   coach.CoachExperience,
		CoachAvailability: coachAvailabilities,
	}

	return coachResponse, nil
}

func (s *coachService) FindByCode(code string) (dto.CoachResponse, error) {
	coach, err := s.coachRepo.FindByCode(code)
	if err != nil {
		return dto.CoachResponse{}, err
	}
	var coachAvailabilities []dto.CoachAvailabilityResponse

	for _, ca := range coach.CoachAvailability {
		coachAvailability := dto.CoachAvailabilityResponse{
			ID:        ca.ID.String(),
			Day:       ca.Day,
			StartTime: ca.StartTime,
			EndTime:   ca.EndTime,
		}
		coachAvailabilities = append(coachAvailabilities, coachAvailability)
	}

	coachResponse := dto.CoachResponse{
		ID:                coach.ID.String(),
		FirstName:         coach.User.FirstName,
		LastName:          coach.User.LastName,
		Position:          coach.Position,
		Code:              coach.Code,
		Price:             coach.Price,
		UserID:            coach.UserID,
		GameID:            coach.GameID,
		CoachExperience:   coach.CoachExperience,
		CoachAvailability: coachAvailabilities,
	}
	return coachResponse, nil
}

func (s *coachService) FindByGameID(gameID string) ([]dto.CoachResponse, error) {
	coaches, err := s.coachRepo.FindByGameID(gameID)
	if err != nil {
		return []dto.CoachResponse{}, err
	}

	var coachResponses []dto.CoachResponse
	for _, coach := range coaches {
		var coachAvailabilities []dto.CoachAvailabilityResponse
		for _, ca := range coach.CoachAvailability {

			coachAvailability := dto.CoachAvailabilityResponse{
				ID:        ca.ID.String(),
				Day:       ca.Day,
				StartTime: ca.StartTime,
				EndTime:   ca.EndTime,
			}
			coachAvailabilities = append(coachAvailabilities, coachAvailability)
		}
		coachResponse := dto.CoachResponse{
			ID:                coach.ID.String(),
			FirstName:         coach.User.FirstName,
			LastName:          coach.User.LastName,
			Position:          coach.Position,
			Code:              coach.Code,
			Price:             coach.Price,
			UserID:            coach.UserID,
			GameID:            coach.GameID,
			CoachExperience:   coach.CoachExperience,
			CoachAvailability: coachAvailabilities,
		}
		coachResponses = append(coachResponses, coachResponse)
	}

	return coachResponses, nil
}

func (s *coachService) Create(coach models.Coach) (dto.CoachResponse, error) {
	var coachResponse dto.CoachResponse
	game, err := s.gameRepo.FindByID(coach.GameID)
	if err != nil {
		return coachResponse, err
	}

	if coach.UserID != "" {
		user, err := s.userRepo.FindByID(coach.UserID)
		if err != nil {
			return coachResponse, err
		}

		first1 := user.FirstName[0:3]
		last1 := user.LastName[len(user.LastName)-2:]
		code := strings.ToUpper(first1 + last1 + game.Name)

		coach.Code = strings.ReplaceAll(code, " ", "")

		coachResponse.FirstName = user.FirstName
		coachResponse.LastName = user.LastName
		coachResponse.Email = user.Email

		user.Role = "Coach"

		_, err = s.userRepo.Update(user, user.ID.String())
		if err != nil {
			return coachResponse, err
		}

	} else {
		first1 := coach.User.FirstName[0:3]
		last1 := coach.User.LastName[len(coach.User.LastName)-3:]
		code := strings.ToUpper(first1 + last1 + game.Name)

		coach.Code = strings.ReplaceAll(code, " ", "")
		coachResponse.FirstName = coach.User.FirstName
		coachResponse.LastName = coach.User.LastName
		coachResponse.Email = coach.User.Email

		coach.User.Role = "Coach"

		if err := coach.User.HashPassword(coach.User.Password); err != nil {
			return coachResponse, err
		}

	}

	coach, err = s.coachRepo.Create(coach)
	if err != nil {
		return coachResponse, err
	}

	var coachAvailabilities []dto.CoachAvailabilityResponse

	for _, ca := range coach.CoachAvailability {
		coachAvailability := dto.CoachAvailabilityResponse{
			ID:        ca.ID.String(),
			Day:       ca.Day,
			StartTime: ca.StartTime,
			EndTime:   ca.EndTime,
		}
		coachAvailabilities = append(coachAvailabilities, coachAvailability)
	}

	coachResponse = dto.CoachResponse{
		ID:                coach.ID.String(),
		FirstName:         coachResponse.FirstName,
		LastName:          coachResponse.LastName,
		Email:             coachResponse.Email,
		Position:          coach.Position,
		Code:              coach.Code,
		Price:             coach.Price,
		UserID:            coach.UserID,
		GameID:            coach.GameID,
		CoachExperience:   coach.CoachExperience,
		CoachAvailability: coachAvailabilities,
	}

	return coachResponse, nil
}

func (s *coachService) Update(coach models.Coach, id string) (dto.CoachResponse, error) {
	coach, err := s.coachRepo.Update(coach, id)
	if err != nil {
		return dto.CoachResponse{}, err
	}

	var coachAvailabilities []dto.CoachAvailabilityResponse

	for _, ca := range coach.CoachAvailability {
		coachAvailability := dto.CoachAvailabilityResponse{
			ID:        ca.ID.String(),
			Day:       ca.Day,
			StartTime: ca.StartTime,
			EndTime:   ca.EndTime,
		}
		coachAvailabilities = append(coachAvailabilities, coachAvailability)
	}

	coachResponse := dto.CoachResponse{
		ID:                coach.ID.String(),
		FirstName:         coach.User.FirstName,
		LastName:          coach.User.LastName,
		Email:             coach.User.Email,
		Position:          coach.Position,
		Code:              coach.Code,
		Price:             coach.Price,
		UserID:            coach.UserID,
		GameID:            coach.GameID,
		CoachExperience:   coach.CoachExperience,
		CoachAvailability: coachAvailabilities,
	}

	return coachResponse, nil
}

func (s *coachService) Delete(id string) (models.Coach, error) {
	coach, err := s.coachRepo.Delete(id)
	if err != nil {
		return coach, err
	}
	return coach, nil
}
