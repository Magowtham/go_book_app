// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: proto/book_app.proto

package book_app

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
	BookAppBackend_SignUpService_FullMethodName = "/book_app.BookAppBackend/SignUpService"
)

// BookAppBackendClient is the client API for BookAppBackend service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookAppBackendClient interface {
	SignUpService(ctx context.Context, in *SignUp, opts ...grpc.CallOption) (*SignUpStatus, error)
}

type bookAppBackendClient struct {
	cc grpc.ClientConnInterface
}

func NewBookAppBackendClient(cc grpc.ClientConnInterface) BookAppBackendClient {
	return &bookAppBackendClient{cc}
}

func (c *bookAppBackendClient) SignUpService(ctx context.Context, in *SignUp, opts ...grpc.CallOption) (*SignUpStatus, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SignUpStatus)
	err := c.cc.Invoke(ctx, BookAppBackend_SignUpService_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookAppBackendServer is the server API for BookAppBackend service.
// All implementations must embed UnimplementedBookAppBackendServer
// for forward compatibility.
type BookAppBackendServer interface {
	SignUpService(context.Context, *SignUp) (*SignUpStatus, error)
	mustEmbedUnimplementedBookAppBackendServer()
}

// UnimplementedBookAppBackendServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBookAppBackendServer struct{}

func (UnimplementedBookAppBackendServer) SignUpService(context.Context, *SignUp) (*SignUpStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUpService not implemented")
}
func (UnimplementedBookAppBackendServer) mustEmbedUnimplementedBookAppBackendServer() {}
func (UnimplementedBookAppBackendServer) testEmbeddedByValue()                        {}

// UnsafeBookAppBackendServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookAppBackendServer will
// result in compilation errors.
type UnsafeBookAppBackendServer interface {
	mustEmbedUnimplementedBookAppBackendServer()
}

func RegisterBookAppBackendServer(s grpc.ServiceRegistrar, srv BookAppBackendServer) {
	// If the following call pancis, it indicates UnimplementedBookAppBackendServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BookAppBackend_ServiceDesc, srv)
}

func _BookAppBackend_SignUpService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookAppBackendServer).SignUpService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookAppBackend_SignUpService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookAppBackendServer).SignUpService(ctx, req.(*SignUp))
	}
	return interceptor(ctx, in, info, handler)
}

// BookAppBackend_ServiceDesc is the grpc.ServiceDesc for BookAppBackend service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookAppBackend_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "book_app.BookAppBackend",
	HandlerType: (*BookAppBackendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUpService",
			Handler:    _BookAppBackend_SignUpService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/book_app.proto",
}
