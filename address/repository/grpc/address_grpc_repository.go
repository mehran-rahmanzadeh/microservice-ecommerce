package repository

import (
	"context"
	domain "github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/base"
	"github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/proto"
)

type addressGrpcRepository struct {
	client proto.UserProfileClient
}

func NewAddressGrpcRepository(client proto.UserProfileClient) domain.AddressGrpcRepository {
	return &addressGrpcRepository{client}
}

func (r addressGrpcRepository) GetUser(ctx context.Context, input *proto.ProfileInput) (*proto.Profile, error) {
	return r.client.GetUserProfile(ctx, input)
}