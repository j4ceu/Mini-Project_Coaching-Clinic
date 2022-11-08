package repositories

import (
	"Mini-Project_Coaching-Clinic/models"

	"gorm.io/gorm"
)

type UserRepositories interface {
	FindByID(id string) (models.User, error)
	Create(user models.User) (models.User, error)
	FindAll() ([]models.User, error)
	Update(user models.User, id string) (models.User, error)
	Delete(id string) (models.User, error)
	FindByEmail(email string) (models.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepositories(db *gorm.DB) *userRepo {
	return &userRepo{db}
}

func (r *userRepo) FindByID(id string) (models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepo) Create(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepo) FindAll() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}

func (r *userRepo) Update(userUpdate models.User, id string) (models.User, error) {
	var user models.User
	err := r.db.Model(&user).Where("id = ?", id).Updates(&userUpdate).Error
	if err != nil {
		return user, err
	}
	return r.FindByID(id)
}

func (r *userRepo) Delete(id string) (models.User, error) {
	var user models.User
	err := r.db.Delete(&user).Where("id = ?", id).Error
	if err != nil {
		return user, err
	}
	return r.FindByID(id)
}

func (r *userRepo) FindByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
