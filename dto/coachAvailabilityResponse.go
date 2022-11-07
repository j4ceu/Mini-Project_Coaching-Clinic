package dto

type CoachAvailabilityResponse struct {
	ID        string `json:"id"`
	CoachID   string `json:"coach_id"`
	Day       string `json:"day"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}
