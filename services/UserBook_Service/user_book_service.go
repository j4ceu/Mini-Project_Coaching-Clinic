package UserBook_Service

import (
	"Mini-Project_Coaching-Clinic/dto"
	"Mini-Project_Coaching-Clinic/dto/payload"
	"Mini-Project_Coaching-Clinic/helper"
	"Mini-Project_Coaching-Clinic/models"
	"Mini-Project_Coaching-Clinic/repositories/CoachAvailability_Repository"
	"Mini-Project_Coaching-Clinic/repositories/Coach_Repository"
	"Mini-Project_Coaching-Clinic/repositories/UserBook_Repository"
	"Mini-Project_Coaching-Clinic/repositories/UserPayment_Repository"
)

type UserBookService interface {
	FindAll() ([]dto.UserBookResponse, error)
	FindByUserID(id string) ([]dto.UserBookResponse, error)
	FindByID(id string) (dto.UserBookResponse, error)
	Create(userBook payload.UserBookPayloadCreate) (dto.UserBookResponse, error)
	Update(userBook payload.UserBookPayloadUpdate, id string) (dto.UserBookResponse, error)
	Delete(id string) (models.UserBook, error)
}

type userBookService struct {
	userBookRepo          UserBook_Repository.UserBookRepository
	userPaymentRepo       UserPayment_Repository.UserPaymentRepository
	coachAvailabilityRepo CoachAvailability_Repository.CoachAvailabilityRepositories
	coachRepo             Coach_Repository.CoachRepositories
}

func NewUserBookServices(userBookRepo UserBook_Repository.UserBookRepository, userPaymentRepo UserPayment_Repository.UserPaymentRepository, coachAvailabilityRepo CoachAvailability_Repository.CoachAvailabilityRepositories, coachRepo Coach_Repository.CoachRepositories) *userBookService {
	return &userBookService{userBookRepo, userPaymentRepo, coachAvailabilityRepo, coachRepo}
}

func (s *userBookService) FindAll() ([]dto.UserBookResponse, error) {
	userBooks, err := s.userBookRepo.FindAll()
	if err != nil {
		return []dto.UserBookResponse{}, err
	}

	var userBookResponses []dto.UserBookResponse

	for _, userBook := range userBooks {
		userBookResponse := dto.UserBookResponse{
			ID:                  userBook.ID.String(),
			Title:               userBook.Title,
			CoachAvailabilityID: userBook.CoachAvailabilityID,
			UserPaymentID:       userBook.UserPaymentID,
			Summary:             userBook.Summary,
			Done:                userBook.Done,
		}
		userBookResponses = append(userBookResponses, userBookResponse)
	}

	return userBookResponses, nil
}

func (s *userBookService) FindByUserID(id string) ([]dto.UserBookResponse, error) {
	userBooks, err := s.userPaymentRepo.FindUserBookByUserID(id)
	if err != nil {
		return []dto.UserBookResponse{}, err
	}
	var userBookResponses []dto.UserBookResponse

	for _, userBook := range userBooks {
		userBookResponse := dto.UserBookResponse{
			ID:                  userBook.ID.String(),
			Title:               userBook.Title,
			CoachAvailabilityID: userBook.CoachAvailabilityID,
			UserPaymentID:       userBook.UserPaymentID,
			Summary:             userBook.Summary,
			Done:                userBook.Done,
		}
		userBookResponses = append(userBookResponses, userBookResponse)
	}
	return userBookResponses, nil
}

func (s *userBookService) FindByID(id string) (dto.UserBookResponse, error) {
	userBook, err := s.userBookRepo.FindByID(id)
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
		UserPaymentID:       userBook.UserPaymentID,
		Done:                userBook.Done,
		CoachAvailability:   &coachAvailabilityResponse,
	}
	return userBookResponse, nil
}

func (s *userBookService) Create(userBookPayload payload.UserBookPayloadCreate) (dto.UserBookResponse, error) {

	userBookModel := models.UserBook{
		Title:               userBookPayload.Title,
		CoachAvailabilityID: userBookPayload.CoachAvailabilityID,
		UserPaymentID:       userBookPayload.UserPaymentID,
	}

	userBook, err := s.userBookRepo.Create(userBookModel)
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
		UserPaymentID:       userBook.UserPaymentID,
		CoachAvailability:   &coachAvailabilityResponse,
	}

	return userBookResponse, nil
}

func (s *userBookService) Update(userBookPayload payload.UserBookPayloadUpdate, id string) (dto.UserBookResponse, error) {

	userBookModel := models.UserBook{
		Title:               userBookPayload.Title,
		CoachAvailabilityID: userBookPayload.CoachAvailabilityID,
		UserPaymentID:       userBookPayload.UserPaymentID,
	}

	if userBookPayload.Done != nil {
		userBookModel.Done = *userBookPayload.Done
	}
	if userBookPayload.Summary != nil {
		fileName, buf, err := helper.OpenFileFromMultipartForm(userBookPayload.Summary)
		if err != nil {
			return dto.UserBookResponse{}, err
		}

		url := helper.UploadFileToFirebase(*buf, fileName)

		userBookModel.Summary = url
	}

	userBook, err := s.userBookRepo.Update(userBookModel, id)
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
		UserPaymentID:       userBook.UserPaymentID,
		CoachAvailability:   &coachAvailabilityResponse,
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
