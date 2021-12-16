package api

import (
	"context"

	"github.com/quarksgroup/paypack-go/paypack"
)

type merchantService struct {
	http *wrapper
}

func (s *merchantService) Me(ctx context.Context) (*paypack.Merchant, error) {

	endpoint := "merchants/me"

	out := new(paypack.Merchant)

	_, err := s.http.do(ctx, "GET", endpoint, nil, out)
	return out, err
}

var _ (paypack.MerchantService) = (*merchantService)(nil)
