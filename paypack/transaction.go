package paypack

import (
	"time"
)

type Option string

// Transaction represent transaction information
type Transaction struct {
	Ref       string     `json:"ref,omitempty"`
	Status    string     `json:"status,omitempty"`
	Amount    float64    `json:"amount,omitempty"`
	Client    string     `json:"client,omitempty"`
	Kind      string     `json:"kind,omitempty"`
	Fee       float64    `json:"fee,omitempty"`
	Provider  string     `json:"provider,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	Processed *time.Time `json:"processed_at,omitempty"`
	Commited  *time.Time `json:"commited_at,omitempty"`
}

// Transactions this reperesent informations of more than tx by using list
type Transactions struct {
	Offset       uint64
	Limit        uint64
	From         *string
	To           *string
	Kind         *string
	Fee          float64
	Cashin       float64
	Cashout      float64
	Total        uint64
	Transactions []Transaction
}

// TransactionRequest represents as single payload required for making transaction
type TransactionRequest struct {
	Amount     float64
	Number     string
	Mode       string
	WebhookIds []string
}

//TransactionResponse represent as single response data created after transaction was commited
type TransactionResponse struct {
	Ref       string     `json:"ref"`
	Status    string     `json:"status"`
	Amount    float64    `json:"amount"`
	Provider  string     `json:"provider"`
	UserRef   string     `json:"user_ref"`
	Kind      string     `json:"kind"`
	CreatedAt *time.Time `json:"created_at"`
}
