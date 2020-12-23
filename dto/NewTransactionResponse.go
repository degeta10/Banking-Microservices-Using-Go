package dto

// NewTransactionResponse ..
type NewTransactionResponse struct {
	TransactionID string `json:"transaction_id"`
	AccountID     string `json:"account_id"`
	Balance       int64  `json:"balance"`
}
