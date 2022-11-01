package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"firstname" validate:"min=3,max=50"`
	LastName  string    `json:"lastname" validate:"min=3,max=50"`
	Email     string    `json:"email" validate:"email"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	return
}
