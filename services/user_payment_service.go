package services

import (
	"Mini-Project_Coaching-Clinic/dto"
	"Mini-Project_Coaching-Clinic/dto/payload"
	"Mini-Project_Coaching-Clinic/helper"
	"Mini-Project_Coaching-Clinic/models"
	"Mini-Project_Coaching-Clinic/repositories"
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserPaymentService interface {
	FindAll() ([]dto.UserPaymentResponse, error)
	FindByPaidAndProofOfPaymentIsNotNull() ([]dto.UserPaymentResponse, error)
	FindByID(id string) (dto.UserPaymentResponse, error)
	FindByInvoiceNumber(invoiceNumber string) (dto.UserPaymentResponse, error)
	Create(userPayment payload.UserPaymentPayload) (dto.UserPaymentResponse, error)
	Update(userPayment payload.UserPaymentPayload, id string, claims jwt.MapClaims) (dto.UserPaymentResponse, error)
	Delete(id string) (models.UserPayment, error)
}

type userPaymentService struct {
	userPaymentRepo repositories.UserPaymentRepository
	userRepo        repositories.UserRepositories
}

func NewUserPaymentServices(userPaymentRepo repositories.UserPaymentRepository, userRepo repositories.UserRepositories) *userPaymentService {
	return &userPaymentService{userPaymentRepo, userRepo}
}

func (s *userPaymentService) FindAll() ([]dto.UserPaymentResponse, error) {
	userPayments, err := s.userPaymentRepo.FindAll()
	if err != nil {
		return []dto.UserPaymentResponse{}, err
	}

	var userPaymentResponses []dto.UserPaymentResponse

	for _, userPayment := range userPayments {

		userPaymentResponse := dto.UserPaymentResponse{
			ID:             userPayment.ID.String(),
			UserID:         userPayment.UserID,
			Email:          userPayment.User.Email,
			Invoice:        userPayment.Invoice,
			ProofOfPayment: userPayment.ProofOfPayment,
			Amount:         userPayment.Amount,
			Paid:           userPayment.Paid,
			ExpiredAt:      userPayment.ExpiredAt,
		}
		userPaymentResponses = append(userPaymentResponses, userPaymentResponse)
	}

	return userPaymentResponses, nil
}

func (s *userPaymentService) FindByInvoiceNumber(invoiceNumber string) (dto.UserPaymentResponse, error) {
	userPayment, err := s.userPaymentRepo.FindByInvoiceNumber(invoiceNumber)
	if err != nil {
		return dto.UserPaymentResponse{}, err
	}

	var userBookResponse []dto.UserBookResponse
	for _, userBook := range userPayment.UserBook {
		userBook := dto.UserBookResponse{
			ID:                  userBook.ID.String(),
			Title:               userBook.Title,
			CoachAvailabilityID: userBook.CoachAvailabilityID,
			Summary:             userBook.Summary,
			Done:                userBook.Done,
		}
		userBookResponse = append(userBookResponse, userBook)
	}

	userPaymentResponse := dto.UserPaymentResponse{
		ID:             userPayment.ID.String(),
		UserID:         userPayment.UserID,
		Email:          userPayment.User.Email,
		Invoice:        userPayment.Invoice,
		ProofOfPayment: userPayment.ProofOfPayment,
		Amount:         userPayment.Amount,
		Paid:           userPayment.Paid,
		ExpiredAt:      userPayment.ExpiredAt,
		UserBook:       &userBookResponse,
	}

	return userPaymentResponse, nil
}

func (s *userPaymentService) FindByPaidAndProofOfPaymentIsNotNull() ([]dto.UserPaymentResponse, error) {
	userPayments, err := s.userPaymentRepo.FindByPaidAndProofOfPaymentIsNotNull()
	if err != nil {
		return []dto.UserPaymentResponse{}, err
	}

	var userPaymentResponses []dto.UserPaymentResponse

	for _, userPayment := range userPayments {
		userPaymentResponse := dto.UserPaymentResponse{
			ID:             userPayment.ID.String(),
			UserID:         userPayment.UserID,
			Email:          userPayment.User.Email,
			Invoice:        userPayment.Invoice,
			ProofOfPayment: userPayment.ProofOfPayment,
			Amount:         userPayment.Amount,
			Paid:           userPayment.Paid,
			ExpiredAt:      userPayment.ExpiredAt,
		}
		userPaymentResponses = append(userPaymentResponses, userPaymentResponse)
	}
	return userPaymentResponses, nil
}

func (s *userPaymentService) FindByID(id string) (dto.UserPaymentResponse, error) {
	userPayment, err := s.userPaymentRepo.FindByID(id)
	if err != nil {
		return dto.UserPaymentResponse{}, err
	}
	var userBookResponse []dto.UserBookResponse

	for _, userBook := range userPayment.UserBook {
		userBook := dto.UserBookResponse{
			ID:                  userBook.ID.String(),
			Title:               userBook.Title,
			CoachAvailabilityID: userBook.CoachAvailabilityID,
			Summary:             userBook.Summary,
			Done:                userBook.Done,
		}
		userBookResponse = append(userBookResponse, userBook)
	}

	userPaymentResponse := dto.UserPaymentResponse{
		ID:             userPayment.ID.String(),
		UserID:         userPayment.UserID,
		Email:          userPayment.User.Email,
		Invoice:        userPayment.Invoice,
		ProofOfPayment: userPayment.ProofOfPayment,
		Amount:         userPayment.Amount,
		Paid:           userPayment.Paid,
		ExpiredAt:      userPayment.ExpiredAt,
		UserBook:       &userBookResponse,
	}
	return userPaymentResponse, nil
}

func (s *userPaymentService) Create(userPayment payload.UserPaymentPayload) (dto.UserPaymentResponse, error) {
	now := time.Now()

	userPaymentModel := models.UserPayment{
		UserID:    userPayment.UserID,
		ExpiredAt: now.Local().Add(time.Minute * 15).Format("2006-01-02 15:04:05"),
	}

	getUser, err := s.userRepo.FindByID(userPayment.UserID)
	if err != nil {
		return dto.UserPaymentResponse{}, err
	}

	userPaymentCreate, err := s.userPaymentRepo.Create(userPaymentModel)
	if err != nil {
		return dto.UserPaymentResponse{}, err
	}

	userPaymentResponse := dto.UserPaymentResponse{
		ID:             userPaymentCreate.ID.String(),
		UserID:         userPaymentCreate.UserID,
		Email:          getUser.Email,
		Invoice:        userPaymentCreate.Invoice,
		ProofOfPayment: userPaymentCreate.ProofOfPayment,
		Amount:         userPaymentCreate.Amount,
		Paid:           userPaymentCreate.Paid,
		ExpiredAt:      userPaymentCreate.ExpiredAt,
	}

	return userPaymentResponse, nil
}

func (s *userPaymentService) Update(userPaymentUpdate payload.UserPaymentPayload, id string, claims jwt.MapClaims) (dto.UserPaymentResponse, error) {
	userID := claims["user_id"].(string)
	role := claims["role"].(string)

	//Get UserPayment by ID
	getUserPayment, err := s.userPaymentRepo.FindByID(id)
	if err != nil {
		return dto.UserPaymentResponse{}, err
	}

	//Check User Id and Role User
	if role == "User" {
		if getUserPayment.UserID != userID {
			return dto.UserPaymentResponse{}, errors.New("unauthorized")
		}
	}

	var userPaymentModel models.UserPayment

	if userPaymentUpdate.Paid != nil {
		if *userPaymentUpdate.Paid == true {
			now := time.Now()
			userPaymentModel.InvoiceNumber = "INV-" + now.Format("20060102150405")
			userPaymentModel.Paid = true

			getUserPayment, err := s.userPaymentRepo.FindByIDWithAllRelationUserBook(id)
			if err != nil {
				return dto.UserPaymentResponse{}, err
			}

			url := helper.GenerateInvoice(getUserPayment, userPaymentModel.InvoiceNumber)
			userPaymentModel.Invoice = url
		}
		userPaymentModel.Paid = false
		log.Println("Paid ", userPaymentModel.Paid)
	}

	if userPaymentUpdate.ProofOfPayment != nil {
		fileName, buf, err := helper.OpenFileFromMultipartForm(userPaymentUpdate.ProofOfPayment)
		if err != nil {
			return dto.UserPaymentResponse{}, err
		}

		url := helper.UploadFileToFirebase(*buf, fileName)

		userPaymentModel.ProofOfPayment = url

	}

	userPayment, err := s.userPaymentRepo.Update(userPaymentModel, id)
	if err != nil {
		return dto.UserPaymentResponse{}, err
	}

	getUser, err := s.userRepo.FindByID(userPayment.UserID)
	if err != nil {
		return dto.UserPaymentResponse{}, err
	}

	userPaymentResponse := dto.UserPaymentResponse{
		ID:             userPayment.ID.String(),
		UserID:         userPayment.UserID,
		Email:          getUser.Email,
		InvoiceNumber:  userPayment.InvoiceNumber,
		Invoice:        userPayment.Invoice,
		ProofOfPayment: userPayment.ProofOfPayment,
		Amount:         userPayment.Amount,
		Paid:           userPayment.Paid,
		ExpiredAt:      userPayment.ExpiredAt,
	}

	return userPaymentResponse, nil
}

func (s *userPaymentService) Delete(id string) (models.UserPayment, error) {
	userPayment, err := s.userPaymentRepo.Delete(id)
	if err != nil {
		return userPayment, err
	}
	return userPayment, nil
}
