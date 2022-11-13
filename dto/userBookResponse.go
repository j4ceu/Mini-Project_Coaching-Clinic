package dto

type UserBookResponse struct {
	ID                  string                     `json:"id"`
	Title               string                     `json:"title"`
	CoachAvailabilityID string                     `json:"coach_availability_id"`
	Summary             string                     `json:"summary,omitempty"`
	Done                bool                       `json:"done"`
	UserPaymentID       string                     `json:"user_payment_id,omitempty"`
	CoachAvailability   *CoachAvailabilityResponse `json:"coach_availability,omitempty"`
}
