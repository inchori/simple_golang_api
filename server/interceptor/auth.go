package interceptor

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
	"log"
)

type JWTInterceptor struct {
	interceptedMethods map[string]bool
}

func NewJWTInterceptor(method []string) *JWTInterceptor {
	interceptor := JWTInterceptor{}
	interceptor.InterceptMethods(method)
	return &interceptor
}

func (ic *JWTInterceptor) InterceptMethods(methods []string) {
	if ic.interceptedMethods == nil {
		ic.interceptedMethods = make(map[string]bool)
	}
	for i := range methods {
		ic.interceptedMethods[methods[i]] = true
	}
}

func (ic *JWTInterceptor) checkMethod(method string) bool {
	return ic.interceptedMethods[method]
}

func (ic *JWTInterceptor) Interceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	isInterceptedMethod := ic.checkMethod(info.FullMethod)
	if !isInterceptedMethod {
		return handler(ctx, req)
	}
	jwtTokenStr, err := grpc_auth.AuthFromMD(ctx, "Bearer")
	if err != nil {
		log.Fatalf("error jwt from header: %v", err)
		return ctx, nil
	}

	token, err := jwt.Parse(jwtTokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		log.Fatalf("error jwt parsing: %v", err)
		return nil, err
	}

	sub, err := token.Claims.GetSubject()
	if err != nil {
		log.Fatalf("invalid sub: %v", err)
		return nil, err
	}

	if !token.Valid {
		log.Fatalf("invalid jwt: %v", err)
		return nil, err
	}

	newCtx := context.WithValue(ctx, ContextKeyAuthenticated{}, sub)
	return newCtx, nil
}
