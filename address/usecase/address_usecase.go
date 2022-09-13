package usecase

import (
	"context"
	domain "github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/base"
	"github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type addressUsecase struct {
	dbRepo domain.AddressDBRepository
	grpcRepo domain.AddressGrpcRepository
}

func NewAddressUsecase(dbRepo domain.AddressDBRepository, grpcRepo domain.AddressGrpcRepository) domain.AddressUsecase {
	return &addressUsecase{dbRepo: dbRepo, grpcRepo: grpcRepo}
}

func (u addressUsecase) AddressList(ctx context.Context, input *proto.ListAddressInput) ([]*proto.Address, error) {
	user, err := u.grpcRepo.GetUser(ctx, &proto.ProfileInput{Token: input.Token})
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	addressListInput := domain.AddressList{
		AddressUser:      domain.AddressUser{UserID: uint(user.Id)},
	}
	addresses, err := u.dbRepo.FetchAddressList(ctx, &addressListInput)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	var protoAddresses []*proto.Address
	for _, address := range addresses {
		protoAddresses = append(protoAddresses, &proto.Address{
			Id:            uint64(address.ID),
			UserId:        uint64(address.UserID),
			PostalAddress: address.PostalAddress,
			PostalCode:    address.PostalCode,
		})
	}
	return protoAddresses, nil
}

func (u addressUsecase) AddressDetail(ctx context.Context, input *proto.RetrieveDeleteAddressInput) (*proto.Address, error) {
	user, err := u.grpcRepo.GetUser(ctx, &proto.ProfileInput{Token: input.Token})
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	retrieveAddressInput := domain.AddressRetrieveDelete{
		AddressUser:                domain.AddressUser{UserID: uint(user.Id)},
		RetrieveDeleteAddressInput: proto.RetrieveDeleteAddressInput{AddressId: input.AddressId},
	}
	address, err := u.dbRepo.GetAddress(ctx, &retrieveAddressInput)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	protoAddress := proto.Address{
		Id:            uint64(address.ID),
		UserId:        uint64(address.UserID),
		PostalAddress: address.PostalAddress,
		PostalCode:    address.PostalCode,
	}
	return &protoAddress, nil
}

func (u addressUsecase) CreateAddress(ctx context.Context, input *proto.CreateUpdateAddressInput) (*proto.Address, error) {
	user, err := u.grpcRepo.GetUser(ctx, &proto.ProfileInput{Token: input.Token})
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	createAddressInput := domain.AddressCreateUpdate{
		AddressUser:              domain.AddressUser{UserID: uint(user.Id)},
		CreateUpdateAddressInput: proto.CreateUpdateAddressInput{
			Token:         input.Token,
			PostalAddress: input.PostalAddress,
			PostalCode:    input.PostalCode,
		},
	}
	address, err := u.dbRepo.CreateAddress(ctx, &createAddressInput)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	protoAddress := proto.Address{
		Id:            uint64(address.ID),
		UserId:        uint64(address.UserID),
		PostalAddress: address.PostalAddress,
		PostalCode:    address.PostalCode,
	}
	return &protoAddress, nil
}

func (u addressUsecase) UpdateAddress(ctx context.Context, input *proto.CreateUpdateAddressInput) (*proto.Address, error) {
	user, err := u.grpcRepo.GetUser(ctx, &proto.ProfileInput{Token: input.Token})
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	createAddressInput := domain.AddressCreateUpdate{
		AddressUser:              domain.AddressUser{UserID: uint(user.Id)},
		CreateUpdateAddressInput: proto.CreateUpdateAddressInput{
			Token:         input.Token,
			AddressId: 	   input.AddressId,
			PostalAddress: input.PostalAddress,
			PostalCode:    input.PostalCode,
		},
	}
	address, err := u.dbRepo.UpdateAddress(ctx, &createAddressInput)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	protoAddress := proto.Address{
		Id:            uint64(address.ID),
		UserId:        uint64(address.UserID),
		PostalAddress: address.PostalAddress,
		PostalCode:    address.PostalCode,
	}
	return &protoAddress, nil
}

func (u addressUsecase) DeleteAddress(ctx context.Context, input *proto.RetrieveDeleteAddressInput) (*proto.AddressDeleted, error) {
	user, err := u.grpcRepo.GetUser(ctx, &proto.ProfileInput{Token: input.Token})
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	deleteAddressInput := domain.AddressRetrieveDelete{
		AddressUser:                domain.AddressUser{UserID: uint(user.Id)},
		RetrieveDeleteAddressInput: proto.RetrieveDeleteAddressInput{
			AddressId: input.AddressId,
		},
	}
	accepted, err := u.dbRepo.DeleteAddress(ctx, &deleteAddressInput)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	protoAccepted := proto.AddressDeleted{Success: accepted}
	return &protoAccepted, nil
}