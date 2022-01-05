package api

import (
	"context"

	"github.com/quarksgroup/paypack-go/paypack"
)

type merchantService struct {
	http *wrapper
}

func (s *merchantService) Me(ctx context.Context) (*paypack.Merchant, error) {

	const endpoint = "merchants/me"

	out := new(merchantResponse)

	_, err := s.http.do(ctx, "GET", endpoint, nil, out)

	res := &paypack.Merchant{
		Name:    out.Name,
		InRate:  out.In,
		OutRate: out.Out,
		Balance: out.Balance,
	}

	return res, err
}

var _ (paypack.MerchantService) = (*merchantService)(nil)
