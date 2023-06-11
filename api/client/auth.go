package client

import (
	"context"

	"github.com/ogen-go/ogen/ogenerrors"
)

// HeaderToken is header base authorization with pre-shared token.
type HeaderToken string

func (ta HeaderToken) HeaderAuth(_ context.Context, _ string) (HeaderAuth, error) {
	return HeaderAuth{
		APIKey: string(ta),
	}, nil
}

// QueryAuth ignored and always returns [ogenerrors.ErrSkipClientSecurity].
func (ta HeaderToken) QueryAuth(_ context.Context, _ string) (QueryAuth, error) {
	return QueryAuth{}, ogenerrors.ErrSkipClientSecurity
}
