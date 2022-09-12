package domain

import (
	"context"
	"github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/proto"
)

type ProfileRepository interface {
	GetUser(ctx context.Context, input *proto.GetUserInfoInput) (*proto.User, error)
}

type ProfileUsecase interface {
	GetUserProfile(ctx context.Context, input *proto.ProfileInput) (*proto.Profile, error)
}
