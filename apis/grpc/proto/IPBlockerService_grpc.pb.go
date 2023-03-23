// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: apis/grpc/proto/IPBlockerService.proto

package proto

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

// IPBlockerServiceClient is the client API for IPBlockerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IPBlockerServiceClient interface {
	HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
	AuthorizeIP(ctx context.Context, in *AuthroizeIPRequest, opts ...grpc.CallOption) (*AuthroizeIPResponse, error)
}

type iPBlockerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIPBlockerServiceClient(cc grpc.ClientConnInterface) IPBlockerServiceClient {
	return &iPBlockerServiceClient{cc}
}

func (c *iPBlockerServiceClient) HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/proto.IPBlockerService/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPBlockerServiceClient) AuthorizeIP(ctx context.Context, in *AuthroizeIPRequest, opts ...grpc.CallOption) (*AuthroizeIPResponse, error) {
	out := new(AuthroizeIPResponse)
	err := c.cc.Invoke(ctx, "/proto.IPBlockerService/AuthorizeIP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IPBlockerServiceServer is the server API for IPBlockerService service.
// All implementations must embed UnimplementedIPBlockerServiceServer
// for forward compatibility
type IPBlockerServiceServer interface {
	HealthCheck(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
	AuthorizeIP(context.Context, *AuthroizeIPRequest) (*AuthroizeIPResponse, error)
	mustEmbedUnimplementedIPBlockerServiceServer()
}

// UnimplementedIPBlockerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIPBlockerServiceServer struct {
}

func (UnimplementedIPBlockerServiceServer) HealthCheck(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedIPBlockerServiceServer) AuthorizeIP(context.Context, *AuthroizeIPRequest) (*AuthroizeIPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthorizeIP not implemented")
}
func (UnimplementedIPBlockerServiceServer) mustEmbedUnimplementedIPBlockerServiceServer() {}

// UnsafeIPBlockerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IPBlockerServiceServer will
// result in compilation errors.
type UnsafeIPBlockerServiceServer interface {
	mustEmbedUnimplementedIPBlockerServiceServer()
}

func RegisterIPBlockerServiceServer(s grpc.ServiceRegistrar, srv IPBlockerServiceServer) {
	s.RegisterService(&IPBlockerService_ServiceDesc, srv)
}

func _IPBlockerService_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPBlockerServiceServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.IPBlockerService/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPBlockerServiceServer).HealthCheck(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPBlockerService_AuthorizeIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthroizeIPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPBlockerServiceServer).AuthorizeIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.IPBlockerService/AuthorizeIP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPBlockerServiceServer).AuthorizeIP(ctx, req.(*AuthroizeIPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IPBlockerService_ServiceDesc is the grpc.ServiceDesc for IPBlockerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IPBlockerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.IPBlockerService",
	HandlerType: (*IPBlockerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _IPBlockerService_HealthCheck_Handler,
		},
		{
			MethodName: "AuthorizeIP",
			Handler:    _IPBlockerService_AuthorizeIP_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apis/grpc/proto/IPBlockerService.proto",
}
