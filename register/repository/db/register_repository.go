package repository

import (
	"context"
	domain "github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/base"
	"github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/model"
	"github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/proto"
	"gorm.io/gorm"
)

type RegisterRepository struct {
	DB *gorm.DB
}

func NewRegisterRepository(db *gorm.DB) domain.RegisterRepository {
	return &RegisterRepository{db}
}

func (r *RegisterRepository) Register(ctx context.Context, input *proto.RegisterInput) (model.User, error) {
	user := model.User{Email: input.Email, Password: input.Password}
	tx := r.DB.Create(&user)
	return user, tx.Error
}
