package domain

import (
	"context"
	"github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/proto"
)

type LoginRepository interface {
	Authenticate(ctx context.Context, input *proto.ValidateCredentialsInput) (bool, error)
}

type LoginUsecase interface {
	LoginUser(ctx context.Context, input *proto.LoginInput) (*proto.Token, error)
	AuthenticateToken(ctx context.Context, input *proto.TokenInput) (*proto.Authentication, error)
}
