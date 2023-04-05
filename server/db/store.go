package db

import (
	"context"
	"errors"
	"time"

	pb "swing/server/pb"

	"github.com/fatih/structs"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (p *GormProvider) GetStore(ctx context.Context, v *ListFilter, pagination *pb.PaginationResponse) (data []*pb.StoreORM, err error) {

	query := p.db_main.Model(&pb.StoreORM{})

	query = query.Scopes(FilterScoope(v.Filter), QueryScoop(v.Query)).Order("updated_at DESC")
	query = query.Scopes(Paginate(data, pagination, query), CustomOrderScoop(v.CustomOrder), Sort(v.Sort), Sort(&pb.Sort{Column: "updated_at", Direction: "DESC"}))

	if err := query.Find(&data).Error; err != nil {
		logrus.Errorf("[db][func: GetStore] Failed when execute query: %s", err)
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}
	}

	return data, nil

}

func (p *GormProvider) GetStoreDetail(ctx context.Context, v *pb.StoreORM) (data *pb.StoreORM, err error) {

	query := p.db_main.Model(&pb.StoreORM{})

	query = query.Where(v)

	if err := query.Find(&data).Error; err != nil {
		logrus.Errorf("[db][func: GetStoreDetail] Failed when execute query: %s", err)
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}

	if data.Id < 1 {
		return nil, nil
	}

	return data, nil

}

func (p *GormProvider) UpdateOrCreateStore(ctx context.Context, data *pb.StoreORM) (*pb.StoreORM, error) {

	if data.Id > 0 {

		model := &pb.StoreORM{
			Id: data.Id,
		}

		m := structs.Map(data)

		if err := p.db_main.Model(&model).Updates(&m).Error; err != nil {
			logrus.Errorf("[db][func: UpdateOrCreateStore] Failed when execute query: %s", err)
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, status.Error(codes.NotFound, "ID Not Found")
			} else {
				return nil, status.Error(codes.Internal, "Internal Error")
			}
		}

		return model, nil

	} else {

		if err := p.db_main.Create(&data).Error; err != nil {
			logrus.Errorf("[db][func: UpdateOrCreateStore] Failed when execute query: %s", err)
			return nil, status.Error(codes.Internal, "Internal Error")
		}

		return data, nil

	}

}

func (p *GormProvider) DeleteStore(ctx context.Context, data *pb.StoreORM) (*pb.StoreORM, error) {

	if data.Id > 0 {

		model := &pb.StoreORM{
			Id: data.Id,
		}

		if err := p.db_main.Model(&model).Update("deleted_at", time.Now()).Error; err != nil {
			logrus.Errorf("[db][func: DeleteStore] Failed when execute query: %s", err)
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, status.Error(codes.NotFound, "ID Not Found")
			} else {
				return nil, status.Error(codes.Internal, "Internal Error")
			}
		}

		return model, nil

	}

	return nil, status.Error(codes.NotFound, "ID Not Found")

}

func (p *GormProvider) ActionStore(ctx context.Context, id int32, isActive bool) (*pb.StoreORM, error) {

	if id > 0 {

		model := &pb.StoreORM{
			Id: id,
		}

		if err := p.db_main.Model(&model).Update("is_active", isActive).Error; err != nil {
			logrus.Errorf("[db][func: ActionSupplier] Failed when execute query: %s", err)
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, status.Error(codes.NotFound, "ID Not Found")
			} else {
				return nil, status.Error(codes.Internal, "Internal Error")
			}
		}

		return model, nil

	}

	return nil, status.Error(codes.NotFound, "ID Not Found")

}
