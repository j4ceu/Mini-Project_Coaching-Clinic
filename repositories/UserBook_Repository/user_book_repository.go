package UserBook_Repository

import (
	"Mini-Project_Coaching-Clinic/models"

	"gorm.io/gorm"
)

type UserBookRepository interface {
	FindAll() ([]models.UserBook, error)
	FindByID(id string) (models.UserBook, error)
	Create(UserBook models.UserBook) (models.UserBook, error)
	Update(UserBook models.UserBook, id string) (models.UserBook, error)
	Delete(id string) (models.UserBook, error)
}

type userBookRepo struct {
	db *gorm.DB
}

func NewUserBookRepository(db *gorm.DB) *userBookRepo {
	return &userBookRepo{db}
}

func (r *userBookRepo) FindAll() ([]models.UserBook, error) {
	var userBooks []models.UserBook
	err := r.db.Find(&userBooks).Error
	if err != nil {
		return userBooks, err
	}
	return userBooks, nil
}

func (r *userBookRepo) FindByID(id string) (models.UserBook, error) {
	var userBook models.UserBook
	err := r.db.Where("id = ?", id).Preload("CoachAvailability").First(&userBook).Error
	if err != nil {
		return userBook, err
	}
	return userBook, nil
}

func (r *userBookRepo) Create(userBook models.UserBook) (models.UserBook, error) {
	err := r.db.Create(&userBook).Error
	if err != nil {
		return userBook, err
	}
	return userBook, nil
}

func (r *userBookRepo) Update(userBookUpdate models.UserBook, id string) (models.UserBook, error) {
	var userBook models.UserBook
	err := r.db.Model(&userBook).Where("id = ?", id).Updates(&userBookUpdate).Error
	if err != nil {
		return userBook, err
	}
	return r.FindByID(id)
}

func (r *userBookRepo) Delete(id string) (models.UserBook, error) {
	var userBook models.UserBook
	err := r.db.Where("id = ?", id).Delete(&userBook).Error
	if err != nil {
		return userBook, err
	}
	return userBook, nil
}
