package api

import (
	"context"
	"fmt"
	"strings"
	"swing/server/db"
	pb "swing/server/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) StoreList(ctx context.Context, req *pb.StoreListRequest) (*pb.StoreListResponse, error) {

	result := &pb.StoreListResponse{
		Error:   false,
		Code:    200,
		Message: "Success",
		Data:    []*pb.Store{},
	}

	result.Pagination = setPagination(req.GetLimit(), req.GetPage())

	storeORMs, err := s.provider.GetStore(ctx, &db.ListFilter{
		Filter: req.GetFilter(),
		Query:  req.GetQuery(),
		Sort: &pb.Sort{
			Column:    req.GetSort(),
			Direction: req.GetDir().String(),
		},
	}, result.GetPagination())
	if err != nil {
		return nil, err
	}

	for _, storeORM := range storeORMs {

		storePB, err := storeORM.ToPB(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}

		result.Data = append(result.Data, &storePB)

	}

	return result, nil

}

func (s *Server) StoreDetail(ctx context.Context, req *pb.StoreDetailRequest) (*pb.StoreDetailResponse, error) {

	result := &pb.StoreDetailResponse{
		Error:   false,
		Code:    200,
		Message: "Success",
	}

	if req.GetId() < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "Id is required")
	}

	storeORM, err := s.provider.GetStoreDetail(ctx, &pb.StoreORM{
		Id: req.GetId(),
	})
	if err != nil {
		return nil, err
	}

	if storeORM == nil {
		return nil, status.Errorf(codes.NotFound, "Data not found")
	}

	storePB, err := storeORM.ToPB(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}

	result.Data = &storePB

	return result, nil

}

func (s *Server) StoreCreate(ctx context.Context, req *pb.StoreCreateRequest) (*pb.StoreCreateResponse, error) {

	result := &pb.StoreCreateResponse{
		Error:   false,
		Code:    200,
		Message: "Success",
	}

	if req.GetData() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Data is required")
	}

	if req.GetData().GetOperationalTimeStart() < 0 || req.GetData().GetOperationalTimeStart() > 24 {
		return nil, status.Errorf(codes.InvalidArgument, "Operational Time Start should be greater than 0 and lesser than 24")
	}

	if req.GetData().GetOperationalTimeEnd() < 0 || req.GetData().GetOperationalTimeEnd() > 24 {
		return nil, status.Errorf(codes.InvalidArgument, "Operational Time End should be greater than 0 and lesser than 24")
	}

	if req.GetData().GetOperationalTimeStart() > req.GetData().GetOperationalTimeEnd() {
		return nil, status.Errorf(codes.InvalidArgument, "Operational Time Start should not be greater than Operational Time End")
	}

	dataStoreORM, err := req.GetData().ToORM(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}

	if req.GetId() > 0 {

		result.Message = "Updated"

		_, err := s.StoreDetail(ctx, &pb.StoreDetailRequest{Id: req.GetId()})
		if err != nil {
			return nil, err
		}

		dataStoreORM.Id = req.GetId()

	}

	singleProductORM, err := s.provider.GetStoreDetail(ctx, &pb.StoreORM{Url: strings.ToLower(req.GetData().GetUrl())})
	if err != nil {
		return nil, err
	}

	if singleProductORM != nil && req.GetId() < 1 {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Duplicate Url: %v", req.GetData().GetUrl()))
	}

	storeORM, err := s.provider.UpdateOrCreateStore(ctx, &dataStoreORM)
	if err != nil {
		return nil, err
	}

	storePB, err := storeORM.ToPB(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}

	result.Data = &storePB

	return result, nil

}

func (s *Server) StoreDelete(ctx context.Context, req *pb.StoreDeleteRequest) (*pb.StoreDeleteResponse, error) {

	result := &pb.StoreDeleteResponse{
		Error:   false,
		Code:    200,
		Message: "Success",
	}

	if req.GetId() < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "Id is required")
	}

	storeDetailRes, err := s.StoreDetail(ctx, &pb.StoreDetailRequest{Id: req.GetId()})
	if err != nil {
		return nil, err
	}

	storeORM, err := storeDetailRes.GetData().ToORM(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}

	_, err = s.provider.DeleteStore(ctx, &storeORM)
	if err != nil {
		return nil, err
	}

	storePB, err := storeORM.ToPB(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}

	result.Data = &storePB

	return result, nil

}
