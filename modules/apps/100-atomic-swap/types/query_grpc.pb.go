// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package types

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

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Params queries all parameters of the ibc-transfer module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// EscrowAddress returns the escrow address for a particular port and channel id.
	EscrowAddress(ctx context.Context, in *QueryEscrowAddressRequest, opts ...grpc.CallOption) (*QueryEscrowAddressResponse, error)
	Orders(ctx context.Context, in *QueryOrdersRequest, opts ...grpc.CallOption) (*QueryOrdersResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/ibc.applications.atomic_swap.v1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) EscrowAddress(ctx context.Context, in *QueryEscrowAddressRequest, opts ...grpc.CallOption) (*QueryEscrowAddressResponse, error) {
	out := new(QueryEscrowAddressResponse)
	err := c.cc.Invoke(ctx, "/ibc.applications.atomic_swap.v1.Query/EscrowAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Orders(ctx context.Context, in *QueryOrdersRequest, opts ...grpc.CallOption) (*QueryOrdersResponse, error) {
	out := new(QueryOrdersResponse)
	err := c.cc.Invoke(ctx, "/ibc.applications.atomic_swap.v1.Query/Orders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations should embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Params queries all parameters of the ibc-transfer module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// EscrowAddress returns the escrow address for a particular port and channel id.
	EscrowAddress(context.Context, *QueryEscrowAddressRequest) (*QueryEscrowAddressResponse, error)
	Orders(context.Context, *QueryOrdersRequest) (*QueryOrdersResponse, error)
}

// UnimplementedQueryServer should be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) EscrowAddress(context.Context, *QueryEscrowAddressRequest) (*QueryEscrowAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EscrowAddress not implemented")
}
func (UnimplementedQueryServer) Orders(context.Context, *QueryOrdersRequest) (*QueryOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Orders not implemented")
}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.applications.atomic_swap.v1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_EscrowAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEscrowAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).EscrowAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.applications.atomic_swap.v1.Query/EscrowAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).EscrowAddress(ctx, req.(*QueryEscrowAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Orders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Orders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.applications.atomic_swap.v1.Query/Orders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Orders(ctx, req.(*QueryOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ibc.applications.atomic_swap.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "EscrowAddress",
			Handler:    _Query_EscrowAddress_Handler,
		},
		{
			MethodName: "Orders",
			Handler:    _Query_Orders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ibc/applications/atomic_swap/v1/query.proto",
}
