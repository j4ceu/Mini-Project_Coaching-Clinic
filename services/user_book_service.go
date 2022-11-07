package services

import (
	"Mini-Project_Coaching-Clinic/dto"
	"Mini-Project_Coaching-Clinic/models"
	"Mini-Project_Coaching-Clinic/repositories"
)

type UserBookService interface {
	FindAll() ([]models.UserBook, error)
	FindByUserID(id string) ([]models.UserBook, error)
	FindByID(id string) (models.UserBook, error)
	Create(UserBook models.UserBook) (dto.UserBookResponse, error)
	Update(UserBook models.UserBook, id string) (dto.UserBookResponse, error)
	Delete(id string) (models.UserBook, error)
}

type userBookService struct {
	userBookRepo          repositories.UserBookRepository
	userPaymentRepo       repositories.UserPaymentRepository
	coachAvailabilityRepo repositories.CoachAvailabilityRepositories
	coachRepo             repositories.CoachRepositories
}

func NewUserBookServices(userBookRepo repositories.UserBookRepository, userPaymentRepo repositories.UserPaymentRepository, coachAvailabilityRepo repositories.CoachAvailabilityRepositories, coachRepo repositories.CoachRepositories) *userBookService {
	return &userBookService{userBookRepo, userPaymentRepo, coachAvailabilityRepo, coachRepo}
}

func (s *userBookService) FindAll() ([]models.UserBook, error) {
	userBooks, err := s.userBookRepo.FindAll()
	if err != nil {
		return userBooks, err
	}
	return userBooks, nil
}

func (s *userBookService) FindByUserID(id string) ([]models.UserBook, error) {
	userBooks, err := s.userPaymentRepo.FindUserBookByUserID(id)
	if err != nil {
		return userBooks, err
	}
	return userBooks, nil
}

func (s *userBookService) FindByID(id string) (models.UserBook, error) {
	userBook, err := s.userBookRepo.FindByID(id)
	if err != nil {
		return userBook, err
	}
	return userBook, nil
}

func (s *userBookService) Create(userBook models.UserBook) (dto.UserBookResponse, error) {
	userBook, err := s.userBookRepo.Create(userBook)
	if err != nil {
		return dto.UserBookResponse{}, err
	}

	// Get Coach Avail
	getCoachAvailability, err := s.coachAvailabilityRepo.FindByID(userBook.CoachAvailabilityID)
	if err != nil {
		return dto.UserBookResponse{}, err
	}

	// Update Coach Avail
	getCoachAvailability.Book = true
	_, err = s.coachAvailabilityRepo.Update(getCoachAvailability, getCoachAvailability.ID.String())
	if err != nil {
		return dto.UserBookResponse{}, err
	}

	userBook.CoachAvailability = getCoachAvailability

	// Get Coach
	getCoach, err := s.coachRepo.FindByID(getCoachAvailability.CoachID)
	if err != nil {
		return dto.UserBookResponse{}, err
	}
	// Get User Payment
	var userPayment models.UserPayment
	userPayment.Amount += getCoach.Price

	// Update User Payment
	_, err = s.userPaymentRepo.Update(userPayment, userBook.UserPaymentID)
	if err != nil {
		return dto.UserBookResponse{}, err
	}

	coachAvailabilityResponse := dto.CoachAvailabilityResponse{
		ID:        getCoachAvailability.ID.String(),
		CoachID:   getCoachAvailability.CoachID,
		Day:       getCoachAvailability.Day,
		StartTime: getCoachAvailability.StartTime,
		EndTime:   getCoachAvailability.EndTime,
	}

	userBookResponse := dto.UserBookResponse{
		ID:                  userBook.ID.String(),
		Title:               userBook.Title,
		CoachAvailabilityID: userBook.CoachAvailabilityID,
		Summary:             userBook.Summary,
		Done:                userBook.Done,
		CoachAvailability:   coachAvailabilityResponse,
	}

	return userBookResponse, nil
}

func (s *userBookService) Update(userBookUpdate models.UserBook, id string) (dto.UserBookResponse, error) {
	userBook, err := s.userBookRepo.Update(userBookUpdate, id)
	if err != nil {
		return dto.UserBookResponse{}, err
	}
	coachAvailabilityResponse := dto.CoachAvailabilityResponse{
		ID:        userBook.CoachAvailability.ID.String(),
		CoachID:   userBook.CoachAvailability.CoachID,
		Day:       userBook.CoachAvailability.Day,
		StartTime: userBook.CoachAvailability.StartTime,
		EndTime:   userBook.CoachAvailability.EndTime,
	}

	userBookResponse := dto.UserBookResponse{
		ID:                  userBook.ID.String(),
		Title:               userBook.Title,
		CoachAvailabilityID: userBook.CoachAvailabilityID,
		Summary:             userBook.Summary,
		Done:                userBook.Done,
		CoachAvailability:   coachAvailabilityResponse,
	}

	return userBookResponse, nil
}

func (s *userBookService) Delete(id string) (models.UserBook, error) {
	userBook, err := s.userBookRepo.Delete(id)
	if err != nil {
		return userBook, err
	}
	return userBook, nil
}
