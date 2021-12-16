package paypack

import (
	"context"
	"time"
)

// Token represents the credentials used to authorize
// the requests to access protected resources.
type Token struct {
	Access  string
	Refresh string
	Expires time.Time
}

// TokenKey is the key to use with the context.WithValue
// function to associate an Token value with a context.
type TokenKey struct{}

// TokenSource returns a token.
type TokenSource interface {
	Token(context.Context) (*Token, error)
}

// AuthService handles authentication to the underlying API
type AuthService interface {
	// Login with cleint_id and client_secret to the underlying API and get an JWT token of paypack api
	Login(context.Context, string, string) (*Token, error)

	// Refresh the access token for revalidated the jwt token
	Refresh(ctx context.Context, token *Token) (*Token, error)
}

// WithContext returns a copy of parent in which the token value is set
func WithContext(parent context.Context, token *Token) context.Context {
	return context.WithValue(parent, TokenKey{}, token)
}

// TokenFrom returns the login token rom the context.
func TokenFrom(ctx context.Context) *Token {
	token, _ := ctx.Value(TokenKey{}).(*Token)
	return token
}
