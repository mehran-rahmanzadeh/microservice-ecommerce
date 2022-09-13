package repository

import (
	"context"
	domain "github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/base"
	"github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/model"
	"gorm.io/gorm"
)

type addressDBRepository struct {
	DB *gorm.DB
}

func NewAddressDBRepository(db *gorm.DB) domain.AddressDBRepository {
	return &addressDBRepository{db}
}

func (r addressDBRepository) FetchAddressList(ctx context.Context, input *domain.AddressList) ([]*model.Address, error) {
	var result []*model.Address
	tx := r.DB.Where("user_id = ?", input.UserID).Find(&result)
	return result, tx.Error
}

func (r addressDBRepository) GetAddress(ctx context.Context, input *domain.AddressRetrieveDelete) (*model.Address, error) {
	var result *model.Address
	tx := r.DB.Where("id = ? AND user_id = ?", input.AddressId, input.UserID).First(&result)
	return result, tx.Error
}

func (r addressDBRepository) CreateAddress(ctx context.Context, input *domain.AddressCreateUpdate) (*model.Address, error) {
	address := model.Address{
		UserID:        input.UserID,
		PostalAddress: input.PostalAddress,
		PostalCode:    input.PostalCode,
	}
	tx := r.DB.Create(&address)
	return &address, tx.Error
}

func (r addressDBRepository) UpdateAddress(ctx context.Context, input *domain.AddressCreateUpdate) (*model.Address, error) {
	address := model.Address{
		Model: gorm.Model{ID: uint(input.AddressId)},
	}
	tx := r.DB.Model(&address).Updates(model.Address{PostalCode: input.PostalCode, PostalAddress: input.PostalAddress})
	return &address, tx.Error
}

func (r addressDBRepository) DeleteAddress(ctx context.Context, input *domain.AddressRetrieveDelete) (bool, error) {
	address := model.Address{Model: gorm.Model{ID: uint(input.AddressId)}}
	tx := r.DB.Delete(&address)
	return tx.RowsAffected > 0, tx.Error
}
