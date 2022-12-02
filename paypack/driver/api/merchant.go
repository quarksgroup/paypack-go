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

	_, err := s.http.do(ctx, "GET", endpoint, nil, out, nil)

	res := &paypack.Merchant{
		ID:            out.ID,
		Name:          out.Name,
		InRate:        out.In,
		OutRate:       out.Out,
		AirtelInRate:  out.AirtelInRate,
		AirtelOutRate: out.AirtelOutRate,
		Balance:       out.Balance,
		AirtelBalance: out.AirtelBalance,
		MtnBalance:    out.MtnBalance,
	}

	return res, err
}

var _ (paypack.MerchantService) = (*merchantService)(nil)
