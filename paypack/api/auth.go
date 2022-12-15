package api

import (
	"context"
	"fmt"
	"time"

	"github.com/quarksgroup/paypack-go/paypack"
)

const (
	loginEndpoint   = "auth/agents/authorize"
	refershEndpoint = "auth/refresh"
)

func (c *Client) Login(ctx context.Context, clientId, clietnSecret string) (*paypack.Token, error) {

	in := loginRequest{
		ClientId:     clientId,
		ClientSecret: clietnSecret,
	}

	out := new(tokenResponse)

	_, err := c.do(ctx, "POST", loginEndpoint, in, out, nil)

	if err != nil {
		return nil, err
	}

	return convertToken(out), nil
}

func (c *Client) Refresh(ctx context.Context, token *paypack.Token) (*paypack.Token, error) {

	// check token has expired or is about to expire soon
	if time.Unix(token.Expires, 0).After(time.Now().UTC().Add(-time.Minute)) {
		return token, nil
	}

	out := new(tokenResponse)

	_, err := c.do(ctx, "GET", fmt.Sprintf("%s/%s", refershEndpoint, token.Refresh), nil, out, nil)

	if err != nil {
		return nil, err
	}

	return convertToken(out), nil
}

func convertToken(tk *tokenResponse) *paypack.Token {
	return &paypack.Token{
		Access:  tk.Access,
		Refresh: tk.Refresh,
		Expires: tk.Expires,
	}
}
