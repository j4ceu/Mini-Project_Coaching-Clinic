package payload

import "mime/multipart"

type UserBookPayloadCreate struct {
	Title               string `json:"title"`
	CoachAvailabilityID string `json:"coach_availability_id"`
	UserPaymentID       string `json:"user_payment_id"`
}

type UserBookPayloadUpdate struct {
	Title               string                `form:"title"`
	CoachAvailabilityID string                `form:"coach_availability_id"`
	UserPaymentID       string                `form:"user_payment_id"`
	Summary             *multipart.FileHeader `form:"summary"`
	Done                *bool                 `form:"done"`
}
