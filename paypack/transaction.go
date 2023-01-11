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
	Amount     float64  // Amount to be transacted
	Number     string   // Phone number that will refunded or withdrawn money from
	Mode       string   // development or production mode which will be used to choose where your callback will be sent
	MetaData   string   // Optional encoded json string
	WebhookIds []string // Optional registered webhook_ids array where callback will be sent
}

//TransactionResponse represent as single response data created after transaction was commited
type TransactionResponse struct {
	Ref       string     `json:"ref"`        // Unique reference of transaction
	Status    string     `json:"status"`     // Status of transaction
	Amount    float64    `json:"amount"`     // Amount of transaction
	Provider  string     `json:"provider"`   // Provider of transaction mtn, airtel
	UserRef   string     `json:"user_ref"`   // User reference of transaction this is unique dynamic provided by user or provided by system if not provided by user
	Kind      string     `json:"kind"`       // Kind of transaction cashin or cashout
	CreatedAt *time.Time `json:"created_at"` // Time when transaction was created
}
