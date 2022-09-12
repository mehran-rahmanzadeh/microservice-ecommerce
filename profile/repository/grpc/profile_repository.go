package repository

import (
	"context"
	domain "github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/base"
	"github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/proto"
)

type profileRepository struct {
	registerClient proto.UserRegisterClient
}

func NewProfileRepository(registerClient proto.UserRegisterClient) domain.ProfileRepository {
	return &profileRepository{
		registerClient: registerClient,
	}
}

func (r profileRepository) GetUser(ctx context.Context, input *proto.GetUserInfoInput) (*proto.User, error) {
	return r.registerClient.GetUserInfo(ctx, input)
}
