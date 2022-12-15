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

//Transaction represent transaction data details information from payments-paypack
type Transaction struct {
	Ref       string    `json:"ref"`
	Status    string    `json:"status"`
	Amount    float64   `json:"amount"`
	Fee       float64   `json:"fee"`
	Kind      string    `json:"kind"`
	Provider  string    `json:"provider"`
	Client    string    `json:"client"`
	Merchant  string    `json:"merchant"`
	Timestamp time.Time `json:"timestamp"`
}

//listTransactions represent transactions list data details information from payments-paypack
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

//listEventResponse represent events data details information from payments-paypack
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

//merchantResponse represent merchant details information from payments-paypack
type merchantResponse struct {
	ID            string  `json:"id,omitempty"`
	Name          string  `json:"name,omitempty"`
	In            float64 `json:"in_rate,omitempty"`
	Out           float64 `json:"out_rate,omitempty"`
	AirtelInRate  float64 `json:"airtel_in_rate,omitempty"`
	AirtelOutRate float64 `json:"airtel_out_rate,omitempty"`
	Balance       float64 `json:"balance,omitempty"`
	AirtelBalance float64 `json:"airtel_balance,omitempty"`
	MtnBalance    float64 `json:"mtn_balance,omitempty"`
}

//checkoutResponse represent checkout details information from payments-paypack
type checkoutResponse struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Merchant     string `json:"merchant"`
	Logo         string `json:"logo"`
	SendEmail    bool   `json:"send_email"`      // This will represent if merchant need to send email when payments successed
	Email        string `json:"email,omitempty"` //This will be support email that will be sent to customer email for reply
	CancelUrl    string `json:"cancel_url"`
	SuccessUrl   string `json:"success_url"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}
