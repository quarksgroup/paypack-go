package oauth

import (
	"net/http"

	"github.com/quarksgroup/paypack-go/paypack"
	"github.com/quarksgroup/paypack-go/paypack/transport/internal"
)

// Supported authentication schemes.
const (
	SchemeBearer = "Bearer"
	SchemeToken  = "token"
)

// Transport is an http.RoundTripper that refreshes oauth
// tokens, wrapping a base RoundTripper and refreshing the
// token if expired.
type Transport struct {
	Scheme string
	Source paypack.TokenSource
	Base   http.RoundTripper
}

// RoundTrip authorizes and authenticates the request with
// an access token from the request context.
func (t *Transport) RoundTrip(r *http.Request) (*http.Response, error) {
	ctx := r.Context()
	token, err := t.Source.Token(ctx)
	if err != nil {
		return nil, err
	}
	if token == nil {
		return t.base().RoundTrip(r)
	}
	r2 := internal.CloneRequest(r)
	if r2.Header.Get("Authorization") == "" {
		r2.Header.Set("Authorization", t.scheme()+" "+token.Access)
	}
	return t.base().RoundTrip(r2)
}

// base returns the base transport. If no base transport
// is configured, the default transport is returned.
func (t *Transport) base() http.RoundTripper {
	if t.Base != nil {
		return t.Base
	}
	return http.DefaultTransport
}

// scheme returns the token scheme. If no scheme is
// configured, the bearer scheme is used.
func (t *Transport) scheme() string {
	if t.Scheme == "" {
		return SchemeBearer
	}
	return t.Scheme
}
