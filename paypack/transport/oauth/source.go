package oauth

import (
	"context"

	"github.com/quarksgroup/paypack-go/paypack"
)

// StaticTokenSource returns a TokenSource that always
// returns the same token. Because the provided token t
// is never refreshed, StaticTokenSource is only useful
// for tokens that never expire.
func StaticTokenSource(t *paypack.Token) paypack.TokenSource {
	return staticTokenSource{t}
}

type staticTokenSource struct {
	token *paypack.Token
}

func (s staticTokenSource) Token(context.Context) (*paypack.Token, error) {
	return s.token, nil
}

// ContextTokenSource returns a TokenSource that returns
// a token from the http.Request context.
func ContextTokenSource() paypack.TokenSource {
	return contextTokenSource{}
}

type contextTokenSource struct {
}

func (s contextTokenSource) Token(ctx context.Context) (*paypack.Token, error) {
	token, _ := ctx.Value(paypack.TokenKey{}).(*paypack.Token)
	return token, nil
}
