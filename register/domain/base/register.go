package domain

import (
	"context"
	domain "github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/model"
	"github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/proto"
)

type RegisterRepository interface {
	Register(ctx context.Context, input *proto.RegisterInput) (domain.User, error)
	Authenticate(ctx context.Context, input *proto.ValidateCredentialsInput) (bool, error)
}

type RegisterUsecase interface {
	RegisterUser(ctx context.Context, input *proto.RegisterInput) (*proto.User, error)
	AuthenticateUser(ctx context.Context, input *proto.ValidateCredentialsInput) (*proto.Validate, error)
}
