package manager

import (
	"context"
	"fmt"
	"strings"
	"swing/server/pb"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type JWTManager struct {
	secretKey     string
	tokenDuration time.Duration
}

type UserClaims struct {
	User *pb.User

	jwt.StandardClaims
}

func NewJWTManager(secretKey string, tokenDuration time.Duration) *JWTManager {

	return &JWTManager{secretKey, tokenDuration}

}

func (manager *JWTManager) Generate(user *pb.User) (string, error) {

	claims := UserClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(manager.tokenDuration).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(manager.secretKey))

}

func (manager *JWTManager) Verify(accessToken string) (*UserClaims, error) {

	token, err := jwt.ParseWithClaims(
		accessToken,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}

			return []byte(manager.secretKey), nil
		},
	)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil

}

func (manager *JWTManager) GetMeFromJWT(ctx context.Context) (*UserClaims, error) {

	accessToken := ""

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		values := md["authorization"]
		if len(values) > 0 {
			split := strings.Split(values[0], " ")
			accessToken = split[0]
			if len(split) > 1 {
				accessToken = split[1]
			}
		}
	}

	if accessToken == "" {
		logrus.Errorf("[manager][func: GetMeFromJWT] Access Token is Empty")
		return nil, status.Error(codes.Unauthenticated, "Access Token is required")
	}

	userClaims, err := manager.Verify(accessToken)
	if err != nil {
		logrus.Errorf("[manager][func: GetMeFromJWT] Failed to Verifiy Token '%s', Error: %v", accessToken, err)
		return nil, status.Errorf(codes.Unauthenticated, "Session Expired")
	}

	now := time.Now()

	// fmt.Printf("[manager][func: GetMeFromJWT] Token Verify Expired: %v|%v|%v", !(now.Unix() <= userClaims.ExpiresAt), now.String(), time.Unix(userClaims.ExpiresAt, 0).String())
	if !(now.Unix() <= userClaims.ExpiresAt) {
		return nil, status.Errorf(codes.Unauthenticated, "Session Expired")
	}

	return userClaims, nil

}
