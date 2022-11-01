package models

type Game struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`        // Game Name
	Description string `json:"description"` // Game Description
}
