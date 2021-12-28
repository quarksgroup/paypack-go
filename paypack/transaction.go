package paypack

import (
	"context"
	"time"
)

type Option string

// Transaction represent transaction information
type Transaction struct {
	Ref     string
	Kind    string
	Fee     float64
	Client  string
	Amount  float64
	Created time.Time
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
	Amount float64
	Number string
}

//TransactionResponse represent as single response data created after transaction was commited
type TransactionResponse struct {
	Ref       string     `json:"ref"`
	Status    string     `json:"status"`
	Amount    float64    `json:"amount"`
	Kind      string     `json:"kind"`
	CreatedAt *time.Time `json:"created_at"`
}

// TransactionService is the transaction engine responsible for transactions
// on the underying third party service of paypack api.
type TransactionService interface {

	// Cashin handles cashin http api request for https://payments.paypack.rw/api/transactions/cashin
	Cashin(context.Context, *TransactionRequest) (*TransactionResponse, error)
	// Cashout handles Cashout http api request for https://payments.paypack.rw/api/transactions/cashout
	Cashout(context.Context, *TransactionRequest) (*TransactionResponse, error)
	// Find handles Find http api request for https://payments.paypack.rw/api/transactions/find/{ref}
	Find(context.Context, string) (*Transaction, error)
	// List handles List http api request for https://payments.paypack.rw/api/transactions/list with paramas
	List(ctx context.Context, options ...Option) (*Transactions, error)
}
