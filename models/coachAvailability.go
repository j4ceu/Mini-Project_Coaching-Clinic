package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CoachAvailability struct {
	ID        uuid.UUID `gorm:"primaryKey" json:"id"`                     // Coach Avalibility ID
	CoachID   string    `json:"coach_id" gorm:"uniqueIndex : time_idx"`   // foreign key
	Day       string    `json:"day"`                                      // Coach Avalibility Day
	StartTime string    `json:"start_time" gorm:"uniqueIndex : time_idx"` // Coach Avalibility Start Time
	EndTime   string    `json:"end_time" gorm:"uniqueIndex : time_idx"`   // Coach Avalibility End Time
	Book      bool      `json:"book"`                                     // Coach Avalibility Book
}

func (coachAvalibility *CoachAvailability) BeforeCreate(tx *gorm.DB) (err error) {
	coachAvalibility.ID = uuid.New()
	return
}
