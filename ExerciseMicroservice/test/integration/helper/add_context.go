package helper

import (
	"ExerciseMicroservice/internal/auth"
	"context"
	"github.com/99designs/gqlgen/client"
)

func AddContext(token string) client.Option {
	return func(req *client.Request) {
		ctx := context.WithValue(req.HTTP.Context(), auth.TokenCtxKey, token)
		req.HTTP = req.HTTP.WithContext(ctx)
	}
}
