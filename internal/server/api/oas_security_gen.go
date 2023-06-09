// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/http"
	"strings"
)

// SecurityHandler is handler for security parameters.
type SecurityHandler interface {
	// HandleHeaderAuth handles HeaderAuth security.
	HandleHeaderAuth(ctx context.Context, operationName string, t HeaderAuth) (context.Context, error)
	// HandleQueryAuth handles QueryAuth security.
	HandleQueryAuth(ctx context.Context, operationName string, t QueryAuth) (context.Context, error)
}

func findAuthorization(h http.Header, prefix string) (string, bool) {
	v, ok := h["Authorization"]
	if !ok {
		return "", false
	}
	for _, vv := range v {
		scheme, value, ok := strings.Cut(vv, " ")
		if !ok || !strings.EqualFold(scheme, prefix) {
			continue
		}
		return value, true
	}
	return "", false
}

func (s *Server) securityHeaderAuth(ctx context.Context, operationName string, req *http.Request) (context.Context, bool, error) {
	var t HeaderAuth
	const parameterName = "X-Api-Key"
	value := req.Header.Get(parameterName)
	if value == "" {
		return ctx, false, nil
	}
	t.APIKey = value
	rctx, err := s.sec.HandleHeaderAuth(ctx, operationName, t)
	if err != nil {
		return nil, false, err
	}
	return rctx, true, err
}
func (s *Server) securityQueryAuth(ctx context.Context, operationName string, req *http.Request) (context.Context, bool, error) {
	var t QueryAuth
	const parameterName = "token"
	q := req.URL.Query()
	if !q.Has(parameterName) {
		return ctx, false, nil
	}
	value := q.Get(parameterName)
	t.APIKey = value
	rctx, err := s.sec.HandleQueryAuth(ctx, operationName, t)
	if err != nil {
		return nil, false, err
	}
	return rctx, true, err
}
