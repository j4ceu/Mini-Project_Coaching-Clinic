package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserPayment struct {
	ID             uuid.UUID  `gorm:"primaryKey" json:"id"`
	UserID         string     `json:"user_id"` // foreign key
	User           User       `gorm:"foreignKey:UserID;references:ID" json:"user"`
	UserBook       []UserBook `gorm:"foreignKey:UserPaymentID;references:ID;" json:"user_book"`
	Invoice        string     `json:"invoice" gorm:"default:null"`
	InvoiceNumber  string     `json:"invoice_number" gorm:"unique; default:null"`
	ProofOfPayment string     `json:"proof_of_payment" gorm:"default:null"`
	Amount         int        `json:"amount"`
	Paid           bool       `json:"paid" gorm:"default:false"`
	ExpiredAt      string     `json:"expired_at"`
}

func (userPayment *UserPayment) BeforeCreate(tx *gorm.DB) (err error) {
	userPayment.ID = uuid.New()
	return
}
