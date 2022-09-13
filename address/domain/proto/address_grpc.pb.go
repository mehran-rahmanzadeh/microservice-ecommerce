// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// UserAddressClient is the client API for UserAddress service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserAddressClient interface {
	ListAddress(ctx context.Context, in *ListAddressInput, opts ...grpc.CallOption) (UserAddress_ListAddressClient, error)
	AddressDetail(ctx context.Context, in *RetrieveDeleteAddressInput, opts ...grpc.CallOption) (*Address, error)
	CreateAddress(ctx context.Context, in *CreateUpdateAddressInput, opts ...grpc.CallOption) (*Address, error)
	UpdateAddress(ctx context.Context, in *CreateUpdateAddressInput, opts ...grpc.CallOption) (*Address, error)
	DeleteAddress(ctx context.Context, in *RetrieveDeleteAddressInput, opts ...grpc.CallOption) (*AddressDeleted, error)
}

type userAddressClient struct {
	cc grpc.ClientConnInterface
}

func NewUserAddressClient(cc grpc.ClientConnInterface) UserAddressClient {
	return &userAddressClient{cc}
}

func (c *userAddressClient) ListAddress(ctx context.Context, in *ListAddressInput, opts ...grpc.CallOption) (UserAddress_ListAddressClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserAddress_ServiceDesc.Streams[0], "/domain.UserAddress/ListAddress", opts...)
	if err != nil {
		return nil, err
	}
	x := &userAddressListAddressClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UserAddress_ListAddressClient interface {
	Recv() (*Address, error)
	grpc.ClientStream
}

type userAddressListAddressClient struct {
	grpc.ClientStream
}

func (x *userAddressListAddressClient) Recv() (*Address, error) {
	m := new(Address)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userAddressClient) AddressDetail(ctx context.Context, in *RetrieveDeleteAddressInput, opts ...grpc.CallOption) (*Address, error) {
	out := new(Address)
	err := c.cc.Invoke(ctx, "/domain.UserAddress/AddressDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAddressClient) CreateAddress(ctx context.Context, in *CreateUpdateAddressInput, opts ...grpc.CallOption) (*Address, error) {
	out := new(Address)
	err := c.cc.Invoke(ctx, "/domain.UserAddress/CreateAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAddressClient) UpdateAddress(ctx context.Context, in *CreateUpdateAddressInput, opts ...grpc.CallOption) (*Address, error) {
	out := new(Address)
	err := c.cc.Invoke(ctx, "/domain.UserAddress/UpdateAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAddressClient) DeleteAddress(ctx context.Context, in *RetrieveDeleteAddressInput, opts ...grpc.CallOption) (*AddressDeleted, error) {
	out := new(AddressDeleted)
	err := c.cc.Invoke(ctx, "/domain.UserAddress/DeleteAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserAddressServer is the server API for UserAddress service.
// All implementations must embed UnimplementedUserAddressServer
// for forward compatibility
type UserAddressServer interface {
	ListAddress(*ListAddressInput, UserAddress_ListAddressServer) error
	AddressDetail(context.Context, *RetrieveDeleteAddressInput) (*Address, error)
	CreateAddress(context.Context, *CreateUpdateAddressInput) (*Address, error)
	UpdateAddress(context.Context, *CreateUpdateAddressInput) (*Address, error)
	DeleteAddress(context.Context, *RetrieveDeleteAddressInput) (*AddressDeleted, error)
	mustEmbedUnimplementedUserAddressServer()
}

// UnimplementedUserAddressServer must be embedded to have forward compatible implementations.
type UnimplementedUserAddressServer struct {
}

func (UnimplementedUserAddressServer) ListAddress(*ListAddressInput, UserAddress_ListAddressServer) error {
	return status.Errorf(codes.Unimplemented, "method ListAddress not implemented")
}
func (UnimplementedUserAddressServer) AddressDetail(context.Context, *RetrieveDeleteAddressInput) (*Address, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddressDetail not implemented")
}
func (UnimplementedUserAddressServer) CreateAddress(context.Context, *CreateUpdateAddressInput) (*Address, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAddress not implemented")
}
func (UnimplementedUserAddressServer) UpdateAddress(context.Context, *CreateUpdateAddressInput) (*Address, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAddress not implemented")
}
func (UnimplementedUserAddressServer) DeleteAddress(context.Context, *RetrieveDeleteAddressInput) (*AddressDeleted, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAddress not implemented")
}
func (UnimplementedUserAddressServer) mustEmbedUnimplementedUserAddressServer() {}

// UnsafeUserAddressServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserAddressServer will
// result in compilation errors.
type UnsafeUserAddressServer interface {
	mustEmbedUnimplementedUserAddressServer()
}

func RegisterUserAddressServer(s grpc.ServiceRegistrar, srv UserAddressServer) {
	s.RegisterService(&UserAddress_ServiceDesc, srv)
}

func _UserAddress_ListAddress_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListAddressInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserAddressServer).ListAddress(m, &userAddressListAddressServer{stream})
}

type UserAddress_ListAddressServer interface {
	Send(*Address) error
	grpc.ServerStream
}

type userAddressListAddressServer struct {
	grpc.ServerStream
}

func (x *userAddressListAddressServer) Send(m *Address) error {
	return x.ServerStream.SendMsg(m)
}

func _UserAddress_AddressDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveDeleteAddressInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAddressServer).AddressDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.UserAddress/AddressDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAddressServer).AddressDetail(ctx, req.(*RetrieveDeleteAddressInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAddress_CreateAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUpdateAddressInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAddressServer).CreateAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.UserAddress/CreateAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAddressServer).CreateAddress(ctx, req.(*CreateUpdateAddressInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAddress_UpdateAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUpdateAddressInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAddressServer).UpdateAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.UserAddress/UpdateAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAddressServer).UpdateAddress(ctx, req.(*CreateUpdateAddressInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAddress_DeleteAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveDeleteAddressInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAddressServer).DeleteAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.UserAddress/DeleteAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAddressServer).DeleteAddress(ctx, req.(*RetrieveDeleteAddressInput))
	}
	return interceptor(ctx, in, info, handler)
}

// UserAddress_ServiceDesc is the grpc.ServiceDesc for UserAddress service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserAddress_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "domain.UserAddress",
	HandlerType: (*UserAddressServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddressDetail",
			Handler:    _UserAddress_AddressDetail_Handler,
		},
		{
			MethodName: "CreateAddress",
			Handler:    _UserAddress_CreateAddress_Handler,
		},
		{
			MethodName: "UpdateAddress",
			Handler:    _UserAddress_UpdateAddress_Handler,
		},
		{
			MethodName: "DeleteAddress",
			Handler:    _UserAddress_DeleteAddress_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListAddress",
			Handler:       _UserAddress_ListAddress_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "domain/proto/address.proto",
}
