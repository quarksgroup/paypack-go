package api

import (
	"encoding/json"
	"time"
)

type loginRequest struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type tokenResponse struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
	Expires int64  `json:"expires"`
}
type transactionRequest struct {
	Amount float64 `json:"amount"`
	Number string  `json:"number"`
}

type Transaction struct {
	Ref       string    `json:"ref"`
	Status    string    `json:"status"`
	Amount    float64   `json:"amount"`
	Fee       float64   `json:"fee"`
	Kind      string    `json:"kind"`
	Client    string    `json:"client"`
	Merchant  string    `json:"merchant"`
	Timestamp time.Time `json:"timestamp"`
}
type listTransactions struct {
	Offset       uint64        `json:"offset"`
	Limit        uint64        `json:"limit"`
	From         *string       `json:"from,omitempty"`
	To           *string       `json:"to,omitempty"`
	Kind         *string       `json:"kind,omitempty"`
	Cashin       float64       `json:"cashin,omitempty"`
	Cashout      float64       `json:"cashout,omitempty"`
	Fee          float64       `json:"fee,omitempty"`
	Total        uint64        `json:"total,omitempty"`
	Transactions []Transaction `json:"transactions"`
}

type EventResponse struct {
	ID        string          `json:"event_id"`
	Kind      string          `json:"event_kind"`
	Data      json.RawMessage `json:"data"`
	CreatedAt string          `json:"created_at"`
}
type listEventResponse struct {
	Ref          *string         `json:"ref,omitempty"`
	Status       *string         `json:"status,omitempty"`
	Kind         *string         `json:"kind,omitempty"`
	Offset       uint64          `json:"offset,omitempty"`
	Limit        uint64          `json:"limit,omitempty"`
	EventKind    *string         `json:"event-kind,omitempty"`
	Total        uint64          `json:"total,omitempty"`
	Transactions []EventResponse `json:"transactions"`
}

type merchantResponse struct {
	ID      string  `json:"id,omitempty"`
	Name    string  `json:"name,omitempty"`
	In      float64 `json:"in_rate,omitempty"`
	Out     float64 `json:"out_rate,omitempty"`
	Balance float64 `json:"balance,omitempty"`
}
