package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CoachAvailibility struct {
	ID        uuid.UUID `gorm:"primaryKey" json:"id"` // Coach Avalibility ID
	CoachID   string    `json:"coach_id"`             // foreign key
	Day       string    `json:"day"`                  // Coach Avalibility Day
	StartTime string    `json:"start_time"`           // Coach Avalibility Start Time
	EndTime   string    `json:"end_time"`             // Coach Avalibility End Time
}

func (coachAvalibility *CoachAvailibility) BeforeCreate(tx *gorm.DB) (err error) {
	coachAvalibility.ID = uuid.New()
	return
}
