package middleware

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	"usermate/internal/service"
)

func Jwt(service *service.CheckTokenService) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (any, error) {
			return handler(ctx, req)
		}
	}
}
