package api

import (
	"context"
	"fmt"
	"time"

	"github.com/quarksgroup/paypack-go/paypack"
)

type authService struct {
	client *wrapper
}

const (
	loginEndpoint   = "auth/agents/authorize"
	refershEndpoint = "/auth/refresh"
	requestRetries  = 5
)

func (s *authService) Login(ctx context.Context, clientId, clietnSecret string) (*paypack.Token, error) {

	in := loginRequest{
		ClientId:     clientId,
		ClientSecret: clietnSecret,
	}

	out := new(tokenResponse)

	_, err := s.client.do(ctx, "POST", loginEndpoint, in, out)

	return convertToken(out), err
}

func (c *authService) Refresh(ctx context.Context, token *paypack.Token) (*paypack.Token, error) {

	// check token has expired or is about to expire soon
	if token.Expires.After(time.Now().UTC().Add(-time.Minute)) {
		return token, nil
	}

	out := new(tokenResponse)

	_, err := c.client.do(ctx, "POST", fmt.Sprintf("%s/%s", refershEndpoint, token.Refresh), nil, out)

	return convertToken(out), err
}

func convertToken(tk *tokenResponse) *paypack.Token {
	return &paypack.Token{
		Access:  tk.Access,
		Refresh: tk.Refresh,
		Expires: time.Unix(tk.Expires, 0),
	}
}
