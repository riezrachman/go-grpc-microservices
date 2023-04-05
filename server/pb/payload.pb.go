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
	Data   *Store    `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
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

func (x *StoreListRequest) GetData() *Store {
	if x != nil {
		return x.Data
	}
	return nil
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

var File_payload_proto protoreflect.FileDescriptor

var file_payload_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x73, 0x77, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x67, 0x6f, 0x72, 0x6d, 0x5f, 0x64, 0x62, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x03, 0x0a, 0x10, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4c, 0x69,
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
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x23, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x77, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x32, 0x0a, 0x04, 0x73, 0x6f, 0x72,
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
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_payload_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_payload_proto_goTypes = []interface{}{
	(*StoreListRequest)(nil),   // 0: swing.v1.StoreListRequest
	(*StoreListResponse)(nil),  // 1: swing.v1.StoreListResponse
	(*Store)(nil),              // 2: swing.v1.Store
	(Direction)(0),             // 3: swing.v1.Direction
	(*PaginationResponse)(nil), // 4: swing.v1.PaginationResponse
}
var file_payload_proto_depIdxs = []int32{
	2, // 0: swing.v1.StoreListRequest.data:type_name -> swing.v1.Store
	3, // 1: swing.v1.StoreListRequest.dir:type_name -> swing.v1.Direction
	2, // 2: swing.v1.StoreListResponse.data:type_name -> swing.v1.Store
	4, // 3: swing.v1.StoreListResponse.pagination:type_name -> swing.v1.PaginationResponse
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_payload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
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
