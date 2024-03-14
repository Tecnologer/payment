package httputils

import "context"

const (
	TokenKey = "token"
)

// ContextWithToken adds a token to the context and returns the new context.
func ContextWithToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, TokenKey, token)
}
