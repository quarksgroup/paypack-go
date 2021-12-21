package paypack

import "context"

type Merchant struct {
	Name    string  `json:"name"`
	RateIn  float64 `json:"rate_in"`
	RateOut float64 `json:"rate_out"`
	Balance float64 `json:"balance"`
}

type MerchantService interface {
	Me(context.Context) (*Merchant, error)
}
