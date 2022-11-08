package services

import (
	"Mini-Project_Coaching-Clinic/dto"
	"Mini-Project_Coaching-Clinic/middlewares"
	"errors"

	"Mini-Project_Coaching-Clinic/models"
	"Mini-Project_Coaching-Clinic/repositories"
)

type UserService interface {
	FindByID(id string) (models.User, error)
	Create(user models.User) (models.User, error)
	FindAll() ([]models.User, error)
	Update(user models.User, id string) (models.User, error)
	Delete(id string) (models.User, error)
	LoginUser(email string, password string) (dto.UserResponse, error)
}

type userService struct {
	repository repositories.UserRepositories
}

func NewUserService(repository repositories.UserRepositories) *userService {
	return &userService{repository}
}

func (s *userService) FindByID(id string) (models.User, error) {
	user, err := s.repository.FindByID(id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *userService) Create(user models.User) (models.User, error) {
	user.Role = "User" // set role default to User

	if err := user.HashPassword(user.Password); err != nil {
		return models.User{}, err
	}

	user, err := s.repository.Create(user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *userService) FindAll() ([]models.User, error) {
	users, err := s.repository.FindAll()
	if err != nil {
		return users, err
	}
	return users, nil
}

func (s *userService) Update(user models.User, id string) (models.User, error) {
	if user.Password != "" {
		if err := user.HashPassword(user.Password); err != nil {
			return models.User{}, err
		}
	}
	user, err := s.repository.Update(user, id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *userService) Delete(id string) (models.User, error) {
	user, err := s.repository.Delete(id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *userService) LoginUser(email string, password string) (dto.UserResponse, error) {
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return dto.UserResponse{}, err
	}

	credentialError := user.CheckPassword(password)
	if credentialError != nil {
		return dto.UserResponse{}, errors.New("Invalid email or password")
	}

	token, err := middlewares.CreateToken(user.ID.String(), user.Role)
	if err != nil {
		return dto.UserResponse{}, err
	}

	userResponse := dto.UserResponse{
		Email: user.Email,
		Token: token,
	}

	return userResponse, nil
}
