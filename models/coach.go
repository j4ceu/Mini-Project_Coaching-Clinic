package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Coach struct {
	ID                uuid.UUID           `gorm:"primaryKey" json:"id"`
	Position          string              `json:"position"`                                                   // Coach Position
	Code              string              `json:"code" gorm:"uniqueIndex:coach_code_idx"`                     // Coach Code
	Price             int                 `json:"price"`                                                      // Coach Price
	UserID            string              `json:"user_id"`                                                    // foreign key
	User              User                `gorm:"foreignKey:UserID;references:ID" json:"user"`                // one to one relationship with user table
	GameID            uint                `json:"game_id"`                                                    // foreign key
	Game              Game                `gorm:"foreignKey:GameID;references:ID" json:"game"`                // one to one relationship with game table
	CoachExperience   []CoachExperience   `gorm:"foreignKey:CoachID;references:ID" json:"coach_experience"`   // one to many relationship with coach_experience table
	CoachAvailibility []CoachAvailibility `gorm:"foreignKey:CoachID;references:ID" json:"coach_availibility"` // one to many relationship with coach_avalibility table
}

func (coach *Coach) BeforeCreate(tx *gorm.DB) (err error) {
	coach.ID = uuid.New()
	return
}
