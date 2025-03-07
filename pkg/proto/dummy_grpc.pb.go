// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: demo.proto

package proto

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
	DemoService_OneReqOneResp_FullMethodName   = "/demo.DemoService/OneReqOneResp"
	DemoService_OneReqManyResp_FullMethodName  = "/demo.DemoService/OneReqManyResp"
	DemoService_ManyReqOneResp_FullMethodName  = "/demo.DemoService/ManyReqOneResp"
	DemoService_ManyReqManyResp_FullMethodName = "/demo.DemoService/ManyReqManyResp"
	DemoService_UnaryFail_FullMethodName       = "/demo.DemoService/UnaryFail"
	DemoService_UnaryDeadline_FullMethodName   = "/demo.DemoService/UnaryDeadline"
)

// DemoServiceClient is the client API for DemoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DemoServiceClient interface {
	// Unary
	// C --> S
	// C <-- S
	OneReqOneResp(ctx context.Context, in *DemoRequest, opts ...grpc.CallOption) (*DemoResponse, error)
	// Server streaming
	// C --> S (once at start)
	// C < - S (stream)
	OneReqManyResp(ctx context.Context, in *DemoRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[DemoResponse], error)
	// Client streaming
	// C - > S (stream)
	// C <-- S (once at end)
	ManyReqOneResp(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[DemoRequest, DemoResponse], error)
	// Bi-directional streaming
	// C - > S (stream)
	// C < - S (stream)
	ManyReqManyResp(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[DemoRequest, DemoResponse], error)
	// Unary but that always fails with error `codes.InvalidArgument`
	UnaryFail(ctx context.Context, in *DemoRequest, opts ...grpc.CallOption) (*DemoResponse, error)
	// Unary that will answer after 5 seconds
	UnaryDeadline(ctx context.Context, in *DemoRequest, opts ...grpc.CallOption) (*DemoResponse, error)
}

type demoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDemoServiceClient(cc grpc.ClientConnInterface) DemoServiceClient {
	return &demoServiceClient{cc}
}

func (c *demoServiceClient) OneReqOneResp(ctx context.Context, in *DemoRequest, opts ...grpc.CallOption) (*DemoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DemoResponse)
	err := c.cc.Invoke(ctx, DemoService_OneReqOneResp_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoServiceClient) OneReqManyResp(ctx context.Context, in *DemoRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[DemoResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &DemoService_ServiceDesc.Streams[0], DemoService_OneReqManyResp_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[DemoRequest, DemoResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type DemoService_OneReqManyRespClient = grpc.ServerStreamingClient[DemoResponse]

func (c *demoServiceClient) ManyReqOneResp(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[DemoRequest, DemoResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &DemoService_ServiceDesc.Streams[1], DemoService_ManyReqOneResp_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[DemoRequest, DemoResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type DemoService_ManyReqOneRespClient = grpc.ClientStreamingClient[DemoRequest, DemoResponse]

func (c *demoServiceClient) ManyReqManyResp(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[DemoRequest, DemoResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &DemoService_ServiceDesc.Streams[2], DemoService_ManyReqManyResp_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[DemoRequest, DemoResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type DemoService_ManyReqManyRespClient = grpc.BidiStreamingClient[DemoRequest, DemoResponse]

func (c *demoServiceClient) UnaryFail(ctx context.Context, in *DemoRequest, opts ...grpc.CallOption) (*DemoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DemoResponse)
	err := c.cc.Invoke(ctx, DemoService_UnaryFail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoServiceClient) UnaryDeadline(ctx context.Context, in *DemoRequest, opts ...grpc.CallOption) (*DemoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DemoResponse)
	err := c.cc.Invoke(ctx, DemoService_UnaryDeadline_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DemoServiceServer is the server API for DemoService service.
// All implementations must embed UnimplementedDemoServiceServer
// for forward compatibility.
type DemoServiceServer interface {
	// Unary
	// C --> S
	// C <-- S
	OneReqOneResp(context.Context, *DemoRequest) (*DemoResponse, error)
	// Server streaming
	// C --> S (once at start)
	// C < - S (stream)
	OneReqManyResp(*DemoRequest, grpc.ServerStreamingServer[DemoResponse]) error
	// Client streaming
	// C - > S (stream)
	// C <-- S (once at end)
	ManyReqOneResp(grpc.ClientStreamingServer[DemoRequest, DemoResponse]) error
	// Bi-directional streaming
	// C - > S (stream)
	// C < - S (stream)
	ManyReqManyResp(grpc.BidiStreamingServer[DemoRequest, DemoResponse]) error
	// Unary but that always fails with error `codes.InvalidArgument`
	UnaryFail(context.Context, *DemoRequest) (*DemoResponse, error)
	// Unary that will answer after 5 seconds
	UnaryDeadline(context.Context, *DemoRequest) (*DemoResponse, error)
	mustEmbedUnimplementedDemoServiceServer()
}

// UnimplementedDemoServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDemoServiceServer struct{}

func (UnimplementedDemoServiceServer) OneReqOneResp(context.Context, *DemoRequest) (*DemoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OneReqOneResp not implemented")
}
func (UnimplementedDemoServiceServer) OneReqManyResp(*DemoRequest, grpc.ServerStreamingServer[DemoResponse]) error {
	return status.Errorf(codes.Unimplemented, "method OneReqManyResp not implemented")
}
func (UnimplementedDemoServiceServer) ManyReqOneResp(grpc.ClientStreamingServer[DemoRequest, DemoResponse]) error {
	return status.Errorf(codes.Unimplemented, "method ManyReqOneResp not implemented")
}
func (UnimplementedDemoServiceServer) ManyReqManyResp(grpc.BidiStreamingServer[DemoRequest, DemoResponse]) error {
	return status.Errorf(codes.Unimplemented, "method ManyReqManyResp not implemented")
}
func (UnimplementedDemoServiceServer) UnaryFail(context.Context, *DemoRequest) (*DemoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnaryFail not implemented")
}
func (UnimplementedDemoServiceServer) UnaryDeadline(context.Context, *DemoRequest) (*DemoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnaryDeadline not implemented")
}
func (UnimplementedDemoServiceServer) mustEmbedUnimplementedDemoServiceServer() {}
func (UnimplementedDemoServiceServer) testEmbeddedByValue()                      {}

// UnsafeDemoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DemoServiceServer will
// result in compilation errors.
type UnsafeDemoServiceServer interface {
	mustEmbedUnimplementedDemoServiceServer()
}

func RegisterDemoServiceServer(s grpc.ServiceRegistrar, srv DemoServiceServer) {
	// If the following call pancis, it indicates UnimplementedDemoServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DemoService_ServiceDesc, srv)
}

func _DemoService_OneReqOneResp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DemoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServiceServer).OneReqOneResp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DemoService_OneReqOneResp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServiceServer).OneReqOneResp(ctx, req.(*DemoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DemoService_OneReqManyResp_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DemoRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DemoServiceServer).OneReqManyResp(m, &grpc.GenericServerStream[DemoRequest, DemoResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type DemoService_OneReqManyRespServer = grpc.ServerStreamingServer[DemoResponse]

func _DemoService_ManyReqOneResp_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DemoServiceServer).ManyReqOneResp(&grpc.GenericServerStream[DemoRequest, DemoResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type DemoService_ManyReqOneRespServer = grpc.ClientStreamingServer[DemoRequest, DemoResponse]

func _DemoService_ManyReqManyResp_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DemoServiceServer).ManyReqManyResp(&grpc.GenericServerStream[DemoRequest, DemoResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type DemoService_ManyReqManyRespServer = grpc.BidiStreamingServer[DemoRequest, DemoResponse]

func _DemoService_UnaryFail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DemoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServiceServer).UnaryFail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DemoService_UnaryFail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServiceServer).UnaryFail(ctx, req.(*DemoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DemoService_UnaryDeadline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DemoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServiceServer).UnaryDeadline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DemoService_UnaryDeadline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServiceServer).UnaryDeadline(ctx, req.(*DemoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DemoService_ServiceDesc is the grpc.ServiceDesc for DemoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DemoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "demo.DemoService",
	HandlerType: (*DemoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OneReqOneResp",
			Handler:    _DemoService_OneReqOneResp_Handler,
		},
		{
			MethodName: "UnaryFail",
			Handler:    _DemoService_UnaryFail_Handler,
		},
		{
			MethodName: "UnaryDeadline",
			Handler:    _DemoService_UnaryDeadline_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "OneReqManyResp",
			Handler:       _DemoService_OneReqManyResp_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ManyReqOneResp",
			Handler:       _DemoService_ManyReqOneResp_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ManyReqManyResp",
			Handler:       _DemoService_ManyReqManyResp_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "demo.proto",
}
