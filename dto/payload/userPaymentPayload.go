package payload

type UserPaymentPayload struct {
	InvoiceNumber  *string `json:"invoice_number"`
	Invoice        *string `json:"invoice"`
	ProofOfPayment *string `json:"proof_of_payment"`
	Amount         *int    `json:"amount"`
	Paid           *bool   `json:"paid"`
	ExpiredAt      *string `json:"expired_at"`
	UserID         string  `json:"user_id"`
}
