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

// UserRegisterClient is the client API for UserRegister service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserRegisterClient interface {
	Register(ctx context.Context, in *RegisterInput, opts ...grpc.CallOption) (*User, error)
	ValidateCredentials(ctx context.Context, in *ValidateCredentialsInput, opts ...grpc.CallOption) (*Validate, error)
	GetUserInfo(ctx context.Context, in *GetUserInfoInput, opts ...grpc.CallOption) (*User, error)
}

type userRegisterClient struct {
	cc grpc.ClientConnInterface
}

func NewUserRegisterClient(cc grpc.ClientConnInterface) UserRegisterClient {
	return &userRegisterClient{cc}
}

func (c *userRegisterClient) Register(ctx context.Context, in *RegisterInput, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/domain.UserRegister/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRegisterClient) ValidateCredentials(ctx context.Context, in *ValidateCredentialsInput, opts ...grpc.CallOption) (*Validate, error) {
	out := new(Validate)
	err := c.cc.Invoke(ctx, "/domain.UserRegister/ValidateCredentials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRegisterClient) GetUserInfo(ctx context.Context, in *GetUserInfoInput, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/domain.UserRegister/GetUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserRegisterServer is the server API for UserRegister service.
// All implementations must embed UnimplementedUserRegisterServer
// for forward compatibility
type UserRegisterServer interface {
	Register(context.Context, *RegisterInput) (*User, error)
	ValidateCredentials(context.Context, *ValidateCredentialsInput) (*Validate, error)
	GetUserInfo(context.Context, *GetUserInfoInput) (*User, error)
	mustEmbedUnimplementedUserRegisterServer()
}

// UnimplementedUserRegisterServer must be embedded to have forward compatible implementations.
type UnimplementedUserRegisterServer struct {
}

func (UnimplementedUserRegisterServer) Register(context.Context, *RegisterInput) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserRegisterServer) ValidateCredentials(context.Context, *ValidateCredentialsInput) (*Validate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateCredentials not implemented")
}
func (UnimplementedUserRegisterServer) GetUserInfo(context.Context, *GetUserInfoInput) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedUserRegisterServer) mustEmbedUnimplementedUserRegisterServer() {}

// UnsafeUserRegisterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserRegisterServer will
// result in compilation errors.
type UnsafeUserRegisterServer interface {
	mustEmbedUnimplementedUserRegisterServer()
}

func RegisterUserRegisterServer(s grpc.ServiceRegistrar, srv UserRegisterServer) {
	s.RegisterService(&UserRegister_ServiceDesc, srv)
}

func _UserRegister_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRegisterServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.UserRegister/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRegisterServer).Register(ctx, req.(*RegisterInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRegister_ValidateCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateCredentialsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRegisterServer).ValidateCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.UserRegister/ValidateCredentials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRegisterServer).ValidateCredentials(ctx, req.(*ValidateCredentialsInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRegister_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRegisterServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.UserRegister/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRegisterServer).GetUserInfo(ctx, req.(*GetUserInfoInput))
	}
	return interceptor(ctx, in, info, handler)
}

// UserRegister_ServiceDesc is the grpc.ServiceDesc for UserRegister service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserRegister_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "domain.UserRegister",
	HandlerType: (*UserRegisterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _UserRegister_Register_Handler,
		},
		{
			MethodName: "ValidateCredentials",
			Handler:    _UserRegister_ValidateCredentials_Handler,
		},
		{
			MethodName: "GetUserInfo",
			Handler:    _UserRegister_GetUserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "domain/proto/register.proto",
}
