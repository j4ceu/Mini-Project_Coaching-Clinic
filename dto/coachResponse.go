package dto

import "Mini-Project_Coaching-Clinic/models"

type CoachResponse struct {
	ID                string                      `json:"id"`
	FirstName         string                      `json:"first_name"`
	LastName          string                      `json:"last_name"`
	Email             string                      `json:"email,omitempty"`
	Position          string                      `json:"position"`
	Code              string                      `json:"code"`
	Price             int                         `json:"price"`
	UserID            string                      `json:"user_id"`
	GameID            uint                        `json:"game_id"`
	CoachExperience   []models.CoachExperience    `json:"coach_experience"`
	CoachAvailability []CoachAvailabilityResponse `json:"coach_availibility"`
}
