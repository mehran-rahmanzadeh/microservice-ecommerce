package server

import (
	"context"
	domain "github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/base"
	"github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/proto"
)

type userAddressServer struct {
	proto.UnimplementedUserAddressServer
	addressUsecase domain.AddressUsecase
}

func NewUserAddressServer(addressUsecase domain.AddressUsecase) *userAddressServer {
	return &userAddressServer{
		addressUsecase: addressUsecase,
	}
}

func (s userAddressServer) ListAddress(input *proto.ListAddressInput, stream proto.UserAddress_ListAddressServer) error {
	addresses, err := s.addressUsecase.AddressList(context.Background(), input)
	if err != nil {
		return err
	}
	for _, address := range addresses {
		if err := stream.Send(address); err != nil {
			return err
		}
	}
	return nil
}
func (s userAddressServer) AddressDetail(ctx context.Context, input *proto.RetrieveDeleteAddressInput) (*proto.Address, error) {
	return s.addressUsecase.AddressDetail(ctx, input)
}
func (s userAddressServer) CreateAddress(ctx context.Context, input *proto.CreateUpdateAddressInput) (*proto.Address, error) {
	return s.addressUsecase.CreateAddress(ctx, input)
}
func (s userAddressServer) UpdateAddress(ctx context.Context, input *proto.CreateUpdateAddressInput) (*proto.Address, error) {
	return s.addressUsecase.UpdateAddress(ctx, input)
}
func (s userAddressServer) DeleteAddress(ctx context.Context, input *proto.RetrieveDeleteAddressInput) (*proto.AddressDeleted, error) {
	return s.addressUsecase.DeleteAddress(ctx, input)
}
