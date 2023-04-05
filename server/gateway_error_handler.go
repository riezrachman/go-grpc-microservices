package main

import (
	"context"
	"encoding/json"
	"net/http"

	pb "swing/server/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
)

func CustomHTTPError(ctx context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {

	const fallback = `{"error": true, "code": 500, "message": "Internal Error"}`

	w.Header().Set("Content-type", "application/json")
	grpcErrorCode := grpc.Code(err)
	httpErrorCode := runtime.HTTPStatusFromCode(grpcErrorCode)
	msg := status.Convert(err).Message()
	w.WriteHeader(httpErrorCode)

	logrus.Printf("[core][func: CustomHTTPError] Error gateway: %v", err)
	logrus.Printf("[core][func: CustomHTTPError] GRPC Error code: %v", grpcErrorCode)
	logrus.Printf("[core][func: CustomHTTPError] HTTP Error code: %v", httpErrorCode)
	logrus.Printf("[core][func: CustomHTTPError] Error Message: %s", msg)

	body := &pb.ErrorBodyResponse{
		Error:   true,
		Code:    uint32(httpErrorCode),
		Message: msg,
	}

	jErr := json.NewEncoder(w).Encode(body)

	if jErr != nil {
		w.Write([]byte(fallback))
	}

}
