package domain

import "github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/proto"

type AddressUser struct {
	UserID uint
}

type AddressList struct {
	AddressUser
	proto.ListAddressInput
}

type AddressRetrieveDelete struct {
	AddressUser
	proto.RetrieveDeleteAddressInput
}

type AddressCreateUpdate struct {
	AddressUser
	proto.CreateUpdateAddressInput
}
