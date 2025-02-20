// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: telegraph.proto

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
	TelegraphService_Dispatch_FullMethodName       = "/local.grpc.telegraph.TelegraphService/Dispatch"
	TelegraphService_DispatchUnary_FullMethodName  = "/local.grpc.telegraph.TelegraphService/DispatchUnary"
	TelegraphService_DispatchStream_FullMethodName = "/local.grpc.telegraph.TelegraphService/DispatchStream"
	TelegraphService_Subscribe_FullMethodName      = "/local.grpc.telegraph.TelegraphService/Subscribe"
)

// TelegraphServiceClient is the client API for TelegraphService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Telegraph service ... telegraph sang a song about the world outside!
type TelegraphServiceClient interface {
	// Dispatch a communique or a stream of 'em ...
	Dispatch(ctx context.Context, in *Communique, opts ...grpc.CallOption) (*Response, error)
	DispatchUnary(ctx context.Context, in *Communique, opts ...grpc.CallOption) (*Response, error)
	DispatchStream(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[Communique, Response], error)
	// Subscribe and receive a stream of publications ...
	Subscribe(ctx context.Context, in *Communique, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Response], error)
}

type telegraphServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTelegraphServiceClient(cc grpc.ClientConnInterface) TelegraphServiceClient {
	return &telegraphServiceClient{cc}
}

func (c *telegraphServiceClient) Dispatch(ctx context.Context, in *Communique, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, TelegraphService_Dispatch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *telegraphServiceClient) DispatchUnary(ctx context.Context, in *Communique, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, TelegraphService_DispatchUnary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *telegraphServiceClient) DispatchStream(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[Communique, Response], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &TelegraphService_ServiceDesc.Streams[0], TelegraphService_DispatchStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Communique, Response]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TelegraphService_DispatchStreamClient = grpc.ClientStreamingClient[Communique, Response]

func (c *telegraphServiceClient) Subscribe(ctx context.Context, in *Communique, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Response], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &TelegraphService_ServiceDesc.Streams[1], TelegraphService_Subscribe_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Communique, Response]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TelegraphService_SubscribeClient = grpc.ServerStreamingClient[Response]

// TelegraphServiceServer is the server API for TelegraphService service.
// All implementations must embed UnimplementedTelegraphServiceServer
// for forward compatibility.
//
// Telegraph service ... telegraph sang a song about the world outside!
type TelegraphServiceServer interface {
	// Dispatch a communique or a stream of 'em ...
	Dispatch(context.Context, *Communique) (*Response, error)
	DispatchUnary(context.Context, *Communique) (*Response, error)
	DispatchStream(grpc.ClientStreamingServer[Communique, Response]) error
	// Subscribe and receive a stream of publications ...
	Subscribe(*Communique, grpc.ServerStreamingServer[Response]) error
	mustEmbedUnimplementedTelegraphServiceServer()
}

// UnimplementedTelegraphServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTelegraphServiceServer struct{}

func (UnimplementedTelegraphServiceServer) Dispatch(context.Context, *Communique) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Dispatch not implemented")
}
func (UnimplementedTelegraphServiceServer) DispatchUnary(context.Context, *Communique) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DispatchUnary not implemented")
}
func (UnimplementedTelegraphServiceServer) DispatchStream(grpc.ClientStreamingServer[Communique, Response]) error {
	return status.Errorf(codes.Unimplemented, "method DispatchStream not implemented")
}
func (UnimplementedTelegraphServiceServer) Subscribe(*Communique, grpc.ServerStreamingServer[Response]) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedTelegraphServiceServer) mustEmbedUnimplementedTelegraphServiceServer() {}
func (UnimplementedTelegraphServiceServer) testEmbeddedByValue()                          {}

// UnsafeTelegraphServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TelegraphServiceServer will
// result in compilation errors.
type UnsafeTelegraphServiceServer interface {
	mustEmbedUnimplementedTelegraphServiceServer()
}

func RegisterTelegraphServiceServer(s grpc.ServiceRegistrar, srv TelegraphServiceServer) {
	// If the following call pancis, it indicates UnimplementedTelegraphServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TelegraphService_ServiceDesc, srv)
}

func _TelegraphService_Dispatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Communique)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelegraphServiceServer).Dispatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TelegraphService_Dispatch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelegraphServiceServer).Dispatch(ctx, req.(*Communique))
	}
	return interceptor(ctx, in, info, handler)
}

func _TelegraphService_DispatchUnary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Communique)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelegraphServiceServer).DispatchUnary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TelegraphService_DispatchUnary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelegraphServiceServer).DispatchUnary(ctx, req.(*Communique))
	}
	return interceptor(ctx, in, info, handler)
}

func _TelegraphService_DispatchStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TelegraphServiceServer).DispatchStream(&grpc.GenericServerStream[Communique, Response]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TelegraphService_DispatchStreamServer = grpc.ClientStreamingServer[Communique, Response]

func _TelegraphService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Communique)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TelegraphServiceServer).Subscribe(m, &grpc.GenericServerStream[Communique, Response]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TelegraphService_SubscribeServer = grpc.ServerStreamingServer[Response]

// TelegraphService_ServiceDesc is the grpc.ServiceDesc for TelegraphService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TelegraphService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "local.grpc.telegraph.TelegraphService",
	HandlerType: (*TelegraphServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Dispatch",
			Handler:    _TelegraphService_Dispatch_Handler,
		},
		{
			MethodName: "DispatchUnary",
			Handler:    _TelegraphService_DispatchUnary_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DispatchStream",
			Handler:       _TelegraphService_DispatchStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Subscribe",
			Handler:       _TelegraphService_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "telegraph.proto",
}
