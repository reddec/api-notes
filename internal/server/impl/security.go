package impl

import (
	"context"
	"crypto/subtle"
	"errors"

	"github.com/reddec/api-notes/internal/server/api"
)

var ErrInvalidToken = errors.New("invalid token")

type StaticToken struct {
	Token string
}

func (st *StaticToken) HandleHeaderAuth(ctx context.Context, _ string, t api.HeaderAuth) (context.Context, error) {
	return st.validateToken(ctx, t.APIKey)
}

func (st *StaticToken) HandleQueryAuth(ctx context.Context, _ string, t api.QueryAuth) (context.Context, error) {
	return st.validateToken(ctx, t.APIKey)
}

func (st *StaticToken) validateToken(ctx context.Context, token string) (context.Context, error) {
	if subtle.ConstantTimeCompare([]byte(st.Token), []byte(token)) == 0 {
		return nil, ErrInvalidToken
	}
	return ctx, nil
}

type AnyToken struct{}

func (ns *AnyToken) HandleHeaderAuth(ctx context.Context, _ string, _ api.HeaderAuth) (context.Context, error) {
	return ctx, nil
}

func (ns *AnyToken) HandleQueryAuth(ctx context.Context, _ string, _ api.QueryAuth) (context.Context, error) {
	return ctx, nil
}
