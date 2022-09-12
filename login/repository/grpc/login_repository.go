package repository

import (
	"context"
	domain "github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/base"
	"github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/proto"
)

type loginRepository struct {
	client proto.UserRegisterClient
}

func NewLoginRepository(client proto.UserRegisterClient) domain.LoginRepository {
	return &loginRepository{client}
}

// Authenticate user based on credentials (uses register grpc service)
func (r *loginRepository) Authenticate(ctx context.Context, input *proto.ValidateCredentialsInput) (bool, error) {
	validation, err := r.client.ValidateCredentials(ctx, input)
	if err != nil {
		return false, err
	}
	return validation.IsValid, nil
}