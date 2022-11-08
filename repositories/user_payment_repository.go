package repositories

import (
	"Mini-Project_Coaching-Clinic/models"

	"gorm.io/gorm"
)

type UserPaymentRepository interface {
	FindAll() ([]models.UserPayment, error)
	FindByInvoiceNumber(invoiceNumber string) (models.UserPayment, error)
	FindByID(id string) (models.UserPayment, error)
	FindByPaidAndProofOfPaymentIsNotNull() ([]models.UserPayment, error)
	FindUserBookByUserID(id string) ([]models.UserBook, error)
	Create(UserPayment models.UserPayment) (models.UserPayment, error)
	Update(UserPayment models.UserPayment, id string) (models.UserPayment, error)
	Delete(id string) (models.UserPayment, error)
	FindByIDWithAllRelationUserBook(id string) (models.UserPayment, error)
}

type userPaymentRepo struct {
	db *gorm.DB
}

func NewUserPaymentRepository(db *gorm.DB) *userPaymentRepo {
	return &userPaymentRepo{db}
}

func (r *userPaymentRepo) FindAll() ([]models.UserPayment, error) {
	var userPayments []models.UserPayment
	err := r.db.Find(&userPayments).Error
	if err != nil {
		return userPayments, err
	}
	return userPayments, nil
}

func (r *userPaymentRepo) FindUserBookByUserID(id string) ([]models.UserBook, error) {
	var userPayment models.UserPayment
	err := r.db.Where("user_id = ?", id).Preload("UserBook").Find(&userPayment).Error
	if err != nil {
		return []models.UserBook{}, err
	}
	return userPayment.UserBook, nil
}

func (r *userPaymentRepo) FindByPaidAndProofOfPaymentIsNotNull() ([]models.UserPayment, error) {
	var userPayments []models.UserPayment
	err := r.db.Where("paid = ? AND proof_of_payment IS NOT NULL", false).Preload("UserBook").Find(&userPayments).Error
	if err != nil {
		return userPayments, err
	}
	return userPayments, nil
}

func (r *userPaymentRepo) FindByID(id string) (models.UserPayment, error) {
	var userPayment models.UserPayment
	err := r.db.Where("id = ?", id).Preload("UserBook").Preload("UserBook.CoachAvailability").First(&userPayment).Error
	if err != nil {
		return userPayment, err
	}
	return userPayment, nil
}

func (r *userPaymentRepo) Create(userPayment models.UserPayment) (models.UserPayment, error) {
	err := r.db.Create(&userPayment).Error
	if err != nil {
		return userPayment, err
	}
	return userPayment, nil
}

func (r *userPaymentRepo) Update(userPaymentUpdate models.UserPayment, id string) (models.UserPayment, error) {
	var userPayment models.UserPayment
	err := r.db.Model(&userPayment).Where("id = ?", id).Updates(userPaymentUpdate).Error
	if err != nil {
		return userPayment, err
	}
	return r.FindByID(id)
}

func (r *userPaymentRepo) Delete(id string) (models.UserPayment, error) {
	var userPayment models.UserPayment
	err := r.db.Where("id = ?", id).Delete(&userPayment).Error
	if err != nil {
		return userPayment, err
	}
	return userPayment, nil
}

func (r *userPaymentRepo) FindByIDWithAllRelationUserBook(id string) (models.UserPayment, error) {
	var userPayment models.UserPayment
	err := r.db.Where("id = ?", id).Preload("User").Preload("UserBook").Preload("UserBook.CoachAvailability").Preload("UserBook.CoachAvailability.Coach").First(&userPayment).Error
	if err != nil {
		return userPayment, err
	}
	return userPayment, nil
}

func (r *userPaymentRepo) FindByInvoiceNumber(invoiceNumber string) (models.UserPayment, error) {
	var userPayment models.UserPayment
	err := r.db.Where("invoice_number = ?", invoiceNumber).Preload("UserBook").First(&userPayment).Error
	if err != nil {
		return userPayment, err
	}
	return userPayment, nil
}
