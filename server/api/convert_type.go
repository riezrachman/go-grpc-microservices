package api

import (
	"context"
	pb "swing/server/pb"
)

func (s *Server) ProductORM_TProduct(ctx context.Context, m *pb.ProductORM) (*pb.TProduct, error) {
	fromPB, err := m.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	var store *pb.TStore
	storeORM, err := s.provider.GetStoreDetail(ctx, &pb.StoreORM{Id: *m.StoreId})
	if err != nil {
		return nil, err
	}
	if storeORM != nil {
		store, err = s.StoreORM_TStore(ctx, storeORM)
		if err != nil {
			return nil, err
		}
	}
	to := pb.TProduct{
		Id:          fromPB.GetId(),
		Title:       fromPB.GetTitle(),
		Url:         fromPB.GetUrl(),
		Price:       fromPB.GetPrice(),
		Description: fromPB.GetDescription(),
		Store:       store,
		CreatedAt:   fromPB.GetCreatedAt(),
		UpdatedAt:   fromPB.GetUpdatedAt(),
	}
	return &to, nil
}

func (s *Server) StoreORM_TStore(ctx context.Context, m *pb.StoreORM) (*pb.TStore, error) {
	fromPB, err := m.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	to := pb.TStore{
		Id:                   fromPB.GetId(),
		Name:                 fromPB.GetName(),
		Url:                  fromPB.GetUrl(),
		Address:              fromPB.GetAddress(),
		Phone:                fromPB.GetPhone(),
		OperationalTimeStart: fromPB.GetOperationalTimeStart(),
		OperationalTimeEnd:   fromPB.GetOperationalTimeEnd(),
		CreatedAt:            fromPB.GetCreatedAt(),
		UpdatedAt:            fromPB.GetUpdatedAt(),
	}
	return &to, nil
}
