package request

import (
	"context"

	"github.com/weekndCN/rw-app/core"
)

type key int

const authKey key = 1

// WithAuth return a copy of parent context
func WithAuth(parent context.Context, auth *core.Auth) context.Context {
	return context.WithValue(parent, authKey, auth)
}

// AuthFrom return a auth fron context
func AuthFrom(ctx context.Context) (*core.Auth, bool) {
	auth, ok := ctx.Value(authKey).(*core.Auth)
	return auth, ok
}
