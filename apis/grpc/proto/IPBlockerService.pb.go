// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: apis/grpc/proto/IPBlockerService.proto

package proto

import (
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

type HealthCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HealthCheckRequest) Reset() {
	*x = HealthCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_grpc_proto_IPBlockerService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckRequest) ProtoMessage() {}

func (x *HealthCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apis_grpc_proto_IPBlockerService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckRequest.ProtoReflect.Descriptor instead.
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return file_apis_grpc_proto_IPBlockerService_proto_rawDescGZIP(), []int{0}
}

type HealthCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceOnline      bool `protobuf:"varint,1,opt,name=ServiceOnline,proto3" json:"ServiceOnline,omitempty"`
	DatabaseAccessable bool `protobuf:"varint,2,opt,name=DatabaseAccessable,proto3" json:"DatabaseAccessable,omitempty"`
}

func (x *HealthCheckResponse) Reset() {
	*x = HealthCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_grpc_proto_IPBlockerService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckResponse) ProtoMessage() {}

func (x *HealthCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apis_grpc_proto_IPBlockerService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckResponse.ProtoReflect.Descriptor instead.
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return file_apis_grpc_proto_IPBlockerService_proto_rawDescGZIP(), []int{1}
}

func (x *HealthCheckResponse) GetServiceOnline() bool {
	if x != nil {
		return x.ServiceOnline
	}
	return false
}

func (x *HealthCheckResponse) GetDatabaseAccessable() bool {
	if x != nil {
		return x.DatabaseAccessable
	}
	return false
}

type AuthroizeIPRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IPAddress      string   `protobuf:"bytes,1,opt,name=IPAddress,proto3" json:"IPAddress,omitempty"`           // required
	ValidCountries []string `protobuf:"bytes,2,rep,name=ValidCountries,proto3" json:"ValidCountries,omitempty"` // required
}

func (x *AuthroizeIPRequest) Reset() {
	*x = AuthroizeIPRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_grpc_proto_IPBlockerService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthroizeIPRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthroizeIPRequest) ProtoMessage() {}

func (x *AuthroizeIPRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apis_grpc_proto_IPBlockerService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthroizeIPRequest.ProtoReflect.Descriptor instead.
func (*AuthroizeIPRequest) Descriptor() ([]byte, []int) {
	return file_apis_grpc_proto_IPBlockerService_proto_rawDescGZIP(), []int{2}
}

func (x *AuthroizeIPRequest) GetIPAddress() string {
	if x != nil {
		return x.IPAddress
	}
	return ""
}

func (x *AuthroizeIPRequest) GetValidCountries() []string {
	if x != nil {
		return x.ValidCountries
	}
	return nil
}

type AuthroizeIPResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Authorized bool `protobuf:"varint,1,opt,name=Authorized,proto3" json:"Authorized,omitempty"`
}

func (x *AuthroizeIPResponse) Reset() {
	*x = AuthroizeIPResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_grpc_proto_IPBlockerService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthroizeIPResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthroizeIPResponse) ProtoMessage() {}

func (x *AuthroizeIPResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apis_grpc_proto_IPBlockerService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthroizeIPResponse.ProtoReflect.Descriptor instead.
func (*AuthroizeIPResponse) Descriptor() ([]byte, []int) {
	return file_apis_grpc_proto_IPBlockerService_proto_rawDescGZIP(), []int{3}
}

func (x *AuthroizeIPResponse) GetAuthorized() bool {
	if x != nil {
		return x.Authorized
	}
	return false
}

var File_apis_grpc_proto_IPBlockerService_proto protoreflect.FileDescriptor

var file_apis_grpc_proto_IPBlockerService_proto_rawDesc = []byte{
	0x0a, 0x26, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x49, 0x50, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x14, 0x0a, 0x12, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x6b, 0x0a, 0x13, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0d,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x6e, 0x6c, 0x69,
	0x6e, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12,
	0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x22, 0x5a, 0x0a, 0x12, 0x41, 0x75, 0x74, 0x68, 0x72, 0x6f, 0x69, 0x7a, 0x65, 0x49,
	0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x50, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x49, 0x50, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0x35,
	0x0a, 0x13, 0x41, 0x75, 0x74, 0x68, 0x72, 0x6f, 0x69, 0x7a, 0x65, 0x49, 0x50, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x65, 0x64, 0x32, 0xa2, 0x01, 0x0a, 0x10, 0x49, 0x50, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0b, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x46, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x49,
	0x50, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x72, 0x6f,
	0x69, 0x7a, 0x65, 0x49, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x72, 0x6f, 0x69, 0x7a, 0x65, 0x49, 0x50,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x18, 0x5a, 0x16, 0x49, 0x50,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apis_grpc_proto_IPBlockerService_proto_rawDescOnce sync.Once
	file_apis_grpc_proto_IPBlockerService_proto_rawDescData = file_apis_grpc_proto_IPBlockerService_proto_rawDesc
)

func file_apis_grpc_proto_IPBlockerService_proto_rawDescGZIP() []byte {
	file_apis_grpc_proto_IPBlockerService_proto_rawDescOnce.Do(func() {
		file_apis_grpc_proto_IPBlockerService_proto_rawDescData = protoimpl.X.CompressGZIP(file_apis_grpc_proto_IPBlockerService_proto_rawDescData)
	})
	return file_apis_grpc_proto_IPBlockerService_proto_rawDescData
}

var file_apis_grpc_proto_IPBlockerService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_apis_grpc_proto_IPBlockerService_proto_goTypes = []interface{}{
	(*HealthCheckRequest)(nil),  // 0: proto.HealthCheckRequest
	(*HealthCheckResponse)(nil), // 1: proto.HealthCheckResponse
	(*AuthroizeIPRequest)(nil),  // 2: proto.AuthroizeIPRequest
	(*AuthroizeIPResponse)(nil), // 3: proto.AuthroizeIPResponse
}
var file_apis_grpc_proto_IPBlockerService_proto_depIdxs = []int32{
	0, // 0: proto.IPBlockerService.HealthCheck:input_type -> proto.HealthCheckRequest
	2, // 1: proto.IPBlockerService.AuthorizeIP:input_type -> proto.AuthroizeIPRequest
	1, // 2: proto.IPBlockerService.HealthCheck:output_type -> proto.HealthCheckResponse
	3, // 3: proto.IPBlockerService.AuthorizeIP:output_type -> proto.AuthroizeIPResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_apis_grpc_proto_IPBlockerService_proto_init() }
func file_apis_grpc_proto_IPBlockerService_proto_init() {
	if File_apis_grpc_proto_IPBlockerService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apis_grpc_proto_IPBlockerService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckRequest); i {
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
		file_apis_grpc_proto_IPBlockerService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckResponse); i {
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
		file_apis_grpc_proto_IPBlockerService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthroizeIPRequest); i {
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
		file_apis_grpc_proto_IPBlockerService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthroizeIPResponse); i {
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
			RawDescriptor: file_apis_grpc_proto_IPBlockerService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apis_grpc_proto_IPBlockerService_proto_goTypes,
		DependencyIndexes: file_apis_grpc_proto_IPBlockerService_proto_depIdxs,
		MessageInfos:      file_apis_grpc_proto_IPBlockerService_proto_msgTypes,
	}.Build()
	File_apis_grpc_proto_IPBlockerService_proto = out.File
	file_apis_grpc_proto_IPBlockerService_proto_rawDesc = nil
	file_apis_grpc_proto_IPBlockerService_proto_goTypes = nil
	file_apis_grpc_proto_IPBlockerService_proto_depIdxs = nil
}
