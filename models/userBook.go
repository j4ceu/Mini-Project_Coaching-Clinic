package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserBook struct {
	gorm.Model
	ID                  uuid.UUID         `gorm:"primaryKey" json:"ID"`
	Title               string            `json:"title"`
	CoachAvailabilityID string            `json:"coach_availability_id"` // foreign key
	CoachAvailability   CoachAvailability `gorm:"foreignKey:CoachAvailabilityID;references:ID" json:"coach_availability"`
	Summary             string            `json:"summary"`
	Done                bool              `json:"done"`
	UserPaymentID       string            `json:"user_payment_id"` // foreign key
}

func (userBook *UserBook) BeforeCreate(tx *gorm.DB) (err error) {
	userBook.ID = uuid.New()
	return
}
