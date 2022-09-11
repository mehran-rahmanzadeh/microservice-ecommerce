package domain

import (
	"context"
	domain "github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/model"
	"github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/proto"
)

type RegisterRepository interface {
	Register(ctx context.Context, input *proto.RegisterInput) (domain.User, error)
}

type RegisterUsecase interface {
	RegisterUser(ctx context.Context, input *proto.RegisterInput) (*proto.User, error)
}
