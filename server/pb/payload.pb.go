// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: payload.proto

package pb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StoreListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32     `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Page   int32     `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Filter string    `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
	Query  string    `protobuf:"bytes,4,opt,name=query,proto3" json:"query,omitempty"`
	Sort   string    `protobuf:"bytes,11,opt,name=sort,proto3" json:"sort,omitempty"`
	Dir    Direction `protobuf:"varint,12,opt,name=dir,proto3,enum=swing.v1.Direction" json:"dir,omitempty"`
}

func (x *StoreListRequest) Reset() {
	*x = StoreListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreListRequest) ProtoMessage() {}

func (x *StoreListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreListRequest.ProtoReflect.Descriptor instead.
func (*StoreListRequest) Descriptor() ([]byte, []int) {
	return file_payload_proto_rawDescGZIP(), []int{0}
}

func (x *StoreListRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *StoreListRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *StoreListRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *StoreListRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *StoreListRequest) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

func (x *StoreListRequest) GetDir() Direction {
	if x != nil {
		return x.Dir
	}
	return Direction_DESC
}

type StoreListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error      bool                `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Code       int32               `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message    string              `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Data       []*Store            `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty"`
	Pagination *PaginationResponse `protobuf:"bytes,5,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *StoreListResponse) Reset() {
	*x = StoreListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payload_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreListResponse) ProtoMessage() {}

func (x *StoreListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payload_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreListResponse.ProtoReflect.Descriptor instead.
func (*StoreListResponse) Descriptor() ([]byte, []int) {
	return file_payload_proto_rawDescGZIP(), []int{1}
}

func (x *StoreListResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *StoreListResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *StoreListResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *StoreListResponse) GetData() []*Store {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *StoreListResponse) GetPagination() *PaginationResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type StoreDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *StoreDetailRequest) Reset() {
	*x = StoreDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payload_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreDetailRequest) ProtoMessage() {}

func (x *StoreDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payload_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreDetailRequest.ProtoReflect.Descriptor instead.
func (*StoreDetailRequest) Descriptor() ([]byte, []int) {
	return file_payload_proto_rawDescGZIP(), []int{2}
}

func (x *StoreDetailRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type StoreDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error   bool   `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Code    int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Data    *Store `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *StoreDetailResponse) Reset() {
	*x = StoreDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payload_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreDetailResponse) ProtoMessage() {}

func (x *StoreDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payload_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreDetailResponse.ProtoReflect.Descriptor instead.
func (*StoreDetailResponse) Descriptor() ([]byte, []int) {
	return file_payload_proto_rawDescGZIP(), []int{3}
}

func (x *StoreDetailResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *StoreDetailResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *StoreDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *StoreDetailResponse) GetData() *Store {
	if x != nil {
		return x.Data
	}
	return nil
}

type StoreCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Data *Store `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *StoreCreateRequest) Reset() {
	*x = StoreCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payload_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreCreateRequest) ProtoMessage() {}

func (x *StoreCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payload_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreCreateRequest.ProtoReflect.Descriptor instead.
func (*StoreCreateRequest) Descriptor() ([]byte, []int) {
	return file_payload_proto_rawDescGZIP(), []int{4}
}

func (x *StoreCreateRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *StoreCreateRequest) GetData() *Store {
	if x != nil {
		return x.Data
	}
	return nil
}

type StoreCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error   bool   `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Code    int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Data    *Store `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *StoreCreateResponse) Reset() {
	*x = StoreCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payload_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreCreateResponse) ProtoMessage() {}

func (x *StoreCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payload_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreCreateResponse.ProtoReflect.Descriptor instead.
func (*StoreCreateResponse) Descriptor() ([]byte, []int) {
	return file_payload_proto_rawDescGZIP(), []int{5}
}

func (x *StoreCreateResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *StoreCreateResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *StoreCreateResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *StoreCreateResponse) GetData() *Store {
	if x != nil {
		return x.Data
	}
	return nil
}

type StoreDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *StoreDeleteRequest) Reset() {
	*x = StoreDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payload_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreDeleteRequest) ProtoMessage() {}

func (x *StoreDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payload_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreDeleteRequest.ProtoReflect.Descriptor instead.
func (*StoreDeleteRequest) Descriptor() ([]byte, []int) {
	return file_payload_proto_rawDescGZIP(), []int{6}
}

func (x *StoreDeleteRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type StoreDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error   bool   `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Code    int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Data    *Store `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *StoreDeleteResponse) Reset() {
	*x = StoreDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payload_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreDeleteResponse) ProtoMessage() {}

func (x *StoreDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payload_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreDeleteResponse.ProtoReflect.Descriptor instead.
func (*StoreDeleteResponse) Descriptor() ([]byte, []int) {
	return file_payload_proto_rawDescGZIP(), []int{7}
}

func (x *StoreDeleteResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *StoreDeleteResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *StoreDeleteResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *StoreDeleteResponse) GetData() *Store {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_payload_proto protoreflect.FileDescriptor

var file_payload_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x73, 0x77, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x67, 0x6f, 0x72, 0x6d, 0x5f, 0x64, 0x62, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x09, 0x0a, 0x10, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x54, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x3e, 0x92, 0x41, 0x3b, 0x32, 0x39, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x70, 0x65, 0x72, 0x20, 0x70, 0x61,
	0x67, 0x65, 0x2e, 0x20, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x20, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x20, 0x27, 0x2d, 0x31, 0x27, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x64, 0x69, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x20, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x55, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x41, 0x92,
	0x41, 0x3e, 0x32, 0x3c, 0x50, 0x61, 0x67, 0x65, 0x20, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x20,
	0x74, 0x6f, 0x20, 0x67, 0x65, 0x74, 0x20, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x20, 0x44, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x27, 0x2d, 0x31, 0x27, 0x20,
	0x66, 0x6f, 0x72, 0x20, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x20, 0x70, 0x61, 0x67, 0x65,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x80, 0x04, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0xe7, 0x03, 0x92, 0x41, 0xe3, 0x03, 0x32, 0xe0,
	0x03, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x27, 0x41, 0x4e, 0x44, 0x27, 0x20, 0x71, 0x75,
	0x65, 0x72, 0x69, 0x65, 0x73, 0x20, 0x0a, 0x20, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x3a, 0x20,
	0x5b, 0x6b, 0x65, 0x79, 0x3a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2c, 0x6b, 0x65, 0x79, 0x3a, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x5d, 0x20, 0x0a, 0x20, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x3a,
	0x20, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x3d, 0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x54, 0x6f, 0x6b,
	0x6f, 0x2c, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x3a, 0x31, 0x20, 0x3d, 0x3e, 0x20,
	0x57, 0x48, 0x45, 0x52, 0x45, 0x20, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x20, 0x3d,
	0x20, 0x31, 0x20, 0x41, 0x4e, 0x44, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x20, 0x3d, 0x20, 0x27, 0x54,
	0x6f, 0x6b, 0x6f, 0x27, 0x20, 0x0a, 0x20, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x3a, 0x20, 0x0a, 0x20, 0x2d, 0x20, 0x27, 0x6b, 0x65, 0x79, 0x3a, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x27, 0x20, 0x3d, 0x3e, 0x20, 0x27, 0x6b, 0x65, 0x79, 0x20, 0x3d, 0x20, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x27, 0x20, 0x0a, 0x20, 0x2d, 0x20, 0x27, 0x6b, 0x65, 0x79, 0x3a, 0x25, 0x25,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x27, 0x20, 0x3d, 0x3e, 0x20, 0x27, 0x6b, 0x65, 0x79, 0x20, 0x4c,
	0x49, 0x4b, 0x45, 0x20, 0x25, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x25, 0x27, 0x20, 0x0a, 0x20, 0x2d,
	0x20, 0x27, 0x6b, 0x65, 0x79, 0x3a, 0x25, 0x21, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x27, 0x20, 0x3d,
	0x3e, 0x20, 0x27, 0x6b, 0x65, 0x79, 0x20, 0x49, 0x4c, 0x49, 0x4b, 0x45, 0x20, 0x25, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x25, 0x27, 0x20, 0x0a, 0x20, 0x2d, 0x20, 0x27, 0x6b, 0x65, 0x79, 0x3a, 0x3e,
	0x3d, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x27, 0x20, 0x3d, 0x3e, 0x20, 0x27, 0x6b, 0x65, 0x79, 0x20,
	0x3e, 0x3d, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x27, 0x20, 0x0a, 0x20, 0x2d, 0x20, 0x27, 0x6b,
	0x65, 0x79, 0x3a, 0x3c, 0x3d, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x27, 0x20, 0x3d, 0x3e, 0x20, 0x27,
	0x6b, 0x65, 0x79, 0x20, 0x3c, 0x3d, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x27, 0x20, 0x0a, 0x20,
	0x2d, 0x20, 0x27, 0x6b, 0x65, 0x79, 0x31, 0x3a, 0x3e, 0x3d, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x31,
	0x2c, 0x6b, 0x65, 0x79, 0x32, 0x3a, 0x3c, 0x3d, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x27, 0x20,
	0x3d, 0x3e, 0x20, 0x27, 0x6b, 0x65, 0x79, 0x31, 0x20, 0x3e, 0x3d, 0x20, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x31, 0x20, 0x41, 0x4e, 0x44, 0x20, 0x6b, 0x65, 0x79, 0x32, 0x20, 0x3c, 0x3d, 0x20, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x32, 0x27, 0x20, 0x42, 0x45, 0x54, 0x57, 0x45, 0x45, 0x4e, 0x20, 0x45,
	0x58, 0x41, 0x4d, 0x50, 0x4c, 0x45, 0x20, 0x0a, 0x20, 0x2d, 0x20, 0x27, 0x6b, 0x65, 0x79, 0x31,
	0x2e, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x31, 0x3a, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x27, 0x20, 0x3d, 0x3e, 0x20, 0x27, 0x6b, 0x65, 0x79, 0x31, 0x2d, 0x3e, 0x3e, 0x6e, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x31, 0x20, 0x3d, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x27, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0xdd, 0x02, 0x0a, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0xc6, 0x02, 0x92, 0x41, 0xc2, 0x02,
	0x32, 0xbf, 0x02, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x27, 0x4f, 0x52, 0x27, 0x20, 0x71,
	0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x20, 0x0a, 0x20, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x3a,
	0x20, 0x5b, 0x6b, 0x65, 0x79, 0x31, 0x2c, 0x6b, 0x65, 0x79, 0x32, 0x2c, 0x6b, 0x65, 0x79, 0x33,
	0x3a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5d, 0x20, 0x0a, 0x20, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x3a, 0x20, 0x71, 0x75, 0x65, 0x72, 0x79, 0x3d, 0x6e, 0x61, 0x6d, 0x65, 0x2c, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x54, 0x6f, 0x6b, 0x6f, 0x20, 0x3d,
	0x3e, 0x20, 0x57, 0x48, 0x45, 0x52, 0x45, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x20, 0x3d, 0x20, 0x27,
	0x54, 0x6f, 0x6b, 0x6f, 0x27, 0x20, 0x4f, 0x52, 0x20, 0x74, 0x79, 0x70, 0x65, 0x20, 0x3d, 0x20,
	0x27, 0x54, 0x6f, 0x6b, 0x6f, 0x27, 0x20, 0x0a, 0x20, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x20, 0x0a, 0x20, 0x2d, 0x20, 0x27, 0x6b, 0x65, 0x79, 0x3a, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x27, 0x20, 0x3d, 0x3e, 0x20, 0x27, 0x6b, 0x65, 0x79, 0x20, 0x3d, 0x20,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x27, 0x20, 0x0a, 0x20, 0x2d, 0x20, 0x27, 0x6b, 0x65, 0x79, 0x3a,
	0x25, 0x25, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x27, 0x20, 0x3d, 0x3e, 0x20, 0x27, 0x6b, 0x65, 0x79,
	0x20, 0x4c, 0x49, 0x4b, 0x45, 0x20, 0x25, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x25, 0x27, 0x20, 0x0a,
	0x20, 0x2d, 0x20, 0x27, 0x6b, 0x65, 0x79, 0x3a, 0x25, 0x21, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x27,
	0x20, 0x3d, 0x3e, 0x20, 0x27, 0x6b, 0x65, 0x79, 0x20, 0x49, 0x4c, 0x49, 0x4b, 0x45, 0x20, 0x25,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x25, 0x27, 0x20, 0x0a, 0x20, 0x2d, 0x20, 0x27, 0x6b, 0x65, 0x79,
	0x31, 0x2e, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x31, 0x3a, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x27, 0x20, 0x3d, 0x3e, 0x20, 0x27, 0x6b, 0x65, 0x79, 0x31, 0x2d, 0x3e, 0x3e, 0x6e,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x31, 0x20, 0x3d, 0x20, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x27, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x32, 0x0a, 0x04, 0x73, 0x6f, 0x72,
	0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1e, 0x92, 0x41, 0x1b, 0x32, 0x19, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x62, 0x79, 0x20, 0x63, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x3b, 0x0a,
	0x03, 0x64, 0x69, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x73, 0x77, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x14, 0x92, 0x41, 0x11, 0x32, 0x0f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x20, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x64, 0x69, 0x72, 0x22, 0xba, 0x01, 0x0a, 0x11, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x77, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3c, 0x0a, 0x0a, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x73, 0x77, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x24, 0x0a, 0x12, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x7e, 0x0a,
	0x13, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x77, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x49, 0x0a,
	0x12, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x77, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x7e, 0x0a, 0x13, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x77, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x24, 0x0a, 0x12, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x7e,
	0x0a, 0x13, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x77, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x06,
	0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_payload_proto_rawDescOnce sync.Once
	file_payload_proto_rawDescData = file_payload_proto_rawDesc
)

func file_payload_proto_rawDescGZIP() []byte {
	file_payload_proto_rawDescOnce.Do(func() {
		file_payload_proto_rawDescData = protoimpl.X.CompressGZIP(file_payload_proto_rawDescData)
	})
	return file_payload_proto_rawDescData
}

var file_payload_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_payload_proto_goTypes = []interface{}{
	(*StoreListRequest)(nil),    // 0: swing.v1.StoreListRequest
	(*StoreListResponse)(nil),   // 1: swing.v1.StoreListResponse
	(*StoreDetailRequest)(nil),  // 2: swing.v1.StoreDetailRequest
	(*StoreDetailResponse)(nil), // 3: swing.v1.StoreDetailResponse
	(*StoreCreateRequest)(nil),  // 4: swing.v1.StoreCreateRequest
	(*StoreCreateResponse)(nil), // 5: swing.v1.StoreCreateResponse
	(*StoreDeleteRequest)(nil),  // 6: swing.v1.StoreDeleteRequest
	(*StoreDeleteResponse)(nil), // 7: swing.v1.StoreDeleteResponse
	(Direction)(0),              // 8: swing.v1.Direction
	(*Store)(nil),               // 9: swing.v1.Store
	(*PaginationResponse)(nil),  // 10: swing.v1.PaginationResponse
}
var file_payload_proto_depIdxs = []int32{
	8,  // 0: swing.v1.StoreListRequest.dir:type_name -> swing.v1.Direction
	9,  // 1: swing.v1.StoreListResponse.data:type_name -> swing.v1.Store
	10, // 2: swing.v1.StoreListResponse.pagination:type_name -> swing.v1.PaginationResponse
	9,  // 3: swing.v1.StoreDetailResponse.data:type_name -> swing.v1.Store
	9,  // 4: swing.v1.StoreCreateRequest.data:type_name -> swing.v1.Store
	9,  // 5: swing.v1.StoreCreateResponse.data:type_name -> swing.v1.Store
	9,  // 6: swing.v1.StoreDeleteResponse.data:type_name -> swing.v1.Store
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_payload_proto_init() }
func file_payload_proto_init() {
	if File_payload_proto != nil {
		return
	}
	file_core_proto_init()
	file_gorm_db_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_payload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreListRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_payload_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_payload_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreDetailRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_payload_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreDetailResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_payload_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreCreateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_payload_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreCreateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_payload_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreDeleteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_payload_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreDeleteResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_payload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_payload_proto_goTypes,
		DependencyIndexes: file_payload_proto_depIdxs,
		MessageInfos:      file_payload_proto_msgTypes,
	}.Build()
	File_payload_proto = out.File
	file_payload_proto_rawDesc = nil
	file_payload_proto_goTypes = nil
	file_payload_proto_depIdxs = nil
}
