package api

import (
	"context"

	"github.com/quarksgroup/paypack-go/paypack"
)

//Profile response to call payment /merchants/me endpoint
func (c *Client) Profile(ctx context.Context) (*paypack.Merchant, error) {

	const endpoint = "merchants/me"

	out := new(merchantResponse)

	_, err := c.do(ctx, "GET", endpoint, nil, out, nil)

	if err != nil {
		return nil, err
	}

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

	return res, nil
}

//Checkout that will query checkout information of a given agent id
func (c *Client) Checkout(ctx context.Context, agent string) (*paypack.Checkout, error) {

	endpoint := "checkouts/find/" + agent

	out := new(checkoutResponse)

	_, err := c.do(ctx, "GET", endpoint, nil, out, nil)

	if err != nil {
		return nil, err
	}

	res := &paypack.Checkout{
		ID:           out.ID,
		Name:         out.Name,
		Merchant:     out.Merchant,
		Logo:         out.Logo,
		Email:        out.Email,
		AppUrl:       out.AppUrl,
		SendEmail:    out.SendEmail,
		ClientId:     out.ClientId,
		ClientSecret: out.ClientSecret,
		CancelUrl:    out.CancelUrl,
		SuccessUrl:   out.SuccessUrl,
	}

	return res, nil
}
