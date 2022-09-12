package server

import (
	"context"
	domain "github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/base"
	"github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/proto"
)

type userProfileServer struct {
	proto.UnimplementedUserProfileServer
	profileUsecase domain.ProfileUsecase
}

func NewUserProfileServer(profileUsecase domain.ProfileUsecase) *userProfileServer {
	return &userProfileServer{profileUsecase: profileUsecase}
}

func (s userProfileServer) GetUserProfile(ctx context.Context, input *proto.ProfileInput) (*proto.Profile, error) {
	return s.profileUsecase.GetUserProfile(ctx, input)
}
