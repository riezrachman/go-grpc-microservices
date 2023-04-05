package api

import (
	"context"
	"crypto/tls"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"swing/server/db"
	manager "swing/server/lib/jwt"
	pb "swing/server/pb"

	"github.com/sirupsen/logrus"
)

const apiServicePath string = "/swing.v1.ApiService/"

type Server struct {
	provider *db.GormProvider
	manager  *manager.JWTManager

	pb.ApiServiceServer
}

func New(jwt_secret string, jwt_duration string, db01 *gorm.DB) *Server {

	secret := jwt_secret

	tokenDuration, err := time.ParseDuration(jwt_duration)
	if err != nil {
		logrus.Panic("[api][func: New] Failed when parsing duration: %v", err)
	}

	return &Server{
		provider: db.NewProvider(db01, tokenDuration),
		manager:  manager.NewJWTManager(secret, tokenDuration),
	}

}

func (s *Server) GetManager() *manager.JWTManager {

	return s.manager

}

func (s *Server) NotImplementedError() error {

	return status.New(codes.Unimplemented, "Not implemented.").Err()

}

func (s *Server) HealthCheck(ctx context.Context, _ *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {

	result := &pb.HealthCheckResponse{
		Error:   false,
		Code:    200,
		Message: "API Running!",
	}

	return result, nil

}

func (s *Server) GetHttpClient(ctx context.Context) (*http.Client, error) {

	var transport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 3 * time.Minute,
		}).Dial,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		TLSHandshakeTimeout: 3 * time.Minute,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   3 * time.Minute,
	}

	return client, nil

}

func HTTPResponseParse(httpRes *http.Response, functionName string) (string, error) {

	logrus.Printf("[api][func: %s] RESPONSE STATUS: %s", functionName, httpRes.Status)

	httpResBodyBuf := new(strings.Builder)
	_, err := io.Copy(httpResBodyBuf, httpRes.Body)
	if err != nil {
		logrus.Printf("[api][func: %s] Error: %v", functionName, err)
		return "", err
	}

	responseBody := httpResBodyBuf.String()

	logrus.Printf("[api][func: %s] RESPONSE BODY: %s", functionName, responseBody)

	return responseBody, nil

}

func setPagination(limit int32, page int32) *pb.PaginationResponse {

	res := &pb.PaginationResponse{
		Limit: 10,
		Page:  1,
	}

	if limit == 0 && page == 0 {
		res.Limit = -1
		res.Page = -1
		return res
	} else {
		res.Limit = limit
		res.Page = page
	}

	if res.Page == 0 {
		res.Page = 1
	}

	switch {
	case res.Limit > 100:
		res.Limit = 100
	case res.Limit <= 0:
		res.Limit = 10
	}

	return res

}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
