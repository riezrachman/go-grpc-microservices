package api

import (
	"context"
	pb "swing/server/pb"
)

func (s *Server) StoreList(ctx context.Context, req *pb.StoreListRequest) (*pb.StoreListResponse, error) {

	result := &pb.StoreListResponse{
		Error:   false,
		Code:    200,
		Message: "Success",
		Data:    []*pb.Store{},
	}

	return result, nil

}
