package paypack

import "context"

type Merchant struct {
	ID      string  `json:"id",omitempty`
	Name    string  `json:"name,omitempty"`
	InRate  float64 `json:"in_rate,omitempty"`
	OutRate float64 `json:"out_rate,omitempty"`
	Balance float64 `json:"balance,omitempty"`
}

type MerchantService interface {
	Me(context.Context) (*Merchant, error)
}
