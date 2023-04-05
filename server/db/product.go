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

func (p *GormProvider) GetProduct(ctx context.Context, v *ListFilter, pagination *pb.PaginationResponse) (data []*pb.ProductORM, err error) {

	query := p.db_main.Model(&pb.ProductORM{})

	if v.Data != nil {
		query = query.Where(v.Data)
	}

	query = query.Scopes(FilterScoope(v.Filter), QueryScoop(v.Query)).Order("updated_at DESC")
	query = query.Scopes(Paginate(data, pagination, query), CustomOrderScoop(v.CustomOrder), Sort(v.Sort), Sort(&pb.Sort{Column: "updated_at", Direction: "DESC"}))

	if err := query.Find(&data).Error; err != nil {
		logrus.Errorf("[db][func: GetProduct] Failed when execute query: %s", err)
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}
	}

	return data, nil

}

func (p *GormProvider) GetProductDetail(ctx context.Context, v *pb.ProductORM) (data *pb.ProductORM, err error) {

	query := p.db_main.Model(&pb.ProductORM{})

	query = query.Where(v)

	if err := query.Find(&data).Error; err != nil {
		logrus.Errorf("[db][func: GetProductDetail] Failed when execute query: %s", err)
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}

	if data.Id < 1 {
		return nil, nil
	}

	return data, nil

}

func (p *GormProvider) UpdateOrCreateProduct(ctx context.Context, data *pb.ProductORM) (*pb.ProductORM, error) {

	if data.Id > 0 {

		model := &pb.ProductORM{
			Id: data.Id,
		}

		m := structs.Map(data)

		for k := range m {
			if k == "CreatedAt" || k == "DeletedAt" {
				delete(m, k)
			} else if k == "UpdatedAt" {
				m[k] = time.Now()
			}
		}

		if err := p.db_main.Model(&model).Updates(&m).Error; err != nil {
			logrus.Errorf("[db][func: UpdateOrCreateProduct] Failed when execute query: %s", err)
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, status.Error(codes.NotFound, "ID Not Found")
			} else {
				return nil, status.Error(codes.Internal, "Internal Error")
			}
		}

		return model, nil

	} else {

		if err := p.db_main.Create(&data).Error; err != nil {
			logrus.Errorf("[db][func: UpdateOrCreateProduct] Failed when execute query: %s", err)
			return nil, status.Error(codes.Internal, "Internal Error")
		}

		return data, nil

	}

}

func (p *GormProvider) DeleteProduct(ctx context.Context, data *pb.ProductORM) (*pb.ProductORM, error) {

	if data.Id > 0 {

		model := &pb.ProductORM{
			Id: data.Id,
		}

		if err := p.db_main.Model(&model).Update("deleted_at", time.Now()).Error; err != nil {
			logrus.Errorf("[db][func: DeleteProduct] Failed when execute query: %s", err)
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

func (p *GormProvider) ActionProduct(ctx context.Context, id int32, isActive bool) (*pb.ProductORM, error) {

	if id > 0 {

		model := &pb.ProductORM{
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
