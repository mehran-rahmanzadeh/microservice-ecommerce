package domain

import (
	"context"
	"github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/model"
	"github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/proto"
)

type AddressDBRepository interface {
	FetchAddressList(ctx context.Context, input *AddressList) ([]*model.Address, error)
	GetAddress(ctx context.Context, input *AddressRetrieveDelete) (*model.Address, error)
	CreateAddress(ctx context.Context, input *AddressCreateUpdate) (*model.Address, error)
	UpdateAddress(ctx context.Context, input *AddressCreateUpdate) (*model.Address, error)
	DeleteAddress(ctx context.Context, input *AddressRetrieveDelete) (bool, error)
}

type AddressGrpcRepository interface {
	GetUser(ctx context.Context, input *proto.ProfileInput) (*proto.Profile, error)
}

type AddressUsecase interface {
	AddressList(ctx context.Context, input *proto.ListAddressInput) ([]*proto.Address, error)
	AddressDetail(ctx context.Context, input *proto.RetrieveDeleteAddressInput) (*proto.Address, error)
	CreateAddress(ctx context.Context, input *proto.CreateUpdateAddressInput) (*proto.Address, error)
	UpdateAddress(ctx context.Context, input *proto.CreateUpdateAddressInput) (*proto.Address, error)
	DeleteAddress(ctx context.Context, input *proto.RetrieveDeleteAddressInput) (*proto.AddressDeleted, error)
}
