package dto

type UserPaymentResponse struct {
	ID             string              `json:"id"`
	UserID         string              `json:"user_id"`
	Email          string              `json:"email,omitempty"`
	Invoice        string              `json:"invoice"`
	ProofOfPayment string              `json:"proof_of_payment"`
	Amount         int                 `json:"amount"`
	Paid           bool                `json:"paid"`
	ExpiredAt      string              `json:"expired_at"`
	InvoiceNumber  string              `json:"invoice_number"`
	UserBook       *[]UserBookResponse `json:"user_book,omitempty"`
}
