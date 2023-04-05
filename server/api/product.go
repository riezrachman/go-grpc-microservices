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

func (s *Server) ProductList(ctx context.Context, req *pb.ProductListRequest) (*pb.ProductListResponse, error) {

	result := &pb.ProductListResponse{
		Error:   false,
		Code:    200,
		Message: "Success",
		Data:    []*pb.TProduct{},
	}

	result.Pagination = setPagination(req.GetLimit(), req.GetPage())

	productORMs, err := s.provider.GetProduct(ctx, &db.ListFilter{
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

	for _, productORM := range productORMs {

		product, err := s.ProductORM_TProduct(ctx, productORM)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}

		result.Data = append(result.Data, product)

	}

	return result, nil

}

func (s *Server) ProductDetail(ctx context.Context, req *pb.ProductDetailRequest) (*pb.ProductDetailResponse, error) {

	result := &pb.ProductDetailResponse{
		Error:   false,
		Code:    200,
		Message: "Success",
	}

	if req.GetId() < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "Id is required")
	}

	productORM, err := s.provider.GetProductDetail(ctx, &pb.ProductORM{
		Id: req.GetId(),
	})
	if err != nil {
		return nil, err
	}

	if productORM == nil {
		return nil, status.Errorf(codes.NotFound, "Data not found")
	}

	product, err := s.ProductORM_TProduct(ctx, productORM)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}

	result.Data = product

	return result, nil

}

func (s *Server) ProductCreate(ctx context.Context, req *pb.ProductCreateRequest) (*pb.ProductCreateResponse, error) {

	result := &pb.ProductCreateResponse{
		Error:   false,
		Code:    200,
		Message: "Success",
	}

	if req.GetData() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Data is required")
	}

	if req.GetStoreId() < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "Store is required")
	}

	storeId := req.GetStoreId()

	storeORM, err := s.provider.GetStoreDetail(ctx, &pb.StoreORM{Id: storeId})
	if err != nil {
		return nil, err
	}

	if storeORM == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Store not found")
	}

	dataProductORM, err := req.GetData().ToORM(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}

	dataProductORM.StoreId = &storeId

	if req.GetId() > 0 {

		result.Message = "Updated"

		_, err := s.ProductDetail(ctx, &pb.ProductDetailRequest{Id: req.GetId()})
		if err != nil {
			return nil, err
		}

		dataProductORM.Id = req.GetId()

	}

	singleProductORM, err := s.provider.GetProductDetail(ctx, &pb.ProductORM{Url: strings.ToLower(req.GetData().GetUrl())})
	if err != nil {
		return nil, err
	}

	if singleProductORM != nil && req.GetId() < 1 {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Duplicate Url: %v", req.GetData().GetUrl()))
	}

	productORM, err := s.provider.UpdateOrCreateProduct(ctx, &dataProductORM)
	if err != nil {
		return nil, err
	}

	product, err := s.ProductORM_TProduct(ctx, productORM)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}

	result.Data = product

	return result, nil

}

func (s *Server) ProductDelete(ctx context.Context, req *pb.ProductDeleteRequest) (*pb.ProductDeleteResponse, error) {

	result := &pb.ProductDeleteResponse{
		Error:   false,
		Code:    200,
		Message: "Success",
	}

	if req.GetId() < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "Id is required")
	}

	productORM, err := s.provider.GetProductDetail(ctx, &pb.ProductORM{Id: req.GetId()})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}

	if productORM == nil {
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}

	_, err = s.provider.DeleteProduct(ctx, productORM)
	if err != nil {
		return nil, err
	}

	product, err := s.ProductORM_TProduct(ctx, productORM)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}

	result.Data = product

	return result, nil

}
