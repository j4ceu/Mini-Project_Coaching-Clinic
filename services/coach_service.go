package services

import (
	"Mini-Project_Coaching-Clinic/dto"
	"Mini-Project_Coaching-Clinic/models"
	"Mini-Project_Coaching-Clinic/repositories"
	"strings"
)

type CoachService interface {
	FindByID(id string) (models.Coach, error)
	FindByGameID(gameID string) ([]models.Coach, error)
	Create(coach models.Coach) (dto.CoachResponse, error)
	Update(coach models.Coach, id string) (dto.CoachResponse, error)
	Delete(id string) (models.Coach, error)
}

type coachService struct {
	coachRepo repositories.CoachRepositories
	gameRepo  repositories.GameRepositories
	userRepo  repositories.UserRepositories
}

func NewCoachService(coachRepo repositories.CoachRepositories, gameRepo repositories.GameRepositories, userRepo repositories.UserRepositories) *coachService {
	return &coachService{coachRepo, gameRepo, userRepo}
}

func (s *coachService) FindByID(id string) (models.Coach, error) {
	coach, err := s.coachRepo.FindByID(id)
	if err != nil {
		return coach, err
	}
	return coach, nil
}

func (s *coachService) FindByGameID(gameID string) ([]models.Coach, error) {
	coaches, err := s.coachRepo.FindByGameID(gameID)
	if err != nil {
		return coaches, err
	}
	return coaches, nil
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

	}

	coach, err = s.coachRepo.Create(coach)
	if err != nil {
		return coachResponse, err
	}

	coachResponse = dto.CoachResponse{
		ID:               coach.ID.String(),
		FirstName:        coachResponse.FirstName,
		LastName:         coachResponse.LastName,
		Email:            coachResponse.Email,
		Position:         coach.Position,
		Code:             coach.Code,
		Price:            coach.Price,
		UserID:           coach.UserID,
		GameID:           coach.GameID,
		CoachExperience:  coach.CoachExperience,
		CoachAvalibility: coach.CoachAvailibility,
	}

	return coachResponse, nil
}

func (s *coachService) Update(coach models.Coach, id string) (dto.CoachResponse, error) {
	coach, err := s.coachRepo.Update(coach, id)
	if err != nil {
		return dto.CoachResponse{}, err
	}

	coachResponse := dto.CoachResponse{
		ID:               coach.ID.String(),
		FirstName:        coach.User.FirstName,
		LastName:         coach.User.LastName,
		Email:            coach.User.Email,
		Position:         coach.Position,
		Code:             coach.Code,
		Price:            coach.Price,
		UserID:           coach.UserID,
		GameID:           coach.GameID,
		CoachExperience:  coach.CoachExperience,
		CoachAvalibility: coach.CoachAvailibility,
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
