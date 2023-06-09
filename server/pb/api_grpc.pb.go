// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: api.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ApiServiceClient is the client API for ApiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiServiceClient interface {
	HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
	StoreList(ctx context.Context, in *StoreListRequest, opts ...grpc.CallOption) (*StoreListResponse, error)
	StoreDetail(ctx context.Context, in *StoreDetailRequest, opts ...grpc.CallOption) (*StoreDetailResponse, error)
	StoreCreate(ctx context.Context, in *StoreCreateRequest, opts ...grpc.CallOption) (*StoreCreateResponse, error)
	StoreDelete(ctx context.Context, in *StoreDeleteRequest, opts ...grpc.CallOption) (*StoreDeleteResponse, error)
	ProductList(ctx context.Context, in *ProductListRequest, opts ...grpc.CallOption) (*ProductListResponse, error)
	ProductDetail(ctx context.Context, in *ProductDetailRequest, opts ...grpc.CallOption) (*ProductDetailResponse, error)
	ProductCreate(ctx context.Context, in *ProductCreateRequest, opts ...grpc.CallOption) (*ProductCreateResponse, error)
	ProductDelete(ctx context.Context, in *ProductDeleteRequest, opts ...grpc.CallOption) (*ProductDeleteResponse, error)
}

type apiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApiServiceClient(cc grpc.ClientConnInterface) ApiServiceClient {
	return &apiServiceClient{cc}
}

func (c *apiServiceClient) HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/swing.v1.ApiService/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) StoreList(ctx context.Context, in *StoreListRequest, opts ...grpc.CallOption) (*StoreListResponse, error) {
	out := new(StoreListResponse)
	err := c.cc.Invoke(ctx, "/swing.v1.ApiService/StoreList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) StoreDetail(ctx context.Context, in *StoreDetailRequest, opts ...grpc.CallOption) (*StoreDetailResponse, error) {
	out := new(StoreDetailResponse)
	err := c.cc.Invoke(ctx, "/swing.v1.ApiService/StoreDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) StoreCreate(ctx context.Context, in *StoreCreateRequest, opts ...grpc.CallOption) (*StoreCreateResponse, error) {
	out := new(StoreCreateResponse)
	err := c.cc.Invoke(ctx, "/swing.v1.ApiService/StoreCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) StoreDelete(ctx context.Context, in *StoreDeleteRequest, opts ...grpc.CallOption) (*StoreDeleteResponse, error) {
	out := new(StoreDeleteResponse)
	err := c.cc.Invoke(ctx, "/swing.v1.ApiService/StoreDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ProductList(ctx context.Context, in *ProductListRequest, opts ...grpc.CallOption) (*ProductListResponse, error) {
	out := new(ProductListResponse)
	err := c.cc.Invoke(ctx, "/swing.v1.ApiService/ProductList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ProductDetail(ctx context.Context, in *ProductDetailRequest, opts ...grpc.CallOption) (*ProductDetailResponse, error) {
	out := new(ProductDetailResponse)
	err := c.cc.Invoke(ctx, "/swing.v1.ApiService/ProductDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ProductCreate(ctx context.Context, in *ProductCreateRequest, opts ...grpc.CallOption) (*ProductCreateResponse, error) {
	out := new(ProductCreateResponse)
	err := c.cc.Invoke(ctx, "/swing.v1.ApiService/ProductCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ProductDelete(ctx context.Context, in *ProductDeleteRequest, opts ...grpc.CallOption) (*ProductDeleteResponse, error) {
	out := new(ProductDeleteResponse)
	err := c.cc.Invoke(ctx, "/swing.v1.ApiService/ProductDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiServiceServer is the server API for ApiService service.
// All implementations must embed UnimplementedApiServiceServer
// for forward compatibility
type ApiServiceServer interface {
	HealthCheck(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
	StoreList(context.Context, *StoreListRequest) (*StoreListResponse, error)
	StoreDetail(context.Context, *StoreDetailRequest) (*StoreDetailResponse, error)
	StoreCreate(context.Context, *StoreCreateRequest) (*StoreCreateResponse, error)
	StoreDelete(context.Context, *StoreDeleteRequest) (*StoreDeleteResponse, error)
	ProductList(context.Context, *ProductListRequest) (*ProductListResponse, error)
	ProductDetail(context.Context, *ProductDetailRequest) (*ProductDetailResponse, error)
	ProductCreate(context.Context, *ProductCreateRequest) (*ProductCreateResponse, error)
	ProductDelete(context.Context, *ProductDeleteRequest) (*ProductDeleteResponse, error)
	mustEmbedUnimplementedApiServiceServer()
}

// UnimplementedApiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedApiServiceServer struct {
}

func (UnimplementedApiServiceServer) HealthCheck(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedApiServiceServer) StoreList(context.Context, *StoreListRequest) (*StoreListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreList not implemented")
}
func (UnimplementedApiServiceServer) StoreDetail(context.Context, *StoreDetailRequest) (*StoreDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreDetail not implemented")
}
func (UnimplementedApiServiceServer) StoreCreate(context.Context, *StoreCreateRequest) (*StoreCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreCreate not implemented")
}
func (UnimplementedApiServiceServer) StoreDelete(context.Context, *StoreDeleteRequest) (*StoreDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreDelete not implemented")
}
func (UnimplementedApiServiceServer) ProductList(context.Context, *ProductListRequest) (*ProductListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductList not implemented")
}
func (UnimplementedApiServiceServer) ProductDetail(context.Context, *ProductDetailRequest) (*ProductDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductDetail not implemented")
}
func (UnimplementedApiServiceServer) ProductCreate(context.Context, *ProductCreateRequest) (*ProductCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductCreate not implemented")
}
func (UnimplementedApiServiceServer) ProductDelete(context.Context, *ProductDeleteRequest) (*ProductDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductDelete not implemented")
}
func (UnimplementedApiServiceServer) mustEmbedUnimplementedApiServiceServer() {}

// UnsafeApiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiServiceServer will
// result in compilation errors.
type UnsafeApiServiceServer interface {
	mustEmbedUnimplementedApiServiceServer()
}

func RegisterApiServiceServer(s grpc.ServiceRegistrar, srv ApiServiceServer) {
	s.RegisterService(&ApiService_ServiceDesc, srv)
}

func _ApiService_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/swing.v1.ApiService/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).HealthCheck(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_StoreList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).StoreList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/swing.v1.ApiService/StoreList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).StoreList(ctx, req.(*StoreListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_StoreDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).StoreDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/swing.v1.ApiService/StoreDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).StoreDetail(ctx, req.(*StoreDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_StoreCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).StoreCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/swing.v1.ApiService/StoreCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).StoreCreate(ctx, req.(*StoreCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_StoreDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).StoreDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/swing.v1.ApiService/StoreDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).StoreDelete(ctx, req.(*StoreDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ProductList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ProductList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/swing.v1.ApiService/ProductList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ProductList(ctx, req.(*ProductListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ProductDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ProductDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/swing.v1.ApiService/ProductDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ProductDetail(ctx, req.(*ProductDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ProductCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ProductCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/swing.v1.ApiService/ProductCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ProductCreate(ctx, req.(*ProductCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ProductDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ProductDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/swing.v1.ApiService/ProductDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ProductDelete(ctx, req.(*ProductDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiService_ServiceDesc is the grpc.ServiceDesc for ApiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "swing.v1.ApiService",
	HandlerType: (*ApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _ApiService_HealthCheck_Handler,
		},
		{
			MethodName: "StoreList",
			Handler:    _ApiService_StoreList_Handler,
		},
		{
			MethodName: "StoreDetail",
			Handler:    _ApiService_StoreDetail_Handler,
		},
		{
			MethodName: "StoreCreate",
			Handler:    _ApiService_StoreCreate_Handler,
		},
		{
			MethodName: "StoreDelete",
			Handler:    _ApiService_StoreDelete_Handler,
		},
		{
			MethodName: "ProductList",
			Handler:    _ApiService_ProductList_Handler,
		},
		{
			MethodName: "ProductDetail",
			Handler:    _ApiService_ProductDetail_Handler,
		},
		{
			MethodName: "ProductCreate",
			Handler:    _ApiService_ProductCreate_Handler,
		},
		{
			MethodName: "ProductDelete",
			Handler:    _ApiService_ProductDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
