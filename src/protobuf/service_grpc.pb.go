// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protobuf

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

// CryptoServiceClient is the client API for CryptoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CryptoServiceClient interface {
	CreateNewCrypto(ctx context.Context, in *NewCryptoRequest, opts ...grpc.CallOption) (*NewCryptoResponse, error)
	GetCryptos(ctx context.Context, in *ListCryptosRequest, opts ...grpc.CallOption) (*ListCryptosResponse, error)
}

type cryptoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCryptoServiceClient(cc grpc.ClientConnInterface) CryptoServiceClient {
	return &cryptoServiceClient{cc}
}

func (c *cryptoServiceClient) CreateNewCrypto(ctx context.Context, in *NewCryptoRequest, opts ...grpc.CallOption) (*NewCryptoResponse, error) {
	out := new(NewCryptoResponse)
	err := c.cc.Invoke(ctx, "/protobuf.CryptoService/CreateNewCrypto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) GetCryptos(ctx context.Context, in *ListCryptosRequest, opts ...grpc.CallOption) (*ListCryptosResponse, error) {
	out := new(ListCryptosResponse)
	err := c.cc.Invoke(ctx, "/protobuf.CryptoService/GetCryptos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CryptoServiceServer is the server API for CryptoService service.
// All implementations must embed UnimplementedCryptoServiceServer
// for forward compatibility
type CryptoServiceServer interface {
	CreateNewCrypto(context.Context, *NewCryptoRequest) (*NewCryptoResponse, error)
	GetCryptos(context.Context, *ListCryptosRequest) (*ListCryptosResponse, error)
	mustEmbedUnimplementedCryptoServiceServer()
}

// UnimplementedCryptoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCryptoServiceServer struct {
}

func (UnimplementedCryptoServiceServer) CreateNewCrypto(context.Context, *NewCryptoRequest) (*NewCryptoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewCrypto not implemented")
}
func (UnimplementedCryptoServiceServer) GetCryptos(context.Context, *ListCryptosRequest) (*ListCryptosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCryptos not implemented")
}
func (UnimplementedCryptoServiceServer) mustEmbedUnimplementedCryptoServiceServer() {}

// UnsafeCryptoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CryptoServiceServer will
// result in compilation errors.
type UnsafeCryptoServiceServer interface {
	mustEmbedUnimplementedCryptoServiceServer()
}

func RegisterCryptoServiceServer(s grpc.ServiceRegistrar, srv CryptoServiceServer) {
	s.RegisterService(&CryptoService_ServiceDesc, srv)
}

func _CryptoService_CreateNewCrypto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewCryptoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).CreateNewCrypto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.CryptoService/CreateNewCrypto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).CreateNewCrypto(ctx, req.(*NewCryptoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_GetCryptos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCryptosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).GetCryptos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.CryptoService/GetCryptos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).GetCryptos(ctx, req.(*ListCryptosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CryptoService_ServiceDesc is the grpc.ServiceDesc for CryptoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CryptoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.CryptoService",
	HandlerType: (*CryptoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNewCrypto",
			Handler:    _CryptoService_CreateNewCrypto_Handler,
		},
		{
			MethodName: "GetCryptos",
			Handler:    _CryptoService_GetCryptos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/service.proto",
}
