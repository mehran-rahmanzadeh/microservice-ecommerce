package server

import (
	"context"
	domain "github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/base"
	"github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/proto"
)

type registerUserServer struct {
	proto.UnimplementedUserRegisterServer
	registerUsecase domain.RegisterUsecase
}

func NewRegisterUserServer(registerUserUsecase domain.RegisterUsecase) *registerUserServer {
	return &registerUserServer{registerUsecase: registerUserUsecase}
}

func (s *registerUserServer) Register(ctx context.Context, input *proto.RegisterInput) (*proto.User, error) {
	return s.registerUsecase.RegisterUser(ctx, input)
}