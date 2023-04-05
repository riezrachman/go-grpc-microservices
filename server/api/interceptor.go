package api

import (
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"context"
	"time"

	"swing/server/db"
)

const AuthTokenKey = "authentication"

var authFilteredModule = []string{
	apiServicePath + "GetUser",
}

var (
	errGrpcUnauthenticated = status.Errorf(codes.Unauthenticated, "Token is missing")
)

type UserCtxKey struct{}

func UnaryInterceptors(authInterceptor *AuthInterceptor) grpc.UnaryServerInterceptor {

	return grpc_middleware.ChainUnaryServer(
		LoggingInterceptor,
		ErrorsInterceptor,
		authInterceptor.Unary(),
	)

}

func StreamInterceptors(authInterceptor *AuthInterceptor) grpc.StreamServerInterceptor {

	return grpc_middleware.ChainStreamServer(
		authInterceptor.Stream(),
	)

}

func ErrorsInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (out interface{}, err error) {

	out, err = handler(ctx, req)

	switch err := err.(type) {
	case db.NotFoundError:
		return out, status.Errorf(codes.NotFound, err.Error())
	}

	return out, err

}

func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (out interface{}, err error) {

	if info.FullMethod == "/grpc.health.v1.Health/Check" {

		out, err = handler(ctx, req)
		return out, err

	} else {

		entry := logrus.WithField("method", info.FullMethod)
		start := time.Now()
		out, err = handler(ctx, req)
		duration := time.Since(start)

		if err != nil {
			entry = entry.WithError(err)
		}

		entry.WithField("duration", duration.String()).Info("finished RPC")
		return out, err

	}

}

func AuthenticationInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (out interface{}, err error) {

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return handler(ctx, req)
	}

	if restricted := contains(authFilteredModule, info.FullMethod); !restricted {
		return handler(ctx, req)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errGrpcUnauthenticated
	}

	tokenString, ok := md[AuthTokenKey]
	if !ok || len(tokenString) < 1 {
		return nil, errGrpcUnauthenticated
	}

	token, err := jwt.Parse(tokenString[0], func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		ctx = context.WithValue(ctx, UserCtxKey{}, claims["un"])

	} else {

		return nil, errGrpcUnauthenticated

	}

	return handler(ctx, req)

}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func int32Contains(s []int32, e int32) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func containsSubs(s []string, e string) bool {
	for _, a := range s {
		c := strings.Contains(e, a)
		if c {
			return true
		}
	}
	return false
}
