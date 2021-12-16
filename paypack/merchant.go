package paypack

import "context"

type Merchant struct {
	ID       string
	Name     string
	Active   bool
	RateIn   float64
	RateOut  float64
	Balance  float64
	Verified bool
}

type MerchantService interface {
	Me(context.Context) (*Merchant, error)
}
