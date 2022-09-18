// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: proto/audit.proto

package audit_logs

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

// AuditServiceClient is the client API for AuditService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuditServiceClient interface {
	Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogResponse, error)
	GetLogByEntity(ctx context.Context, in *LogByEntityRequest, opts ...grpc.CallOption) (*LogsResponse, error)
	GetLogByEntityPaginated(ctx context.Context, in *LogByEntityPagedRequest, opts ...grpc.CallOption) (*LogsResponse, error)
	GetLogByEntityID(ctx context.Context, in *LogByEntityIDRequest, opts ...grpc.CallOption) (*LogsResponse, error)
	GetLogByEntityIDPaginated(ctx context.Context, in *LogByEntityIDPagedRequest, opts ...grpc.CallOption) (*LogsResponse, error)
}

type auditServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuditServiceClient(cc grpc.ClientConnInterface) AuditServiceClient {
	return &auditServiceClient{cc}
}

func (c *auditServiceClient) Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogResponse, error) {
	out := new(LogResponse)
	err := c.cc.Invoke(ctx, "/audit.AuditService/Log", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auditServiceClient) GetLogByEntity(ctx context.Context, in *LogByEntityRequest, opts ...grpc.CallOption) (*LogsResponse, error) {
	out := new(LogsResponse)
	err := c.cc.Invoke(ctx, "/audit.AuditService/GetLogByEntity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auditServiceClient) GetLogByEntityPaginated(ctx context.Context, in *LogByEntityPagedRequest, opts ...grpc.CallOption) (*LogsResponse, error) {
	out := new(LogsResponse)
	err := c.cc.Invoke(ctx, "/audit.AuditService/GetLogByEntityPaginated", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auditServiceClient) GetLogByEntityID(ctx context.Context, in *LogByEntityIDRequest, opts ...grpc.CallOption) (*LogsResponse, error) {
	out := new(LogsResponse)
	err := c.cc.Invoke(ctx, "/audit.AuditService/GetLogByEntityID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auditServiceClient) GetLogByEntityIDPaginated(ctx context.Context, in *LogByEntityIDPagedRequest, opts ...grpc.CallOption) (*LogsResponse, error) {
	out := new(LogsResponse)
	err := c.cc.Invoke(ctx, "/audit.AuditService/GetLogByEntityIDPaginated", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuditServiceServer is the server API for AuditService service.
// All implementations should embed UnimplementedAuditServiceServer
// for forward compatibility
type AuditServiceServer interface {
	Log(context.Context, *LogRequest) (*LogResponse, error)
	GetLogByEntity(context.Context, *LogByEntityRequest) (*LogsResponse, error)
	GetLogByEntityPaginated(context.Context, *LogByEntityPagedRequest) (*LogsResponse, error)
	GetLogByEntityID(context.Context, *LogByEntityIDRequest) (*LogsResponse, error)
	GetLogByEntityIDPaginated(context.Context, *LogByEntityIDPagedRequest) (*LogsResponse, error)
}

// UnimplementedAuditServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAuditServiceServer struct {
}

func (UnimplementedAuditServiceServer) Log(context.Context, *LogRequest) (*LogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Log not implemented")
}
func (UnimplementedAuditServiceServer) GetLogByEntity(context.Context, *LogByEntityRequest) (*LogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLogByEntity not implemented")
}
func (UnimplementedAuditServiceServer) GetLogByEntityPaginated(context.Context, *LogByEntityPagedRequest) (*LogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLogByEntityPaginated not implemented")
}
func (UnimplementedAuditServiceServer) GetLogByEntityID(context.Context, *LogByEntityIDRequest) (*LogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLogByEntityID not implemented")
}
func (UnimplementedAuditServiceServer) GetLogByEntityIDPaginated(context.Context, *LogByEntityIDPagedRequest) (*LogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLogByEntityIDPaginated not implemented")
}

// UnsafeAuditServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuditServiceServer will
// result in compilation errors.
type UnsafeAuditServiceServer interface {
	mustEmbedUnimplementedAuditServiceServer()
}

func RegisterAuditServiceServer(s grpc.ServiceRegistrar, srv AuditServiceServer) {
	s.RegisterService(&AuditService_ServiceDesc, srv)
}

func _AuditService_Log_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuditServiceServer).Log(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/audit.AuditService/Log",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuditServiceServer).Log(ctx, req.(*LogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuditService_GetLogByEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogByEntityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuditServiceServer).GetLogByEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/audit.AuditService/GetLogByEntity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuditServiceServer).GetLogByEntity(ctx, req.(*LogByEntityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuditService_GetLogByEntityPaginated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogByEntityPagedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuditServiceServer).GetLogByEntityPaginated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/audit.AuditService/GetLogByEntityPaginated",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuditServiceServer).GetLogByEntityPaginated(ctx, req.(*LogByEntityPagedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuditService_GetLogByEntityID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogByEntityIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuditServiceServer).GetLogByEntityID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/audit.AuditService/GetLogByEntityID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuditServiceServer).GetLogByEntityID(ctx, req.(*LogByEntityIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuditService_GetLogByEntityIDPaginated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogByEntityIDPagedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuditServiceServer).GetLogByEntityIDPaginated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/audit.AuditService/GetLogByEntityIDPaginated",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuditServiceServer).GetLogByEntityIDPaginated(ctx, req.(*LogByEntityIDPagedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuditService_ServiceDesc is the grpc.ServiceDesc for AuditService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuditService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "audit.AuditService",
	HandlerType: (*AuditServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Log",
			Handler:    _AuditService_Log_Handler,
		},
		{
			MethodName: "GetLogByEntity",
			Handler:    _AuditService_GetLogByEntity_Handler,
		},
		{
			MethodName: "GetLogByEntityPaginated",
			Handler:    _AuditService_GetLogByEntityPaginated_Handler,
		},
		{
			MethodName: "GetLogByEntityID",
			Handler:    _AuditService_GetLogByEntityID_Handler,
		},
		{
			MethodName: "GetLogByEntityIDPaginated",
			Handler:    _AuditService_GetLogByEntityIDPaginated_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/audit.proto",
}
