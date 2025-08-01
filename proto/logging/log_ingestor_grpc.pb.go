// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.31.1
// source: proto/log_ingestor.proto

package logging

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	LogIngestor_CreateLogParserFormat_FullMethodName = "/logging.LogIngestor/CreateLogParserFormat"
)

// LogIngestorClient is the client API for LogIngestor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogIngestorClient interface {
	CreateLogParserFormat(ctx context.Context, in *ParserFormatRequest, opts ...grpc.CallOption) (*ParserFormatResponse, error)
}

type logIngestorClient struct {
	cc grpc.ClientConnInterface
}

func NewLogIngestorClient(cc grpc.ClientConnInterface) LogIngestorClient {
	return &logIngestorClient{cc}
}

func (c *logIngestorClient) CreateLogParserFormat(ctx context.Context, in *ParserFormatRequest, opts ...grpc.CallOption) (*ParserFormatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ParserFormatResponse)
	err := c.cc.Invoke(ctx, LogIngestor_CreateLogParserFormat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogIngestorServer is the server API for LogIngestor service.
// All implementations must embed UnimplementedLogIngestorServer
// for forward compatibility.
type LogIngestorServer interface {
	CreateLogParserFormat(context.Context, *ParserFormatRequest) (*ParserFormatResponse, error)
	mustEmbedUnimplementedLogIngestorServer()
}

// UnimplementedLogIngestorServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLogIngestorServer struct{}

func (UnimplementedLogIngestorServer) CreateLogParserFormat(context.Context, *ParserFormatRequest) (*ParserFormatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLogParserFormat not implemented")
}
func (UnimplementedLogIngestorServer) mustEmbedUnimplementedLogIngestorServer() {}
func (UnimplementedLogIngestorServer) testEmbeddedByValue()                     {}

// UnsafeLogIngestorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogIngestorServer will
// result in compilation errors.
type UnsafeLogIngestorServer interface {
	mustEmbedUnimplementedLogIngestorServer()
}

func RegisterLogIngestorServer(s grpc.ServiceRegistrar, srv LogIngestorServer) {
	// If the following call pancis, it indicates UnimplementedLogIngestorServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LogIngestor_ServiceDesc, srv)
}

func _LogIngestor_CreateLogParserFormat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParserFormatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogIngestorServer).CreateLogParserFormat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogIngestor_CreateLogParserFormat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogIngestorServer).CreateLogParserFormat(ctx, req.(*ParserFormatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LogIngestor_ServiceDesc is the grpc.ServiceDesc for LogIngestor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogIngestor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "logging.LogIngestor",
	HandlerType: (*LogIngestorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLogParserFormat",
			Handler:    _LogIngestor_CreateLogParserFormat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/log_ingestor.proto",
}
