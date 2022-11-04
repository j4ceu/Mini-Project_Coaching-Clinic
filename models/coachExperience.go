package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CoachExperience struct {
	ID          uuid.UUID `gorm:"primaryKey" json:"id"` // Coach Experience ID
	CoachID     string    `json:"coach_id"`             // foreign key
	Title       string    `json:"title"`                // Coach Experience Title
	Description string    `json:"description"`          // Coach Experience Description
}

func (coachExperience *CoachExperience) BeforeCreate(tx *gorm.DB) (err error) {
	coachExperience.ID = uuid.New()
	return
}
