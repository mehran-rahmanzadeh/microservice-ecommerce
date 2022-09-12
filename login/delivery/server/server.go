package server

import (
	"context"
	domain "github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/base"
	"github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/proto"
)

type loginUserServer struct {
	proto.UnimplementedUserLoginServer
	loginUsecase domain.LoginUsecase
}

func NewLoginUserServer(loginUsecase domain.LoginUsecase) *loginUserServer {
	return &loginUserServer{loginUsecase: loginUsecase}
}

func (s loginUserServer) Login(ctx context.Context, input *proto.LoginInput) (*proto.Token, error) {
	return s.loginUsecase.LoginUser(ctx, input)
}

func (s loginUserServer) Authenticate(ctx context.Context, input *proto.TokenInput) (*proto.Authentication, error) {
	return s.loginUsecase.AuthenticateToken(ctx, input)
}
