package api

import (
	"context"
	"net/http"

	"github.com/quarksgroup/paypack-go/paypack"
)

//Profile response to call payment /merchants/me endpoint
func (c *Client) Profile(ctx context.Context) (*paypack.Merchant, error) {

	const endpoint = "merchants/me"

	out := new(merchantResponse)

	_, err := c.do(ctx, "GET", endpoint, nil, out, nil)

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

//FindCheckout that will query checkout information of a given agent id
func (c *Client) FindCheckout(ctx context.Context, agent string) (*paypack.Checkout, error) {

	endpoint := "checkout/find/" + agent

	out := new(checkoutResponse)

	header := http.Header{
		"X-security": []string{"secret"}, // help for more security
	}
	_, err := c.do(ctx, "GET", endpoint, nil, out, header)

	res := &paypack.Checkout{
		ID:           out.ID,
		Name:         out.Name,
		Merchant:     out.Merchant,
		Logo:         out.Logo,
		SendEmail:    out.SendEmail,
		CancelUrl:    out.CancelUrl,
		SuccessUrl:   out.SuccessUrl,
		ClientId:     out.ClientId,
		ClientSecret: out.ClientSecret,
	}

	return res, err
}
